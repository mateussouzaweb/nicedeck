package console

import (
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// Emulator struct
type Emulator struct {
	Type          string `json:"type"`
	Name          string `json:"name"`
	Program       string `json:"program"`
	Available     bool   `json:"available"`
	Installed     bool   `json:"installed"`
	Executable    string `json:"executable"`
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
// This list is almost a copy of ES-DE systems
func GetPlatforms(options *Options) ([]*Platform, error) {

	platforms := []*Platform{}

	platforms = append(platforms, &Platform{
		Name:    "3DS",
		Console: "Nintendo 3DS",
		Folder:  "3DS",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Azahar",
			Program:       "azahar",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".3ds .3dsx .app .axf .cci .cxi .elf .7z .zip",
			LaunchOptions: "${ROM}", // No full-screen option
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "DC",
		Console: "Sega Dreamcast",
		Folder:  "DC",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Flycast",
			Program:       "flycast",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".chd .cdi .iso .elf .cue .gdi .lst .dat .m3u .7z .zip",
			LaunchOptions: "-config window:fullscreen=yes ${ROM}",
		}, {
			Type:          "program",
			Name:          "Redream",
			Program:       "redream",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".chd .cdi .cue .gdi .7z",
			LaunchOptions: "-b -e ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "GBA",
		Console: "Nintendo Game Boy Advance",
		Folder:  "GBA",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "MGBA",
			Program:       "mgba",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".agb .bin .cgb .dmg .gb .gba .gbc .sgb .7z .zip",
			LaunchOptions: "-f ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "GC",
		Console: "Nintendo GameCube",
		Folder:  "GC",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Dolphin Emulator",
			Program:       "dolphin",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
			LaunchOptions: "-b -e ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "N64",
		Console: "Nintendo 64",
		Folder:  "N64",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Gopher64",
			Program:       "gopher64",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".bin .d64 .n64 .ndd .u1 .v64 .z64 .7z .zip",
			LaunchOptions: "-f ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "NDS",
		Console: "Nintendo DS",
		Folder:  "NDS",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "MelonDS",
			Program:       "melonds",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".app .bin .nds .7z .zip",
			LaunchOptions: "-f ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS1",
		Console: "Sony PlayStation 1",
		Folder:  "PS1",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "DuckStation",
			Program:       "duckstation",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".bin .cbn .ccd .chd .cue .ecm .exe .img .iso .m3u .mdf .mds .minipsf .pbp .psexe .psf .toc .z .znx .7z .zip",
			LaunchOptions: "-batch -fullscreen ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS2",
		Console: "Sony PlayStation 2",
		Folder:  "PS2",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "PCSX2",
			Program:       "pcsx2",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".bin .chd .ciso .cso .dump .elf .gz .m3u .mdf .img .iso .isz .ngr",
			LaunchOptions: "-batch -nogui -fullscreen ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS3",
		Console: "Sony PlayStation 3",
		Folder:  "PS3",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "RPCS3",
			Program:       "rpcs3",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".ps3 .ps3dir .iso",
			LaunchOptions: "--no-gui ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PS4",
		Console: "Sony PlayStation 4",
		Folder:  "PS4",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "ShadPS4",
			Program:       "shadps4",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".ps4",
			LaunchOptions: "-d -g ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PSP",
		Console: "Sony PlayStation Portable",
		Folder:  "PSP",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "PPSSPP",
			Program:       "ppsspp",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".elf .iso .cso .prx .pbp .7z .zip",
			LaunchOptions: "-f -g ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "PSVITA",
		Console: "Sony PlayStation Vita",
		Folder:  "PSVITA",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Vita3K",
			Program:       "vita3k",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".psvita",
			LaunchOptions: "-F -r ${CONTENT}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "SWITCH",
		Console: "Nintendo Switch",
		Folder:  "SWITCH",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Eden",
			Program:       "eden",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".nca .nro .nso .nsp .xci",
			LaunchOptions: "-f -g ${ROM}",
		}, {
			Type:          "program",
			Name:          "Ryujinx",
			Program:       "ryujinx",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".nca .nro .nso .nsp .xci",
			LaunchOptions: "--fullscreen ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "WII",
		Console: "Nintendo Wii",
		Folder:  "WII",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Dolphin",
			Program:       "dolphin",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip",
			LaunchOptions: "-b -e ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "WIIU",
		Console: "Nintendo Wii U",
		Folder:  "WIIU",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Cemu",
			Program:       "cemu",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".rpx .wua .wud .wux",
			LaunchOptions: "-f -g ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "XBOX",
		Console: "Microsoft Xbox",
		Folder:  "XBOX",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Xemu",
			Program:       "xemu",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".iso",
			LaunchOptions: "-full-screen -dvd_path ${ROM}",
		}},
	})

	platforms = append(platforms, &Platform{
		Name:    "X360",
		Console: "Microsoft Xbox 360",
		Folder:  "X360",
		Emulators: []*Emulator{{
			Type:          "program",
			Name:          "Xenia",
			Program:       "xenia",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".iso .zar",
			LaunchOptions: "--fullscreen=true ${ROM}",
		}},
	})

	// Fill information on each emulator based on their attributes
	for _, platform := range platforms {
		for _, emulator := range platform.Emulators {

			// Emulators that are based on program state
			if emulator.Type == "program" {
				program, err := programs.GetProgramByID(emulator.Program)
				if err != nil {
					return platforms, err
				}

				installed, err := program.Package.Installed()
				if err != nil {
					return platforms, err
				}

				emulator.Installed = installed
				emulator.Available = program.Package.Available()
				emulator.Executable = program.Package.Executable()
			}

		}
	}

	return platforms, nil
}
