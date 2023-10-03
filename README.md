# NiceDeck

Steck Deck customizations for a nice experience targered **for advanced users only**.
This process will automate basic installation without the need of extra softwares.

All softwares and emulators are flatpak based, because it is official and easier to use / install / update.
Like softwares, each emulator is flatpak based. You can choose only the desired emulators to install.


## Installation and Usage

Open Steam Deck in ``Desktop Mode``, launch ``Konsole`` and type the command below to install ``nicedeck``.

```bash
curl https://mateussouzaweb.github.io/nicedeck/install.sh | bash -
```

After the program has been installed, you can install all nice deck experience programs on just the ones that you desire:

```bash
# Install all
nicedeck 

# Install specific programs
nicedeck --install=citra,yuzu
```

## Folders and Structure

- NiceDeck will create the ``$HOME/Games`` folder with basic structure for emulation.
- You can optionally map the MicroSD card path in the install process with symlink to keep this data separated from main drive installation (this is important because by using symbolic links, we can avoid some permissions issues with flatpak along the way).
- Once you have installed the desired emulators, place the ``BIOS`` and ``ROMs`` for each emulator in their respective folders.

## Application Launcher

Is time to organize the system launcher in Desktop Mode. Right-click on the ``Application Laucher`` icon and select the option ``Edit Applications``. The opened software will allows deep customization of the installed applications on the launcher, including the support for adding new categories (like **Emulators**), renaming apps, removing the not desired ones, sorting and many more. Make the customizations based in your needs.

You can also take the opportunity and add favorites programs to the special favorites section. Open the launcher, right-click in the application and select the ``Add to Favorites`` option.

## Add Programs to Steam

You should now add the softwares and emulators to the Steam. This will allow the use of these programs directly from the ``Gaming Mode``, with all shortcuts and special menus that is available on Steam Deck. For each downloaded software or emulator, open the launcher menu, right-click on the app and choose the ``Add to Steam`` option. Once the apps are added, we can grab their images for nice UI on Steam.

## Controller Layout

To configure controller layout for each program, make sure you are in ``Gaming Mode`` first:

The most easier layout are for browser like softwares like ``Google Chrome``:

- Open the library, select the software that you desire and open the controller menu. 
- Set the default model as ``Web Browser`` and close the menu.

For emulators, you will need to set some custom layouts to allow both gamepad and mouse usage:

- Open the library, select the emulator that you desire and open the controller menu.
- Set the default template as ``Gamepad with Mouse Trackpad``.
- Click on ``Edit Layout``.
- Set buttons as ``L4 - Enter Key``, ``L5 - Escape Key``, ``R4 - Left Click`` and ``R5 - Rigth Click``.
- Set ``Right Trackpad Behavior`` as ``Mouse`` with click as ``Left Mouse Click``.
- Set ``Left Trackpad Behavior`` as ``Scroll Whell`` with click as ``Right Mouse Click``.
- Click on the ``Gear Icon`` and use the option ``Export Layout``.
- Name it as ``Nice Deck`` with ``...`` as description and confirm save.
- Finally, close the menu to conclude changes.

Once you have it configured on each program, it time to have a nice experience :D.

## Final Tips

- Open the software and configure it on ``Gaming Mode`` first to see how it goes.
- Some programs will require a secondary switch to ``Desktop Mode`` in order to configure advanced settings given the limitations of ``Gaming Mode``.
- Make sure gamepad mapping are correctly configured inside the software. You may also create additional commands for such specific program.
- Make sure to use the directories path from home always - even in case of symlinks from MicroSD card.
- Enjoy!
