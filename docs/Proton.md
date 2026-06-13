# Running Windows Applications with Proton on Linux

NiceDeck offers a very robust Proton layer to run Windows native games and applications.

- In general terms, the NiceDeck Proton layer tries to give you a single Windows environment to run your favorite games and applications. The NiceDeck Proton layer is universal and is able to accommodate as many applications as you need.
- In this environment, you will have the NiceDeck provided applications through automated installation such as **Epic Games Launcher**, **EA App**, **Ubisoft Connect** and many more. In fact, the NiceDeck Proton layer is activated once you try to install one of these applications with NiceDeck.
- The NiceDeck Proton layer is independent and provides its own executable. You can set it up, remove NiceDeck, and would still have access to the Windows native applications inside it.
- The Proton environment content is completely visible at your ``$HOME/Games/Proton`` folder, so you don't need to search for the prefix folder inside Steam folders - it's one for all.
- The Proton prefix folder is portable, so you can use it to copy save files and important files more easily.
- The Proton implementation will still be provided to you by Steam. Keep in mind that NiceDeck will only set up the Proton environment for you using the tools provided by Valve.

*That is great! How can I use this Proton layer?*

You should install Steam and the ``Proton - Experimental`` package from Steam on your device first. Make sure to follow these requirements before trying to use the NiceDeck Proton layer.

Then, just install any Windows native only application with NiceDeck, such as **Epic Games Launcher**, **EA App**, **Ubisoft Connect**, **Amazon Games**, **Battle.Net**, **Rockstar Games Launcher** and **GOG Galaxy**. After installation you will see that the Proton executable and the ``$HOME/Games/Proton`` folder are available on your device, ready for use!

## Running Games with Proton

To run games and applications inside your Proton environment without the need to create a new NiceDeck library shortcut, you can just open the system terminal and run the command as described below:

```bash
"$HOME/Games/Proton/run.sh" "$HOME/Games/path/to/app.exe"
```

This command will start the designated program executable in the Proton environment and is very useful for installing drivers or trying new things without creating a shortcut.

If everything goes right for the application or game that you are trying to run, you can later point a new shortcut for it inside NiceDeck for easier access with nice covers, etc.

For easier and more advanced cases, you can create a context menu entry in the KDE environment to let the system automatically create and run this command in the terminal for you. The process is described in the [Additional Tips](./Additional%20Tips.md) section.

## Manually Adding Games

Here are the steps to manually add and play games that you own outside store launchers with the Proton layer:

- Although not required for Steam native applications under Linux, it's a good practice to put your Windows native games and applications inside your ``$HOME/Games`` folder (required in flatpak).
- You can create a new folder such as ``$HOME/Games/Windows`` just to put these Windows native games inside it.
- Copy the content that you have for the game and place it inside your ``$HOME/Games/Windows`` folder.
- Follow the explanation in the **"Running Games with Proton"** section above to run such games or applications directly, or create a new shortcut from NiceDeck and run it from the GUI.

## Installing Drivers

Here are the steps to follow if you need to install additional drivers for your Proton environment (such as ``.NET 8.0``, ``VcRedist 2015+``, etc.):

- Download the driver from a trusted source.
- Place the executable inside your ``$HOME/Games`` folder.
- Follow the explanation in the **"Running Games with Proton"** section above to run such driver applications without the need to create a shortcut for it.

## Tweaking the Environment with Wine

Proton includes Wine components that are very powerful, and you can use them to tweak settings and apply optimizations to the Proton environment such as:

- Setting DPI for fractional scaling on 1440p or 4K screen resolutions.
- Emulating a virtual desktop with a custom screen size.
- Switching the graphical rendering API.
- Opening the Windows Explorer, RegEdit, and Task Manager implementations.
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
"$HOME/Games/Proton/run.sh" wine reg add "HKCU\\Control Panel\\Desktop" /v LogPixels /t REG_DWORD /d "192" /f
```

For more details, please check the official [Wine documentation](https://gitlab.winehq.org/wine/wine/-/wikis/Commands).

## Tweaking the Environment with Extras

The Proton run script available at `$HOME/Games/Proton/run.sh` will be updated every time you install a program with Proton.
If you want to customize the run process before launching it, create the file `$HOME/Games/Proton/extra.sh` and put your customizations inside it.

For reference, here are a few customizations that you can do in the `$HOME/Games/Proton/extra.sh` file:

```bash
#!/bin/bash

# Activate GameMode
# Debug: gamemoded -s && gamemodelist
if [[ "$1" != "wine" ]]; then
  ARGUMENTS=("gamemoderun" "${ARGUMENTS[@]}")
fi

# Replace Proton Experimental with Proton-CachyOS Latest
PROTON_SEARCH="steamapps/common/Proton - Experimental"
PROTON_REPLACE="compatibilitytools.d/Proton-CachyOS Latest"
COMMAND=("${COMMAND[@]/$PROTON_SEARCH/$PROTON_REPLACE}")
ARGUMENTS=("${ARGUMENTS[@]/$PROTON_SEARCH/$PROTON_REPLACE}")
```

Once you set the content in the file, the Proton layer will automatically load it before running the final launch command.

## Games Compatibility with Proton

Proton is always evolving, but there are a few caveats when compared to a native Windows OS. Here is a simple example:

- Some games will crash if ``NVAPI`` is enabled, due to compatibility issues.
- Disabling ``NVAPI`` is required for the GRID racing series, for example (**Race Driver: GRID**, **GRID 2**, **GRID Autosport**, ...).
- You can tweak the launch options in the shortcut to disable ``NVAPI``. To disable NVAPI, edit the shortcut from NiceDeck and set launch options like: ``PROTON_DISABLE_NVAPI=1 %command% ...``.
- This is the same process you may already know from other guides on the internet.

Some games may need specific tweaks or may not work at all. For additional tweaks, you can see the provided documentation in the [Proton repository](https://github.com/ValveSoftware/Proton?tab=readme-ov-file#runtime-config-options) and check the community experience at [ProtonDB](https://www.protondb.com).

## IMPORTANT: Proton on Steam Flatpak

When running the Steam application with flatpak, you will have a few limitations due to the sandbox environment:

1 - Steam / Proton will see only the content that is inside your ``$HOME/Games`` folder. **This is very important to understand**, otherwise you will be lost - always put your Windows games and applications inside that folder.

2 - Symbolic links will not work for content that is outside of the ``$HOME/Games`` folder. If you need to let Steam and Proton access content outside of this folder, you need to run the following command to add support for additional locations:

```bash
# Specific folder
flatpak override --user --filesystem=/path/to/folder com.valvesoftware.Steam

# External drive
flatpak override --user --filesystem=/path/to/mount/point com.valvesoftware.Steam
```

3 - If you are not familiar with terminal commands, you can use [Flatseal](https://flathub.org/en/apps/com.github.tchx84.Flatseal) or the system integrated flatpak permissions management (when available) to add additional locations to the Steam flatpak application.

## IMPORTANT: Layer Implementation Details

Here are a few important details that are good to know about the Proton layer provided by NiceDeck:

- Windows native programs provided with the Proton layer implementation from NiceDeck always use the ``Proton - Experimental`` version provided to you by Steam. You are encouraged not to change the Proton version, even if you know how to do it.
- Once you create a library shortcut to directly launch games, you don't need to (and should not) configure the Proton version inside Steam shortcuts, as this is automatically handled by the NiceDeck Proton layer internally *(doing so will result in an error)*.
- If the game doesn't work with the NiceDeck Proton layer implementation due to compatibility issues with the Proton version used, it's recommended to remove the NiceDeck Proton shortcut for the game and try using it natively with Steam by adding it to the Steam library as a *non-Steam shortcut* and manually selecting another Proton version inside Steam.