# NiceDeck

Steam Deck companion app for console emulation, better gaming experience and additional software support.

![Nice!](./nice.jpg?raw=true)

NiceDeck is an alternative for automated installations softwares like [EmuDeck](https://github.com/dragoonDorise/EmuDeck) and [RetroDeck](https://github.com/XargonWan/RetroDECK). It tries to keep things as independent as possible by just focusing on installing the programs, while also adding it to the ``Steam Library`` with nice cover images. 

Some features of NiceDeck:

- Automatic installation of recommended softwares for general usage, gaming and console emulation (see list below).
- Ability to choose the softwares and emulators to install.
- Official installation via flatpak on available applications (because is easier to use and updates automatically).
- Each software is independent and is maintained / updated directly by their developers.
- Simplified structure for emulators, where you should see only the ``ROMs`` and ``BIOS`` folders for the emulators that you installed.
- Installed programs will be available on the ``Steam Library``, allowing usage on ``Gaming Mode`` at Steam Deck. Nice covers are also expected.
- Built-in parser to grab information and add ROMs to the ``Steam Library`` automatically.
- Built-in tool to backup and restore saved games progress and states on each emulator.
- A correct and workable ``EmulationStation DE`` settings, with systems and finder rules to run games with flatpak in all emulators.
- Originally created for Steam Deck, but works on almost any Linux distro.

## Installation and Usage

Open Steam Deck in ``Desktop Mode``, launch ``Konsole`` and type the command below to install ``NiceDeck``:

```bash
# Install the program and run it.
# You can customized GUI and version with the parameters as:
# curl ... | bash -s [gui] [version] 
curl https://mateussouzaweb.github.io/nicedeck/install | bash -s
```

The program will open and you can finally run the initial setup process. After the setup process, install the desired programs, parse your ROMs or manage Steam shortcuts from GUI.

*Note:* Restart Steam or the device to changes take effect into ``Steam Library``.

## Important Notes

Folders and Structure:

- NiceDeck will create the ``$HOME/Games`` folder with basic structure for emulation.
- You can optionally map the MicroSD card path in the install process with symlinks on home to keep data separated from main drive.
- Once you have installed the desired emulators, place the ``BIOS`` and ``ROMs`` for each emulator in their respective folders.
- Make sure to read the [ROMs documentation](docs/ROMs.md) to learn how to organize and parser your ROMs.

Programs:

- With the exception of ``EmulationStation DE``, NiceDeck **will not pre-configure additional softwares and emulators**.
- This means that you should run configuration process of the emulation, including placing BIOS files and tweak settings before using it. 
- Consult the official guide of each program if you need assistance to correctly configure it.
- Some programs will require a secondary switch to ``Desktop Mode`` in order to configure advanced settings given the limitations of ``Gaming Mode``. 

Controller Layout:

- NiceDeck includes a custom ``Controller Template`` in Steam for general usage, but mainly target for emulators. The template is called ``[NICEDECK] - Gamepad``.
- You should set the best controller layout for each application before running it (for browser like softwares like ``Google Chrome`` use the ``Web Browser`` template for example).
- Open the ``Steam Library``, select the program that you desire and click on ``Controller Icon`` to reveal the customization menu. From the menu, select the template and save changes.

Application Launcher:

- You can also tweak the system launcher in ``Desktop Mode`` for more customizations.
- Right-click on the ``Application Launcher`` icon and select the option ``Edit Applications``.
- The opened software allows deep customization of the available applications on the menu, including the support for adding new categories (like **Emulators**), renaming apps, removing the not desired ones, sorting and many more. Make the customizations based in your needs.
- If you want to add a program to the favorites section, open the system launcher, right-click in the desired application and select the ``Add to Favorites`` option.

Once you have configured the controller layout on each program and run through the setup process, it's time to have a nice experience!

Enjoy!

## Available Softwares

Browsers:

- [Brave Browser](https://brave.com)
- [Firefox](https://www.mozilla.org/en-US/firefox)
- [Google Chrome](https://www.google.com/intl/en_us/chrome)
- [Microsoft Edge](https://www.microsoft.com/en-us/edge)

Streaming:

- [Chiaki](https://chiaki.re)
- [GeForce NOW](https://www.nvidia.com/geforce-now)
- [Jellyfin Media Player](https://jellyfin.org)
- [Moonlight Game Streaming](https://moonlight-stream.org)
- [Xbox Cloud Gaming](https://www.xbox.com/cloud-gaming)

Game Launchers:

- [Bottles](https://usebottles.com)
- [EmulationStation DE](https://es-de.org)
- [Heroic Games Launcher](https://heroicgameslauncher.com)
- [Lutris](https://lutris.net)

Utilities:

- [ProtonPlus](https://github.com/Vysp3r/ProtonPlus)

Console Emulators:

- Microsoft Xbox - [Xemu](https://xemu.app)
- Nintendo 3DS - [Lime3DS](https://lime3ds.github.io)
- Nintendo 3DS - [Citra](https://citra-emu.org)
- Nintendo 64 - [Simple64](https://simple64.github.io)
- Nintendo DS - [MelonDS](https://melonds.kuribo64.net)
- Nintendo Game Boy Advance - [mGBA](https://mgba.io)
- Nintendo GameCube - [Dolphin](https://dolphin-emu.org)
- Nintendo Switch - [Ryujinx](https://ryujinx.org)
- Nintendo Switch - [Yuzu](https://yuzu-emu.org)
- Nintendo Wii - [Dolphin](https://dolphin-emu.org)
- Nintendo Wii U - [Cemu](https://cemu.info)
- Sega Dreamcast - [Flycast](https://github.com/flyinghead/flycast)
- Sony Playstation 1 - [DuckStation](https://www.duckstation.org)
- Sony Playstation 2 - [PCSX2](https://pcsx2.net)
- Sony Playstation 3 - [RPCS3](https://rpcs3.net)
- Sony Playstation Portable - [PPSSPP](https://www.ppsspp.org)

Please not that NiceDeck will not offer support for all emulation softwares out there - we focus only on emulators for single consoles. If you want to emulate older consoles, please consider [RetroArch](https://www.retroarch.com) or something else.

## Using NiceDeck Outside Steam Deck

You can run NiceDeck in any Linux distribution that supports ``flatpak`` with [Flathub](https://flathub.org), just make sure to **install and setup Steam first** in order to have the necessary folders of Steam in your system. 

If you installed Steam via flatpak too, don't worry, NiceDeck will set the necessary settings required to bypass the sandbox limitations of flatpak making you able to launch other applications with Steam - requires ``flatpak-xdg-utils`` package installed.
