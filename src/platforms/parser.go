package platforms

import (
	"os"
	"path/filepath"
	"regexp"
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
	Executable    string `json:"executable"`
	LaunchOptions string `json:"launchOptions"`
}

// Find ROMs in folder and return the list of detected games
func ParseROMs(options *Options) ([]*ROM, error) {

	var results []*ROM

	// Get ROMs path
	separator := string(os.PathSeparator)
	root := fs.ExpandPath("$ROMS")
	realRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return results, err
	}

	// Retrieve platforms
	platforms, err := GetPlatforms(options)
	if err != nil {
		return results, err
	}

	// Fill exclude list
	// Files on these folders will be ignored
	exclude := []string{
		fs.NormalizePath("/Updates/"), // Updates folder
		fs.NormalizePath("/Mods/"),    // Mods folder
		fs.NormalizePath("/DLCs/"),    // DLCs folder
		fs.NormalizePath("/Others/"),  // Folder with games to ignore
		fs.NormalizePath("/Ignore/"),  // Folder with games to ignore
		fs.NormalizePath("/Hide/"),    // Folder with games to ignore
	}

	// Files with these name patterns will be ignored
	excludeRegex := []*regexp.Regexp{
		regexp.MustCompile("(?i)Disc 0?[2-9]"),     // Disc 02 - 09 of some games
		regexp.MustCompile("(?i)Disc [1-9][0-9]"),  // Disc 10 - 99 of some games
		regexp.MustCompile("(?i)Track 0?[1-9]"),    // Track 01 - 09 of some games
		regexp.MustCompile("(?i)Track [1-9][0-9]"), // Track 10 - 99 of some games
	}

	cli.Printf(cli.ColorNotice, "Checking for ROMs available at: %s\n", realRoot)

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

		// Parse basic ROM information
		directory := filepath.Dir(path)
		file := filepath.Base(path)
		extension := filepath.Ext(path)
		name := strings.TrimSuffix(file, extension)

		// Ensure a valid relative path
		relativePath := strings.Replace(path, root+separator, "", 1)
		relativePath = strings.Replace(relativePath, realRoot+separator, "", 1)
		cli.Printf(cli.ColorWarn, "Detected: %s\n", relativePath)

		// Check against exclusion list
		// Verification is simple and consider if path contains given term
		for _, pattern := range exclude {
			if strings.Contains(strings.ToLower(path), strings.ToLower(pattern)) {
				cli.Printf(cli.ColorWarn, "Skipped: file is in the exclude list\n")
				return nil
			}
		}

		// Check against regex exclusion list
		for _, pattern := range excludeRegex {
			if pattern.MatchString(path) {
				cli.Printf(cli.ColorWarn, "Skipped: file is in the exclude list\n")
				return nil
			}
		}

		// Platform and emulator are determined by the folder initial path
		// This model also solve cases for games in sub-folders
		platform := &Platform{}
		emulator := &Emulator{}

		for _, item := range platforms {
			// Ignore platforms without emulators
			if len(item.Emulators) == 0 {
				continue
			}

			// Skip if platform folder prefix is not present in path
			// Means that the ROM belongs to another platform...
			// Please note that is important to check folder with path separator
			if !strings.HasPrefix(strings.ToLower(relativePath), strings.ToLower(item.Folder+separator)) {
				continue
			}

			// Special case to enforce an specific emulator of the platform
			// The condition is to have the emulator name as subfolder
			// Please note that is important to check subfolder with path separator
			for _, itemEmulator := range item.Emulators {
				subfolder := strings.ReplaceAll(itemEmulator.Name, " ", "-")
				subfolder = filepath.Join(item.Folder, subfolder)
				subfolder = strings.ToLower(subfolder + separator)

				if strings.HasPrefix(strings.ToLower(relativePath), subfolder) {
					platform = item
					emulator = itemEmulator
					break
				}
			}
			if emulator.Name != "" {
				break
			}

			// Default case that will use the main platform emulator
			// Using the first emulator that is available for the system
			for _, itemEmulator := range item.Emulators {
				program, err := programs.GetProgramByID(itemEmulator.Program)
				if err != nil {
					return err
				} else if program.Package.Available() {
					platform = item
					emulator = itemEmulator
					break
				}
			}
			if emulator.Name != "" {
				break
			}
		}

		// Ignore if could not detect the emulator
		if emulator.Name == "" {
			cli.Printf(cli.ColorWarn, "Skipped: no emulator found for ROM\n")
			return nil
		}

		// Validate if extension is in the valid list
		valid := strings.Split(emulator.Extensions, " ")
		if !slices.Contains(valid, strings.ToLower(extension)) {
			cli.Printf(cli.ColorWarn, "Skipped: invalid ROM format for %s emulator\n", emulator.Name)
			return nil
		}

		// Check if same ROM already was found with another extension
		// This will prevent multiple results for the same ROM
		for _, item := range results {
			if item.Platform == platform.Name && item.Name == name {
				cli.Printf(cli.ColorWarn, "Skipped: multiple results detected for %s\n", name)
				return nil
			}
		}

		// Find target program from the emulator
		// Ignore when program package is not available for the system
		program, err := programs.GetProgramByID(emulator.Program)
		if err != nil {
			return err
		} else if !program.Package.Available() {
			cli.Printf(cli.ColorWarn, "Skipped: %s emulator program not available\n", emulator.Name)
			return nil
		}

		// Put ROM path in launch options
		executable := program.Package.Executable()
		launchOptions := strings.Replace(emulator.LaunchOptions, "${ROM}", path, 1)

		rom := ROM{
			Path:          path,
			RelativePath:  relativePath,
			Directory:     directory,
			File:          file,
			Extension:     extension,
			Name:          name,
			Console:       platform.Console,
			Platform:      platform.Name,
			Emulator:      emulator.Name,
			Executable:    executable,
			LaunchOptions: launchOptions,
		}

		cli.Printf(cli.ColorSuccess, "Valid: ROM is valid for %s emulator\n", emulator.Name)
		results = append(results, &rom)
		return nil
	})

	if err != nil {
		return results, err
	}

	return results, nil
}
