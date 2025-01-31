package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Installer for Cemu
func Cemu() *Program {
	return &Program{
		ID:              "cemu",
		Name:            "Cemu",
		Description:     "Emulator for Nintendo Wii U",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/WIIU", "$BIOS/WIIU"},
		Website:         "https://cemu.info",
		IconURL:         "https://cdn2.steamgriddb.com/icon/9308b0d6e5898366a4a986bc33f3d3e7.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/c7a9f13a6c0940277d46706c7ca32601.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/9454c84816d82ed1092f2fe2919a3a8e.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/86fb4d9e1de18ebdb6fc534de828d605.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/d5da28d4865fb92720359db84e0dd0dd.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "info.cemu.Cemu",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "cemu.portable",
			AppName: "$EMULATORS/Cemu.AppImage",
		}, &packaging.MacOS{
			AppID:   "cemu.portable",
			AppName: "$EMULATORS/Cemu.app",
		}, &packaging.Windows{
			AppID:  "Cemu.Portable",
			AppExe: "$EMULATORS\\Cemu\\Cemu.exe",
		}),
	}
}

// Installer for Citra
func Citra() *Program {
	return &Program{
		ID:              "citra",
		Name:            "Citra",
		Description:     "Emulator for Nintendo 3DS",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/3DS", "$BIOS/3DS"},
		Website:         "https://github.com/PabloMK7/citra",
		IconURL:         "https://cdn2.steamgriddb.com/icon/713586fe8b2dd639aac846e8ac1536a2.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/30c08c3bbfac55eba7678594e5da022e.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/336fd95d2fd675836a5b72a581072934.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/585191595ac24404854bbce59d0f54d2.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/1d0ba3d7eb612a216c3e4d002deabdb7.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.citra_emu.citra",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "citra.portable",
			AppName: "$EMULATORS/Citra/citra-qt.AppImage",
		}, &packaging.MacOS{
			AppID:   "citra.portable",
			AppName: "$EMULATORS/Citra/citra-qt.app",
		}, &packaging.Windows{
			AppID:  "Citra.Portable",
			AppExe: "$EMULATORS\\Citra\\citra.exe",
		}),
	}
}

// Installer for Citron
func Citron() *Program {
	return &Program{
		ID:              "citron",
		Name:            "Citron",
		Description:     "Emulator for Nintendo Switch",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/SWITCH", "$BIOS/SWITCH"},
		Website:         "https://citron-emu.org",
		IconURL:         "https://cdn2.steamgriddb.com/icon/e47cf05ff3fa2a1a4a4ee22e02ade796.png",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/bc14e96f34edcda0aa5d04b3634405d2.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/f9065c4db2e5945e8e71e94234119a62.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/c4d3e48c9b104390b762019ccd9174e5.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&packaging.AppImage{
			AppID:   "citron.portable",
			AppName: "$EMULATORS/Citron.AppImage",
		}, &packaging.Windows{
			AppID:  "Citron.Portable",
			AppExe: "$EMULATORS\\Citron\\citron.exe",
		}),
	}
}

// Installer for Dolphin
func Dolphin() *Program {
	return &Program{
		ID:              "dolphin",
		Name:            "Dolphin Emulator",
		Description:     "Emulator for Nintendo GameCube and Nintendo Wii",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/GC", "$BIOS/GC", "$ROMS/WII", "$BIOS/WII"},
		Website:         "https://dolphin-emu.org",
		IconURL:         "https://cdn2.steamgriddb.com/icon/7d2a383e54274888b4b73b97e1aaa491.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/5b5bbd3170c560829391c3db7265ee9b.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/8a07e4382e18e3b9f5d2713aeaefc29b.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/cbec7ddbb30e261abd365bf9f814647d.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/018b1d3ea470dbb00e3dd6438af19bfb.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.DolphinEmu.dolphin-emu",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "dolphin.portable",
			AppName: "$EMULATORS/Dolphin.AppImage",
		}, &packaging.MacOS{
			AppID:   "dolphin.portable",
			AppName: "$EMULATORS/Dolphin.app",
		}, &packaging.Windows{
			AppID:  "Dolphin.Portable",
			AppExe: "$EMULATORS\\Dolphin\\Dolphin.exe",
		}),
	}
}

// Installer for DuckStation
func DuckStation() *Program {
	return &Program{
		ID:              "duckstation",
		Name:            "DuckStation",
		Description:     "Emulator for Sony Playtation 1",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/PS1", "$BIOS/PS1"},
		Website:         "https://www.duckstation.org",
		IconURL:         "https://cdn2.steamgriddb.com/icon/ff0abbcc0227c9124a804b084d161a2d.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/96a0d70498272acfee21d3dbae846113.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/1f7c9b9e37afcbd79ebff19b17837cad.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/9c94e659c62b84bf7b39c599b61bc7d3.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/127f12c937b4baf0a8922eb1384391cf.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.duckstation.DuckStation",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "duckstation.portable",
			AppName: "$EMULATORS/DuckStation.AppImage",
		}, &packaging.MacOS{
			AppID:   "duckstation.portable",
			AppName: "$EMULATORS/DuckStation.app",
		}, &packaging.Windows{
			AppID:  "DuckStation.Portable",
			AppExe: "$EMULATORS\\DuckStation\\duckstation-qt-x64-ReleaseLTCG.exe",
		}),
	}
}

// Installer for Flycast
func Flycast() *Program {
	return &Program{
		ID:              "flycast",
		Name:            "Flycast",
		Description:     "Emulator for Sega Dreamcast",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/DC", "$BIOS/DC"},
		Website:         "https://github.com/flyinghead/flycast",
		IconURL:         "https://cdn2.steamgriddb.com/icon/abebb7c39f4b5e46bbcfab2b565ef32b.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/b9b0c8b6beb69bd0c5a213b9422459ce.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/51cf6e65f8242f989f354bf9dfe5a019.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/46b3feb0521b4d823847ebbd4dd58ea6.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.flycast.Flycast",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "flycast.portable",
			AppName: "$EMULATORS/Flycast.AppImage",
		}, &packaging.MacOS{
			AppID:   "flycast.portable",
			AppName: "$EMULATORS/Flycast.app",
		}, &packaging.Windows{
			AppID:  "Flycast.Portable",
			AppExe: "$EMULATORS\\Flycast\\flycast.exe",
		}),
	}
}

// Installer for Lime3DS
func Lime3DS() *Program {
	return &Program{
		ID:              "lime3ds",
		Name:            "Lime3DS",
		Description:     "Emulator for Nintendo 3DS",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/3DS", "$BIOS/3DS"},
		Website:         "https://lime3ds.github.io",
		IconURL:         "https://cdn2.steamgriddb.com/icon/0dc64a4b9b4c8d205734751c155d528f.png",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/9e6cafbef4b54b72de537851e6aaf6b8.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/012c10e6c703bc4a009d10d95dbd95be.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/1cdcecbcc8ce18ffdb147b29928b5781.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/ae8c643004d25250b521d4f7fc01c354.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "io.github.lime3ds.Lime3DS",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "lime3ds.portable",
			AppName: "$EMULATORS/Lime3DS/lime3ds.AppImage",
		}, &packaging.MacOS{
			AppID:   "lime3ds.portable",
			AppName: "$EMULATORS/Lime3DS/lime3ds.app",
		}, &packaging.Windows{
			AppID:  "Lime3DS.Portable",
			AppExe: "$EMULATORS\\Lime3DS\\lime3ds.exe",
		}),
	}
}

// Installer for MelonDS
func MelonDS() *Program {
	return &Program{
		ID:              "melonds",
		Name:            "MelonDS",
		Description:     "Emulator for Nintendo DS",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/NDS", "$BIOS/NDS"},
		Website:         "https://melonds.kuribo64.net",
		IconURL:         "https://cdn2.steamgriddb.com/icon/9c156653d889d37811915236feed8660.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/173f798d1316395cce2c8ecf98aed4d5.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/3b397c602f7c9226cbcb907b3d5e7d5e.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/0ec19bac435cd0ab3fcd2160491b0c7b.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "net.kuribo64.melonDS",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "melonds.portable",
			AppName: "$EMULATORS/MelonDS.AppImage",
		}, &packaging.MacOS{
			AppID:   "melonds.portable",
			AppName: "$EMULATORS/MelonDS.app",
		}, &packaging.Windows{
			AppID:  "MelonDS.Portable",
			AppExe: "$EMULATORS\\MelonDS\\melonDS.exe",
		}),
	}
}

// Installer for mGBA
func MGBA() *Program {
	return &Program{
		ID:              "mgba",
		Name:            "mGBA",
		Description:     "Emulator for Nintendo Game Boy Advance",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/GBA", "$BIOS/GBA"},
		Website:         "https://mgba.io",
		IconURL:         "https://cdn2.steamgriddb.com/icon/5b46370c9fd40a27ce2b2abc281064de.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/e262b1f197f1a9cca59e0868f1e5c94b.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/d280a227a8ef77d87a5d18037c52776a.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/7088b9d5b6a444224cf6380dcfe61554.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/d470133ccf31f9bfdc1dcb45a30c73b1.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "io.mgba.mGBA",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "mgba.portable",
			AppName: "$EMULATORS/MGBA.AppImage",
		}, &packaging.MacOS{
			AppID:   "mgba.portable",
			AppName: "$EMULATORS/MGBA.app",
		}, &packaging.Windows{
			AppID:  "MGBA.Portable",
			AppExe: "$EMULATORS\\MGBA\\mGBA.exe",
		}),
	}
}

// Installer for PCSX2
func PCSX2() *Program {
	return &Program{
		ID:              "pcsx2",
		Name:            "PCSX2",
		Description:     "Emulator for Sony Playstation 2",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/PS2", "$BIOS/PS2"},
		Website:         "https://pcsx2.net",
		IconURL:         "https://cdn2.steamgriddb.com/icon/9a32ff36c65e8ba30915a21b7bd76506.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/7123c9e46f34491cf4f8eb1a813d8f6e.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/3123b87d2cede1c04e380a71701ddfe8.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/f3a71cf60765edd14269d28819d15327.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/9cc25407f209e031babdac7d3c520ccb.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "net.pcsx2.PCSX2",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "pcsx2.portable",
			AppName: "$EMULATORS/PCSX2.AppImage",
		}, &packaging.MacOS{
			AppID:   "pcsx2.portable",
			AppName: "$EMULATORS/PCSX2.app",
		}, &packaging.Windows{
			AppID:  "PCSX2.Portable",
			AppExe: "$EMULATORS\\PCSX2\\pcsx2-qt.exe",
		}),
	}
}

// Installer for PPSSPP
func PPSSPP() *Program {
	return &Program{
		ID:              "ppsspp",
		Name:            "PPSSPP",
		Description:     "Emulator for Sony Playstation Portable",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/PSP", "$BIOS/PSP"},
		Website:         "https://www.ppsspp.org",
		IconURL:         "https://cdn2.steamgriddb.com/icon/2ba3c4b9390cc43edb94e42144729d33.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/e242660df1b69b74dcc7fde711f924ff.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/cf476046d346e8091393001a40a523dc.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/88a52c0d85339a377918fdc1ae9dc922.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/b51ecba56e03d4181e0006ff1e8a5355.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.ppsspp.PPSSPP",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "ppsspp.portable",
			AppName: "$EMULATORS/PPSSPP.AppImage",
		}, &packaging.MacOS{
			AppID:   "ppsspp.portable",
			AppName: "$EMULATORS/PPSSPP.app",
		}, &packaging.Windows{
			AppID:  "PPSSPP.Portable",
			AppExe: "$EMULATORS\\PPSSPP\\PPSSPPWindows64.exe",
		}),
	}
}

// Installer for Redream
func Redream() *Program {
	return &Program{
		ID:              "redream",
		Name:            "Redream",
		Description:     "Emulator for Sega Dreamcast",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/DC", "$BIOS/DC"},
		Website:         "https://redream.io",
		IconURL:         "https://cdn2.steamgriddb.com/icon/e1878879f60985631df0dc2da79396a0.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/6c11cb78b7bbb5c22d5f5271b5494381.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/dd5fbbf85c3198ece6dcd86166c58439.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/2e834824cdba6141dcb14688597a26fa.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/4853deb8a06838c502fc6cb6ce91f704.png",
		Package: packaging.Best(&packaging.Linux{
			AppID:  "redream.portable",
			AppBin: "$EMULATORS/Redream/redream",
		}, &packaging.MacOS{
			AppID:   "redream.portable",
			AppName: "$EMULATORS/Redream.app",
		}, &packaging.Windows{
			AppID:  "Redream.Portable",
			AppExe: "$EMULATORS\\Redream\\redream.exe",
		}),
	}
}

// Installer for RPCS3
func RPCS3() *Program {
	return &Program{
		ID:              "rpcs3",
		Name:            "RPCS3",
		Description:     "Emulator for Sony Playstation 3",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/PS3", "$BIOS/PS3"},
		Website:         "https://rpcs3.net",
		IconURL:         "https://cdn2.steamgriddb.com/icon/add5aebfcb33a2206b6497d53bc4f309.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/bffc98347ee35b3ead06728d6f073c68.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/ace27c5277ecc8da47cd53ff5c82cb4f.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/cddaf8b03288749c50afecad7ac3c9a4.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/15c58997f6690dddb7c501e062a2d1ab.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "net.rpcs3.RPCS3",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "rpcs3.portable",
			AppName: "$EMULATORS/RPCS3.AppImage",
		}, &packaging.MacOS{
			AppID:   "rpcs3.portable",
			AppName: "$EMULATORS/RPCS3.app",
		}, &packaging.Windows{
			AppID:  "RPCS3.Portable",
			AppExe: "$EMULATORS\\RPCS3\\rpcs3.exe",
		}),
	}
}

// Installer for Ryujinx
func Ryujinx() *Program {
	return &Program{
		ID:              "ryujinx",
		Name:            "Ryujinx",
		Description:     "Emulator for Nintendo Switch",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/SWITCH", "$BIOS/SWITCH"},
		Website:         "https://github.com/Ryubing/Ryujinx",
		IconURL:         "https://cdn2.steamgriddb.com/icon/6c7cd904122e623ce625613d6af337c4.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/b948aa07167c9acb17487657e96870e5.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/550d4a283baa604976e81d35d29124df.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/3931532d087eeb1b1c1a96aba6261802.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.ryujinx.Ryujinx",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "ryujinx.portable",
			AppName: "$EMULATORS/Ryujinx.AppImage",
		}, &packaging.MacOS{
			AppID:   "ryujinx.portable",
			AppName: "$EMULATORS/Ryujinx.app",
		}, &packaging.Windows{
			AppID:  "Ryujinx.Portable",
			AppExe: "$EMULATORS\\Ryujinx\\Ryujinx.exe",
		}),
	}
}

// Installer for ShadPS4
func ShadPS4() *Program {
	return &Program{
		ID:              "shadps4",
		Name:            "ShadPS4",
		Description:     "Emulator for Sony Playstation 4",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/PS4", "$BIOS/PS4"},
		Website:         "https://shadps4.net",
		IconURL:         "https://cdn2.steamgriddb.com/icon/3c7941e8f5200be6925e75ed4063311a.png",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/6c70dccf452364ce8e5a9c44c88dd6c1.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/72251a01ac19b84c2208c2a6f18a17da.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/21483d9d9aca5bd442f292cef7bab951.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/cc5e9cea0a79b962c20a9231e65a06ef.jpg",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "net.shadps4.shadPS4",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "shadps4.portable",
			AppName: "$EMULATORS/ShadPS4.AppImage",
		}, &packaging.MacOS{
			AppID:   "shadps4.portable",
			AppName: "$EMULATORS/ShadPS4.app",
		}, &packaging.Windows{
			AppID:  "ShadPS4.Portable",
			AppExe: "$EMULATORS\\ShadPS4\\shadPS4.exe",
		}),
	}
}

// Installer for Simple64
func Simple64() *Program {
	return &Program{
		ID:              "simple64",
		Name:            "Simple64",
		Description:     "Emulator for Nintendo 64",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/N64", "$BIOS/N64"},
		Website:         "https://simple64.github.io",
		IconURL:         "https://cdn2.steamgriddb.com/icon/0ace2e260c8163925254bc878b9eb8ca.png",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/8f6bf2012d96ef9678f8d3a8f27ce358.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/48eeb385ea71aadccce10e2d294879b0.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/e128d1f12ec88795b0a5853d7c754608.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/73888d1bde775303c1749e63e3312a64.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "io.github.simple64.simple64",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.Windows{
			AppID:  "Simple64.Portable",
			AppExe: "$EMULATORS\\Simple64\\simple64-gui.exe",
		}),
	}
}

// Installer for Vita3K
func Vita3K() *Program {
	return &Program{
		ID:              "vita3k",
		Name:            "Vita3k",
		Description:     "Emulator for Sony Playstation Vita",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/PSVITA", "$BIOS/PSVITA"},
		Website:         "https://vita3k.org",
		IconURL:         "https://cdn2.steamgriddb.com/icon/1a4a5f89e71e4bb9973355c964a950b4.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/654798fc20b6d08b12236106fff87920.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/d371697094a73577074c10fb6688f2ff.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/a1357c62042fedf5f0a71ebacfe5987d.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/5e98be1eed79374e1edd72f4b1d838b4.png",
		Package: packaging.Best(&packaging.Linux{
			AppID:  "vita3k.portable",
			AppBin: "$EMULATORS/Vita3K/Vita3K",
		}, &packaging.MacOS{
			AppID:   "vita3k.portable",
			AppName: "$EMULATORS/Vita3K.app",
		}, &packaging.Windows{
			AppID:  "Vita3K.Portable",
			AppExe: "$EMULATORS\\Vita3K\\Vita3K.exe",
		}),
	}
}

// Installer for Xemu
func Xemu() *Program {
	return &Program{
		ID:              "xemu",
		Name:            "Xemu",
		Description:     "Emulator for Microsoft Xbox",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/XBOX", "$BIOS/XBOX"},
		Website:         "https://xemu.app",
		IconURL:         "https://cdn2.steamgriddb.com/icon/fac7fead96dafceaf80c1daffeae82a4.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/a42b7cddd7ebb7c1bced17bddc568d2f.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/b6cd95d53810282d6a734fbb073e9479.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/5b74752b25bd07933b10b2098970f990.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/aa0994c4263018600494efceae69087a.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "app.xemu.xemu",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "xemu.portable",
			AppName: "$EMULATORS/Xemu.AppImage",
		}, &packaging.MacOS{
			AppID:   "xemu.portable",
			AppName: "$EMULATORS/Xemu.app",
		}, &packaging.Windows{
			AppID:  "Xemu.Portable",
			AppExe: "$EMULATORS\\Xemu\\xemu.exe",
		}),
	}
}

// Installer for Xenia
func Xenia() *Program {
	return &Program{
		ID:              "xenia",
		Name:            "Xenia",
		Description:     "Emulator for Microsoft Xbox 360",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/XBOX360", "$BIOS/XBOX360"},
		Website:         "https://xenia.jp",
		IconURL:         "https://cdn2.steamgriddb.com/icon/420c841038c492fed4d19999a813009d.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/fac05328668f599efe18e76cdb284aab.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/e43e55468f8cfee48d517b2c49cecd08.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/1962bcb00dc1bf1b5bcb334257ff3701.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/2958ef004a18f50b380a87d1cfe5366d.png",
		Package: packaging.Best(&packaging.Windows{
			AppID:  "Xenia.Portable",
			AppExe: "$EMULATORS\\Xenia\\xenia_canary.exe",
		}, &packaging.Windows{
			AppID:  "Xenia.Portable",
			AppExe: "$EMULATORS\\Xenia\\xenia.exe",
		}),
	}
}

// Installer for Yuzu
func Yuzu() *Program {
	return &Program{
		ID:              "yuzu",
		Name:            "Yuzu",
		Description:     "Emulator for Nintendo Switch",
		Category:        "Emulators",
		Tags:            []string{"Gaming", "Emulator"},
		RequiredFolders: []string{"$ROMS/SWITCH", "$BIOS/SWITCH"},
		Website:         "https://yuzu-mirror.github.io",
		IconURL:         "https://cdn2.steamgriddb.com/icon/2cfa3753d6a524711acb5fce38eeca1a.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/55d46c8717ed1cb7ac23556df1745b4b.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/75aba7a51147cb571a641b8b9f10385e.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/dd66229e57c186b4c13e52a8b3f274b2.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.yuzu_emu.yuzu",
			Overrides: []string{"--filesystem=host"},
		}, &packaging.AppImage{
			AppID:   "yuzu.portable",
			AppName: "$EMULATORS/Yuzu.AppImage",
		}, &packaging.Windows{
			AppID:  "Yuzu.Portable",
			AppExe: "$EMULATORS\\Yuzu\\yuzu.exe",
		}),
	}
}
