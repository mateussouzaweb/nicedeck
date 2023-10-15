package roms

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// ROM struct
type ROM struct {
	Path          string `json:"path"`
	RelativePath  string `json:"relativePath"`
	Directory     string `json:"directory"`
	File          string `json:"file"`
	Extension     string `json:"extension"`
	Name          string `json:"name"`
	Platform      string `json:"platform"`
	Console       string `json:"console"`
	Emulator      string `json:"emulator"`
	LaunchCommand string `json:"launchCommand"`
}

// Find ROMs in folder and return the list of detected games
func ParseROMs() ([]*ROM, error) {

	var results []*ROM

	// Get ROMs path
	root := os.ExpandEnv("$HOME/Games/ROMs")
	realRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return results, err
	}

	// Fill exclude list
	// Files on these folders will be ignored
	exclude := []string{
		"/Updates/", // Updates folder
		"/Mods/",    // Mods folder
		"/Others/",  // Folder with games to ignore
		"(Disc 2)",  // Disc 2 of some games
		"(Disc 3)",  // Disc 3 of some games
		"(Track 2)", // Track 2 of some games
		"(Track 3)", // Track 3 of some games
		"(Track 4)", // Track 4 of some games
		"(Track 5)", // Track 5 of some games
		"(Track 6)", // Track 6 of some games
	}

	// Note: walkDir does not follow symbolic links
	err = filepath.WalkDir(realRoot, func(path string, dir os.DirEntry, err error) error {

		// Stop in case of errors
		if err != nil {
			return err
		}

		// Ignore directories
		if dir.IsDir() {
			return nil
		}

		directory := filepath.Dir(path)
		file := filepath.Base(path)
		extension := filepath.Ext(path)
		name := strings.TrimSuffix(file, extension)

		// Platform is determined by the initial path
		// This model also should solve cases for games in subfolders
		relativePath := strings.Replace(path, root+"/", "", 1)
		relativePath = strings.Replace(relativePath, realRoot+"/", "", 1)
		pathKeys := strings.Split(relativePath, "/")
		platform := pathKeys[0]

		console := ""
		emulator := ""
		extensions := ""
		launchCommand := ""

		// Fill data based on platform
		// This list is almost a copy of EmulationStation DE systems
		switch platform {
		case "DC":
			console = "Sega Dreamcast"
			emulator = "org.flycast.Flycast"
			extensions = ".chd .cdi .iso .elf .cue .gdi .lst .dat .m3u .7z .zip"
			launchCommand = "${EMULATOR} -config window:fullscreen=yes ${ROM}"
		case "GBA":
			console = "Nintendo Game Boy Advance"
			emulator = "io.mgba.mGBA"
			extensions = ".agb .bin .cgb .dmg .gb .gba .gbc .sgb .7z .zip"
			launchCommand = "${EMULATOR} -f ${ROM}"
		case "GC":
			console = "Nintendo GameCube"
			emulator = "org.DolphinEmu.dolphin-emu"
			extensions = ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip"
			launchCommand = "${EMULATOR} -b -e ${ROM}"
		case "3DS":
			console = "Nintendo 3DS"
			emulator = "org.citra_emu.citra"
			extensions = ".3ds .3dsx .app .axf .cci .cxi .elf .7z .zip"
			launchCommand = "${EMULATOR} ${ROM}" // No fullscreen option yet
		case "NDS":
			console = "Nintendo DS"
			emulator = "net.kuribo64.melonDS"
			extensions = ".app .bin .nds .7z .zip"
			launchCommand = "${EMULATOR} -f ${ROM}"
		case "PS2":
			console = "Sony PlayStation 2"
			emulator = "net.pcsx2.PCSX2"
			extensions = ".bin .chd .ciso .cso .dump .elf .gz .m3u .mdf .img .iso .isz .ngr"
			launchCommand = "${EMULATOR} -batch -nogui -fullscreen ${ROM}"
		case "PS3":
			console = "Sony PlayStation 3"
			emulator = "net.rpcs3.RPCS3"
			extensions = ".desktop .ps3 .ps3dir"
			launchCommand = "${EMULATOR} --no-gui ${ROM}"
		case "PSP":
			console = "Sony PlayStation Portable"
			emulator = "org.ppsspp.PPSSPP"
			extensions = ".elf .iso .cso .prx .pbp .7z .zip"
			launchCommand = "${EMULATOR} -f -g ${ROM}"
		case "SWITCH":
			console = "Nintendo Switch"
			emulator = "org.yuzu_emu.yuzu" // or org.ryujinx.Ryujinx
			extensions = "nca .nro .nso .nsp .xci"
			launchCommand = "${EMULATOR} -f -g ${ROM}" // or --fullscreen
		case "WII":
			console = "Nintendo Wii"
			emulator = "org.DolphinEmu.dolphin-emu"
			extensions = ".ciso .dff .dol .elf .gcm .gcz .iso .json .m3u .rvz .tgc .wad .wbfs .wia .7z .zip"
			launchCommand = "${EMULATOR} -b -e ${ROM}"
		case "WIIU":
			console = "Nintendo Wii U"
			emulator = "info.cemu.Cemu"
			extensions = ".rpx .wua .wud .wux"
			launchCommand = "${EMULATOR} -f -g ${ROM}"
		case "XBOX":
			console = "Microsoft Xbox"
			emulator = "app.xemu.xemu"
			extensions = ".iso"
			launchCommand = "${EMULATOR} -full-screen -dvd_path ${ROM}"
		}

		// Ignore if could not detect the emulator
		if emulator == "" {
			return nil
		}

		// Validate if extension is in the valid list
		valid := strings.Split(extensions, " ")
		if !slices.Contains(valid, strings.ToLower(extension)) {
			return nil
		}

		// Check against exclusion list
		// Verification is simple and consider if path contains given term
		for _, pattern := range exclude {
			if strings.Contains(strings.ToLower(path), strings.ToLower(pattern)) {
				return nil
			}
		}

		// Since all emulators are flatpak based
		// Is easy to fill the launch command
		launchCommand = strings.Replace(launchCommand, "${EMULATOR}", "/var/lib/flatpak/exports/bin/"+emulator, 1)
		launchCommand = strings.Replace(launchCommand, "${ROM}", "'"+path+"'", 1)

		rom := ROM{
			Path:          path,
			RelativePath:  relativePath,
			Directory:     directory,
			File:          file,
			Extension:     extension,
			Name:          name,
			Console:       console,
			Platform:      platform,
			Emulator:      emulator,
			LaunchCommand: launchCommand,
		}

		results = append(results, &rom)

		return nil
	})

	if err != nil {
		return results, err
	}

	return results, nil
}
