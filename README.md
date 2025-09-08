# NiceDeck

Steam companion app for console emulation, better gaming experience and additional software support.

![Nice!](./nice.jpg?raw=true)

[PROGRAM SCREENSHOTS](./screenshots/)

NiceDeck is a solid alternative for automated installations programs like [EmuDeck](https://github.com/dragoonDorise/EmuDeck) and [RetroDeck](https://github.com/XargonWan/RetroDECK). It tries to keep things as simple as possible by focusing on installing the recommended programs and emulators, managing your ROMs library by providing shortcuts, automatically adding it to the ``Steam Library`` and handling your gaming state with backups. 

NiceDeck was originally created for Steam Deck, but works on Linux, Windows and MacOS.

Some features of NiceDeck:

- Automatic installation of recommended softwares for general usage, gaming and console emulation (see list below).
- Ability to choose the softwares and emulators to install.
- Installation for applications and softwares using the best official packaging source.
- Each software is independent and is maintained / updated directly by their developers.
- Simplified structure for emulators, where you should see only the ``ROMs`` and ``BIOS`` folders for the emulators that you installed.
- Installed programs available on the ``Steam Library`` (allowing usage on ``Gaming Mode`` at Steam Deck and ``Big Picture`` mode on Desktop).
- Built-in parser to grab information and add ROMs to the ``Steam Library`` automatically.
- Beautiful and automated covers images for shortcuts in the ``Steam Library``.
- Built-in tool to backup and restore saved games progress and states on each emulator.
- A correct and workable ``ES-DE`` settings, with systems and rules to run games using the installed emulators.

## System Requirements

In general, the NiceDeck program needs that you **install and setup Steam first** in order to have the necessary folders of Steam in your system. Nothing else is required on Steam Deck devices.

- You can also run NiceDeck in any Linux distribution that supports [Flatpak](https://flatpak.org/) with [Flathub](https://flathub.org) repository enable, but make sure to install the ``flatpak-xdg-utils`` package too.
- For MacOS systems, you must have the [Homebrew](https://brew.sh/) package manager to be able to manage programs. 
- On Windows, you must have the new [WinGet](https://github.com/microsoft/winget-cli) package manager, which is automatically included in Windows 11.

Once you have solved the system dependencies, just download and run NiceDeck!

## Installation and Usage

NOTE: You need go into ``Desktop Mode`` in Steam Deck to follow these instructions.

Go to the project [RELEASES](https://github.com/mateussouzaweb/nicedeck/releases) page and download the latest version of NiceDeck for your operating system:

- Steam Deck: ``nicedeck-linux-amd64``.
- Linux x86: ``nicedeck-linux-amd64``.
- Linux ARM: ``nicedeck-linux-arm64``.
- MacOS Apple Silicon: ``nicedeck-macos-arm64``.
- MacOS Intel: ``nicedeck-macos-amd64``.
- Windows x86: ``nicedeck-windows-amd64.exe``.
- Windows ARM: ``nicedeck-windows-arm64.exe``.

On Linux and MacOS, make sure that the file executable: 

- From file navigator, open the file properties and check the *executable* field.
- From terminal, run the command  like ``chmod +x $FILE``.

Once you follow these instructions, double click on the program to start it. With the program running you can finally install the desired programs, parse your ROMs or manage Steam shortcuts from the GUI.

## Important Notes

Folders and Structure:

- NiceDeck will create the ``$HOME/Games`` folder with basic structure for emulation.
- You can optionally map external disks or MicroSD cards with symbolic links on the games folder to keep data separated from main drive.
- Once you have installed the desired emulators, place the ``BIOS`` and ``ROMs`` for each emulator in their respective folders.
- Make sure to read the [ROMs documentation](docs/ROMs.md) to learn how to organize and parser your ROMs.

Programs and Emulators:

- With the exception of ``ES-DE``, NiceDeck **will not pre-configure additional softwares and emulators**.
- This means that you should run configuration process of the emulation, including placing BIOS files and tweak settings before using it. 
- Consult the official guide of each program if you need assistance to correctly configure it.
- On Steam Deck, some programs will require a secondary switch to ``Desktop Mode`` in order to tweak advanced settings given the limitations of ``Gaming Mode``. 

Steam Library:

- You need to restart Steam or the Steam Deck device to changes take effect into your ``Steam Library``.
- After programs and ROMs were available in the ``Steam Library``, you can use the collections feature to better filter and manage your games.

Controller Layout:

- NiceDeck includes a custom ``Controller Template`` in Steam for general usage, but mainly target for emulators. The template is called ``[NICEDECK] - Gamepad``.
- You should set the best controller layout for each application before running it (for browser like softwares like ``Google Chrome`` use the ``Web Browser`` template for example).
- Open the ``Steam Library``, select the program that you desire and click on ``Controller Icon`` to reveal the customization menu. From the menu, select the template and save changes.
- Please note that this controller layout is available only on Steam Deck devices.

Enjoy!

## Available Softwares

The availability of software depends on the operational system that you are using.

Console Emulators:

- Microsoft Xbox - [Xemu](https://xemu.app)
- Microsoft Xbox 360 - [Xenia](https://xenia.jp)
- Nintendo 3DS - [Azahar](https://azahar-emu.org)
- Nintendo 64 - [Simple64](https://simple64.github.io)
- Nintendo DS - [MelonDS](https://melonds.kuribo64.net)
- Nintendo Game Boy Advance - [mGBA](https://mgba.io)
- Nintendo GameCube - [Dolphin](https://dolphin-emu.org)
- Nintendo Switch - [Ryujinx](https://ryujinx.app)
- Nintendo Switch - [Citron](https://citron-emu.org)
- Nintendo Switch - [Eden](https://eden-emu.dev)
- Nintendo Wii - [Dolphin](https://dolphin-emu.org)
- Nintendo Wii U - [Cemu](https://cemu.info)
- Sega Dreamcast - [Flycast](https://github.com/flyinghead/flycast)
- Sega Dreamcast - [Redream](https://redream.io)
- Sony Playstation 1 - [DuckStation](https://www.duckstation.org)
- Sony Playstation 2 - [PCSX2](https://pcsx2.net)
- Sony Playstation 3 - [RPCS3](https://rpcs3.net)
- Sony Playstation 4 - [ShadPS4](https://shadps4.net)
- Sony Playstation Portable - [PPSSPP](https://www.ppsspp.org)
- Sony Playstation Vita - [Vita3k](https://vita3k.org)

Game Launchers and Stores:

- [Amazon Games](https://gaming.amazon.com)
- [Battle.net](https://us.shop.battle.net)
- [Bottles](https://usebottles.com)
- [EA App](https://www.ea.com/ea-app)
- [Epic Games](https://store.epicgames.com)
- [ES-DE](https://es-de.org)
- [GOG Galaxy](https://www.gog.com/galaxy)
- [Heroic Games Launcher](https://heroicgameslauncher.com)
- [Lutris](https://lutris.net)
- [Steam](https://store.steampowered.com)
- [Ubisoft Connect](https://www.ubisoft.com/ubisoft-connect)

Streaming:

- [Chiaki NG](https://streetpea.github.io/chiaki-ng)
- [GeForce NOW](https://www.nvidia.com/geforce-now)
- [Moonlight Game Streaming](https://moonlight-stream.org)
- [Xbox Cloud Gaming](https://www.xbox.com/cloud-gaming)

Utilities:

- [Discord](https://discord.com)
- [ProtonPlus](https://github.com/Vysp3r/ProtonPlus)

Browsers:

- [Brave Browser](https://brave.com)
- [Firefox](https://www.mozilla.org/en-US/firefox)
- [Google Chrome](https://www.google.com/intl/en_us/chrome)
- [Microsoft Edge](https://www.microsoft.com/en-us/edge)

Please not that NiceDeck will not offer support for all emulation softwares out there - we focus only on emulators for single consoles. If you want to emulate older consoles, please consider [RetroArch](https://www.retroarch.com), [OpenEmu](https://openemu.org/) or something else.
