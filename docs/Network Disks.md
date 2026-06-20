# Network Disks

Using a Network Attached Storage (NAS) device to centralize your gaming library is an excellent way to play across multiple systems without duplicating files. However, sharing games over a network requires knowing the limitations and choosing the right process for a fluid experience.

## Limitations & Protocol Considerations

When streaming games directly over a home network, choosing the correct network sharing protocol is vital for stability:

- While NFS (Network File System) is lightweight and popular for Linux-only server tasks, it lacks consistent Windows client support and can run into locking and permission issues in mixed environments, so it is **not recommended for network gaming storage**.
- SMB natively resolves these issues while offering transparent cross-platform support.

Please be aware that gaming with network disks has its own limitations due to network speeds:

- A standard 1 Gbps connection caps out at ~115 MB/s, so you need to calculate how many seconds the game assets will take to load.
- Small files, like retro games (GBA, PS1, ...) tend to work well on 1 Gbps because the assets are small and load almost instantly into RAM. Bandwidth is rarely saturated in such cases.
- Massive or uncompressed files, such as modern AAA titles or games with installs larger than 50 GB, will require a 2.5 Gbps to 10 Gbps network. Modern assets stream dynamically from storage while playing, so a slow network will cause texture pop-in or stuttering.
- Wi-Fi can also greatly affect performance — stick to Wi-Fi 6 or use cables when possible.
- While playing, network bandwidth can be saturated, especially on bigger games. When the network is saturated, it will affect the network speed for all devices.

If you're an advanced user with an SMB-capable NAS and a network that can handle the job, follow the instructions below to use network disks and folders in your setup.

## Setup on Linux

We will use SMB folder mount on systemd to handle the shared folders. This method works across traditional, server, and immutable distributions, but requires `cifs-utils` to be installed on your distribution. Once the required package is installed, use this script to set up the network folder with SMB:

```bash
# SMB remote details
SMB_SOURCE="//192.168.0.100/Shared"
SMB_WORKGROUP="WORKGROUP"
SMB_USERNAME="username"
SMB_PASSWORD="password"

# SMB host details
SMB_NAME="var-mnt-shared"
SMB_DESTINATION="/var/mnt/shared"
SMB_CREDENTIALS="/etc/samba/credentials"

# SMB options
SMB_OPTIONS="credentials=$SMB_CREDENTIALS,iocharset=utf8"
SMB_OPTIONS="$SMB_OPTIONS,uid=1000,gid=1000"
SMB_OPTIONS="$SMB_OPTIONS,file_mode=0775,dir_mode=0775"

# Only uncomment on SELinux-enabled distros (Fedora, RHEL, openSUSE, etc).
# On AppArmor distros (Ubuntu, Debian) this option isn't supported and
# the mount will fail with "Invalid argument" if it's included.
# SMB_OPTIONS="$SMB_OPTIONS,context=system_u:object_r:cifs_t:s0"

# Make sure destination exists
sudo mkdir -p "$SMB_DESTINATION"

# Create credentials file
sudo mkdir -p "$(dirname "$SMB_CREDENTIALS")"
sudo touch "$SMB_CREDENTIALS"
sudo chmod 600 "$SMB_CREDENTIALS"
cat << EOF | sudo tee "$SMB_CREDENTIALS" > /dev/null
domain=$SMB_WORKGROUP
username=$SMB_USERNAME
password=$SMB_PASSWORD
EOF

# Create systemd mount
cat << EOF | sudo tee "/etc/systemd/system/$SMB_NAME.mount" > /dev/null
[Unit]
Description=Mount SMB for $SMB_NAME
After=network-online.target
Wants=network-online.target

[Mount]
What=$SMB_SOURCE
Where=$SMB_DESTINATION
Type=cifs
Options=$SMB_OPTIONS

[Install]
WantedBy=multi-user.target
EOF

# Create systemd automount
cat << EOF | sudo tee "/etc/systemd/system/$SMB_NAME.automount" > /dev/null
[Unit]
Description=Automount SMB for $SMB_NAME

[Automount]
Where=$SMB_DESTINATION

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd to recognize both new unit files
sudo systemctl daemon-reload
```

Now, choose the best mount method based on your machine and start the service (pick only one of the two options below):

```bash
# Enable automount service - useful for desktop
# When using desktop, better to use automount
sudo systemctl disable "$SMB_NAME.mount"
sudo systemctl enable --now "$SMB_NAME.automount"
systemctl status "$SMB_NAME.automount"

# Enable mount service - useful for servers
sudo systemctl disable "$SMB_NAME.automount"
sudo systemctl enable --now "$SMB_NAME.mount"
systemctl status "$SMB_NAME.mount"
```

The folder now is the path of your network disk. To finish, create aliases to the target folder in your system account:

```bash
# Ensure the parent directory exists
mkdir -p $HOME/Games

# Create the symlinks cleanly
ln -s /var/mnt/shared/Games/ROMs $HOME/Games/ROMs
ln -s /var/mnt/shared/Games/BIOS $HOME/Games/BIOS
ln -s /var/mnt/shared/Games/State $HOME/Games/State
```

That is it! NiceDeck will be able to detect and parse your games from the network-mounted folders.

## Setup on Windows

Using the standard "**Map Network Drive**" feature in *File Explorer* often causes Windows to freeze during login if your network adapter isn't fast enough to secure an IP address immediately. The following script relies on **Windows Task Scheduler** and **PowerShell** to dynamically map your storage only after a successful system network handshake is validated.

Run **PowerShell** as an **Administrator** and execute the following script to set up the network disk:

```powershell
# SMB remote details
$SMBSource = "\\192.168.0.100\Shared"
$SMBServer = "192.168.0.100"
$SMBUsername = "username"
$SMBPassword = "password"

# SMB host details
$DriveLetter = "Z:"
$TaskScriptDir = "C:\Scripts"
$TaskScriptPath = "$TaskScriptDir\mount-smb-shared.ps1"
$TaskName = "Automount SMB Shared"

# Save credentials
Write-Host "Saving SMB credentials securely to Windows Credential Manager..." -ForegroundColor Cyan
cmdkey /add:$SMBServer /user:$SMBUsername /pass:$SMBPassword | Out-Null

# Create folders
Write-Host "Creating mount script directory..." -ForegroundColor Cyan
New-Item -ItemType Directory -Force -Path $TaskScriptDir | Out-Null

# Create the background mounting script
$ScriptContent = @"
# Clear existing or broken mappings on this drive letter
If (Get-PSDrive -Name ("$DriveLetter" -replace ':','') -ErrorAction SilentlyContinue) {
    Remove-PSDrive -Name ("$DriveLetter" -replace ':','') -Force | Out-Null
}

# Use net use to permanently bind the drive letter globally across user sessions
net use "$DriveLetter" "$SMBSource" /persistent:yes
"@

Write-Host "Generating background mount script..." -ForegroundColor Cyan
Set-Content -Path $TaskScriptPath -Value $ScriptContent -Force

Write-Host "Making mount script executable..." -ForegroundColor Cyan
Unblock-File -Path $TaskScriptPath

# Define the action: Run the PowerShell script silently in the background
$Action = New-ScheduledTaskAction -Execute "PowerShell.exe" -Argument "-NoProfile -WindowStyle Hidden -File `"$TaskScriptPath`""
$Trigger = New-ScheduledTaskTrigger -AtLogOn
$Settings = New-ScheduledTaskSettingsSet -AllowStartIfOnBatteries -DontStopIfGoingOnBatteries -RunOnlyIfNetworkAvailable

Write-Host "Configuring Task Scheduler..." -ForegroundColor Cyan
Register-ScheduledTask -TaskName $TaskName -Action $Action -Trigger $Trigger -Settings $Settings -Force | Out-Null

# Start it immediately
Write-Host "Starting the automount task" -ForegroundColor Green
Start-ScheduledTask -TaskName $TaskName
Write-Host "Done! Check 'This PC' in File Explorer. Your $DriveLetter drive is now persistently automated." -ForegroundColor Green
```

The folder now is the path of your network disk. To finish, create aliases to the target folder in your system account.
Please note that we will use the links with the command below to ensure a valid folder on Windows format:

```powershell
# Ensure the parent directory exists
New-Item -ItemType Directory -Force -Path "$env:USERPROFILE\Games" | Out-Null

# Create the symlinks cleanly
cmd /c mklink /d "$env:USERPROFILE\Games\ROMs" "Z:\Games\ROMs"
cmd /c mklink /d "$env:USERPROFILE\Games\BIOS" "Z:\Games\BIOS"
cmd /c mklink /d "$env:USERPROFILE\Games\State" "Z:\Games\State"
```

That is it! NiceDeck will be able to detect and parse your games from the network-mounted folders.

## Setup on MacOS

MacOS doesn't ship with an automounter feature, but `launchd` (Apple's service manager) combined with `mount_smbfs` achieves the same result: the share mounts itself in the background after login, once the NAS is actually reachable on the network.

Open **Terminal** and run the following script to set up the network folder with SMB:

```bash
# SMB remote details
SMB_SERVER="192.168.0.100"
SMB_SHARE="Shared"
SMB_USERNAME="username"
SMB_PASSWORD="password"

# Host details
SMB_DESTINATION="/Volumes/Shared"
SCRIPT_DIR="$HOME/Scripts"
SCRIPT_PATH="$SCRIPT_DIR/mount-smb-shared.sh"
PLIST_NAME="com.samba.mount-smb-shared"
PLIST_PATH="$HOME/Library/LaunchAgents/$PLIST_NAME.plist"

# Make sure folders exist
mkdir -p "$SCRIPT_DIR"
mkdir -p "$SMB_DESTINATION"

# Create the background mounting script
cat << EOF > "$SCRIPT_PATH"
#!/bin/bash
# Wait until the NAS responds before attempting to mount
until ping -c1 -t1 "$SMB_SERVER" &>/dev/null; do
  sleep 2
done

# Only mount if not already mounted
if ! mount | grep -q "$SMB_DESTINATION"; then
  mount_smbfs "//$SMB_USERNAME:$SMB_PASSWORD@$SMB_SERVER/$SMB_SHARE" "$SMB_DESTINATION"
fi
EOF

chmod +x "$SCRIPT_PATH"

# Create the LaunchAgent definition
cat << EOF > "$PLIST_PATH"
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>$PLIST_NAME</string>
    <key>ProgramArguments</key>
    <array>
        <string>/bin/bash</string>
        <string>$SCRIPT_PATH</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/tmp/mount-smb-shared.log</string>
    <key>StandardErrorPath</key>
    <string>/tmp/mount-smb-shared.err</string>
</dict>
</plist>
EOF

# Load the agent now (it will also run automatically on every future login)
launchctl unload "$PLIST_PATH" 2>/dev/null
launchctl load "$PLIST_PATH"
```

You can confirm the mount succeeded by running `mount | grep Shared`, or by checking `/Volumes/Shared` in Finder.

The folder now is the path of your network disk. To finish, create aliases to the target folder in your system account:

```bash
# Ensure the parent directory exists
mkdir -p $HOME/Games

# Create the symlinks cleanly
ln -s /Volumes/Shared/Games/ROMs $HOME/Games/ROMs
ln -s /Volumes/Shared/Games/BIOS $HOME/Games/BIOS
ln -s /Volumes/Shared/Games/State $HOME/Games/State
```

That is it! NiceDeck will be able to detect and parse your games from the network-mounted folders.