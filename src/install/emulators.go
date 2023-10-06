package install

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

// Install emulator for Nintendo Wii U - Cemu
func Cemu() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub info.cemu.Cemu
		mkdir -p $HOME/Games/BIOS/WIIU
		mkdir -p $HOME/Games/ROMs/WIIU
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Cemu",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/info.cemu.Cemu.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=/app/bin/Cemu-wrapper info.cemu.Cemu",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/9308b0d6e5898366a4a986bc33f3d3e7.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/c7a9f13a6c0940277d46706c7ca32601.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/9454c84816d82ed1092f2fe2919a3a8e.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/86fb4d9e1de18ebdb6fc534de828d605.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/d5da28d4865fb92720359db84e0dd0dd.png",
	})

	return err
}

// Install emulator for Nintendo 3DS - Citra
func Citra() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub org.citra_emu.citra
		mkdir -p $HOME/Games/BIOS/3DS
		mkdir -p $HOME/Games/ROMs/3DS
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Citra",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.citra_emu.citra.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=citra-qt --file-forwarding org.citra_emu.citra @@ %f @@",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/713586fe8b2dd639aac846e8ac1536a2.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/30c08c3bbfac55eba7678594e5da022e.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/336fd95d2fd675836a5b72a581072934.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/585191595ac24404854bbce59d0f54d2.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/1d0ba3d7eb612a216c3e4d002deabdb7.png",
	})

	return err
}

// Install emulator for Nintendo GameCube and Wii - Dolphin
func Dolphin() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub org.DolphinEmu.dolphin-emu
		mkdir -p $HOME/Games/BIOS/GC
		mkdir -p $HOME/Games/BIOS/WII
		mkdir -p $HOME/Games/ROMs/GC
		mkdir -p $HOME/Games/ROMs/WII
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Dolphin Emulator",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.DolphinEmu.dolphin-emu.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=/app/bin/dolphin-emu-wrapper org.DolphinEmu.dolphin-emu",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/7d2a383e54274888b4b73b97e1aaa491.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/5b5bbd3170c560829391c3db7265ee9b.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/8a07e4382e18e3b9f5d2713aeaefc29b.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/cbec7ddbb30e261abd365bf9f814647d.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/018b1d3ea470dbb00e3dd6438af19bfb.png",
	})

	return err
}

// Install emulator for Sega Dreamcast - Flycast
func Flycast() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub org.flycast.Flycast
		mkdir -p $HOME/Games/BIOS/DC
		mkdir -p $HOME/Games/ROMs/DC
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Flycast",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.flycast.Flycast.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=flycast --file-forwarding org.flycast.Flycast @@ %f @@",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/abebb7c39f4b5e46bbcfab2b565ef32b.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/b9b0c8b6beb69bd0c5a213b9422459ce.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/51cf6e65f8242f989f354bf9dfe5a019.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/46b3feb0521b4d823847ebbd4dd58ea6.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	})

	return err
}

// Install emulator for Nintendo DS - MelonDS
func MelonDS() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub net.kuribo64.melonDS
		mkdir -p $HOME/Games/BIOS/NDS
		mkdir -p $HOME/Games/ROMs/NDS
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "MelonDS",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/net.kuribo64.melonDS.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=melonDS --file-forwarding net.kuribo64.melonDS @@ %f @@",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/9c156653d889d37811915236feed8660.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/173f798d1316395cce2c8ecf98aed4d5.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3b397c602f7c9226cbcb907b3d5e7d5e.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/0ec19bac435cd0ab3fcd2160491b0c7b.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	})

	return err
}

// Install emulator for Nintendo Game Boy Advance - mGBA
func MGBA() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub io.mgba.mGBA
		mkdir -p $HOME/Games/BIOS/GBA
		mkdir -p $HOME/Games/ROMs/GBA
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "MGBA",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/io.mgba.mGBA.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=mgba-qt --file-forwarding io.mgba.mGBA @@ %f @@",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/5b46370c9fd40a27ce2b2abc281064de.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/e262b1f197f1a9cca59e0868f1e5c94b.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d280a227a8ef77d87a5d18037c52776a.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/7088b9d5b6a444224cf6380dcfe61554.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/d470133ccf31f9bfdc1dcb45a30c73b1.png",
	})

	return err
}

// Install emulator for Sony Playstation 2 - PCSX2
func PCSX2() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub net.pcsx2.PCSX2
		mkdir -p $HOME/Games/BIOS/PS2
		mkdir -p $HOME/Games/ROMs/PS2
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "PCSX2",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/net.pcsx2.PCSX2.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=pcsx2-qt net.pcsx2.PCSX2",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/9a32ff36c65e8ba30915a21b7bd76506.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/7123c9e46f34491cf4f8eb1a813d8f6e.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3123b87d2cede1c04e380a71701ddfe8.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/f3a71cf60765edd14269d28819d15327.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/9cc25407f209e031babdac7d3c520ccb.png",
	})

	return err
}

// Install emulator for Sony Playstation Portable - PPSSPP
func PPSSPP() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub org.ppsspp.PPSSPP
		mkdir -p $HOME/Games/BIOS/PSP
		mkdir -p $HOME/Games/ROMs/PSP
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "PPSSPP",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.ppsspp.PPSSPP.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=PPSSPPSDL org.ppsspp.PPSSPP",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/2ba3c4b9390cc43edb94e42144729d33.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/e242660df1b69b74dcc7fde711f924ff.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/cf476046d346e8091393001a40a523dc.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/88a52c0d85339a377918fdc1ae9dc922.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/b51ecba56e03d4181e0006ff1e8a5355.png",
	})

	return err
}

// Install emulator for Sony Playstation 3 - RPCS3
func RPCS3() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub net.rpcs3.RPCS3
		mkdir -p $HOME/Games/BIOS/PS3
		mkdir -p $HOME/Games/ROMs/PS3
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "RPCS3",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/net.rpcs3.RPCS3.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=rpcs3 --file-forwarding net.rpcs3.RPCS3 @@ %f @@",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/add5aebfcb33a2206b6497d53bc4f309.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/bffc98347ee35b3ead06728d6f073c68.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/ace27c5277ecc8da47cd53ff5c82cb4f.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/cddaf8b03288749c50afecad7ac3c9a4.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/15c58997f6690dddb7c501e062a2d1ab.png",
	})

	return err
}

// Install emulator for Nintendo Switch - Ryujinx
func Ryujinx() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub org.ryujinx.Ryujinx
		mkdir -p $HOME/Games/BIOS/SWITCH
		mkdir -p $HOME/Games/ROMs/SWITCH
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Ryujinx",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.ryujinx.Ryujinx.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=ryujinx-wrapper --file-forwarding org.ryujinx.Ryujinx @@ %f @@",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/6c7cd904122e623ce625613d6af337c4.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/b948aa07167c9acb17487657e96870e5.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/550d4a283baa604976e81d35d29124df.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3931532d087eeb1b1c1a96aba6261802.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	})

	return err
}

// Install emulator Microsoft Xbox - Xemu
func Xemu() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub app.xemu.xemu
		mkdir -p $HOME/Games/BIOS/XBOX
		mkdir -p $HOME/Games/ROMs/XBOX
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Xemu",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/app.xemu.xemu.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=xemu app.xemu.xemu",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/fac7fead96dafceaf80c1daffeae82a4.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/a42b7cddd7ebb7c1bced17bddc568d2f.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/b6cd95d53810282d6a734fbb073e9479.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/5b74752b25bd07933b10b2098970f990.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/aa0994c4263018600494efceae69087a.png",
	})

	return err
}

// Install emulator for Nintendo Switch - Yuzu
func Yuzu() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub org.yuzu_emu.yuzu
		mkdir -p $HOME/Games/BIOS/SWITCH
		mkdir -p $HOME/Games/ROMs/SWITCH
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Yuzu",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.yuzu_emu.yuzu.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=yuzu-launcher org.yuzu_emu.yuzu",
		Tags:          []string{"EMULATORS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/2cfa3753d6a524711acb5fce38eeca1a.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/55d46c8717ed1cb7ac23556df1745b4b.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/75aba7a51147cb571a641b8b9f10385e.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/dd66229e57c186b4c13e52a8b3f274b2.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	})

	return err
}
