# Managing ROMs with NiceDeck

NiceDeck has the ability to automatically find and add your ROMs to the Steam Library making every game available on the Steam UI. Please read this guide to learn how to do it.

## ROMs Path and Organization

To start, you must copy your ROMs to the Steam Deck. Use the table below to guide you where you should put your ROMs based on their consoles:

NOTE: ``$ROMS`` represents your ROMs directory, located at ``$HOME/Games/ROMs``.

| Console                   | Emulator        | ROMs Folder      |
|---------------------------|-----------------|------------------|
| Microsoft Xbox            | Xemu            | ``$ROMS/XBOX``   |
| Nintendo 3DS              | Citra           | ``$ROMS/3DS``    |
| Nintendo 64               | Simple64        | ``$ROMS/N64``    |
| Nintendo DS               | MelonDS         | ``$ROMS/DS``     |
| Nintendo Game Boy Advance | mGBA            | ``$ROMS/GBA``    |
| Nintendo GameCube         | Dolphin         | ``$ROMS/GC``     |
| Nintendo Switch           | Ryujinx / Yuzu  | ``$ROMS/SWITCH`` |
| Nintendo Wii              | Dolphin         | ``$ROMS/WII``    |
| Nintendo Wii U            | Cemu            | ``$ROMS/WIIU``   |
| Sega Dreamcast            | Flycast         | ``$ROMS/DC``     |
| Sony Playstation 1        | DuckStation     | ``$ROMS/PS1``    |
| Sony Playstation 2        | PCSX2           | ``$ROMS/PS2``    |
| Sony Playstation 3        | RPCS3           | ``$ROMS/PS3``    |
| Sony Playstation Portable | PPSSPP          | ``$ROMS/PSP``    |

Please note that it's very important to have the ROMs in the correct location. Any ROM outside of these directories will not be parsed by NiceDeck and consequently will not be available on the Steam Library as direct shortcut to the game.

If you want to enforce an specific emulator for a subset of ROMs, you should create a subfolder with the emulator name to enforce it:

- ``$ROMS/SWITCH/Ryujinx`` - Games that always should use the Ryujinx emulator
- ``$ROMS/SWITCH/Yuzu`` - Games that always should use the Yuzu emulator
- ``$ROMS/SWITCH`` - Games that should use the default emulator for that platform

Another important aspect for the ROMs organization are the exclude patterns. Please keep in mind that the parser will ignore any content where the path follows the following patterns:

- ``$ROMS/$PLATFORM/Updates`` - Updates folder
- ``$ROMS/$PLATFORM/Mods`` - Mods folder
- ``$ROMS/$PLATFORM/DLCs`` - DLCs folder
- ``$ROMS/$PLATFORM/Ignore`` - Literally a folder to ignore
- ``$ROMS/$PLATFORM/Others`` - Another special folder to ignore

You also must know that **every available ROM** inside the included folders will be added to the Steam Library. If you want to put only a few games in the Steam Library, you MUST organize your ROMs. Take for example the following organization using the ``GBA`` platform to parse only the favorite games:

- ``$ROMS/GBA/Favorites`` - Games that will be included on parser
- ``$ROMS/GBA/Others`` - Others non-favorite games that will be ignored by the parser

Once you decided the best ROM organization for you and copied your ROMs to the Steam Deck, it's time to run the parser.

## Using the Parser

Simply open the program and run the process to parse your ROMs, scrape data and create the ROM shortcut inside the Steam Library. You also should run this same process again when you update your ROMs content doing actions like adding new games, renaming or removing one of your ROMs.

The parser can take some time to finish based on the size of your ROMs library in the first run. When you need to run it again, don't worry, the parser will consider only the new ROMs in the catalog, making the process fast.

Wait for the conclusion of the process and we are DONE! You can start gaming!

## Organizing Your Collection inside Steam Library

After running the parser and opening Steam again, you will notice that the ROMs will be available on the "uncategorized" collection. That is ok for some people, but if you want to make it better, you need to do some manual work - an very easy work...

Open Steam in the **Desktop Mode** and simply use the search bar to filter your collection with the platform identification. NiceDeck by default will append the platform key on every ROM to identify the console relationship for that ROM (some games where released in multi-platforms, so you will find it very easier to identify from what platform each ROM belongs).

- Search for ``[GBA]``. 
- Select all matched games.
- Right click in one of the games.
- Under the menu, go to ``Add To`` and select ``New Collection``.
- Write the name of the collection - ``Nintendo Game Boy Advance`` - and save.
- When saved, all ROMS of the platform will be added to the created collection.
- Repeat the same process on each platform that you are emulating - ``[3DS]``, ``[PS2]``, ``[XBOX]``...

I do recommend using an Mouse and Keyboard connected to the Steam Deck to run it faster. If can also use the same process to add some games to the favorites.

## Troubleshooting

**- When running the parser, I see the message: "Could not detect ROM information. Skipping..."**\
It means that the system could not detect which is the game of the ROM based on their file name. The easiest fix is to rename the file with the correct game name. For example:

Before: ``MSR - Metropolis Street Racer (USA) (Rev A).cue``\
After: ``Metropolis Street Racer (USA).cue``

**- The scrapped details for my game is incorrect...**\
Make sure you have the correct ROM file name, like in the previous question. You can also search for the correct game name on Steam Grid DB if you have doubts.

**- I got some error while running the parser...**\
That can happen because the Steam Grid DB service cannot handle all of your request with a valid response time. Simply run the command again and see if the error persist - when you run the command again, the parse will continue from where it has been stopped.

**- The parser runs fine, but some of my games does not have cover images...**\
That is expected in non popular games. Steam Grid DB does not have images for all games in the world and someone will need to push artworks for that game in the service. If you can contribute to the project, please consider submit the game cover images to Steam Grid DB via the following link: <https://www.steamgriddb.com/upload>

**- I don't want all of my games in Steam Library, how to manage which game will go to Steam Library?**\
As you know, this process is automated and does not have an interface to manage games individually. If you don't want to a game appears in the Library, I should start saying that then you should not had put the games that you do not want to play in the ROMs folder. If by somehow you still want to play a game, but do not want them to appear in the Steam Library, read the organization section again to learn how to do it.

**- The controller does not work as intended or do not match the one configured on the emulator...**\
You have to manually configure the controller layout in the ROM shortcut like you did on the main emulator. The fastest way to do it, it by launching the ROM. Once the ROM is running, press the Steam Menu button and select the menu "Controller Settings" on the running program; select the correct controller layout and apply it - the change will apply on the fly.