package native

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Parse application details based on the given path
func Parse(path string, options *Options) (*App, error) {

	app := &App{}

	// Parse basic application information
	directory := filepath.Dir(path)
	file := filepath.Base(path)
	extension := filepath.Ext(path)
	name := strings.TrimSuffix(file, extension)

	// Replace placeholder with real data on given value
	replaceData := func(value string) string {
		value = strings.Replace(value, "${PATH}", cli.Quote(path), 1)
		value = strings.Replace(value, "${DIRECTORY}", cli.Quote(directory), 1)
		return value
	}

	// Find runtime based on file data
	for _, platform := range options.Platforms {
		valid := strings.Split(platform.Extensions, " ")
		if !slices.Contains(valid, strings.ToLower(extension)) {
			continue
		}

		// Valid, fill application data with platform
		app.Reference = name
		app.Name = name
		app.Description = fmt.Sprintf("%s executable", platform.Runtime)
		app.Runtime = platform.Runtime
		app.Platform = platform.Name
		app.Path = path
		app.Directory = directory
		app.File = file
		app.Extension = extension
		app.StartDirectory = replaceData(platform.StartDirectory)
		app.Executable = replaceData(platform.Executable)
		app.LaunchOptions = replaceData(platform.LaunchOptions)
		break
	}

	return app, nil
}

// Find applications in folder and return the list of detected applications
func Find(options *Options) ([]*App, error) {
	var results []*App
	return results, nil
}

// Filter applications that match given requirements and return the list to process
func Filter(list []*App, existing []string, options *Options) []*App {

	var toProcess []*App

	// Rebuild option will include every application in the platform
	rebuild := slices.Contains(options.Preferences, "rebuild")

	// Fill the list of applications to process
	for _, item := range list {
		addToList := false

		// Add to the list if application matches platform condition
		if len(options.Platforms) == 0 {
			addToList = true
		} else if slices.Contains(options.Include, item.Platform) {
			addToList = true
		} else {
			cli.Debug("No platform match. Application skipped to the process list: %s\n", item.Reference)
		}

		// When is not rebuilding, include only new detected applications
		if !rebuild {
			if slices.Contains(existing, item.Reference) {
				cli.Debug("Existing. Application skipped to the process list: %s\n", item.Reference)
				addToList = false
			}
		}

		// Finally, if valid, add to the list of applications to process
		if addToList {
			cli.Debug("Application added to the process list: %s\n", item.Reference)
			toProcess = append(toProcess, item)
		}
	}

	return toProcess
}
