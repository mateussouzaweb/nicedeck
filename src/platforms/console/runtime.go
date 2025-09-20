package console

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// Runtime struct
type Runtime struct {
	Platform *Platform
	Emulator *Emulator
	Program  *programs.Program
}

// Find runtime specs for ROM based on their path
func FindRuntime(romPath string, options *Options) (*Runtime, error) {

	result := &Runtime{
		Platform: &Platform{},
		Emulator: &Emulator{},
		Program:  &programs.Program{},
	}

	romPath = strings.ToLower(romPath)

	// Retrieve platforms
	platforms, err := GetPlatforms(options)
	if err != nil {
		return result, err
	}

	// Search in every retrieved platform
	// Platform and emulator are determined by the folder initial path
	// Program is determined by installation status or availability
	// This model also work in cases when games are in sub-folders
	for _, platform := range platforms {
		separator := string(os.PathSeparator)
		mainFolder := strings.ToLower(platform.Folder + separator)

		// Skip platform without emulators
		if len(platform.Emulators) == 0 {
			continue
		}

		// Skip if platform folder prefix is not present in path
		// Means that the ROM belongs to another platform...
		// Please note that is important to check folder with path separator
		if !strings.HasPrefix(romPath, mainFolder) {
			continue
		}

		// Special case to enforce an specific emulator of the platform
		// The condition is to have the emulator name as subfolder
		// Please note that is important to check subfolder with path separator
		for _, emulator := range platform.Emulators {
			subFolder := strings.ReplaceAll(emulator.Name, " ", "-")
			subFolder = filepath.Join(mainFolder, subFolder)
			subFolder = strings.ToLower(subFolder + separator)

			if !strings.HasPrefix(romPath, subFolder) {
				continue
			}

			program, err := programs.GetProgramByID(emulator.Program)
			if err != nil {
				return result, err
			} else {
				result.Platform = platform
				result.Emulator = emulator
				result.Program = program
				return result, nil
			}
		}

		// Default case that will use the installed emulator
		// Check and use the first emulator that is installed for the platform
		for _, emulator := range platform.Emulators {
			program, err := programs.GetProgramByID(emulator.Program)
			if err != nil {
				return result, err
			}

			installed, err := program.Package.Installed()
			if err != nil {
				return result, err
			} else if installed {
				result.Platform = platform
				result.Emulator = emulator
				result.Program = program
				return result, nil
			}
		}

		// Last case that will use the available emulator
		// Check and use the first emulator that is available for the platform
		for _, emulator := range platform.Emulators {
			program, err := programs.GetProgramByID(emulator.Program)
			if err != nil {
				return result, err
			} else if program.Package.Available() {
				result.Platform = platform
				result.Emulator = emulator
				result.Program = program
				return result, nil
			}
		}
	}

	return result, nil
}
