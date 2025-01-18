package platforms

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

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
	root := os.ExpandEnv("$ROMS")
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
		"/Updates/", // Updates folder
		"/Mods/",    // Mods folder
		"/DLCs/",    // DLCs folder
		"/Others/",  // Folder with games to ignore
		"/Ignore/",  // Folder with games to ignore
		"/Hide/",    // Folder with games to ignore
	}

	// Files with these name patterns will be ignored
	excludeRegex := []*regexp.Regexp{
		regexp.MustCompile("(?i)Disc 0?[2-9]"),     // Disc 02 - 09 of some games
		regexp.MustCompile("(?i)Disc [1-9][0-9]"),  // Disc 10 - 99 of some games
		regexp.MustCompile("(?i)Track 0?[1-9]"),    // Track 01 - 09 of some games
		regexp.MustCompile("(?i)Track [1-9][0-9]"), // Track 10 - 99 of some games
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

		// Check against exclusion list
		// Verification is simple and consider if path contains given term
		for _, pattern := range exclude {
			if strings.Contains(strings.ToLower(path), strings.ToLower(pattern)) {
				return nil
			}
		}

		// Check against regex exclusion list
		for _, pattern := range excludeRegex {
			if pattern.MatchString(path) {
				return nil
			}
		}

		directory := filepath.Dir(path)
		file := filepath.Base(path)
		extension := filepath.Ext(path)
		name := strings.TrimSuffix(file, extension)

		relativePath := strings.Replace(path, root+"/", "", 1)
		relativePath = strings.Replace(relativePath, realRoot+"/", "", 1)

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
			if !strings.HasPrefix(strings.ToLower(relativePath), strings.ToLower(item.Folder)) {
				continue
			}

			// Special case to enforce an specific emulator of the platform
			// The condition is to have the emulator name as subfolder
			for _, itemEmulator := range item.Emulators {
				subfolder := strings.ReplaceAll(itemEmulator.Name, " ", "-")
				folder := strings.ToLower(item.Folder + subfolder + "/")
				if strings.HasPrefix(strings.ToLower(relativePath), folder) {
					platform = item
					emulator = itemEmulator
					break
				}
			}
			if emulator.Name != "" {
				break
			}

			// Default case that will use the main platform emulator
			platform = item
			emulator = item.Emulators[0]
			break
		}

		// Ignore if could not detect the emulator
		if emulator.Name == "" {
			return nil
		}

		// Validate if extension is in the valid list
		valid := strings.Split(emulator.Extensions, " ")
		if !slices.Contains(valid, strings.ToLower(extension)) {
			return nil
		}

		// Check if same ROM already was found with another extension
		// This will prevent multiple results for the same ROM
		for _, item := range results {
			if item.Platform == platform.Name && item.Name == name {
				return nil
			}
		}

		// Find target program from the emulator
		program, err := programs.GetProgramByID(emulator.Program)
		if err != nil {
			return err
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

		results = append(results, &rom)

		return nil
	})

	if err != nil {
		return results, err
	}

	return results, nil
}
