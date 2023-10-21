# Managing ROMs with NiceDeck

NiceDeck has the ability to automatically find and add your ROMs to the Steam Library making every game available on the Steam UI. Please read this guide to learn how to do it.

## ROMs Path and Organization

To start, you must copy your ROMs to the Steam Deck. Use the table below to guide you where you should put your ROMs based on their consoles:

| Console                   | Emulator | ROMs Folder                 |
|---------------------------|----------|-----------------------------|
| Microsoft Xbox            | Xemu     | ``$HOME/Games/ROMs/XBOX``   |
| Nintendo 3DS              | Citra    | ``$HOME/Games/ROMs/3DS``    |
| Nintendo DS               | MelonDS  | ``$HOME/Games/ROMs/DS``     |
| Nintendo Game Boy Advance | mGBA     | ``$HOME/Games/ROMs/GBA``    |
| Nintendo GameCube         | Dolphin  | ``$HOME/Games/ROMs/GC``     |
| Nintendo Switch           | Yuzu     | ``$HOME/Games/ROMs/SWITCH`` |
| Nintendo Wii              | Dolphin  | ``$HOME/Games/ROMs/WII``    |
| Nintendo Wii U            | Cemu     | ``$HOME/Games/ROMs/WIIU``   |
| Sega Dreamcast            | Flycast  | ``$HOME/Games/ROMs/DC``     |
| Sony Playstation 2        | PCSX2    | ``$HOME/Games/ROMs/PS2``    |
| Sony Playstation 3        | RPCS3    | ``$HOME/Games/ROMs/PS3``    |
| Sony Playstation Portable | PPSSPP   | ``$HOME/Games/ROMs/PSP``    |

Please note that it's very important to have the ROMs in the correct location. Any ROM outside of these directories will not be parsed by NiceDeck and consequently will not be available on the Steam Library.

Another important aspect for the ROMs organization, are the exclude patterns inside the ROMs folder. Please keep in mind that the parser will ignore any content where the path follows the following patterns:

- ``$HOME/Games/ROMS/$PLATFORM/Updates`` - Updates folder
- ``$HOME/Games/ROMS/$PLATFORM/Mods`` - Mods folder
- ``$HOME/Games/ROMS/$PLATFORM/Ignore`` - Literaly a folder to ignore
- ``$HOME/Games/ROMS/$PLATFORM/Others`` - Another special folder to ignore

You also must know that **every available ROM** inside will be added to the Steam Library. If you want to put only a few games in the Steam Library, you MUST organize your ROMs. Take for example the following organization using the ``GBA`` platform to parse only the favorite games:

- ``$HOME/Games/ROMS/GBA/Favorites`` - Games that will be included on parser
- ``$HOME/Games/ROMS/GBA/Others`` - Others non-favorite games that will be ignored by the parser

Once you decided the best ROM organization for you and copied your ROMs to the Steam Deck, it's time to run the parser.

## Using the Parser

Simply run one of the following commands to parse your ROMs, scrape data and create the ROM shortcut inside the Steam Library. You also should run this commands again when you update your ROMs content doing actions like adding new games, renaming or removing one of your ROMs:

```bash
# Parse and scrape data for all platforms / folders
nicedeck roms

# Parse and scrape data for specific platform(s) only
nicedeck roms GBA,3DS
nicedeck roms --platforms=GBA,3DS

# Rebuild everything by parsing again 
nicedeck roms --rebuild

# Run the parse, but uses Ryujinx for Nintendo Switch emulation
# Note: Ryujinx is not enabled by the default, you must specify the preference if want to run Nintendo Switch games with Ryujinx.
nicedeck roms --platforms=SWITCH --preferences=use-ryujinx
```

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
- When saved, all ROMS of the platform will be added to the created colection.
- Repeat the same process on each platform that you are emulating - ``[3DS]``, ``[PS2]``, ``[XBOX]``...

I do recommend using an Mouse and Keyboard connected to the Steam Deck to run it faster. If can also use the same process to add some games to the favorites.

## Troubeshotting

**When running the parser, I see the message: "Could not detect ROM information. Skipping..."**\
This means that the system could not detect which is the game of the ROM based on their file name. The easiest fix is to rename the file with the correct game name. For example:

Before: ``MSR - Metropolis Street Racer (USA) (Rev A).cue``\
After: ``Metropolis Street Racer (USA).cue``

**I got some error while running the parser...**\
This can happen because the Steam Grid DB service cannot handle all of your request with a valid response time. Simply run the command again and see if the error persist - when you run the command again, the parse will continue from where it has been stopped.

**The parser runs fine, but some of my games does not have cover images...**\
This can happen because Steam Grid DB does not have images for all games in the world and someone will need to push artworks for that game in the service. If you can contribute to the project, please consider submit the game cover images to Steam Grid DB via the following link: <https://www.steamgriddb.com/upload>

**I don't want all of my games in Steam Library, how to manage which game will go to Steam Library?**\
As you know, this process is automated and does not have an interface to manage games individually. If you don't want to a game appears in the Library, I should start saying that then you should not had put the games that you do not want to play in the ROMs folder. If by somehow you still want to play a game, but do not want them to appear in the Steam Library, read the organization section again to learn how to do it.

**The controller does not work as intended or do not match the one configured on the emulator...**\
You have to manually configure the controller layout in the ROM shortcut like you did on the main emulator. The fastest way to do it, it by launching the ROM. Once the ROM is running, press the Steam Menu button and select the menu "Controller Settings" on the running program; select the correct controller layout and apply it - this change will  on the fly.