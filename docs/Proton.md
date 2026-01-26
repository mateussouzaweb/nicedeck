# Running Windows Applications with Proton on Linux

NiceDeck offers a very robust Proton layer to run Windows native games and applications, but why?

- In general terms, the NiceDeck proton layer tries to give you a single Windows environment to run your favorite games and applications. NiceDeck Proton layer is universal and is able to accommodate as many applications as you need.
- In this environment, you will have the NiceDeck provided applications through automated installation such as **Epic Games Launcher**, **EA App**, **Ubisoft Connect** and many more. In fact, NiceDeck proton layer is activated once you try to install one of these applications with NiceDeck.
- NiceDeck Proton layer is independent and provide its own executable. You can setup it, remove NiceDeck and would still have access to the Windows native applications inside it.
- The Proton environment content is completely visible at your ``$HOME/Games/Proton`` folder, so you don't need to search for the prefix folder inside Steam folders - it's one for all. 
- The Proton prefix folder is portable, or, you can use it to copy save files and important files more easily.
- Proton implementation will still be provided to you by Steam. Keep in mind that NiceDeck will only setup the Proton environment for you using the tools provided by Valve.

*That is great! How can I use this Proton layer?*

You should install Steam and the ``Proton - Experimental`` package from Steam in your device first. Make sure to follow these requirements before trying to use the NiceDeck Proton layer.

Then, Just install any Windows native only application with NiceDeck, such as **Epic Games Launcher**, **EA App**, **Ubisoft Connect**, **Amazon Games**, **Battle.Net** and **GOG Galaxy**. After installation you will see that the Proton executable and the ``$HOME/Games/Proton`` folder is available at your device, ready for use! 

## Manually Adding Games

Here are the steps to manually add and play games that you own outside store launchers with the Proton layer:

- Although not required for Steam native application under Linux, it's a good practice to put your Windows native games and applications inside your ``$HOME/Games`` folder (required in flatpak).
- You can create a new folder such as ``$HOME/Games/Windows`` just to put these Windows native games inside it.
- Copy the content that you have for the game and place it inside your ``$HOME/Games`` folder.
- Follow the explanation in the **"Running Games with Proton"** section below to run such games or applications directly, or create a new shortcut from NiceDeck and run it from the GUI.

## Installing Drivers

Here is the steps to follow if you need to install additional drivers for your Proton environment (such as ``.NET 8.0``, ``VcRedist 2015+``, ``...``):

- Download the driver from the trusted source.
- Place the executable inside your ``$HOME/Games`` folder.
- Follow the explanation in the **"Running Games with Proton"** section below to run such driver applications without the need to create a shortcut for it.

## Tweaking Environment with Wine

Proton includes Wine components that are very powerful and you can use it to tweak settings and apply optimizations to the Proton environment such as:

- Set DPI for fractional scaling on 1440p or 4k screen resolutions.
- Emulate a virtual desktop with custom screen size.
- Switch graphical rendering API.
- Open Windows Explorer, RegEdit, Task Manager implementations.
- Many more.

To execute Wine programs, open the terminal and run commands like:

```bash
# Request to open Wine programs
"$HOME/Games/Proton/run.sh" wine winecfg # Open WineConfig
"$HOME/Games/Proton/run.sh" wine regedit # Open RegEdit
"$HOME/Games/Proton/run.sh" wine taskmgr # Open Task Manager
"$HOME/Games/Proton/run.sh" wine explorer # Open Explorer

# Manipulate DPI size directly from CLI
# Valid values for DPI are: 96, 120, 144, 192
"$HOME/Games/Proton/run.sh" wine reg add "HCU\\Control Panel\\Desktop" /v LogPixels /t REG_DWORD /d "192" /f
```

For more details, please check the official [Wine documentation](https://gitlab.winehq.org/wine/wine/-/wikis/Commands).

## Running Games with Proton

To run games and applications inside your Proton environment without the need to create a new NiceDeck library shortcut, you can just open the system terminal and run the command as describe below:

```bash
"$HOME/Games/Proton/run.sh" "$HOME/Games/path/to/app.exe"
```

This command will start the designed program executable in the Proton environment and is very useful to install drivers or try new things without creating a shortcut.

If everything goes right for the application or gaming that you are trying to running, you can later point a new shortcut for it inside NiceDeck for easier access with nice covers and etc.

For more easier and advanced cases, you can create a context service menu in KDE environment to let the system automatically create and run this command in the terminal for you. The process is described in the [Additional Tips](./Additional%20Tips.md) section.

## Games Compatibility with Proton

Proton is always evolving, but the are a few caveat when compared to Windows native OS. Here is a simple example:

- Some games will crash if ``NVAPI`` is enabled due to compatibility issues.
- Disabling ``NVAPI`` is required in the GRID racing series for example (**Race Drive GRID**, **GRID 2**, **GRID AutoSports**, ...)
- You can tweak the launch options in the shortcut to disable ``NVAPI``. To disable NVAPI, edit the shortcut from NiceDeck and set launch options like: ``PROTON_DISABLE_NVAPI=1 %command% ...``.
- This is the same process as you already know for others guides in the internet.

Some games may need specific tweaks or won't will work at all. To know the additional options for more tweaks, you can see the provided documentation on the [Proton repository](https://github.com/ValveSoftware/Proton?tab=readme-ov-file#runtime-config-options) and check the community experience at [ProtonDB](https://www.protondb.com).

## IMPORTANT: Proton on Steam Flatpak

When running Steam application with flatpak, you will have a few limitations due to the sandbox environment:

1 - Steam / Proton will see only the content that is inside your ``$HOME/Games`` folder. **This is very important to understand**, otherwise you will be lost in a big hole - always put the Windows games and applications inside that folder.

2 - Symbolic Links will not work for content that is outside of the ``$HOME/Games`` folder. If you need to let Steam and Proton access content outside of this folder, you need to run the following command to add support for additional locations:

```bash
# Specific folder
flatpak override --user --filesystem=/path/to/folder com.valvesoftware.Steam

# External drive
flatpak override --user --filesystem=/path/to/mount/point com.valvesoftware.Steam
```

3 - If you are not familiar com terminal commands, you can use [Flatseal](https://flathub.org/en/apps/com.github.tchx84.Flatseal) or the system integrated flatpak permissions management when available to add additional locations to the Steam flatpak application.

## IMPORTANT: Layer Implementation Details

Here are a few important details that is good to know about the Proton layer provided by NiceDeck:

- Windows native programs provided with Proton layer implementation from NiceDeck always use ``Proton - Experimental`` version provided to you by Steam. You are encouraged to not change the Proton version, even if you know how to do it.
- Once you create a library shortcut to directly launch games, you don't need or should configure the Proton version inside Steam shortcuts as this is automatically handled by NiceDeck Proton layer internally *(making such action will result in error)*.
- If the game don't work with NiceDeck Proton layer implementation due to compatibility issues with the used Proton version, it's recommended to remove the NiceDeck Proton shortcut of the game and try to use it natively with Steam, adding it to the Steam library as *non-steam-shortcut* and manually selecting another Proton version inside Steam.
