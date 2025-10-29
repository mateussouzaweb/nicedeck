package console

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// ROM struct
type ROM struct {
	Path          string `json:"path"`
	RelativePath  string `json:"relativePath"`
	Directory     string `json:"directory"`
	File          string `json:"file"`
	Extension     string `json:"extension"`
	Name          string `json:"name"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Platform      string `json:"platform"`
	Console       string `json:"console"`
	Emulator      string `json:"emulator"`
	Program       string `json:"program"`
	Executable    string `json:"executable"`
	LaunchOptions string `json:"launchOptions"`
}

// Find ROM details based on the given path
func ParseROM(path string, options *Options) (*ROM, error) {

	rom := &ROM{}

	// Get root paths
	root := options.RootPath
	realRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return rom, err
	}

	// Parse basic ROM information
	directory := filepath.Dir(path)
	file := filepath.Base(path)
	extension := filepath.Ext(path)
	name := strings.TrimSuffix(file, extension)
	separator := string(os.PathSeparator)

	// Ensure a valid final and relative path
	// Final path can be represented via symbolic links
	finalPath := strings.Replace(path, realRoot, root, 1)
	relativePath := strings.Replace(finalPath, root+separator, "", 1)

	// Check against exclusion list
	if options.ShouldExclude(relativePath) {
		cli.Debug("Skipped: file is in the exclude list.\n")
		return rom, nil
	}

	// Retrieve runtime detail
	runtime, err := FindRuntime(relativePath, options)
	if err != nil {
		return rom, err
	}

	// Ignore if could not detect the emulator
	if runtime.Emulator.Name == "" {
		cli.Debug("Skipped: no emulator found for ROM.\n")
		return rom, nil
	}

	// Validate if extension is in the valid list
	valid := strings.Split(runtime.Emulator.Extensions, " ")
	if !slices.Contains(valid, strings.ToLower(extension)) {
		cli.Debug("Skipped: invalid ROM format for %s emulator.\n", runtime.Emulator.Name)
		return rom, nil
	}

	// Valid, fill ROM data with runtime
	executable := runtime.Program.Package.Executable()
	launchOptions := runtime.Emulator.LaunchOptions
	launchOptions = strings.Replace(launchOptions, "${ROM}", cli.Quote(finalPath), 1)

	title := name + " [" + runtime.Platform.Name + "]"
	description := "ROM for " + runtime.Platform.Name

	rom.Path = finalPath
	rom.RelativePath = relativePath
	rom.Directory = directory
	rom.File = file
	rom.Extension = extension
	rom.Name = name
	rom.Title = title
	rom.Description = description
	rom.Console = runtime.Platform.Console
	rom.Platform = runtime.Platform.Name
	rom.Emulator = runtime.Emulator.Name
	rom.Program = runtime.Emulator.Program
	rom.Executable = executable
	rom.LaunchOptions = launchOptions

	return rom, nil
}

// Find ROMs in folder and return the list of detected games
func ParseROMs(options *Options) ([]*ROM, error) {

	var results []*ROM

	// Get ROMs path
	root := options.RootPath
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

		// Ignore directories without a directory extension
		// Please note that some emulators like PS3/PS4 use folders as ROM
		// These folders will always have a directory extension
		// So we only skip folders without an extension
		if dir.IsDir() && filepath.Ext(realPath) == "" {
			return nil
		}

		// Parse individual ROM file
		cli.Debug("Detected: %s\n", realPath)
		rom, err := ParseROM(realPath, options)
		if err != nil {
			return err
		} else if rom.Name == "" {
			return nil
		}

		// Check if same ROM already was found with another extension
		// This will prevent multiple results for the same ROM
		for _, item := range results {
			if item.Platform == rom.Platform && item.Name == rom.Name {
				cli.Debug("Skipped: multiple results detected for %s.\n", rom.Name)
				return nil
			}
		}

		cli.Debug("Valid: ROM is valid for %s emulator.\n", rom.Emulator)
		results = append(results, rom)
		return nil
	})

	return results, err
}

// Filter ROMs that match given requirements and return the list to process
func FilterROMs(roms []*ROM, existing []string, options *Options) []*ROM {

	var toProcess []*ROM

	// Rebuild option will include every ROM in the platform
	rebuild := slices.Contains(options.Preferences, "rebuild")

	// Fill the list of ROMs to process
	for _, rom := range roms {
		addToList := false

		// Add to the list if ROM matches platform condition
		if len(options.Platforms) == 0 {
			addToList = true
		} else if slices.Contains(options.Platforms, rom.Platform) {
			addToList = true
		} else {
			cli.Debug("No platform match. ROM skipped to the process list: %s\n", rom.RelativePath)
		}

		// When is not rebuilding, include only new detected ROMs
		if !rebuild {
			if slices.Contains(existing, rom.RelativePath) {
				cli.Debug("Existing. ROM skipped to the process list: %s\n", rom.RelativePath)
				addToList = false
			}
		}

		// Finally, if valid, add to the list of ROMs to process
		if addToList {
			cli.Debug("ROM added to the process list: %s\n", rom.RelativePath)
			toProcess = append(toProcess, rom)
		}
	}

	return toProcess
}
