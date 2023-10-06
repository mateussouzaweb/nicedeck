# NiceDeck

Steck Deck customizations for a **nice experience** targered **for advanced users only** (experimental software):

- Automatic installation of recomended softwares for general usage, gaming and console emulation (see below).
- You can choose the softwares and emulators to install.
- Oficial installation via flatpak on available applications (because is easier to use and automatically updates).
- Opinated and simplified structure for emulators, where you should see only the ``ROMs`` and ``BIOS`` folders for the emulators that you installed.
- Installed programs will be available on the ``Steam Library``, allowing usage on ``Gaming Mode`` at Steam Deck.

## Installation and Usage

Open Steam Deck in ``Desktop Mode``, launch ``Konsole`` and type the command below to install ``nicedeck``.

```bash
# Install the program
curl https://mateussouzaweb.github.io/nicedeck/install | bash -

# Make it available on current shell
export PATH="$PATH:$HOME/.local/bin"
```

After the program has been installed, you can run the setup process:

```bash
# Usage help
nicedeck help

# Install all programs
nicedeck setup

# Install specific programs only
nicedeck install --programs=citra,yuzu
```

*Note:* Restart Steam to changes take effect into ``Steam Library``.

## Important Notes

Folders and Structure:

- NiceDeck will create the ``$HOME/Games`` folder with basic structure for emulation.
- You can optionally map the MicroSD card path in the install process with symlinks to keep this data separated from main drive (this is important because by using symbolic links, we can avoid some permissions issues with flatpak along the way).
- Once you have installed the desired emulators, place the ``BIOS`` and ``ROMs`` for each emulator in their respective folders.

Programs:

- With the exception of ``EmulationStation DE``, NiceDeck **will not** pre-configure additional softwares and emulators.
- This means that you should run configuration process of the emulation, including placing BIOS files and tweak settings before using it.
- Some programs will require a secondary switch to ``Desktop Mode`` in order to configure advanced settings given the limitations of ``Gaming Mode``. Consult the official guide of each program if you need assistance.
- Make sure to use the directories path from home always - even in case of symlinks from MicroSD card.

Controller Layout:

- NiceDeck includes a custom ``Controller Template`` in Steam for general usage, but mainly target for emulators. The template is called ``[NICEDECK] - Gamepad``.
- You should set the best controller layout for each application before running it (for browser like softwares like ``Google Chrome`` use the ``Web Browser`` template for example).
- Open the ``Steam Library``, select the program that you desire and click on ``Controller Icon`` to reveal the customization menu. From the menu, select the template and save changes.

Application Launcher:

- You can also tweak the system launcher in ``Desktop Mode`` for more customizations.
- Right-click on the ``Application Launcher`` icon and select the option ``Edit Applications``.
- The opened software allows deep customization of the available applications on the menu, including the support for adding new categories (like **Emulators**), renaming apps, removing the not desired ones, sorting and many more. Make the customizations based in your needs.
- You can also take the opportunity and add favorites programs to the special favorites section. Open the launcher, right-click in the desired application and select the ``Add to Favorites`` option.

Once you have configured the controller layout on each program and run through the setup process, it's time to have a nice experience!

## Available Programs:

Softwares:

- Firefox
- Google Chrome
- Jellyfin Media Player
- Moonlight Game Streaming

Game Launchers:

- Bottles
- EmulationStation DE
- Heroic Games Launcher
- Lutris

Emulators:

- Microsoft Xbox - Xemu
- Nintendo 3DS - Citra
- Nintendo DS - MelonDS
- Nintendo Game Boy Advance - mGBA
- Nintendo GameCube - Dolphin
- Nintendo Switch - Ryujinx
- Nintendo Switch - Yuzu
- Nintendo Wii - Dolphin
- Nintendo Wii U - Cemu
- Sega Dreamcast - Flycast
- Sony Playstation 2 - PCSX2
- Sony Playstation 3 - RPCS3
- Sony Playstation Portable - PPSSPP