# Nice Deck

Steck Deck customizations for a nice experience.

**WARNING**: For advanced users only.

## Setup

Open Steam Deck in ``Desktop Mode`` and connect a Bluetooth Mouse and Keyboard. Now, open ``Konsole`` and type the commands from each section below.

## Folders and Structure

To start, we need some basic directory structure to put games and related data. The process is easy like just creating new folders:

```bash
# Go to home or MicroSD path
# Create basic structure
cd $HOME 
mkdir -p Games/BIOS
mkdir -p Games/ROMs
mkdir -p Games/Save
```

If you want to put your data inside the MicroSD card, repeat the steps above in the MicroSD path and create symbolic links to these folders on the home directory. This is important because by using symbolic links, we can avoid some permissions issues with flatpak along the way:

```bash
# Remove folders in home to create symlink
rm -r $HOME/Games/BIOS
rm -r $HOME/Games/ROMs
rm -r $HOME/Games/Save

# Create symlinks
ln -s /run/media/MicroSD/Games/BIOS $HOME/Games/BIOS
ln -s /run/media/MicroSD/Games/ROMs $HOME/Games/ROMs
ln -s /run/media/MicroSD/Games/Save $HOME/Games/Save

# Go back to home
cd $HOME
```

## Softwares

All softwares are flatpak based, because it is official and easier to use / install / update:

```bash
# Google Chrome
flatpak install -y flathub com.google.Chrome

# Firefox
flatpak install -y flathub org.mozilla.firefox

# Moonlight
flatpak install -y flathub com.moonlight_stream.Moonlight

# Heroic Games Launcher
flatpak install -y flathub com.heroicgameslauncher.hgl
```

## Emulators

Like softwares, each emulator is flatpak based. Choose the desired emulators to install and create related structures with the following commands:

```bash
# TODO: N64, PSX, PSVITA, XBOX360

# Microsoft Xbox
flatpak install -y flathub app.xemu.xemu
mkdir -p Games/BIOS/XBOX
mkdir -p Games/ROMs/XBOX

# Nintendo 3DS
flatpak install -y flathub org.citra_emu.citra
mkdir -p Games/BIOS/3DS
mkdir -p Games/ROMs/3DS

# Nintendo DS
flatpak install -y flathub net.kuribo64.melonDS
mkdir -p Games/BIOS/NDS
mkdir -p Games/ROMs/NDS

# Nintendo Game Boy Advance
flatpak install -y flathub io.mgba.mGBA
mkdir -p Games/BIOS/GBA
mkdir -p Games/ROMs/GBA

# Nintendo GameCube / Wii
flatpak install -y flathub org.DolphinEmu.dolphin-emu
mkdir -p Games/BIOS/GC
mkdir -p Games/BIOS/WII
mkdir -p Games/ROMs/GC
mkdir -p Games/ROMs/WII

# Nintendo Switch
flatpak install -y flathub org.yuzu_emu.yuzu
flatpak install -y flathub org.ryujinx.Ryujinx
mkdir -p Games/BIOS/SWITCH
mkdir -p Games/ROMs/SWITCH

# Nintendo Wii U
flatpak install -y flathub info.cemu.Cemu
mkdir -p Games/BIOS/WIIU
mkdir -p Games/ROMs/WIIU

# Sega Dreamcast
flatpak install -y flathub org.flycast.Flycast
mkdir -p Games/BIOS/DC
mkdir -p Games/ROMs/DC

# Sony Playstation 2
flatpak install -y flathub net.pcsx2.PCSX2
mkdir -p Games/BIOS/PS2
mkdir -p Games/ROMs/PS2

# Sony Playstation 3
flatpak install -y flathub net.rpcs3.RPCS3
mkdir -p Games/BIOS/PS3
mkdir -p Games/ROMs/PS3

# Sony Playstation Portable
flatpak install -y flathub org.ppsspp.PPSSPP
mkdir -p Games/BIOS/PSP
mkdir -p Games/ROMs/PSP
```

Once you have installed the desired emulators, place the ``BIOS`` and ``ROMs`` for each emulator in their respective folders.

## Application Launcher

Is time to organize the system launcher in Desktop Mode. Right-click on the ``Application Laucher`` icon and select the option ``Edit Applications``. The opened software will allows deep customization of the installed applications on the launcher, including the support for adding new categories (like **Emulators**), renaming apps, removing the not desired ones, sorting and many more. Make the customizations based in your needs.

You can also take the opportunity and add favorites programs to the special favorites section. Open the launcher, right-click in the application and select the ``Add to Favorites`` option.

## Add Programs to Steam

You should now add the softwares and emulators to the Steam. This will allow the use of these programs directly from the ``Gaming Mode``, with all shortcuts and special menus that is available on Steam Deck. For each downloaded software or emulator, open the launcher menu, right-click on the app and choose the ``Add to Steam`` option. Once the apps are added, we can grab their images for nice UI on Steam.

## Adding Steam Images

Run the following set of commands to automatically download the best cover images for the installed programs and wait for its conclusion:

```bash
# If you dont see the images on Steam
# Use the following command to get updated application ids
# grep -i "<game-title>" ~/.local/share/Steam/steamapps/appmanifest_*.acf

cd ~/.local/share/Steam/userdata/${USER_ID}/config
mkdir -p grid && cd grid/

# Required format
# Icon: ${APPID}.ico
# Banner: ${APPID}.png
# Cover: ${APPID}p.png
# Hero: ${APPID}_hero_.png
# Logo: ${APPID}_logo.png
ICONS="https://cdn2.steamgriddb.com/file/sgdb-cdn/icon"
BANNERS="https://cdn2.steamgriddb.com/file/sgdb-cdn/grid"
COVERS="https://cdn2.steamgriddb.com/file/sgdb-cdn/grid"
HEROS="https://cdn2.steamgriddb.com/file/sgdb-cdn/hero"
LOGOS="https://cdn2.steamgriddb.com/file/sgdb-cdn/logo"

# Google Chrome (id 4210646725)
# wget -q -O 4210646725.ico ${ICONS}/3941c4358616274ac2436eacf67fae05.ico
wget -q -O 4210646725.png ${BANNERS}/d40c243072a2d2957b3484e775f1f925.png
wget -q -O 4210646725p.png ${COVERS}/d45c26607db83f6f14b09dd70123913b.png
wget -q -O 4210646725_hero.png ${HEROS}/cae83cfcb1d8a2a4bb17bd1446fb1cee.png
wget -q -O 4210646725_logo.png ${LOGOS}/3b049d0f6cbf5421d399f156807b8657.png

# Firefox (id 3384410319)
wget -q -O 3384410319.png ${BANNERS}/9384fe92aef7ea0128be2c916ed07cea.png
wget -q -O 3384410319p.png ${COVERS}/4529f985441a035ae4a107b8862ba4dd.png
wget -q -O 3384410319_hero.png ${HEROS}/a318166b8539611449bf21ddc297a783.png
wget -q -O 3384410319_logo.png ${LOGOS}/43285a8b542fcdc35377439e05dcb04f.png

# Moonlight (id 2258966675)
wget -q -O 2258966675.png ${BANNERS}/8a8f67cacf3e3d2d63614f515a2079b8.png
wget -q -O 2258966675p.png ${COVERS}/030d60c36d51783da9e4cbb6aa5abd2c.png
wget -q -O 2258966675_hero.png ${HEROS}/0afefa2281c2f8b0b86d6332e2cdbe7d.png
wget -q -O 2258966675_logo.png ${LOGOS}/beb5ad322e679d0a6045c6cfc56e8b92.png

# Heroic Games Launcher (id 2728092030)
wget -q -O 2728092030.png ${BANNERS}/94e8e64cdefe77dcc168855c54f14acd.png
wget -q -O 2728092030p.png ${COVERS}/2b1c6cedeaf9571589e3dc9d51ba20e5.png
wget -q -O 2728092030_hero.png ${HEROS}/bee5ca2551bf346f067a3ac16057bc40.png
wget -q -O 2728092030_logo.png ${LOGOS}/6eebc030d78d41b6cbcf9067aeda9198.png

# Yuzu (id 2259668265)
wget -q -O 2259668265.png ${BANNERS}/dd66229e57c186b4c13e52a8b3f274b2.png
wget -q -O 2259668265p.png ${COVERS}/75aba7a51147cb571a641b8b9f10385e.png
wget -q -O 2259668265_hero.png ${HEROS}/c24f9ae141fa02c7fa1deea7e1149557.png
wget -q -O 2259668265_logo.png ${LOGOS}/55d46c8717ed1cb7ac23556df1745b4b.png

# MelonDS (id 2541270363)
wget -q -O 2541270363.png ${BANNERS}/0ec19bac435cd0ab3fcd2160491b0c7b.png
wget -q -O 2541270363p.png ${COVERS}/3b397c602f7c9226cbcb907b3d5e7d5e.png
wget -q -O 2541270363_hero.png ${HEROS}/c24f9ae141fa02c7fa1deea7e1149557.png
wget -q -O 2541270363_logo.png ${LOGOS}/173f798d1316395cce2c8ecf98aed4d5.png

# MGBA (id 3243913981)
wget -q -O 3243913981.png ${BANNERS}/7088b9d5b6a444224cf6380dcfe61554.png
wget -q -O 3243913981p.png ${COVERS}/d280a227a8ef77d87a5d18037c52776a.png
wget -q -O 3243913981_hero.png ${HEROS}/d470133ccf31f9bfdc1dcb45a30c73b1.png
wget -q -O 3243913981_logo.png ${LOGOS}/e262b1f197f1a9cca59e0868f1e5c94b.png

# Cemu (id 3647450655)
wget -q -O 3647450655.png ${BANNERS}/86fb4d9e1de18ebdb6fc534de828d605.png
wget -q -O 3647450655p.png ${COVERS}/9454c84816d82ed1092f2fe2919a3a8e.png
wget -q -O 3647450655_hero.png ${HEROS}/d5da28d4865fb92720359db84e0dd0dd.png
wget -q -O 3647450655_logo.png ${LOGOS}/c7a9f13a6c0940277d46706c7ca32601.png

# Ryujinx (id 3765673273)
wget -q -O 3765673273.png ${BANNERS}/3931532d087eeb1b1c1a96aba6261802.png
wget -q -O 3765673273p.png ${COVERS}/550d4a283baa604976e81d35d29124df.png
wget -q -O 3765673273_hero.png ${HEROS}/c24f9ae141fa02c7fa1deea7e1149557.png
wget -q -O 3765673273_logo.png ${LOGOS}/b948aa07167c9acb17487657e96870e5.png

# Dolphin (id 4088724280)
wget -q -O 4088724280.png ${BANNERS}/cbec7ddbb30e261abd365bf9f814647d.png
wget -q -O 4088724280p.png ${COVERS}/8a07e4382e18e3b9f5d2713aeaefc29b.png
wget -q -O 4088724280_hero.png ${HEROS}/018b1d3ea470dbb00e3dd6438af19bfb.png
wget -q -O 4088724280_logo.png ${LOGOS}/5b5bbd3170c560829391c3db7265ee9b.png

# Citra (id 2736076325)
wget -q -O 2736076325.png ${BANNERS}/585191595ac24404854bbce59d0f54d2.png
wget -q -O 2736076325p.png ${COVERS}/336fd95d2fd675836a5b72a581072934.png
wget -q -O 2736076325_hero.png ${HEROS}/1d0ba3d7eb612a216c3e4d002deabdb7.png
wget -q -O 2736076325_logo.png ${LOGOS}/30c08c3bbfac55eba7678594e5da022e.png

# RPCS3 (id 3610084102)
wget -q -O 3610084102.png ${BANNERS}/cddaf8b03288749c50afecad7ac3c9a4.png
wget -q -O 3610084102p.png ${COVERS}/ace27c5277ecc8da47cd53ff5c82cb4f.png
wget -q -O 3610084102_hero.png ${HEROS}/15c58997f6690dddb7c501e062a2d1ab.png
wget -q -O 3610084102_logo.png ${LOGOS}/bffc98347ee35b3ead06728d6f073c68.png

# Xemu (id 3182720503)
wget -q -O 3182720503.png ${BANNERS}/5b74752b25bd07933b10b2098970f990.png
wget -q -O 3182720503p.png ${COVERS}/b6cd95d53810282d6a734fbb073e9479.png
wget -q -O 3182720503_hero.png ${HEROS}/aa0994c4263018600494efceae69087a.png
wget -q -O 3182720503_logo.png ${LOGOS}/a42b7cddd7ebb7c1bced17bddc568d2f.png

# PCSX2 (id 4159621629)
wget -q -O 4159621629.png ${BANNERS}/f3a71cf60765edd14269d28819d15327.png
wget -q -O 4159621629p.png ${COVERS}/3123b87d2cede1c04e380a71701ddfe8.png
wget -q -O 4159621629_hero.png ${HEROS}/60312efd57a8cd64fb7f54d5d8e4c2dd.png
wget -q -O 4159621629_logo.png ${LOGOS}/7123c9e46f34491cf4f8eb1a813d8f6e.png

# Flycast (id 2561959160)
wget -q -O 2561959160.png ${BANNERS}/46b3feb0521b4d823847ebbd4dd58ea6.png
wget -q -O 2561959160p.png ${COVERS}/51cf6e65f8242f989f354bf9dfe5a019.png
wget -q -O 2561959160_hero.png ${HEROS}/c24f9ae141fa02c7fa1deea7e1149557.png
wget -q -O 2561959160_logo.png ${LOGOS}/b9b0c8b6beb69bd0c5a213b9422459ce.png

# PPSSPP (id 2676695813)
wget -q -O 2676695813.png ${BANNERS}/88a52c0d85339a377918fdc1ae9dc922.png
wget -q -O 2676695813p.png ${COVERS}/cf476046d346e8091393001a40a523dc.png
wget -q -O 2676695813_hero.png ${HEROS}/b51ecba56e03d4181e0006ff1e8a5355.png
wget -q -O 2676695813_logo.png ${LOGOS}/e242660df1b69b74dcc7fde711f924ff.png
```

Once all images have been downloaded to the Steam, return to ``Gaming Mode`` to finally set the controller layout.

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
- Name it as ``NiceDeck`` with ``...`` as description and confirm save.
- Finally, close the menu to conclude changes.

Once you have it configured on each program, it time to have a nice experience :D.

## Final Tips

- Open the software and configure it on ``Gaming Mode`` first to see how it goes.
- Some programs will require a secondary switch to ``Desktop Mode`` in order to configure advanced settings given the limitations of ``Gaming Mode``.
- Make sure gamepad mapping are correctly configured inside the software. You may also create additional commands for such specific program.
- Make sure to use the directories path from always - in case of symlink from MicroSD card.
- Enjoy!
