package platforms

// Emulator struct
type Emulator struct {
	Name          string `json:"name"`
	Program       string `json:"program"`
	Extensions    string `json:"extensions"`
	LaunchOptions string `json:"launchOptions"`
}

// Platform struct
type Platform struct {
	Name      string      `json:"name"`
	Console   string      `json:"console"`
	Folder    string      `json:"folder"`
	Emulators []*Emulator `json:"emulators"`
}

// Retrieve system platform specs.
// This list is almost a copy of EmulationStation DE systems
func GetPlatforms(options *Options) []*Platform {

	platforms := []*Platform{}

	platforms = append(platforms, &Platform{
		Name:    "3DS",
		Console: "Nintendo 3DS",
		Folder:  "3DS/",
		Emulators: []*Emulator{{
			Name:          "Lime3DS",
			Program:       "/var/lib/flatpak/exports/bin/io.github.lime3ds.Lime3DS",
			Extensions:    ".3ds .3dsx .app .axf .cci .cxi .elf .7z .zip",
			LaunchOptions: "\"${ROM}\"", // No full-screen option
		}, {
			Name:          "Citra",
			Program:       "/var/lib/flatpak/exports/bin/org.citra_emu.citra",
			Extensions:    ".3ds .3dsx .app .axf .cci .cxi .elf .7z .zip",
			LaunchOptions: "\"${ROM}\"", // No full-screen option
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "DC",
		Console: "Sega Dreamcast",
		Folder:  "DC/",
		Emulators: []*Emulator{{
			Name:          "Flycast",
			Program:       "/var/lib/flatpak/exports/bin/org.flycast.Flycast",
			Extensions:    ".chd .cdi .iso .elf .cue .gdi .lst .dat .m3u .7z .zip",
			LaunchOptions: "-config window:fullscreen=yes \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "GBA",
		Console: "Nintendo Game Boy Advance",
		Folder:  "GBA/",
		Emulators: []*Emulator{{
			Name:          "MGBA",
			Program:       "/var/lib/flatpak/exports/bin/io.mgba.mGBA",
			Extensions:    ".agb .bin .cgb .dmg .gb .gba .gbc .sgb .7z .zip",
			LaunchOptions: "-f \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "GC",
		Console: "Nintendo GameCube",
		Folder:  "GC/",
		Emulators: []*Emulator{{
			Name:          "Dolphin Emulator",
			Program:       "/var/lib/flatpak/exports/bin/org.DolphinEmu.dolphin-emu",
			Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
			LaunchOptions: "-b -e \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "N64",
		Console: "Nintendo 64",
		Folder:  "N64/",
		Emulators: []*Emulator{{
			Name:          "Simple64",
			Program:       "/var/lib/flatpak/exports/bin/io.github.simple64.simple64",
			Extensions:    ".bin .d64 .n64 .ndd .u1 .v64 .z64 .7z .zip",
			LaunchOptions: "--nogui \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "NDS",
		Console: "Nintendo DS",
		Folder:  "NDS/",
		Emulators: []*Emulator{{
			Name:          "MelonDS",
			Program:       "/var/lib/flatpak/exports/bin/net.kuribo64.melonDS",
			Extensions:    ".app .bin .nds .7z .zip",
			LaunchOptions: "-f \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS1",
		Console: "Sony PlayStation 1",
		Folder:  "PS1/",
		Emulators: []*Emulator{{
			Name:          "DuckStation",
			Program:       "/var/lib/flatpak/exports/bin/org.duckstation.DuckStation",
			Extensions:    ".bin .cbn .ccd .chd .cue .ecm .exe .img .iso .m3u .mdf .mds .minipsf .pbp .psexe .psf .toc .z .znx .7z .zip",
			LaunchOptions: "-batch -fullscreen \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS2",
		Console: "Sony PlayStation 2",
		Folder:  "PS2/",
		Emulators: []*Emulator{{
			Name:          "PCSX2",
			Program:       "/var/lib/flatpak/exports/bin/net.pcsx2.PCSX2",
			Extensions:    ".bin .chd .ciso .cso .dump .elf .gz .m3u .mdf .img .iso .isz .ngr",
			LaunchOptions: "-batch -nogui -fullscreen \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS3",
		Console: "Sony PlayStation 3",
		Folder:  "PS3/",
		Emulators: []*Emulator{{
			Name:          "RPCS3",
			Program:       "/var/lib/flatpak/exports/bin/net.rpcs3.RPCS3",
			Extensions:    ".desktop .ps3 .ps3dir",
			LaunchOptions: "--no-gui \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PSP",
		Console: "Sony PlayStation Portable",
		Folder:  "PSP/",
		Emulators: []*Emulator{{
			Name:          "PPSSPP",
			Program:       "/var/lib/flatpak/exports/bin/org.ppsspp.PPSSPP",
			Extensions:    ".elf .iso .cso .prx .pbp .7z .zip",
			LaunchOptions: "-f -g \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "SWITCH",
		Console: "Nintendo Switch",
		Folder:  "SWITCH/",
		Emulators: []*Emulator{{
			Name:          "Ryujinx",
			Program:       "/var/lib/flatpak/exports/bin/org.ryujinx.Ryujinx",
			Extensions:    "nca .nro .nso .nsp .xci",
			LaunchOptions: "--fullscreen \"${ROM}\"",
		}, {
			Name:          "Yuzu",
			Program:       "/var/lib/flatpak/exports/bin/org.yuzu_emu.yuzu",
			Extensions:    "nca .nro .nso .nsp .xci",
			LaunchOptions: "-f -g \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "WII",
		Console: "Nintendo Wii",
		Folder:  "WII/",
		Emulators: []*Emulator{{
			Name:          "Dolphin",
			Program:       "/var/lib/flatpak/exports/bin/org.DolphinEmu.dolphin-emu",
			Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
			LaunchOptions: "-b -e \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "WIIU",
		Console: "Nintendo Wii U",
		Folder:  "WIIU/",
		Emulators: []*Emulator{{
			Name:          "Cemu",
			Program:       "/var/lib/flatpak/exports/bin/info.cemu.Cemu",
			Extensions:    ".rpx .wua .wud .wux",
			LaunchOptions: "-f -g \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "XBOX",
		Console: "Microsoft Xbox",
		Folder:  "XBOX/",
		Emulators: []*Emulator{{
			Name:          "Xemu",
			Program:       "/var/lib/flatpak/exports/bin/app.xemu.xemu",
			Extensions:    ".iso",
			LaunchOptions: "-full-screen -dvd_path \"${ROM}\"",
		}},
	})

	return platforms
}
