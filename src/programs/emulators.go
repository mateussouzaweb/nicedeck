package programs

// Installer for Cemu
func Cemu() *Program {
	return &Program{
		ID:               "cemu",
		Name:             "Cemu",
		Description:      "Emulator for Nintendo Wii U",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/WIIU", "$BIOS/WIIU"},
		FlatpakAppID:     "info.cemu.Cemu",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/9308b0d6e5898366a4a986bc33f3d3e7.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/c7a9f13a6c0940277d46706c7ca32601.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/9454c84816d82ed1092f2fe2919a3a8e.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/86fb4d9e1de18ebdb6fc534de828d605.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/d5da28d4865fb92720359db84e0dd0dd.png",
	}
}

// Installer for Citra
func Citra() *Program {
	return &Program{
		ID:               "citra",
		Name:             "Citra",
		Description:      "Emulator for Nintendo 3DS",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/3DS", "$BIOS/3DS"},
		FlatpakAppID:     "org.citra_emu.citra",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/713586fe8b2dd639aac846e8ac1536a2.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/30c08c3bbfac55eba7678594e5da022e.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/336fd95d2fd675836a5b72a581072934.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/585191595ac24404854bbce59d0f54d2.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/1d0ba3d7eb612a216c3e4d002deabdb7.png",
	}
}

// Installer for Dolphin
func Dolphin() *Program {
	return &Program{
		ID:               "dolphin",
		Name:             "Dolphin Emulator",
		Description:      "Emulator for Nintendo GameCube and Nintendo Wii",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/GC", "$BIOS/GC", "$ROMS/WII", "$BIOS/WII"},
		FlatpakAppID:     "org.DolphinEmu.dolphin-emu",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/7d2a383e54274888b4b73b97e1aaa491.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/5b5bbd3170c560829391c3db7265ee9b.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/8a07e4382e18e3b9f5d2713aeaefc29b.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/cbec7ddbb30e261abd365bf9f814647d.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/018b1d3ea470dbb00e3dd6438af19bfb.png",
	}
}

// Installer for DuckStation
func DuckStation() *Program {
	return &Program{
		ID:               "duckstation",
		Name:             "DuckStation",
		Description:      "Emulator for Sony Playtation 1",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/PS1", "$BIOS/PS1"},
		FlatpakAppID:     "org.duckstation.DuckStation",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/ff0abbcc0227c9124a804b084d161a2d.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/96a0d70498272acfee21d3dbae846113.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/1f7c9b9e37afcbd79ebff19b17837cad.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/9c94e659c62b84bf7b39c599b61bc7d3.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/127f12c937b4baf0a8922eb1384391cf.png",
	}
}

// Installer for Flycast
func Flycast() *Program {
	return &Program{
		ID:               "flycast",
		Name:             "Flycast",
		Description:      "Emulator for Sega Dreamcast",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/DC", "$BIOS/DC"},
		FlatpakAppID:     "org.flycast.Flycast",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/abebb7c39f4b5e46bbcfab2b565ef32b.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/b9b0c8b6beb69bd0c5a213b9422459ce.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/51cf6e65f8242f989f354bf9dfe5a019.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/46b3feb0521b4d823847ebbd4dd58ea6.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	}
}

// Installer for MelonDS
func MelonDS() *Program {
	return &Program{
		ID:               "melonds",
		Name:             "MelonDS",
		Description:      "Emulator for Nintendo DS",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/NDS", "$BIOS/NDS"},
		FlatpakAppID:     "net.kuribo64.melonDS",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/9c156653d889d37811915236feed8660.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/173f798d1316395cce2c8ecf98aed4d5.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/3b397c602f7c9226cbcb907b3d5e7d5e.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/0ec19bac435cd0ab3fcd2160491b0c7b.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	}
}

// Installer for mGBA
func MGBA() *Program {
	return &Program{
		ID:               "mgba",
		Name:             "mGBA",
		Description:      "Emulator for Nintendo Game Boy Advance",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/GBA", "$BIOS/GBA"},
		FlatpakAppID:     "io.mgba.mGBA",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/5b46370c9fd40a27ce2b2abc281064de.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/e262b1f197f1a9cca59e0868f1e5c94b.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/d280a227a8ef77d87a5d18037c52776a.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/7088b9d5b6a444224cf6380dcfe61554.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/d470133ccf31f9bfdc1dcb45a30c73b1.png",
	}
}

// Installer for PCSX2
func PCSX2() *Program {
	return &Program{
		ID:               "pcsx2",
		Name:             "PCSX2",
		Description:      "Emulator for Sony Playstation 2",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/PS2", "$BIOS/PS2"},
		FlatpakAppID:     "net.pcsx2.PCSX2",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/9a32ff36c65e8ba30915a21b7bd76506.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/7123c9e46f34491cf4f8eb1a813d8f6e.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/3123b87d2cede1c04e380a71701ddfe8.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/f3a71cf60765edd14269d28819d15327.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/9cc25407f209e031babdac7d3c520ccb.png",
	}
}

// Installer for PPSSPP
func PPSSPP() *Program {
	return &Program{
		ID:               "ppsspp",
		Name:             "PPSSPP",
		Description:      "Emulator for Sony Playstation Portable",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/PSP", "$BIOS/PSP"},
		FlatpakAppID:     "org.ppsspp.PPSSPP",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/2ba3c4b9390cc43edb94e42144729d33.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/e242660df1b69b74dcc7fde711f924ff.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/cf476046d346e8091393001a40a523dc.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/88a52c0d85339a377918fdc1ae9dc922.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/b51ecba56e03d4181e0006ff1e8a5355.png",
	}
}

// Installer for RPCS3
func RPCS3() *Program {
	return &Program{
		ID:               "rpcs3",
		Name:             "RPCS3",
		Description:      "Emulator for Sony Playstation 3",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/PS3", "$BIOS/PS3"},
		FlatpakAppID:     "net.rpcs3.RPCS3",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/add5aebfcb33a2206b6497d53bc4f309.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/bffc98347ee35b3ead06728d6f073c68.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/ace27c5277ecc8da47cd53ff5c82cb4f.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/cddaf8b03288749c50afecad7ac3c9a4.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/15c58997f6690dddb7c501e062a2d1ab.png",
	}
}

// Installer for Ryujinx
func Ryujinx() *Program {
	return &Program{
		ID:               "ryujinx",
		Name:             "Ryujinx",
		Description:      "Emulator for Nintendo Switch",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/SWITCH", "$BIOS/SWITCH"},
		FlatpakAppID:     "org.ryujinx.Ryujinx",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/6c7cd904122e623ce625613d6af337c4.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/b948aa07167c9acb17487657e96870e5.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/550d4a283baa604976e81d35d29124df.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/3931532d087eeb1b1c1a96aba6261802.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	}
}

// Installer for Simple64
func Simple64() *Program {
	return &Program{
		ID:               "simple64",
		Name:             "Simple64",
		Description:      "Emulator for Nintendo 64",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/N64", "$BIOS/N64"},
		FlatpakAppID:     "io.github.simple64.simple64",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/0ace2e260c8163925254bc878b9eb8ca.png",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/8f6bf2012d96ef9678f8d3a8f27ce358.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/48eeb385ea71aadccce10e2d294879b0.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/e128d1f12ec88795b0a5853d7c754608.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/73888d1bde775303c1749e63e3312a64.png",
	}
}

// Installer for Xemu
func Xemu() *Program {
	return &Program{
		ID:               "xemu",
		Name:             "Xemu",
		Description:      "Emulator for Microsoft Xbox",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/XBOX", "$BIOS/XBOX"},
		FlatpakAppID:     "app.xemu.xemu",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/fac7fead96dafceaf80c1daffeae82a4.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/a42b7cddd7ebb7c1bced17bddc568d2f.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/b6cd95d53810282d6a734fbb073e9479.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/5b74752b25bd07933b10b2098970f990.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/aa0994c4263018600494efceae69087a.png",
	}
}

// Installer for Yuzu
func Yuzu() *Program {
	return &Program{
		ID:               "yuzu",
		Name:             "Yuzu",
		Description:      "Emulator for Nintendo Switch",
		Category:         "Emulators",
		Tags:             []string{"Gaming", "Emulator"},
		RequiredFolders:  []string{"$ROMS/SWITCH", "$BIOS/SWITCH"},
		FlatpakAppID:     "org.yuzu_emu.yuzu",
		FlatpakOverrides: []string{"--filesystem=host"},
		IconURL:          "https://cdn2.steamgriddb.com/icon/2cfa3753d6a524711acb5fce38eeca1a.ico",
		LogoURL:          "https://cdn2.steamgriddb.com/logo/55d46c8717ed1cb7ac23556df1745b4b.png",
		CoverURL:         "https://cdn2.steamgriddb.com/grid/75aba7a51147cb571a641b8b9f10385e.png",
		BannerURL:        "https://cdn2.steamgriddb.com/grid/dd66229e57c186b4c13e52a8b3f274b2.png",
		HeroURL:          "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
	}
}
