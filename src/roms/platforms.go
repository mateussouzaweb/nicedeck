package roms

import "slices"

// Emulator struct
type Emulator struct {
	Name          string `json:"name"`
	Program       string `json:"program"`
	Folders       string `json:"folders"`
	Extensions    string `json:"extensions"`
	LaunchOptions string `json:"launchOptions"`
}

// Platform struct
type Platform struct {
	Name      string      `json:"name"`
	Console   string      `json:"console"`
	Emulators []*Emulator `json:"emulators"`
}

// Retrieve system platform specs.
// This list is almost a copy of EmulationStation DE systems
func GetPlatforms(options *Options) []*Platform {

	platforms := []*Platform{}

	platforms = append(platforms, &Platform{
		Name:    "DC",
		Console: "Sega Dreamcast",
		Emulators: []*Emulator{{
			Name:          "Flycast",
			Program:       "/var/lib/flatpak/exports/bin/org.flycast.Flycast",
			Folders:       "DC/",
			Extensions:    ".chd .cdi .iso .elf .cue .gdi .lst .dat .m3u .7z .zip",
			LaunchOptions: "-config window:fullscreen=yes \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "GBA",
		Console: "Nintendo Game Boy Advance",
		Emulators: []*Emulator{{
			Name:          "MGBA",
			Program:       "/var/lib/flatpak/exports/bin/io.mgba.mGBA",
			Folders:       "GBA/",
			Extensions:    ".agb .bin .cgb .dmg .gb .gba .gbc .sgb .7z .zip",
			LaunchOptions: "-f \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "GC",
		Console: "Nintendo GameCube",
		Emulators: []*Emulator{{
			Name:          "Dolphin Emulator",
			Program:       "/var/lib/flatpak/exports/bin/org.DolphinEmu.dolphin-emu",
			Folders:       "GC/",
			Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
			LaunchOptions: "-b -e \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "3DS",
		Console: "Nintendo 3DS",
		Emulators: []*Emulator{{
			Name:          "Citra",
			Program:       "/var/lib/flatpak/exports/bin/org.citra_emu.citra",
			Folders:       "3DS/",
			Extensions:    ".3ds .3dsx .app .axf .cci .cxi .elf .7z .zip",
			LaunchOptions: "\"${ROM}\"", // No full-screen option
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "N64",
		Console: "Nintendo 64",
		Emulators: []*Emulator{{
			Name:          "Simple64",
			Program:       "/var/lib/flatpak/exports/bin/io.github.simple64.simple64",
			Folders:       "N64/",
			Extensions:    ".bin .d64 .n64 .ndd .u1 .v64 .z64 .7z .zip",
			LaunchOptions: "--nogui \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "NDS",
		Console: "Nintendo DS",
		Emulators: []*Emulator{{
			Name:          "MelonDS",
			Program:       "/var/lib/flatpak/exports/bin/net.kuribo64.melonDS",
			Folders:       "NDS/",
			Extensions:    ".app .bin .nds .7z .zip",
			LaunchOptions: "-f \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS1",
		Console: "Sony PlayStation 1",
		Emulators: []*Emulator{{
			Name:          "DuckStation",
			Program:       "/var/lib/flatpak/exports/bin/org.duckstation.DuckStation",
			Folders:       "PS1/",
			Extensions:    ".bin .cbn .ccd .chd .cue .ecm .exe .img .iso .m3u .mdf .mds .minipsf .pbp .psexe .psf .toc .z .znx .7z .zip",
			LaunchOptions: "-batch -fullscreen \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS2",
		Console: "Sony PlayStation 2",
		Emulators: []*Emulator{{
			Name:          "PCSX2",
			Program:       "/var/lib/flatpak/exports/bin/net.pcsx2.PCSX2",
			Folders:       "PS2/",
			Extensions:    ".bin .chd .ciso .cso .dump .elf .gz .m3u .mdf .img .iso .isz .ngr",
			LaunchOptions: "-batch -nogui -fullscreen \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS3",
		Console: "Sony PlayStation 3",
		Emulators: []*Emulator{{
			Name:          "RPCS3",
			Program:       "/var/lib/flatpak/exports/bin/net.rpcs3.RPCS3",
			Folders:       "PS3/",
			Extensions:    ".desktop .ps3 .ps3dir",
			LaunchOptions: "--no-gui \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PSP",
		Console: "Sony PlayStation Portable",
		Emulators: []*Emulator{{
			Name:          "PPSSPP",
			Program:       "/var/lib/flatpak/exports/bin/org.ppsspp.PPSSPP",
			Folders:       "PSP/",
			Extensions:    ".elf .iso .cso .prx .pbp .7z .zip",
			LaunchOptions: "-f -g \"${ROM}\"",
		}},
	})

	if slices.Contains(options.Preferences, "use-ryujinx") {
		platforms = append(platforms, &Platform{
			Name:    "SWITCH",
			Console: "Nintendo Switch",
			Emulators: []*Emulator{{
				Name:          "Ryujinx",
				Program:       "/var/lib/flatpak/exports/bin/org.ryujinx.Ryujinx",
				Folders:       "SWITCH/",
				Extensions:    "nca .nro .nso .nsp .xci",
				LaunchOptions: "--fullscreen \"${ROM}\"",
			}},
		})
	} else {
		platforms = append(platforms, &Platform{
			Name:    "SWITCH",
			Console: "Nintendo Switch",
			Emulators: []*Emulator{{
				Name:          "Yuzu",
				Program:       "/var/lib/flatpak/exports/bin/org.yuzu_emu.yuzu",
				Folders:       "SWITCH/",
				Extensions:    "nca .nro .nso .nsp .xci",
				LaunchOptions: "-f -g \"${ROM}\"",
			}},
		})
	}

	platforms = append(platforms, &Platform{
		Name:    "WII",
		Console: "Nintendo Wii",
		Emulators: []*Emulator{{
			Name:          "Dolphin Emulator",
			Program:       "/var/lib/flatpak/exports/bin/org.DolphinEmu.dolphin-emu",
			Folders:       "WII/",
			Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
			LaunchOptions: "-b -e \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "WIIU",
		Console: "Nintendo Wii U",
		Emulators: []*Emulator{{
			Name:          "Cemu",
			Program:       "/var/lib/flatpak/exports/bin/info.cemu.Cemu",
			Folders:       "WIIU/",
			Extensions:    ".rpx .wua .wud .wux",
			LaunchOptions: "-f -g \"${ROM}\"",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "XBOX",
		Console: "Microsoft Xbox",
		Emulators: []*Emulator{{
			Name:          "Xemu",
			Program:       "/var/lib/flatpak/exports/bin/app.xemu.xemu",
			Folders:       "XBOX/",
			Extensions:    ".iso",
			LaunchOptions: "-full-screen -dvd_path \"${ROM}\"",
		}},
	})

	return platforms
}
