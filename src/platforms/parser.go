package platforms

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/programs"
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
	Program       string `json:"program"`
	Executable    string `json:"executable"`
	LaunchOptions string `json:"launchOptions"`
}

// Runtime struct
type Runtime struct {
	Platform *Platform
	Emulator *Emulator
	Program  *programs.Program
}

// Find runtime specs for ROM based on their path
func FindRuntimeForROM(romPath string, options *Options) (*Runtime, error) {

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

// Find ROMs in folder and return the list of detected games
func ParseROMs(options *Options) ([]*ROM, error) {

	var results []*ROM

	// Get ROMs path
	root := fs.ExpandPath("$ROMS")
	realRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return results, err
	}

	cli.Printf(cli.ColorNotice, "Checking for ROMs available at: %s\n", root)

	// Note: walkDir does not follow symbolic links
	err = filepath.WalkDir(realRoot, func(realPath string, dir os.DirEntry, err error) error {

		// Stop in case of errors
		if err != nil {
			return err
		}

		// Ignore directories
		if dir.IsDir() {
			return nil
		}

		// Parse basic ROM information
		directory := filepath.Dir(realPath)
		file := filepath.Base(realPath)
		extension := filepath.Ext(realPath)
		name := strings.TrimSuffix(file, extension)
		separator := string(os.PathSeparator)

		// Ensure a valid final and relative path
		// Final path can be represented via symbolic links
		finalPath := strings.Replace(realPath, realRoot, root, 1)
		relativePath := strings.Replace(finalPath, root+separator, "", 1)

		cli.Debug("Detected: %s\n", relativePath)

		// Check against exclusion list
		if options.ShouldExclude(relativePath) {
			cli.Debug("Skipped: file is in the exclude list.\n")
			return nil
		}

		// Retrieve runtime detail
		runtime, err := FindRuntimeForROM(relativePath, options)
		if err != nil {
			return err
		}

		// Ignore if could not detect the emulator
		if runtime.Emulator.Name == "" {
			cli.Debug("Skipped: no emulator found for ROM.\n")
			return nil
		}

		// Validate if extension is in the valid list
		valid := strings.Split(runtime.Emulator.Extensions, " ")
		if !slices.Contains(valid, strings.ToLower(extension)) {
			cli.Debug("Skipped: invalid ROM format for %s emulator.\n", runtime.Emulator.Name)
			return nil
		}

		// Check if same ROM already was found with another extension
		// This will prevent multiple results for the same ROM
		for _, item := range results {
			if item.Platform == runtime.Platform.Name && item.Name == name {
				cli.Debug("Skipped: multiple results detected for %s.\n", name)
				return nil
			}
		}

		// Put ROM path in launch options
		executable := runtime.Program.Package.Executable()
		launchOptions := runtime.Emulator.LaunchOptions
		launchOptions = strings.Replace(launchOptions, "${ROM}", cli.Quote(finalPath), 1)

		rom := ROM{
			Path:          finalPath,
			RelativePath:  relativePath,
			Directory:     directory,
			File:          file,
			Extension:     extension,
			Name:          name,
			Console:       runtime.Platform.Console,
			Platform:      runtime.Platform.Name,
			Emulator:      runtime.Emulator.Name,
			Program:       runtime.Emulator.Program,
			Executable:    executable,
			LaunchOptions: launchOptions,
		}

		cli.Debug("Valid: ROM is valid for %s emulator.\n", runtime.Emulator.Name)
		results = append(results, &rom)
		return nil
	})

	if err != nil {
		return results, err
	}

	return results, nil
}
