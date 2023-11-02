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
	LaunchOptions string `json:"launchOptions"`
}

// Find ROMs in folder and return the list of detected games
func ParseROMs(options *Options) ([]*ROM, error) {

	var results []*ROM

	// Get ROMs path
	root := os.ExpandEnv("$HOME/Games/ROMs")
	realRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return results, err
	}

	// Retrieve platforms
	platforms := GetPlatforms(options)

	// Fill exclude list
	// Files on these folders will be ignored
	exclude := []string{
		"/Updates/",  // Updates folder
		"/Mods/",     // Mods folder
		"/DLCs/",     // DLCs folder
		"/Others/",   // Folder with games to ignore
		"/Ignore/",   // Folder with games to ignore
		"/Hide/",     // Folder with games to ignore
		"(Disc 2)",   // Disc 2 of some games
		"(Disc 3)",   // Disc 3 of some games
		"(Disc 4)",   // Disc 4 of some games
		"(Disc 5)",   // Disc 5 of some games
		"(Disc 6)",   // Disc 6 of some games
		"(Track 1)",  // Track 1 of some games
		"(Track 2)",  // Track 2 of some games
		"(Track 3)",  // Track 3 of some games
		"(Track 4)",  // Track 4 of some games
		"(Track 5)",  // Track 5 of some games
		"(Track 6)",  // Track 6 of some games
		"(Track 7)",  // Track 7 of some games
		"(Track 8)",  // Track 8 of some games
		"(Track 9)",  // Track 9 of some games
		"(Track 10)", // Track 10 of some games
		"(Track 11)", // Track 11 of some games
		"(Track 12)", // Track 12 of some games
		"(Track 13)", // Track 13 of some games
		"(Track 14)", // Track 14 of some games
		"(Track 15)", // Track 15 of some games
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

		relativePath := strings.Replace(path, root+"/", "", 1)
		relativePath = strings.Replace(relativePath, realRoot+"/", "", 1)

		// Platform is determined by the initial path
		// This model is simple and also solve cases for games in subfolders
		pathKeys := strings.Split(relativePath, "/")
		platform := &Platform{}

		for _, item := range platforms {
			if pathKeys[0] == item.Name {
				platform = item
				break
			}
		}

		// Ignore if could not detect the emulator
		if platform.Emulator == "" {
			return nil
		}

		// Validate if extension is in the valid list
		valid := strings.Split(platform.Extensions, " ")
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

		// Put ROM path in launch options
		launchOptions := strings.Replace(platform.LaunchOptions, "${ROM}", path, 1)

		rom := ROM{
			Path:          path,
			RelativePath:  relativePath,
			Directory:     directory,
			File:          file,
			Extension:     extension,
			Name:          name,
			Console:       platform.Console,
			Platform:      platform.Name,
			Emulator:      platform.Emulator,
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
