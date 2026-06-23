package state

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

type Synchronizable struct {
	Platform    string
	Type        string
	Source      *fs.Info
	Destination *fs.Info
}

// Return synchronizable states based on given options
func GetSynchronizables(options *Options) ([]*Synchronizable, error) {

	var result []*Synchronizable

	// Process each state
	for _, state := range GetStates(options) {

		// Check if should process this platform
		if !slices.Contains(options.Platforms, state.Platform) {
			continue
		}

		// By default, destination is the path defined in state as safe location
		// In restore action, destination and source path are inverted
		destination, err := fs.GetInfo(state.Path)
		if err != nil {
			return result, err
		}

		// When destination path exist, ensure that it has the expected type to be processable
		// This is important to avoid synchronization issues
		if destination.Exist {
			if state.Type == "file" && destination.Type != "file" {
				cli.Debug("Skipping file with invalid type: %s\n", destination.Path)
				continue
			} else if state.Type == "folder" && destination.Type != "folder" {
				cli.Debug("Skipping folder with invalid type: %s\n", destination.Path)
				continue
			}
		}

		// Source are in multiple locations due to multiple runtimes and operating systems
		// To ensure compatibility, we process just the first valid location for source
		for _, sourcePath := range state.Source.Paths() {

			source, err := fs.GetInfo(sourcePath)
			if err != nil {
				return result, err
			}

			// Ensure that source exist to be processable
			if state.Type == "file" && source.Exist == false {
				cli.Debug("Skipping file not detected: %s\n", source.Path)
				continue
			} else if state.Type == "folder" && source.Exist == false {
				cli.Debug("Skipping folder not detected: %s\n", source.Path)
				continue
			}

			// Ensure that source has the expected type to be processable
			if state.Type == "file" && source.Type != "file" {
				cli.Debug("Skipping file with invalid type: %s\n", source.Path)
				continue
			} else if state.Type == "folder" && source.Type != "folder" {
				cli.Debug("Skipping folder with invalid type: %s\n", source.Path)
				continue
			}

			// Append synchronizable information
			result = append(result, &Synchronizable{
				Platform:    state.Platform,
				Type:        state.Type,
				Source:      source,
				Destination: destination,
			})

			// Ensure that only the first valid result will be processed
			break

		}
	}

	return result, nil
}

// Sync state of each platform
func SyncState(action string, options *Options) error {

	// Get synchronizable information based on state and options
	synchronizable, err := GetSynchronizables(options)
	if err != nil {
		return err
	}

	// Process each synchronizable item
	for _, item := range synchronizable {

		// Fill source and destination information
		source := item.Source.Path
		destination := item.Destination.Path

		// Default action is copy from source to destination path as backup method
		// However, user can choose to restore state with optional preference
		// When using restore method, invert path information
		if action == "restore" {
			source = item.Destination.Path
			destination = item.Source.Path
		}

		// Process file or folder state
		switch item.Type {
		case "file":

			// Copy file
			cli.Printf(cli.ColorNotice, "Synchronizing file from %s to %s...\n", source, destination)
			err = fs.CopyFile(source, destination, true)
			if err != nil {
				return err
			}

		case "folder":

			// Recursive copy content
			cli.Printf(cli.ColorNotice, "Synchronizing folder from %s to %s...\n", source, destination)
			err = fs.CopyDirectory(source, destination)
			if err != nil {
				return err
			}

		}
	}

	cli.Printf(cli.ColorNotice, "State synchronized.\n")
	return nil
}
