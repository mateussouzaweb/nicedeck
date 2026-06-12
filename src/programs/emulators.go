package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/programs/assets"
	"github.com/mateussouzaweb/nicedeck/src/programs/forgejo"
	"github.com/mateussouzaweb/nicedeck/src/programs/github"
	"github.com/mateussouzaweb/nicedeck/src/programs/gitlab"
	"github.com/mateussouzaweb/nicedeck/src/programs/website"
)

// Installer for Azahar
func Azahar() *packaging.Program {
	return &packaging.Program{
		ID:          "azahar",
		Name:        "Azahar",
		Description: "Emulator for Nintendo 3DS",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Azahar", "$ROMS/3DS", "$BIOS/3DS"},
		Website:     "https://azahar-emu.org",
		IconURL:     assets.Icon("5f35c349bea2a27d9759fac65580a098.png"),
		LogoURL:     assets.Logo("549c5a7673e36671a74ffba405036141.png"),
		CoverURL:    assets.Cover("4ff70a1c13ee2cf27853c7ae06425bc4.png"),
		BannerURL:   assets.Banner("1f4181a701cdfb56675cc7c7f766d60d.png"),
		HeroURL:     assets.Hero("bc6f714aa3dfeef9320a838b79515c2d.png"),
		Package: packaging.Best(&linux.AppImage{
			AppID:     "azahar",
			Launcher:  "$EMULATORS/Azahar/azahar.AppImage",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/azahar-emu/azahar",
				cli.ArchVariant(
					"azahar.AppImage", // amd64
					"azahar.AppImage", // arm64 (WIP)
				),
			),
		}, &macos.Application{
			AppID:     "azahar",
			Launcher:  "$EMULATORS/Azahar/Azahar.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/azahar-emu/azahar",
				cli.ArchVariant(
					"azahar-*-macos-x86_64.zip", // amd64
					"azahar-*-macos-arm64.zip",  // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "Azahar",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Azahar/azahar.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/azahar-emu/azahar",
				cli.ArchVariant(
					"azahar-*-windows-msvc.zip", // amd64
					"azahar-*-windows-msvc.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for Cemu
func Cemu() *packaging.Program {
	return &packaging.Program{
		ID:          "cemu",
		Name:        "Cemu",
		Description: "Emulator for Nintendo Wii U",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Cemu", "$ROMS/WIIU", "$BIOS/WIIU"},
		Website:     "https://cemu.info",
		IconURL:     assets.Icon("2c790f933dcb0c7a747741780c6b435d.png"),
		LogoURL:     assets.Logo("c7a9f13a6c0940277d46706c7ca32601.png"),
		CoverURL:    assets.Cover("9454c84816d82ed1092f2fe2919a3a8e.png"),
		BannerURL:   assets.Banner("86fb4d9e1de18ebdb6fc534de828d605.png"),
		HeroURL:     assets.Hero("d5da28d4865fb92720359db84e0dd0dd.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "info.cemu.Cemu",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "cemu",
			Launcher:  "$EMULATORS/Cemu/Cemu.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/cemu-project/Cemu",
				cli.ArchVariant(
					"cemu-*-macos-*-x64.dmg", // amd64
					"cemu-*-macos-*-x64.dmg", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "Cemu",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Cemu/Cemu.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/cemu-project/Cemu",
				cli.ArchVariant(
					"cemu-*-windows-x64.zip", // amd64
					"cemu-*-windows-x64.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for Dolphin
func Dolphin() *packaging.Program {
	return &packaging.Program{
		ID:          "dolphin",
		Name:        "Dolphin Emulator",
		Description: "Emulator for Nintendo GameCube and Nintendo Wii",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Dolphin", "$ROMS/GC", "$BIOS/GC", "$ROMS/WII", "$BIOS/WII"},
		Website:     "https://dolphin-emu.org",
		IconURL:     assets.Icon("52ec3baefcf93a558d994e1bcd3b5c3d.png"),
		LogoURL:     assets.Logo("5b5bbd3170c560829391c3db7265ee9b.png"),
		CoverURL:    assets.Cover("8a07e4382e18e3b9f5d2713aeaefc29b.png"),
		BannerURL:   assets.Banner("cbec7ddbb30e261abd365bf9f814647d.png"),
		HeroURL:     assets.Hero("c24f9ae141fa02c7fa1deea7e1149557.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "org.DolphinEmu.dolphin-emu",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "dolphin",
			Launcher:  "$EMULATORS/Dolphin/Dolphin.app",
			Arguments: packaging.NoArguments(),
			Source: website.Release(
				"https://dolphin-emu.org/download/", "",
				cli.ArchVariant(
					"https://dl.dolphin-emu.org/releases/*/dolphin-*-universal.dmg", // amd64
					"https://dl.dolphin-emu.org/releases/*/dolphin-*-universal.dmg", // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "Dolphin",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Dolphin/Dolphin.exe",
			Arguments:   packaging.NoArguments(),
			Source: website.Release(
				"https://dolphin-emu.org/download/", "",
				cli.ArchVariant(
					"https://dl.dolphin-emu.org/releases/*/dolphin-*-x64.7z",   // amd64
					"https://dl.dolphin-emu.org/releases/*/dolphin-*-ARM64.7z", // arm64
				),
			),
		}),
	}
}

// Installer for DuckStation
func DuckStation() *packaging.Program {
	return &packaging.Program{
		ID:          "duckstation",
		Name:        "DuckStation",
		Description: "Emulator for Sony Playtation 1",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/DuckStation", "$ROMS/PS1", "$BIOS/PS1"},
		Website:     "https://www.duckstation.org",
		IconURL:     assets.Icon("f985f43b4ba330d5282dfd9be8003e62.png"),
		LogoURL:     assets.Logo("96a0d70498272acfee21d3dbae846113.png"),
		CoverURL:    assets.Cover("1f7c9b9e37afcbd79ebff19b17837cad.png"),
		BannerURL:   assets.Banner("9c94e659c62b84bf7b39c599b61bc7d3.png"),
		HeroURL:     assets.Hero("127f12c937b4baf0a8922eb1384391cf.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "org.duckstation.DuckStation",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "duckstation",
			Launcher:  "$EMULATORS/DuckStation/DuckStation.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/stenzek/duckstation",
				cli.ArchVariant(
					"duckstation-mac-release.zip", // amd64
					"duckstation-mac-release.zip", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "DuckStation",
			Installer:   "",
			Uninstaller: "",
			Launcher: cli.ArchVariant(
				"$EMULATORS/DuckStation/duckstation-qt-x64-ReleaseLTCG.exe",   // amd64
				"$EMULATORS/DuckStation/duckstation-qt-ARM64-ReleaseLTCG.exe", // arm64
			),
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/stenzek/duckstation",
				cli.ArchVariant(
					"duckstation-windows-x64-release.zip",   // amd64
					"duckstation-windows-arm64-release.zip", // arm64
				),
			),
		}),
	}
}

// Installer for Eden
func Eden() *packaging.Program {
	return &packaging.Program{
		ID:          "eden",
		Name:        "Eden Emulator",
		Description: "Emulator for Nintendo Switch",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Eden", "$ROMS/SWITCH", "$BIOS/SWITCH"},
		Website:     "https://eden-emu.dev",
		IconURL:     assets.Icon("2080d47c88a0816da0a7e58e0cd7ad50.png"),
		LogoURL:     assets.Logo("6c0edd40bc18715488dd4ed8abb93a60.png"),
		CoverURL:    assets.Cover("174b4233c093b0bf83e7c6fca65fae2a.png"),
		BannerURL:   assets.Banner("f0ba96a506d7109bd0ec7c26bc957911.png"),
		HeroURL:     assets.Hero("a960ee65d36125cfe5f126bd326ff75b.png"),
		Package: packaging.Best(&linux.AppImage{
			AppID:     "eden",
			Launcher:  "$EMULATORS/Eden/Eden.AppImage",
			Arguments: packaging.NoArguments(),
			Source: forgejo.Release(
				"https://git.eden-emu.dev",
				"eden-emu/eden",
				cli.ArchVariant(
					"Eden-Linux-*-amd64-gcc-standard.AppImage",   // amd64
					"Eden-Linux-*-aarch64-gcc-standard.AppImage", // arm64
				),
			),
		}, &macos.Application{
			AppID:     "eden",
			Launcher:  "$EMULATORS/Eden/Eden.app",
			Arguments: packaging.NoArguments(),
			Source: forgejo.Release(
				"https://git.eden-emu.dev",
				"eden-emu/eden",
				cli.ArchVariant(
					"Eden-macOS-*.tar.gz", // amd64 (WIP)
					"Eden-macOS-*.tar.gz", // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "Eden",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Eden/eden.exe",
			Arguments:   packaging.NoArguments(),
			Source: forgejo.Release(
				"https://git.eden-emu.dev",
				"eden-emu/eden",
				cli.ArchVariant(
					"Eden-Windows-*-amd64-msvc-standard.zip", // amd64
					"Eden-Windows-*-arm64-msvc-standard.zip", // arm64
				),
			),
		}),
	}
}

// Installer for Flycast
func Flycast() *packaging.Program {
	return &packaging.Program{
		ID:          "flycast",
		Name:        "Flycast",
		Description: "Emulator for Sega Dreamcast",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Flycast", "$ROMS/DC", "$BIOS/DC"},
		Website:     "https://github.com/flyinghead/flycast",
		IconURL:     assets.Icon("858a2be748405d1cf063e97622abc791.png"),
		LogoURL:     assets.Logo("b9b0c8b6beb69bd0c5a213b9422459ce.png"),
		CoverURL:    assets.Cover("51cf6e65f8242f989f354bf9dfe5a019.png"),
		BannerURL:   assets.Banner("46b3feb0521b4d823847ebbd4dd58ea6.png"),
		HeroURL:     assets.Hero("c24f9ae141fa02c7fa1deea7e1149557.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "org.flycast.Flycast",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "flycast",
			Launcher:  "$EMULATORS/Flycast/Flycast.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/flyinghead/flycast",
				cli.ArchVariant(
					"flycast-macOS-*.zip", // amd64
					"flycast-macOS-*.zip", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "Flycast",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Flycast/flycast.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/flyinghead/flycast",
				cli.ArchVariant(
					"flycast-win64-*.zip", // amd64
					"flycast-win64-*.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for MelonDS
func MelonDS() *packaging.Program {
	return &packaging.Program{
		ID:          "melonds",
		Name:        "MelonDS",
		Description: "Emulator for Nintendo DS",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/MelonDS", "$ROMS/NDS", "$BIOS/NDS"},
		Website:     "https://melonds.kuribo64.net",
		IconURL:     assets.Icon("8ad297da4fc2cc28dfa3c0cb7df8ae63.png"),
		LogoURL:     assets.Logo("173f798d1316395cce2c8ecf98aed4d5.png"),
		CoverURL:    assets.Cover("3b397c602f7c9226cbcb907b3d5e7d5e.png"),
		BannerURL:   assets.Banner("0ec19bac435cd0ab3fcd2160491b0c7b.png"),
		HeroURL:     assets.Hero("c24f9ae141fa02c7fa1deea7e1149557.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "net.kuribo64.melonDS",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "melonds",
			Launcher:  "$EMULATORS/MelonDS/melonDS.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/melonDS-emu/melonDS",
				cli.ArchVariant(
					"macOS-universal.zip", // amd64
					"macOS-universal.zip", // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "MelonDS",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/MelonDS/melonDS.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/melonDS-emu/melonDS",
				cli.ArchVariant(
					"melonDS-windows-x86_64.zip", // amd64
					"melonDS-windows-x86_64.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for mGBA
func MGBA() *packaging.Program {
	return &packaging.Program{
		ID:          "mgba",
		Name:        "MGBA",
		Description: "Emulator for Nintendo Game Boy Advance",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/MGBA", "$ROMS/GBA", "$BIOS/GBA"},
		Website:     "https://mgba.io",
		IconURL:     assets.Icon("7d5fe9a6097c89cda3ce6a5b1de4dd15.png"),
		LogoURL:     assets.Logo("e262b1f197f1a9cca59e0868f1e5c94b.png"),
		CoverURL:    assets.Cover("d280a227a8ef77d87a5d18037c52776a.png"),
		BannerURL:   assets.Banner("7088b9d5b6a444224cf6380dcfe61554.png"),
		HeroURL:     assets.Hero("d470133ccf31f9bfdc1dcb45a30c73b1.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "io.mgba.mGBA",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "mgba",
			Launcher:  "$EMULATORS/MGBA/mGBA.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/mgba-emu/mgba",
				cli.ArchVariant(
					"mGBA-*-macos.dmg", // amd64
					"mGBA-*-macos.dmg", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "MGBA",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/MGBA/mGBA.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/mgba-emu/mgba",
				cli.ArchVariant(
					"mGBA-*-win64.7z", // amd64
					"mGBA-*-win64.7z", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for PCSX2
func PCSX2() *packaging.Program {
	return &packaging.Program{
		ID:          "pcsx2",
		Name:        "PCSX2",
		Description: "Emulator for Sony Playstation 2",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/PCSX2", "$ROMS/PS2", "$BIOS/PS2"},
		Website:     "https://pcsx2.net",
		IconURL:     assets.Icon("22ec872ee633043cc5aece5adb261367.png"),
		LogoURL:     assets.Logo("7123c9e46f34491cf4f8eb1a813d8f6e.png"),
		CoverURL:    assets.Cover("3123b87d2cede1c04e380a71701ddfe8.png"),
		BannerURL:   assets.Banner("f3a71cf60765edd14269d28819d15327.png"),
		HeroURL:     assets.Hero("9cc25407f209e031babdac7d3c520ccb.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "net.pcsx2.PCSX2",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "pcsx2",
			Launcher:  "$EMULATORS/PCSX2/PCSX2.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/PCSX2/pcsx2",
				cli.ArchVariant(
					"pcsx2-*-macos-Qt.tar.xz", // amd64
					"pcsx2-*-macos-Qt.tar.xz", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "PCSX2",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/PCSX2/pcsx2-qt.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/PCSX2/pcsx2",
				cli.ArchVariant(
					"pcsx2-*-windows-x64-Qt.7z", // amd64
					"pcsx2-*-windows-x64-Qt.7z", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for PPSSPP
func PPSSPP() *packaging.Program {
	return &packaging.Program{
		ID:          "ppsspp",
		Name:        "PPSSPP",
		Description: "Emulator for Sony Playstation Portable",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/PPSSPP", "$ROMS/PSP", "$BIOS/PSP"},
		Website:     "https://www.ppsspp.org",
		IconURL:     assets.Icon("bf0552607a32015c010c166f9771efe9.png"),
		LogoURL:     assets.Logo("e242660df1b69b74dcc7fde711f924ff.png"),
		CoverURL:    assets.Cover("cf476046d346e8091393001a40a523dc.png"),
		BannerURL:   assets.Banner("88a52c0d85339a377918fdc1ae9dc922.png"),
		HeroURL:     assets.Hero("b51ecba56e03d4181e0006ff1e8a5355.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "org.ppsspp.PPSSPP",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "ppsspp",
			Launcher:  "$EMULATORS/PPSSPP/PPSSPPSDL.app",
			Arguments: packaging.NoArguments(),
			Source: website.Release(
				"https://www.ppsspp.org/download/", "",
				cli.ArchVariant(
					"https://www.ppsspp.org/files/*/PPSSPP_macOS.dmg", // amd64
					"https://www.ppsspp.org/files/*/PPSSPP_macOS.dmg", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "PPSSPP",
			Installer:   "",
			Uninstaller: "",
			Launcher: cli.ArchVariant(
				"$EMULATORS/PPSSPP/PPSSPPWindows64.exe",    // amd64
				"$EMULATORS/PPSSPP/PPSSPPWindowsARM64.exe", // arm64
			),
			Arguments: packaging.NoArguments(),
			Source: website.Release(
				"https://www.ppsspp.org/download/", "",
				cli.ArchVariant(
					"https://www.ppsspp.org/files/*/ppsspp_win.zip",         // amd64
					"https://www.ppsspp.org/files/*/PPSSPPWindowsARM64.zip", // arm64
				),
			),
		}),
	}
}

// Installer for Redream
func Redream() *packaging.Program {
	return &packaging.Program{
		ID:          "redream",
		Name:        "Redream",
		Description: "Emulator for Sega Dreamcast",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Redream", "$ROMS/DC", "$BIOS/DC"},
		Website:     "https://redream.io",
		IconURL:     assets.Icon("5cc085288d7afc9d76f6aa846b7e5d5f.png"),
		LogoURL:     assets.Logo("6c11cb78b7bbb5c22d5f5271b5494381.png"),
		CoverURL:    assets.Cover("dd5fbbf85c3198ece6dcd86166c58439.png"),
		BannerURL:   assets.Banner("2e834824cdba6141dcb14688597a26fa.png"),
		HeroURL:     assets.Hero("4853deb8a06838c502fc6cb6ce91f704.png"),
		Package: packaging.Best(&linux.Binary{
			AppID:     "redream",
			Launcher:  "$EMULATORS/Redream/redream",
			Arguments: packaging.NoArguments(),
			Source: website.Release(
				"https://redream.io/download",
				"https://redream.io/",
				cli.ArchVariant(
					"download/redream.x86_64-linux-v*-*-*.tar.gz", // amd64
					"download/redream.x86_64-linux-v*-*-*.tar.gz", // arm64 (WIP)
				),
			),
		}, &macos.Application{
			AppID:     "redream",
			Launcher:  "$EMULATORS/Redream/redream.app",
			Arguments: packaging.NoArguments(),
			Source: website.Release(
				"https://redream.io/download",
				"https://redream.io/",
				cli.ArchVariant(
					"download/redream.universal-mac-v*-*-*.tar.gz", // amd64
					"download/redream.universal-mac-v*-*-*.tar.gz", // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "Redream",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Redream/redream.exe",
			Arguments:   packaging.NoArguments(),
			Source: website.Release(
				"https://redream.io/download",
				"https://redream.io/",
				cli.ArchVariant(
					"download/redream.x86_64-windows-v*-*-*.zip", // amd64
					"download/redream.x86_64-windows-v*-*-*.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for RPCS3
func RPCS3() *packaging.Program {
	return &packaging.Program{
		ID:          "rpcs3",
		Name:        "RPCS3",
		Description: "Emulator for Sony Playstation 3",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/RPCS3", "$ROMS/PS3", "$BIOS/PS3"},
		Website:     "https://rpcs3.net",
		IconURL:     assets.Icon("7f40bec3df593b31feaf13dd4a696415.png"),
		LogoURL:     assets.Logo("bffc98347ee35b3ead06728d6f073c68.png"),
		CoverURL:    assets.Cover("ace27c5277ecc8da47cd53ff5c82cb4f.png"),
		BannerURL:   assets.Banner("cddaf8b03288749c50afecad7ac3c9a4.png"),
		HeroURL:     assets.Hero("15c58997f6690dddb7c501e062a2d1ab.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "net.rpcs3.RPCS3",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "rpcs3",
			Launcher:  "$EMULATORS/RPCS3/RPCS3.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				cli.ArchVariant(
					"https://github.com/RPCS3/rpcs3-binaries-mac",       // amd64
					"https://github.com/RPCS3/rpcs3-binaries-mac-arm64", // arm64
				),
				cli.ArchVariant(
					"rpcs3-*_macos.7z",       // amd64
					"rpcs3-*_macos_arm64.7z", // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "RPCS3",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/RPCS3/rpcs3.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/RPCS3/rpcs3-binaries-win",
				cli.ArchVariant(
					"rpcs3-*_win64.7z", // amd64
					"rpcs3-*_win64.7z", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for Ryujinx
func Ryujinx() *packaging.Program {
	return &packaging.Program{
		ID:          "ryujinx",
		Name:        "Ryujinx",
		Description: "Emulator for Nintendo Switch",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Ryujinx", "$ROMS/SWITCH", "$BIOS/SWITCH"},
		Website:     "https://ryujinx.app",
		IconURL:     assets.Icon("36e8ce541ad3b87a4c89c0c1d76a7e08.png"),
		LogoURL:     assets.Logo("fc6f229f34410b0ae0af5cb7df22e470.png"),
		CoverURL:    assets.Cover("1633727e16b29e084edf3da658e392d0.png"),
		BannerURL:   assets.Banner("9024d61574fcc58378aedbad631674f9.png"),
		HeroURL:     assets.Hero("573185e7a57bcdcd68d7895cf83ffe66.png"),
		Package: packaging.Best(&linux.AppImage{
			AppID:     "ryujinx",
			Launcher:  "$EMULATORS/Ryujinx/Ryujinx.AppImage",
			Arguments: packaging.NoArguments(),
			Source: gitlab.Release(
				"https://git.ryujinx.app", "1",
				cli.ArchVariant(
					"ryujinx-*-x64.AppImage",   // amd64
					"ryujinx-*-arm64.AppImage", // arm64
				),
			),
		}, &macos.Application{
			AppID:     "ryujinx",
			Launcher:  "$EMULATORS/Ryujinx/Ryujinx.app",
			Arguments: packaging.NoArguments(),
			Source: gitlab.Release(
				"https://git.ryujinx.app", "1",
				cli.ArchVariant(
					"ryujinx-*-macos_universal.app.tar.gz", // amd64
					"ryujinx-*-macos_universal.app.tar.gz", // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "Ryujinx",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Ryujinx/Ryujinx.exe",
			Arguments:   packaging.NoArguments(),
			Source: gitlab.Release(
				"https://git.ryujinx.app", "1",
				cli.ArchVariant(
					"ryujinx-*-win_x64.zip", // amd64
					"ryujinx-*-win_x64.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for ShadPS4
func ShadPS4() *packaging.Program {
	return &packaging.Program{
		ID:          "shadps4",
		Name:        "ShadPS4",
		Description: "Emulator for Sony Playstation 4",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/ShadPS4", "$ROMS/PS4", "$BIOS/PS4"},
		Website:     "https://shadps4.net",
		IconURL:     assets.Icon("3c7941e8f5200be6925e75ed4063311a.png"),
		LogoURL:     assets.Logo("6c70dccf452364ce8e5a9c44c88dd6c1.png"),
		CoverURL:    assets.Cover("72251a01ac19b84c2208c2a6f18a17da.png"),
		BannerURL:   assets.Banner("21483d9d9aca5bd442f292cef7bab951.png"),
		HeroURL:     assets.Hero("cc5e9cea0a79b962c20a9231e65a06ef.jpg"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "net.shadps4.shadPS4",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "shadps4",
			Launcher:  "$EMULATORS/ShadPS4/shadPS4QtLauncher.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/shadps4-emu/shadps4-qtlauncher",
				cli.ArchVariant(
					"shadPS4QtLauncher-macos-qt-*.zip", // amd64
					"shadPS4QtLauncher-macos-qt-*.zip", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "ShadPS4",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/ShadPS4/shadPS4QtLauncher.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/shadps4-emu/shadps4-qtlauncher",
				cli.ArchVariant(
					"shadPS4QtLauncher-win64-qt-*.zip", // amd64
					"shadPS4QtLauncher-win64-qt-*.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for Simple64
func Simple64() *packaging.Program {
	return &packaging.Program{
		ID:          "simple64",
		Name:        "Simple64",
		Description: "Emulator for Nintendo 64",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Simple64", "$ROMS/N64", "$BIOS/N64"},
		Website:     "https://github.com/simple64/simple64",
		IconURL:     assets.Icon("0ace2e260c8163925254bc878b9eb8ca.png"),
		LogoURL:     assets.Logo("8f6bf2012d96ef9678f8d3a8f27ce358.png"),
		CoverURL:    assets.Cover("48eeb385ea71aadccce10e2d294879b0.png"),
		BannerURL:   assets.Banner("e128d1f12ec88795b0a5853d7c754608.png"),
		HeroURL:     assets.Hero("73888d1bde775303c1749e63e3312a64.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "io.github.simple64.simple64",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &windows.Executable{
			AppID:       "Simple64",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Simple64/simple64-gui.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/simple64/simple64",
				cli.ArchVariant(
					"simple64-win64-*.zip", // amd64
					"simple64-win64-*.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for Vita3K
func Vita3K() *packaging.Program {
	return &packaging.Program{
		ID:          "vita3k",
		Name:        "Vita3k",
		Description: "Emulator for Sony Playstation Vita",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Vita3K", "$ROMS/PSVITA", "$BIOS/PSVITA"},
		Website:     "https://vita3k.org",
		IconURL:     assets.Icon("39f351988d304b68b3bac5bdc5cd955e.png"),
		LogoURL:     assets.Logo("654798fc20b6d08b12236106fff87920.png"),
		CoverURL:    assets.Cover("d371697094a73577074c10fb6688f2ff.png"),
		BannerURL:   assets.Banner("a1357c62042fedf5f0a71ebacfe5987d.png"),
		HeroURL:     assets.Hero("5e98be1eed79374e1edd72f4b1d838b4.png"),
		Package: packaging.Best(&linux.AppImage{
			AppID:     "vita3k",
			Launcher:  "$EMULATORS/Vita3K/Vita3K.AppImage",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/Vita3K/Vita3K",
				cli.ArchVariant(
					"Vita3K-x86_64.AppImage", // amd64
					"Vita3K-x86_64.AppImage", // arm64 (WIP)
				),
			),
		}, &macos.Application{
			AppID:     "vita3k",
			Launcher:  "$EMULATORS/Vita3K/Vita3K.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/Vita3K/Vita3K",
				cli.ArchVariant(
					"macos-latest.dmg", // amd64
					"macos-latest.dmg", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "Vita3K",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Vita3K/Vita3K.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/Vita3K/Vita3K",
				cli.ArchVariant(
					"windows-latest.zip", // amd64
					"windows-latest.zip", // arm64 (WIP)
				),
			),
		}),
	}
}

// Installer for Xemu
func Xemu() *packaging.Program {
	return &packaging.Program{
		ID:          "xemu",
		Name:        "Xemu",
		Description: "Emulator for Microsoft Xbox",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Xemu", "$ROMS/XBOX", "$BIOS/XBOX"},
		Website:     "https://xemu.app",
		IconURL:     assets.Icon("53fa398e3a888d8f115b72a55aa8c7de.png"),
		LogoURL:     assets.Logo("a42b7cddd7ebb7c1bced17bddc568d2f.png"),
		CoverURL:    assets.Cover("b6cd95d53810282d6a734fbb073e9479.png"),
		BannerURL:   assets.Banner("5b74752b25bd07933b10b2098970f990.png"),
		HeroURL:     assets.Hero("aa0994c4263018600494efceae69087a.png"),
		Package: packaging.Best(&linux.Flatpak{
			AppID:     "app.xemu.xemu",
			Namespace: "system",
			Overrides: []string{"--filesystem=host"},
			Arguments: packaging.NoArguments(),
		}, &macos.Application{
			AppID:     "xemu",
			Launcher:  "$EMULATORS/Xemu/xemu.app",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/xemu-project/xemu",
				cli.ArchVariant(
					"xemu-macos-universal-release.zip", // amd64
					"xemu-macos-universal-release.zip", // arm64
				),
			),
		}, &windows.Executable{
			AppID:       "Xemu",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Xemu/xemu.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/xemu-project/xemu",
				cli.ArchVariant(
					"xemu-win-x86_64-release.zip",  // amd64
					"xemu-win-aarch64-release.zip", // arm64
				),
			),
		}),
	}
}

// Installer for Xenia
func Xenia() *packaging.Program {
	return &packaging.Program{
		ID:          "xenia",
		Name:        "Xenia",
		Description: "Emulator for Microsoft Xbox 360",
		Category:    "Emulators",
		Tags:        []string{"Gaming", "Emulator"},
		Flags:       []string{},
		Folders:     []string{"$EMULATORS", "$STATE/Xenia", "$ROMS/X360", "$BIOS/X360"},
		Website:     "https://xenia.jp",
		IconURL:     assets.Icon("9775efcc70ff0918ad952cc9c48a511a.png"),
		LogoURL:     assets.Logo("fac05328668f599efe18e76cdb284aab.png"),
		CoverURL:    assets.Cover("e43e55468f8cfee48d517b2c49cecd08.png"),
		BannerURL:   assets.Banner("1962bcb00dc1bf1b5bcb334257ff3701.png"),
		HeroURL:     assets.Hero("2958ef004a18f50b380a87d1cfe5366d.png"),
		Package: packaging.Best(&linux.Binary{
			AppID:     "xenia",
			Launcher:  "$EMULATORS/Xenia/xenia_canary",
			Arguments: packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/xenia-canary/xenia-canary-releases",
				cli.ArchVariant(
					"xenia_canary_linux.tar.xz", // amd64
					"xenia_canary_linux.tar.xz", // arm64 (WIP)
				),
			),
		}, &windows.Executable{
			AppID:       "Xenia",
			Installer:   "",
			Uninstaller: "",
			Launcher:    "$EMULATORS/Xenia/xenia_canary.exe",
			Arguments:   packaging.NoArguments(),
			Source: github.Release(
				"https://github.com/xenia-canary/xenia-canary-releases",
				cli.ArchVariant(
					"xenia_canary_windows.zip", // amd64
					"xenia_canary_windows.zip", // arm64 (WIP)
				),
			),
		}),
	}
}
