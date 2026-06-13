# NiceDeck

Steam companion app for console emulation, better gaming experience and additional software support.

![Nice!](./nice.jpg?raw=true)

[PROGRAM SCREENSHOTS](./screenshots/)

NiceDeck is a solid alternative for automated installation programs like [EmuDeck](https://github.com/dragoonDorise/EmuDeck) and [RetroDeck](https://github.com/XargonWan/RetroDECK). It tries to keep things as simple as possible by focusing on installing the recommended programs and emulators, managing your ROMs library by providing shortcuts, automatically adding them to the ``Steam Library`` and handling your gaming state with backups.

NiceDeck was originally created for Steam Deck, but works on Linux, Windows and MacOS.

Some features of NiceDeck:

- Automatic installation of recommended software for general usage, gaming and console emulation (see list below).
- Ability to choose the software and emulators to install.
- Installation for applications and software using the best official packaging source.
- Each piece of software is independent and is maintained / updated directly by its developers.
- Simplified structure for emulators, where you should see only the ``ROMs`` and ``BIOS`` folders for the emulators that you installed.
- Installed programs available on the ``Steam Library`` (allowing usage on ``Gaming Mode`` at Steam Deck and ``Big Picture`` mode on Desktop).
- Built-in parser to grab information and add ROMs to the ``Steam Library`` automatically.
- Beautiful and automated cover images for shortcuts in the ``Steam Library``.
- Built-in tool to backup and restore saved game progress and states on each emulator.
- Correct and workable ``ES-DE`` settings, with systems and rules to run games using the installed emulators.
- Linux only: support for additional stores and Windows native games or applications through a custom Proton layer.

## System Requirements

Here are the requirement details based on your device or operating system:

Steam OS:
- No additional requirements, everything is included on Steam OS devices like the Steam Deck.

Linux:
- You can also run NiceDeck in any Linux distribution that supports [Flatpak](https://flatpak.org/) with the [Flathub](https://flathub.org) repository enabled, but make sure to install the ``flatpak-xdg-utils`` package too.
- 7zip is also required if your distribution does not include it: ``sudo apt install p7zip-full``.

MacOS:
- For MacOS systems, you must have the [Homebrew](https://brew.sh/) package manager to be able to manage programs.
- You also need the 7zip program to extract archive files: ``brew install p7zip``.

Windows:
- On Windows, you must have the new [WinGet](https://github.com/microsoft/winget-cli) package manager that is already included in Windows 11.
- You also need the 7zip program to extract archive files: ``winget install -e --id 7zip.7zip``.

Once you have solved the system dependencies, just download and run NiceDeck!

## Installation and Usage

NOTE: You need to go into ``Desktop Mode`` in Steam Deck to follow these instructions.

In general, we recommend that you **install and set up Steam first** in order to have the necessary Steam folders in your system to sync the library before trying to install any other additional software within NiceDeck.

Go to the project [RELEASES](https://github.com/mateussouzaweb/nicedeck/releases) page and download the latest version of NiceDeck for your operating system:

- Steam Deck: ``nicedeck-linux-amd64``.
- Linux x86: ``nicedeck-linux-amd64``.
- Linux ARM: ``nicedeck-linux-arm64``.
- MacOS Apple Silicon: ``nicedeck-macos-universal.zip``.
- MacOS Intel: ``nicedeck-macos-universal.zip``.
- Windows x86: ``nicedeck-windows-amd64.exe``.
- Windows ARM: ``nicedeck-windows-arm64.exe``.

On Linux, make sure that the file is executable:

- From the file navigator, open the file properties and check the *executable* field.
- From the terminal, run a command like ``chmod +x $FILE``.

On MacOS, extract the downloaded zip file to obtain the application: ``NiceDeck.app``.

Once you follow these instructions, double click on the program to start it. With the program running you can finally install the desired programs, parse your ROMs or manage Steam shortcuts from the GUI.

## Important Notes

Folders and Structure:

- NiceDeck will create the ``$HOME/Games`` folder with a basic structure for emulation and general gaming.
- You can optionally map external disks or MicroSD cards with symbolic links in the games folder to keep data separated from the main drive.
- Once you have installed the desired emulators, place the ``BIOS`` and ``ROMs`` for each emulator in their respective folders.
- Make sure to read the [ROMs documentation](docs/ROMs.md) to learn how to organize and parse your ROMs.
- For additional gaming on Linux, make sure to read the [Proton documentation](docs/Proton.md) to learn how to use the Proton layer.

Programs and Emulators:

- With the exception of ``ES-DE``, NiceDeck **will not pre-configure additional software and emulators**.
- This means that you should run the configuration process for each emulator, including placing BIOS files and tweaking settings before using it.
- Consult the official guide of each program if you need assistance to correctly configure it.
- On Steam Deck, some programs will require a secondary switch to ``Desktop Mode`` in order to tweak advanced settings given the limitations of ``Gaming Mode``.

Steam Library:

- You need to restart Steam or the Steam Deck device for changes to take effect in your ``Steam Library``.
- Once programs and ROMs are available in the ``Steam Library``, you can use the collections feature to better filter and manage your games.

Controller Layout:

- NiceDeck includes a custom ``Controller Template`` in Steam for general usage, but mainly targeted for emulators. The template is called ``[NICEDECK] - Gamepad``.
- You should set the best controller layout for each application before running it (for browser-based software like ``Google Chrome``, use the ``Web Browser`` template, for example).
- Open the ``Steam Library``, select the program that you desire and click on ``Controller Icon`` to reveal the customization menu. From the menu, select the template and save changes.
- Please note that this controller layout is available only on Steam Deck devices.

Enjoy!

## Available Software

The availability of software depends on the operating system that you are using.

Console Emulators:

- Microsoft Xbox - [Xemu](https://xemu.app)
- Microsoft Xbox 360 - [Xenia](https://xenia.jp)
- Nintendo 3DS - [Azahar](https://azahar-emu.org)
- Nintendo 64 - [Simple64](https://github.com/simple64/simple64)
- Nintendo DS - [MelonDS](https://melonds.kuribo64.net)
- Nintendo Game Boy Advance - [mGBA](https://mgba.io)
- Nintendo GameCube - [Dolphin](https://dolphin-emu.org)
- Nintendo Switch - [Eden](https://eden-emu.dev)
- Nintendo Switch - [Ryujinx](https://ryujinx.app)
- Nintendo Wii - [Dolphin](https://dolphin-emu.org)
- Nintendo Wii U - [Cemu](https://cemu.info)
- Sega Dreamcast - [Flycast](https://github.com/flyinghead/flycast)
- Sega Dreamcast - [Redream](https://redream.io)
- Sony Playstation 1 - [DuckStation](https://www.duckstation.org)
- Sony Playstation 2 - [PCSX2](https://pcsx2.net)
- Sony Playstation 3 - [RPCS3](https://rpcs3.net)
- Sony Playstation 4 - [ShadPS4](https://shadps4.net)
- Sony Playstation Portable - [PPSSPP](https://www.ppsspp.org)
- Sony Playstation Vita - [Vita3K](https://vita3k.org)

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
- [Rockstar Games Launcher](https://socialclub.rockstargames.com/rockstar-games-launcher)

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

Please note that NiceDeck will not offer support for all emulation software out there - we focus only on emulators for single consoles. If you want to emulate older consoles, please consider [RetroArch](https://www.retroarch.com), [Ares](https://ares-emu.net), [OpenEmu](https://openemu.org) or something else.