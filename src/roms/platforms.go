package roms

// Platform struct
type Platform struct {
	Name          string `json:"name"`
	Console       string `json:"console"`
	Emulator      string `json:"emulator"`
	Extensions    string `json:"extensions"`
	LaunchOptions string `json:"launchOptions"`
}

// Retrieve system platform specs.
// This list is almost a copy of EmulationStation DE systems
func GetPlatforms(options *Options) []*Platform {

	platforms := []*Platform{}

	platforms = append(platforms, &Platform{
		Name:          "DC",
		Console:       "Sega Dreamcast",
		Emulator:      "org.flycast.Flycast",
		Extensions:    ".chd .cdi .iso .elf .cue .gdi .lst .dat .m3u .7z .zip",
		LaunchOptions: "-config window:fullscreen=yes \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "GBA",
		Console:       "Nintendo Game Boy Advance",
		Emulator:      "io.mgba.mGBA",
		Extensions:    ".agb .bin .cgb .dmg .gb .gba .gbc .sgb .7z .zip",
		LaunchOptions: "-f \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "GC",
		Console:       "Nintendo GameCube",
		Emulator:      "org.DolphinEmu.dolphin-emu",
		Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
		LaunchOptions: "-b -e \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "3DS",
		Console:       "Nintendo 3DS",
		Emulator:      "org.citra_emu.citra",
		Extensions:    ".3ds .3dsx .app .axf .cci .cxi .elf .7z .zip",
		LaunchOptions: "\"${ROM}\"", // No fullscreen option yet
	})
	platforms = append(platforms, &Platform{
		Name:          "N64",
		Console:       "Nintendo 64",
		Emulator:      "io.github.simple64.simple64",
		Extensions:    ".bin .d64 .n64 .ndd .u1 .v64 .z64 .7z .zip",
		LaunchOptions: "--nogui \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "NDS",
		Console:       "Nintendo DS",
		Emulator:      "net.kuribo64.melonDS",
		Extensions:    ".app .bin .nds .7z .zip",
		LaunchOptions: "-f \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "PS1",
		Console:       "Sony PlayStation 1",
		Emulator:      "org.duckstation.DuckStation",
		Extensions:    ".bin .cbn .ccd .chd .cue .ecm .exe .img .iso .m3u .mdf .mds .minipsf .pbp .psexe .psf .toc .z .znx .7z .zip",
		LaunchOptions: "-batch -fullscreen \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "PS2",
		Console:       "Sony PlayStation 2",
		Emulator:      "net.pcsx2.PCSX2",
		Extensions:    ".bin .chd .ciso .cso .dump .elf .gz .m3u .mdf .img .iso .isz .ngr",
		LaunchOptions: "-batch -nogui -fullscreen \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "PS3",
		Console:       "Sony PlayStation 3",
		Emulator:      "net.rpcs3.RPCS3",
		Extensions:    ".desktop .ps3 .ps3dir",
		LaunchOptions: "--no-gui \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "PSP",
		Console:       "Sony PlayStation Portable",
		Emulator:      "org.ppsspp.PPSSPP",
		Extensions:    ".elf .iso .cso .prx .pbp .7z .zip",
		LaunchOptions: "-f -g \"${ROM}\"",
	})

	platforms = append(platforms, &Platform{
		Name:          "SWITCH",
		Console:       "Nintendo Switch",
		Emulator:      "org.yuzu_emu.yuzu",
		Extensions:    "nca .nro .nso .nsp .xci",
		LaunchOptions: "-f -g \"${ROM}\"",
	})

	if !options.UseRyujinx {
		platforms = append(platforms, &Platform{
			Name:          "SWITCH",
			Console:       "Nintendo Switch",
			Emulator:      "org.yuzu_emu.yuzu",
			Extensions:    "nca .nro .nso .nsp .xci",
			LaunchOptions: "-f -g \"${ROM}\"",
		})
	} else {
		platforms = append(platforms, &Platform{
			Name:          "SWITCH",
			Console:       "Nintendo Switch",
			Emulator:      "org.ryujinx.Ryujinx",
			Extensions:    "nca .nro .nso .nsp .xci",
			LaunchOptions: "--fullscreen \"${ROM}\"",
		})
	}

	platforms = append(platforms, &Platform{
		Name:          "WII",
		Console:       "Nintendo Wii",
		Emulator:      "org.DolphinEmu.dolphin-emu",
		Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
		LaunchOptions: "-b -e \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "WIIU",
		Console:       "Nintendo Wii U",
		Emulator:      "info.cemu.Cemu",
		Extensions:    ".rpx .wua .wud .wux",
		LaunchOptions: "-f -g \"${ROM}\"",
	})
	platforms = append(platforms, &Platform{
		Name:          "XBOX",
		Console:       "Microsoft Xbox",
		Emulator:      "app.xemu.xemu",
		Extensions:    ".iso",
		LaunchOptions: "-full-screen -dvd_path \"${ROM}\"",
	})

	return platforms
}
