package console

import (
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// Emulator struct
type Emulator struct {
	Name          string `json:"name"`
	Program       string `json:"program"`
	Available     bool   `json:"available"`
	Installed     bool   `json:"installed"`
	Executable    string `json:"executable"`
	Extensions    string `json:"extensions"`
	LaunchOptions string `json:"launchOptions"`
}

// Custom emulator struct
type CustomEmulator struct {
	Name          string `json:"name"`
	Platform      string `json:"platform"`
	Program       string `json:"program"`
	Executable    string `json:"executable"`
	Extensions    string `json:"extensions"`
	LaunchOptions string `json:"launchOptions"`
}

// Platform struct
type Platform struct {
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	Console   string      `json:"console"`
	Folder    string      `json:"folder"`
	Emulators []*Emulator `json:"emulators"`
}

// Custom platform struct
type CustomPlatform struct {
	Name    string `json:"name"`
	Console string `json:"console"`
	Folder  string `json:"folder"`
}

// Retrieve system platform specs.
// This list is almost a copy of ES-DE systems
func GetPlatforms() ([]*Platform, error) {

	platforms := []*Platform{}

	platforms = append(platforms, &Platform{
		Type:    "native",
		Name:    "3DS",
		Console: "Nintendo 3DS",
		Folder:  "3DS",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "DC",
		Console: "Sega Dreamcast",
		Folder:  "DC",
		Emulators: []*Emulator{{
			Name:          "Flycast",
			Program:       "flycast",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".chd .cdi .iso .elf .cue .gdi .lst .dat .m3u .7z .zip",
			LaunchOptions: "-config window:fullscreen=yes ${ROM}",
		}, {
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
		Type:    "native",
		Name:    "GBA",
		Console: "Nintendo Game Boy Advance",
		Folder:  "GBA",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "GC",
		Console: "Nintendo GameCube",
		Folder:  "GC",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "N64",
		Console: "Nintendo 64",
		Folder:  "N64",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "NDS",
		Console: "Nintendo DS",
		Folder:  "NDS",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "PS1",
		Console: "Sony PlayStation 1",
		Folder:  "PS1",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "PS2",
		Console: "Sony PlayStation 2",
		Folder:  "PS2",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "PS3",
		Console: "Sony PlayStation 3",
		Folder:  "PS3",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "PS4",
		Console: "Sony PlayStation 4",
		Folder:  "PS4",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "PSP",
		Console: "Sony PlayStation Portable",
		Folder:  "PSP",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "PSVITA",
		Console: "Sony PlayStation Vita",
		Folder:  "PSVITA",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "SWITCH",
		Console: "Nintendo Switch",
		Folder:  "SWITCH",
		Emulators: []*Emulator{{
			Name:          "Eden",
			Program:       "eden",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".nca .nro .nso .nsp .xci",
			LaunchOptions: "-f -g ${ROM}",
		}, {
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
		Type:    "native",
		Name:    "WII",
		Console: "Nintendo Wii",
		Folder:  "WII",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "WIIU",
		Console: "Nintendo Wii U",
		Folder:  "WIIU",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "XBOX",
		Console: "Microsoft Xbox",
		Folder:  "XBOX",
		Emulators: []*Emulator{{
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
		Type:    "native",
		Name:    "X360",
		Console: "Microsoft Xbox 360",
		Folder:  "X360",
		Emulators: []*Emulator{{
			Name:          "Xenia",
			Program:       "xenia",
			Available:     false,
			Installed:     false,
			Executable:    "",
			Extensions:    ".iso .zar",
			LaunchOptions: "--fullscreen=true ${ROM}",
		}},
	})

	// Fill information on each emulator based on their program attributes
	for _, platform := range platforms {
		for _, emulator := range platform.Emulators {
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

	// Read custom platforms from configuration file
	customPlatformsFile := fs.ExpandPath("$APPLICATIONS/NiceDeck/custom/platforms.json")
	customPlatforms := make([]CustomPlatform, 0)
	err := fs.ReadJSON(customPlatformsFile, &customPlatforms)
	if err != nil {
		return platforms, err
	}

	// Merge custom platforms with the built-in platforms
	for _, customPlatform := range customPlatforms {
		platform := &Platform{
			Type:      "custom",
			Name:      customPlatform.Name,
			Console:   customPlatform.Console,
			Folder:    customPlatform.Folder,
			Emulators: []*Emulator{},
		}

		platforms = append(platforms, platform)
	}

	// Read custom emulators from configuration file
	customEmulatorsFile := fs.ExpandPath("$APPLICATIONS/NiceDeck/custom/emulators.json")
	customEmulators := make([]CustomEmulator, 0)
	err = fs.ReadJSON(customEmulatorsFile, &customEmulators)
	if err != nil {
		return platforms, err
	}

	// Merge custom emulators into platforms
	// Custom emulators are prioritized over the default ones
	for _, customEmulator := range customEmulators {
		emulator := &Emulator{
			Name:          customEmulator.Name,
			Program:       customEmulator.Program,
			Available:     true, // Custom emulators are always available
			Installed:     true, // Custom emulators are always installed
			Executable:    customEmulator.Executable,
			Extensions:    customEmulator.Extensions,
			LaunchOptions: customEmulator.LaunchOptions,
		}

		for _, platform := range platforms {
			if customEmulator.Platform == platform.Name {
				platform.Emulators = append([]*Emulator{emulator}, platform.Emulators...)
			}
		}
	}

	return platforms, nil
}
