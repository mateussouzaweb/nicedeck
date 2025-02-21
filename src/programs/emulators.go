package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/programs/github"
)

// Installer for Cemu
func Cemu() *Program {
	return &Program{
		ID:          "cemu",
		Name:        "Cemu",
		Description: "Emulator for Nintendo Wii U",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Cemu", "$ROMS/WIIU", "$BIOS/WIIU"},
		Website:     "https://cemu.info",
		IconURL:     "https://cdn2.steamgriddb.com/icon/2c790f933dcb0c7a747741780c6b435d.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/c7a9f13a6c0940277d46706c7ca32601.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/9454c84816d82ed1092f2fe2919a3a8e.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/86fb4d9e1de18ebdb6fc534de828d605.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/d5da28d4865fb92720359db84e0dd0dd.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "info.cemu.Cemu",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "cemu",
			AppName: "$EMULATORS/Cemu/Cemu.AppImage",
			Source: github.Release(
				"https://github.com/cemu-project/Cemu",
				"Cemu-*-x86_64.AppImage",
			),
		}, &macos.Application{
			AppID:   "cemu",
			AppName: "$EMULATORS/Cemu/Cemu.app",
			Source: github.Release(
				"https://github.com/cemu-project/Cemu",
				"cemu-*-macos-12-x64.dmg",
			),
		}, &windows.Executable{
			AppID:  "Cemu",
			AppExe: "$EMULATORS\\Cemu\\Cemu.exe",
			Source: github.Release(
				"https://github.com/cemu-project/Cemu",
				"cemu-*-windows-x64.zip",
			),
		}),
	}
}

// Installer for Citra
func Citra() *Program {
	return &Program{
		ID:          "citra",
		Name:        "Citra",
		Description: "Emulator for Nintendo 3DS",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Citra", "$ROMS/3DS", "$BIOS/3DS"},
		Website:     "https://github.com/PabloMK7/citra",
		IconURL:     "https://cdn2.steamgriddb.com/icon/9191e0c0fc4f0ff1d9e4bae7e118944e.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/30c08c3bbfac55eba7678594e5da022e.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/336fd95d2fd675836a5b72a581072934.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/585191595ac24404854bbce59d0f54d2.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/1d0ba3d7eb612a216c3e4d002deabdb7.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "org.citra_emu.citra",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "citra",
			AppName: "$EMULATORS/Citra/citra-qt.AppImage",
			Source: github.Release(
				"https://github.com/PabloMK7/citra",
				"citra-linux-appimage-*.tar.gz",
			),
		}, &macos.Application{
			AppID:   "citra",
			AppName: "$EMULATORS/Citra/citra-qt.app",
			Source: github.Release(
				"https://github.com/PabloMK7/citra",
				"citra-macos-universal-*.tar.gz",
			),
		}, &windows.Executable{
			AppID:  "Citra",
			AppExe: "$EMULATORS\\Citra\\citra-qt.exe",
			Source: github.Release(
				"https://github.com/PabloMK7/citra",
				"citra-windows-msvc-*.zip",
			),
		}),
	}
}

// Installer for Citron
func Citron() *Program {
	return &Program{
		ID:          "citron",
		Name:        "Citron",
		Description: "Emulator for Nintendo Switch",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Citron", "$ROMS/SWITCH", "$BIOS/SWITCH"},
		Website:     "https://citron-emu.org",
		IconURL:     "https://cdn2.steamgriddb.com/icon/e47cf05ff3fa2a1a4a4ee22e02ade796.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/bc14e96f34edcda0aa5d04b3634405d2.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/f9065c4db2e5945e8e71e94234119a62.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/c4d3e48c9b104390b762019ccd9174e5.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&linux.AppImage{
			AppID:   "citron",
			AppName: "$EMULATORS/Citron/Citron.AppImage",
			Source: github.Release(
				"https://github.com/pkgforge-dev/Citron-AppImage",
				"Citron-*-anylinux-x86_64.AppImage",
			),
		}, &windows.Executable{
			AppID:  "Citron",
			AppExe: "$EMULATORS\\Citron\\citron.exe",
			Source: &packaging.Source{
				URL:    "https://git.citron-emu.org/Citron/Citron/releases/download/v0.5-canary-refresh/Citron-Windows-Canary-Refresh_0.5.zip",
				Format: "zip",
			},
		}),
	}
}

// Installer for Dolphin
func Dolphin() *Program {
	return &Program{
		ID:          "dolphin",
		Name:        "Dolphin Emulator",
		Description: "Emulator for Nintendo GameCube and Nintendo Wii",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Dolphin", "$ROMS/GC", "$BIOS/GC", "$ROMS/WII", "$BIOS/WII"},
		Website:     "https://dolphin-emu.org",
		IconURL:     "https://cdn2.steamgriddb.com/icon/52ec3baefcf93a558d994e1bcd3b5c3d.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/5b5bbd3170c560829391c3db7265ee9b.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/8a07e4382e18e3b9f5d2713aeaefc29b.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/cbec7ddbb30e261abd365bf9f814647d.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/018b1d3ea470dbb00e3dd6438af19bfb.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "org.DolphinEmu.dolphin-emu",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "dolphin",
			AppName: "$EMULATORS/Dolphin/Dolphin.AppImage",
			Source: github.Release(
				"https://github.com/pkgforge-dev/Dolphin-emu-AppImage",
				"Dolphin_Emulator-*-anylinux-x86_64.AppImage",
			),
		}, &macos.Application{
			AppID:   "dolphin",
			AppName: "$EMULATORS/Dolphin/Dolphin.app",
			Source: &packaging.Source{
				URL:    "https://dl.dolphin-emu.org/releases/2412/dolphin-2412-universal.dmg",
				Format: "dmg",
			},
		}, &windows.Executable{
			AppID:  "Dolphin",
			AppExe: "$EMULATORS\\Dolphin\\Dolphin.exe",
			Source: &packaging.Source{
				URL:    "https://dl.dolphin-emu.org/releases/2412/dolphin-2412-x64.7z",
				Format: "7z",
			},
		}),
	}
}

// Installer for DuckStation
func DuckStation() *Program {
	return &Program{
		ID:          "duckstation",
		Name:        "DuckStation",
		Description: "Emulator for Sony Playtation 1",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/DuckStation", "$ROMS/PS1", "$BIOS/PS1"},
		Website:     "https://www.duckstation.org",
		IconURL:     "https://cdn2.steamgriddb.com/icon/f985f43b4ba330d5282dfd9be8003e62.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/96a0d70498272acfee21d3dbae846113.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/1f7c9b9e37afcbd79ebff19b17837cad.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/9c94e659c62b84bf7b39c599b61bc7d3.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/127f12c937b4baf0a8922eb1384391cf.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "org.duckstation.DuckStation",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "duckstation",
			AppName: "$EMULATORS/DuckStation/DuckStation.AppImage",
			Source: github.Release(
				"https://github.com/stenzek/duckstation",
				"DuckStation-x64.AppImage",
			),
		}, &macos.Application{
			AppID:   "duckstation",
			AppName: "$EMULATORS/DuckStation/DuckStation.app",
			Source: github.Release(
				"https://github.com/stenzek/duckstation",
				"duckstation-mac-release.zip",
			),
		}, &windows.Executable{
			AppID:  "DuckStation",
			AppExe: "$EMULATORS\\DuckStation\\duckstation-qt-x64-ReleaseLTCG.exe",
			Source: github.Release(
				"https://github.com/stenzek/duckstation",
				"duckstation-windows-x64-release.zip",
			),
		}),
	}
}

// Installer for Flycast
func Flycast() *Program {
	return &Program{
		ID:          "flycast",
		Name:        "Flycast",
		Description: "Emulator for Sega Dreamcast",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Flycast", "$ROMS/DC", "$BIOS/DC"},
		Website:     "https://github.com/flyinghead/flycast",
		IconURL:     "https://cdn2.steamgriddb.com/icon/858a2be748405d1cf063e97622abc791.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/b9b0c8b6beb69bd0c5a213b9422459ce.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/51cf6e65f8242f989f354bf9dfe5a019.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/46b3feb0521b4d823847ebbd4dd58ea6.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "org.flycast.Flycast",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "flycast",
			AppName: "$EMULATORS/Flycast/Flycast.AppImage",
			Source: github.Release(
				"https://github.com/flyinghead/flycast",
				"flycast-x86_64.AppImage",
			),
		}, &macos.Application{
			AppID:   "flycast",
			AppName: "$EMULATORS/Flycast/Flycast.app",
			Source: github.Release(
				"https://github.com/flyinghead/flycast",
				"flycast-macOS-*.zip",
			),
		}, &windows.Executable{
			AppID:  "Flycast",
			AppExe: "$EMULATORS\\Flycast\\flycast.exe",
			Source: github.Release(
				"https://github.com/flyinghead/flycast",
				"flycast-win64-*.zip",
			),
		}),
	}
}

// Installer for Lime3DS
func Lime3DS() *Program {
	return &Program{
		ID:          "lime3ds",
		Name:        "Lime3DS",
		Description: "Emulator for Nintendo 3DS",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Lime3DS", "$ROMS/3DS", "$BIOS/3DS"},
		Website:     "https://lime3ds.github.io",
		IconURL:     "https://cdn2.steamgriddb.com/icon/0dc64a4b9b4c8d205734751c155d528f.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/9e6cafbef4b54b72de537851e6aaf6b8.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/012c10e6c703bc4a009d10d95dbd95be.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/1cdcecbcc8ce18ffdb147b29928b5781.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/ae8c643004d25250b521d4f7fc01c354.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "io.github.lime3ds.Lime3DS",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "lime3ds",
			AppName: "$EMULATORS/Lime3DS/lime3ds.AppImage",
			Source: github.Release(
				"https://github.com/Lime3DS/lime3ds-archive",
				"lime3ds-*-linux-appimage.tar.gz",
			),
		}, &macos.Application{
			AppID:   "lime3ds",
			AppName: "$EMULATORS/Lime3DS/lime3ds.app",
			Source: github.Release(
				"https://github.com/Lime3DS/lime3ds-archive",
				"lime3ds-*-macos-universal.zip",
			),
		}, &windows.Executable{
			AppID:  "Lime3DS",
			AppExe: "$EMULATORS\\Lime3DS\\lime3ds.exe",
			Source: github.Release(
				"https://github.com/Lime3DS/lime3ds-archive",
				"lime3ds-*-windows-msvc.zip",
			),
		}),
	}
}

// Installer for MelonDS
func MelonDS() *Program {
	return &Program{
		ID:          "melonds",
		Name:        "MelonDS",
		Description: "Emulator for Nintendo DS",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/MelonDS", "$ROMS/NDS", "$BIOS/NDS"},
		Website:     "https://melonds.kuribo64.net",
		IconURL:     "https://cdn2.steamgriddb.com/icon/8ad297da4fc2cc28dfa3c0cb7df8ae63.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/173f798d1316395cce2c8ecf98aed4d5.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/3b397c602f7c9226cbcb907b3d5e7d5e.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/0ec19bac435cd0ab3fcd2160491b0c7b.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "net.kuribo64.melonDS",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "melonds",
			AppName: "$EMULATORS/MelonDS/melonDS.AppImage",
			Source: github.Release(
				"https://github.com/melonDS-emu/melonDS",
				"melonDS-appimage-x86_64.zip",
			),
		}, &macos.Application{
			AppID:   "melonds",
			AppName: "$EMULATORS/MelonDS/melonDS.app",
			Source: github.Release(
				"https://github.com/melonDS-emu/melonDS",
				"macOS-universal.zip",
			),
		}, &windows.Executable{
			AppID:  "MelonDS",
			AppExe: "$EMULATORS\\MelonDS\\melonDS.exe",
			Source: github.Release(
				"https://github.com/melonDS-emu/melonDS",
				"melonDS-windows-x86_64.zip",
			),
		}),
	}
}

// Installer for mGBA
func MGBA() *Program {
	return &Program{
		ID:          "mgba",
		Name:        "MGBA",
		Description: "Emulator for Nintendo Game Boy Advance",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/MGBA", "$ROMS/GBA", "$BIOS/GBA"},
		Website:     "https://mgba.io",
		IconURL:     "https://cdn2.steamgriddb.com/icon/7d5fe9a6097c89cda3ce6a5b1de4dd15.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/e262b1f197f1a9cca59e0868f1e5c94b.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/d280a227a8ef77d87a5d18037c52776a.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/7088b9d5b6a444224cf6380dcfe61554.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/d470133ccf31f9bfdc1dcb45a30c73b1.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "io.mgba.mGBA",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "mgba",
			AppName: "$EMULATORS/MGBA/mGBA.AppImage",
			Source: github.Release(
				"https://github.com/mgba-emu/mgba",
				"mGBA-*-appimage-x64.appimage",
			),
		}, &macos.Application{
			AppID:   "mgba",
			AppName: "$EMULATORS/MGBA/mGBA.app",
			Source: github.Release(
				"https://github.com/mgba-emu/mgba",
				"mGBA-*-macos.dmg",
			),
		}, &windows.Executable{
			AppID:  "MGBA",
			AppExe: "$EMULATORS\\MGBA\\mGBA.exe",
			Source: github.Release(
				"https://github.com/mgba-emu/mgba",
				"mGBA-*-win64.7z",
			),
		}),
	}
}

// Installer for PCSX2
func PCSX2() *Program {
	return &Program{
		ID:          "pcsx2",
		Name:        "PCSX2",
		Description: "Emulator for Sony Playstation 2",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/PCSX2", "$ROMS/PS2", "$BIOS/PS2"},
		Website:     "https://pcsx2.net",
		IconURL:     "https://cdn2.steamgriddb.com/icon/22ec872ee633043cc5aece5adb261367.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/7123c9e46f34491cf4f8eb1a813d8f6e.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/3123b87d2cede1c04e380a71701ddfe8.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/f3a71cf60765edd14269d28819d15327.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/9cc25407f209e031babdac7d3c520ccb.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "net.pcsx2.PCSX2",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "pcsx2",
			AppName: "$EMULATORS/PCSX2/PCSX2.AppImage",
			Source: github.Release(
				"https://github.com/PCSX2/pcsx2",
				"pcsx2-*-linux-appimage-x64-Qt.AppImage",
			),
		}, &macos.Application{
			AppID:   "pcsx2",
			AppName: "$EMULATORS/PCSX2/PCSX2.app",
			Source: github.Release(
				"https://github.com/PCSX2/pcsx2",
				"pcsx2-*-macos-Qt.tar.xz",
			),
		}, &windows.Executable{
			AppID:  "PCSX2",
			AppExe: "$EMULATORS\\PCSX2\\pcsx2-qt.exe",
			Source: github.Release(
				"https://github.com/PCSX2/pcsx2",
				"pcsx2-*-windows-x64-Qt.7z",
			),
		}),
	}
}

// Installer for PPSSPP
func PPSSPP() *Program {
	return &Program{
		ID:          "ppsspp",
		Name:        "PPSSPP",
		Description: "Emulator for Sony Playstation Portable",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/PPSSPP", "$ROMS/PSP", "$BIOS/PSP"},
		Website:     "https://www.ppsspp.org",
		IconURL:     "https://cdn2.steamgriddb.com/icon/bf0552607a32015c010c166f9771efe9.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/e242660df1b69b74dcc7fde711f924ff.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/cf476046d346e8091393001a40a523dc.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/88a52c0d85339a377918fdc1ae9dc922.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/b51ecba56e03d4181e0006ff1e8a5355.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "org.ppsspp.PPSSPP",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "ppsspp",
			AppName: "$EMULATORS/PPSSPP/PPSSPP.AppImage",
			Source: github.Release(
				"https://github.com/pkgforge-dev/PPSSPP-AppImage",
				"ppsspp-*-anylinux-x86_64.AppImage",
			),
		}, &macos.Application{
			AppID:   "ppsspp",
			AppName: "$EMULATORS/PPSSPP/PPSSPPSDL.app",
			Source: github.Release(
				"https://github.com/hrydgard/ppsspp",
				"PPSSPPSDL-macOS-v1.18.1.zip",
			),
		}, &windows.Executable{
			AppID:  "PPSSPP",
			AppExe: "$EMULATORS\\PPSSPP\\PPSSPPWindows64.exe",
			Source: &packaging.Source{
				URL:    "https://www.ppsspp.org/files/1_18_1/ppsspp_win.zip",
				Format: "zip",
			},
		}),
	}
}

// Installer for Redream
func Redream() *Program {
	return &Program{
		ID:          "redream",
		Name:        "Redream",
		Description: "Emulator for Sega Dreamcast",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Redream", "$ROMS/DC", "$BIOS/DC"},
		Website:     "https://redream.io",
		IconURL:     "https://cdn2.steamgriddb.com/icon/5cc085288d7afc9d76f6aa846b7e5d5f.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/6c11cb78b7bbb5c22d5f5271b5494381.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/dd5fbbf85c3198ece6dcd86166c58439.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/2e834824cdba6141dcb14688597a26fa.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/4853deb8a06838c502fc6cb6ce91f704.png",
		Package: packaging.Best(&linux.Binary{
			AppID:  "redream",
			AppBin: "$EMULATORS/Redream/redream",
			Source: &packaging.Source{
				URL:    "https://redream.io/download/redream.x86_64-linux-v1.5.0-1133-g03c2ae9.tar.gz",
				Format: "tar.gz",
			},
		}, &macos.Application{
			AppID:   "redream",
			AppName: "$EMULATORS/Redream/Redream.app",
			Source: &packaging.Source{
				URL:    "https://redream.io/download/redream.universal-mac-v1.5.0-1133-g03c2ae9.tar.gz",
				Format: "tar.gz",
			},
		}, &windows.Executable{
			AppID:  "Redream",
			AppExe: "$EMULATORS\\Redream\\redream.exe",
			Source: &packaging.Source{
				URL:    "https://redream.io/download/redream.x86_64-windows-v1.5.0-1133-g03c2ae9.zip",
				Format: "zip",
			},
		}),
	}
}

// Installer for RPCS3
func RPCS3() *Program {
	return &Program{
		ID:          "rpcs3",
		Name:        "RPCS3",
		Description: "Emulator for Sony Playstation 3",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/RPCS3", "$ROMS/PS3", "$BIOS/PS3"},
		Website:     "https://rpcs3.net",
		IconURL:     "https://cdn2.steamgriddb.com/icon/7f40bec3df593b31feaf13dd4a696415.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/bffc98347ee35b3ead06728d6f073c68.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/ace27c5277ecc8da47cd53ff5c82cb4f.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/cddaf8b03288749c50afecad7ac3c9a4.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/15c58997f6690dddb7c501e062a2d1ab.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "net.rpcs3.RPCS3",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "rpcs3",
			AppName: "$EMULATORS/RPCS3/RPCS3.AppImage",
			Source: github.Release(
				"https://github.com/RPCS3/rpcs3-binaries-linux",
				"rpcs3-*_linux64.AppImage",
			),
		}, &macos.Application{
			AppID:   "rpcs3",
			AppName: "$EMULATORS/RPCS3/RPCS3.app",
			Source: github.Release(
				"https://github.com/RPCS3/rpcs3-binaries-mac-arm64",
				"rpcs3-*_macos_arm64.7z",
			),
		}, &windows.Executable{
			AppID:  "RPCS3",
			AppExe: "$EMULATORS\\RPCS3\\rpcs3.exe",
			Source: github.Release(
				"https://github.com/RPCS3/rpcs3-binaries-win",
				"rpcs3-*_win64.7z",
			),
		}),
	}
}

// Installer for Ryujinx
func Ryujinx() *Program {
	return &Program{
		ID:          "ryujinx",
		Name:        "Ryujinx",
		Description: "Emulator for Nintendo Switch",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Ryujinx", "$ROMS/SWITCH", "$BIOS/SWITCH"},
		Website:     "https://github.com/Ryubing/Ryujinx",
		IconURL:     "https://cdn2.steamgriddb.com/icon/446e520ce36e073a153fed039e6d55fe.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/b948aa07167c9acb17487657e96870e5.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/550d4a283baa604976e81d35d29124df.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/3931532d087eeb1b1c1a96aba6261802.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/c24f9ae141fa02c7fa1deea7e1149557.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "org.ryujinx.Ryujinx",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "ryujinx",
			AppName: "$EMULATORS/Ryujinx/Ryujinx.AppImage",
			Source: github.Release(
				"https://github.com/Ryubing/Ryujinx",
				"ryujinx-*-x64.AppImage",
			),
		}, &macos.Application{
			AppID:   "ryujinx",
			AppName: "$EMULATORS/Ryujinx/Ryujinx.app",
			Source: github.Release(
				"https://github.com/Ryubing/Ryujinx",
				"ryujinx-*-macos_universal.app.tar.gz",
			),
		}, &windows.Executable{
			AppID:  "Ryujinx",
			AppExe: "$EMULATORS\\Ryujinx\\Ryujinx.exe",
			Source: github.Release(
				"https://github.com/Ryubing/Ryujinx",
				"ryujinx-*-win_x64.zip",
			),
		}),
	}
}

// Installer for ShadPS4
func ShadPS4() *Program {
	return &Program{
		ID:          "shadps4",
		Name:        "ShadPS4",
		Description: "Emulator for Sony Playstation 4",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/ShadPS4", "$ROMS/PS4", "$BIOS/PS4"},
		Website:     "https://shadps4.net",
		IconURL:     "https://cdn2.steamgriddb.com/icon/3c7941e8f5200be6925e75ed4063311a.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/6c70dccf452364ce8e5a9c44c88dd6c1.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/72251a01ac19b84c2208c2a6f18a17da.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/21483d9d9aca5bd442f292cef7bab951.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/cc5e9cea0a79b962c20a9231e65a06ef.jpg",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "net.shadps4.shadPS4",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "shadps4",
			AppName: "$EMULATORS/ShadPS4/ShadPS4.AppImage",
			Source: github.Release(
				"https://github.com/shadps4-emu/shadPS4",
				"shadps4-linux-qt-*.zip",
			),
		}, &macos.Application{
			AppID:   "shadps4",
			AppName: "$EMULATORS/ShadPS4/ShadPS4.app",
			Source: github.Release(
				"https://github.com/shadps4-emu/shadPS4",
				"shadps4-macos-qt-*.zip",
			),
		}, &windows.Executable{
			AppID:  "ShadPS4",
			AppExe: "$EMULATORS\\ShadPS4\\shadPS4.exe",
			Source: github.Release(
				"https://github.com/shadps4-emu/shadPS4",
				"shadps4-win64-qt-*.zip",
			),
		}),
	}
}

// Installer for Simple64
func Simple64() *Program {
	return &Program{
		ID:          "simple64",
		Name:        "Simple64",
		Description: "Emulator for Nintendo 64",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Simple64", "$ROMS/N64", "$BIOS/N64"},
		Website:     "https://simple64.github.io",
		IconURL:     "https://cdn2.steamgriddb.com/icon/0ace2e260c8163925254bc878b9eb8ca.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/8f6bf2012d96ef9678f8d3a8f27ce358.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/48eeb385ea71aadccce10e2d294879b0.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/e128d1f12ec88795b0a5853d7c754608.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/73888d1bde775303c1749e63e3312a64.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "io.github.simple64.simple64",
			Overrides: []string{"--filesystem=host"},
		}, &windows.Executable{
			AppID:  "Simple64",
			AppExe: "$EMULATORS\\Simple64\\simple64-gui.exe",
			Source: github.Release(
				"https://github.com/simple64/simple64",
				"simple64-win64-*.zip",
			),
		}),
	}
}

// Installer for Vita3K
func Vita3K() *Program {
	return &Program{
		ID:          "vita3k",
		Name:        "Vita3k",
		Description: "Emulator for Sony Playstation Vita",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Vita3K", "$ROMS/PSVITA", "$BIOS/PSVITA"},
		Website:     "https://vita3k.org",
		IconURL:     "https://cdn2.steamgriddb.com/icon/39f351988d304b68b3bac5bdc5cd955e.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/654798fc20b6d08b12236106fff87920.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/d371697094a73577074c10fb6688f2ff.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/a1357c62042fedf5f0a71ebacfe5987d.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/5e98be1eed79374e1edd72f4b1d838b4.png",
		Package: packaging.Best(&linux.AppImage{
			AppID:   "vita3k",
			AppName: "$EMULATORS/Vita3K/Vita3K.AppImage",
			Source: github.Release(
				"https://github.com/Vita3K/Vita3K",
				"Vita3K-x86_64.AppImage",
			),
		}, &macos.Application{
			AppID:   "vita3k",
			AppName: "$EMULATORS/Vita3K/Vita3K.app",
			Source: github.Release(
				"https://github.com/Vita3K/Vita3K",
				"macos-latest.dmg",
			),
		}, &windows.Executable{
			AppID:  "Vita3K",
			AppExe: "$EMULATORS\\Vita3K\\Vita3K.exe",
			Source: github.Release(
				"https://github.com/Vita3K/Vita3K",
				"windows-latest.zip",
			),
		}),
	}
}

// Installer for Xemu
func Xemu() *Program {
	return &Program{
		ID:          "xemu",
		Name:        "Xemu",
		Description: "Emulator for Microsoft Xbox",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Xemu", "$ROMS/XBOX", "$BIOS/XBOX"},
		Website:     "https://xemu.app",
		IconURL:     "https://cdn2.steamgriddb.com/icon/53fa398e3a888d8f115b72a55aa8c7de.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/a42b7cddd7ebb7c1bced17bddc568d2f.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/b6cd95d53810282d6a734fbb073e9479.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/5b74752b25bd07933b10b2098970f990.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/aa0994c4263018600494efceae69087a.png",
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "app.xemu.xemu",
			Overrides: []string{"--filesystem=host"},
		}, &linux.AppImage{
			AppID:   "xemu",
			AppName: "$EMULATORS/Xemu/Xemu.AppImage",
			Source: github.Release(
				"https://github.com/xemu-project/xemu",
				"xemu-*-x86_64.AppImage",
			),
		}, &macos.Application{
			AppID:   "xemu",
			AppName: "$EMULATORS/Xemu/Xemu.app",
			Source: github.Release(
				"https://github.com/xemu-project/xemu",
				"xemu-macos-universal-release.zip",
			),
		}, &windows.Executable{
			AppID:  "Xemu",
			AppExe: "$EMULATORS\\Xemu\\xemu.exe",
			Source: github.Release(
				"https://github.com/xemu-project/xemu",
				"xemu-win-release.zip",
			),
		}),
	}
}

// Installer for Xenia
func Xenia() *Program {
	return &Program{
		ID:          "xenia",
		Name:        "Xenia",
		Description: "Emulator for Microsoft Xbox 360",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Folders:     []string{"$EMULATORS", "$STATE/Xenia", "$ROMS/X360", "$BIOS/X360"},
		Website:     "https://xenia.jp",
		IconURL:     "https://cdn2.steamgriddb.com/icon/9775efcc70ff0918ad952cc9c48a511a.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/fac05328668f599efe18e76cdb284aab.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/e43e55468f8cfee48d517b2c49cecd08.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/1962bcb00dc1bf1b5bcb334257ff3701.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/2958ef004a18f50b380a87d1cfe5366d.png",
		Package: packaging.Best(&linux.Binary{
			AppID:  "xenia",
			AppBin: "$EMULATORS/Xenia/xenia_canary",
			Source: github.Release(
				"https://github.com/xenia-canary/xenia-canary-releases",
				"xenia_canary_linux.tar.gz",
			),
		}, &windows.Executable{
			AppID:  "Xenia",
			AppExe: "$EMULATORS\\Xenia\\xenia_canary.exe",
			Source: github.Release(
				"https://github.com/xenia-canary/xenia-canary-releases",
				"xenia_canary_windows.zip",
			),
		}),
	}
}
