# Managing ROMs with NiceDeck

NiceDeck has the ability to automatically find and add your ROMs to the ``Steam Library``, making every game available on the Steam UI. Please read this guide to learn how to do it.

## ROMs Path and Organization

To start, you must copy your ROMs to the Steam Deck. Use the table below to guide you on where you should put your ROMs based on their consoles:

NOTE: ``$ROMS`` represents your ROMs directory, located at ``$HOME/Games/ROMs``.

| Console                   | Emulator                | ROMs Folder       |
|---------------------------|-------------------------|-------------------|
| Microsoft Xbox            | Xemu                    | ``$ROMS/XBOX``    |
| Microsoft Xbox 360        | Xenia                   | ``$ROMS/X360``    |
| Nintendo 3DS              | Azahar                  | ``$ROMS/3DS``     |
| Nintendo 64               | Simple64                | ``$ROMS/N64``     |
| Nintendo DS               | MelonDS                 | ``$ROMS/DS``      |
| Nintendo Game Boy Advance | mGBA                    | ``$ROMS/GBA``     |
| Nintendo GameCube         | Dolphin                 | ``$ROMS/GC``      |
| Nintendo Switch           | Eden / Ryujinx          | ``$ROMS/SWITCH``  |
| Nintendo Wii              | Dolphin                 | ``$ROMS/WII``     |
| Nintendo Wii U            | Cemu                    | ``$ROMS/WIIU``    |
| Sega Dreamcast            | Flycast / Redream       | ``$ROMS/DC``      |
| Sony Playstation 1        | DuckStation             | ``$ROMS/PS1``     |
| Sony Playstation 2        | PCSX2                   | ``$ROMS/PS2``     |
| Sony Playstation 3        | RPCS3                   | ``$ROMS/PS3``     |
| Sony Playstation 4        | ShadPS4                 | ``$ROMS/PS4``     |
| Sony Playstation Portable | PPSSPP                  | ``$ROMS/PSP``     |
| Sony Playstation Vita     | Vita3K                  | ``$ROMS/PSVITA``  |

Please note that it's very important to have the ROMs in the correct location. Any ROM outside of these directories will not be parsed by NiceDeck and consequently will not be available in the ``Steam Library`` as a direct shortcut to the game.

If you want to enforce a specific emulator for a subset of ROMs, you should create a subfolder with the emulator name to enforce it:

- ``$ROMS/SWITCH/Eden`` - Games that should always use the Eden emulator
- ``$ROMS/SWITCH/Ryujinx`` - Games that should always use the Ryujinx emulator
- ``$ROMS/SWITCH`` - Games that should use the default emulator for that platform

Another important aspect of ROM organization is the exclude patterns. Please keep in mind that the parser will ignore any content where the path follows these patterns:

- ``$ROMS/$PLATFORM/Updates`` - Updates folder
- ``$ROMS/$PLATFORM/Mods`` - Mods folder
- ``$ROMS/$PLATFORM/DLCs`` - DLCs folder
- ``$ROMS/$PLATFORM/Ignore`` - Literally a folder to ignore
- ``$ROMS/$PLATFORM/Others`` - Another special folder to ignore

You also must know that **every available ROM** inside the included folders will be added to the ``Steam Library``. If you want to put only a few games in the ``Steam Library``, you MUST organize your ROMs. Take for example the following organization using the ``GBA`` platform to parse only the favorite games:

- ``$ROMS/GBA/Favorites`` - Games that will be included by the parser
- ``$ROMS/GBA/Others`` - Other non-favorite games that will be ignored by the parser

Once you've decided on the best ROM organization for you and copied your ROMs to the Steam Deck, it's time to run the parser.

## Using the Parser

Simply open the program and run the process to parse your ROMs, scrape data, and create the ROM shortcuts inside the Steam Library. You should also run this same process again whenever you update your ROMs content by adding new games, renaming, or removing any of your ROMs.

The parser can take some time to finish on its first run, based on the size of your ROM library. When you need to run it again, don't worry - the parser will consider only the new ROMs in the catalog, making the process fast.

Wait for the process to complete and we are DONE! You can start gaming!

## Organizing Your Collection inside Steam Library

After running the parser and opening Steam again, you will notice that the ROMs will be available in the "Uncategorized" collection. That is fine for some people, but if you want to make it better, you'll need to do some manual work.

Open Steam in **Desktop Mode** and simply use the search bar to filter your collection by the platform identifier. By default, NiceDeck will append the platform key to every ROM to identify its console relationship (since some games were released on multiple platforms, this makes it much easier to identify which platform each ROM belongs to).

- Search for ``[GBA]``.
- Select all matched games.
- Right click on one of the games.
- Under the menu, go to ``Add To`` and select ``New Collection``.
- Write the name of the collection - ``Nintendo Game Boy Advance`` - and save.
- Once saved, all ROMs for that platform will be added to the created collection.
- Repeat the same process for each platform you are emulating - ``[3DS]``, ``[PS2]``, ``[XBOX]``...

On Steam Deck, we recommend using a mouse and keyboard connected to the device to perform this process faster.
You can also use the same process to add some games to your favorites.

## Troubleshooting

**- When running the parser, I see the message: "Could not detect ROM information. Skipping..."**\
This means that the system could not detect the game based on its file name. The easiest fix is to rename the file with the correct game name. For example:

Before: ``MSR - Metropolis Street Racer (USA) (Rev A).cue``\
After: ``Metropolis Street Racer (USA).cue``

**- The scraped details for my game are incorrect...**\
Make sure you have the correct ROM file name, as described in the previous question. You can also search for the correct game name on SteamGridDB if you have doubts.

**- I got an error while running the parser...**\
This can happen because the SteamGridDB service cannot handle all requests with a valid response time. Simply run the command again and see if the error persists - when you run the command again, the parser will continue from where it stopped.

**- The parser runs fine, but some of my games do not have cover images...**\
This is expected for less popular games. SteamGridDB does not have images for every game in the world, and someone needs to submit artwork for that game to the service. If you'd like to contribute to the project, please consider submitting cover images for that game to SteamGridDB via the following link: <https://www.steamgriddb.com/upload>

**- I don't want all of my games in the Steam Library. How do I manage which games will go to the Steam Library?**\
As mentioned, this process is automated and does not have an interface to manage games individually. If you don't want a game to appear in the Library, you simply should not have put that game in the ROMs folder in the first place. If for some reason you still want to keep a game in your ROMs folder but don't want it to appear in the Steam Library, read the organization section again to learn how to do this.

**- The controller does not work as intended, or does not match the one configured in the emulator...**\
You have to manually configure the controller layout in the ROM shortcut, just as you did for the main emulator. The fastest way to do this is by launching the ROM. Once the ROM is running, press the Steam Menu button and open the "Controller Settings" menu for the running program; select the correct controller layout and apply it - the change will take effect immediately.