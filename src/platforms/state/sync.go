package state

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Synchronizable struct
type Synchronizable struct {
	Platform    string   `json:"platform"`
	Type        string   `json:"type"`
	Source      *fs.Info `json:"source"`
	Destination *fs.Info `json:"destination"`
	Recommended bool     `json:"recommended"`
}

// Return synchronizable items based on given options
func GetSynchronizables(options *Options) ([]*Synchronizable, error) {

	var result []*Synchronizable

	// Process each state
	for _, state := range GetStates(options) {

		// Check if should process this platform
		if !slices.Contains(options.Platforms, state.Platform) {
			continue
		}

		// Backup action copy from platform source to state destination
		if options.Action == "backup" {
			destination, err := fs.GetInfo(state.Destination)
			if err != nil {
				return result, err
			}

			// Check destination
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

				// Check if sync is recommended
				recommended := false
				if source.ModifiedTime > destination.ModifiedTime {
					recommended = true
				}

				// Append synchronizable information
				result = append(result, &Synchronizable{
					Platform:    state.Platform,
					Type:        state.Type,
					Source:      source,
					Destination: destination,
					Recommended: recommended,
				})

				// Ensure that only the first valid result will be processed
				break
			}

			continue
		}

		// Restore action copy from state destination to platform source
		// This is the inverse of the backup action
		if options.Action == "restore" {
			source, err := fs.GetInfo(state.Destination)
			if err != nil {
				return result, err
			}

			// Check source
			if source.Exist {
				if state.Type == "file" && source.Type != "file" {
					cli.Debug("Skipping file with invalid type: %s\n", source.Path)
					continue
				} else if state.Type == "folder" && source.Type != "folder" {
					cli.Debug("Skipping folder with invalid type: %s\n", source.Path)
					continue
				}
			}

			// Destination are in multiple locations due to multiple runtimes and operating systems
			// To ensure compatibility, we process just the first valid location for destination
			for _, destinationPath := range state.Source.Paths() {

				destination, err := fs.GetInfo(destinationPath)
				if err != nil {
					return result, err
				}

				// Ensure that destination exist to be processable
				if state.Type == "file" && destination.Exist == false {
					cli.Debug("Skipping file not detected: %s\n", destination.Path)
					continue
				} else if state.Type == "folder" && destination.Exist == false {
					cli.Debug("Skipping folder not detected: %s\n", destination.Path)
					continue
				}

				// Ensure that destination has the expected type to be processable
				if state.Type == "file" && destination.Type != "file" {
					cli.Debug("Skipping file with invalid type: %s\n", destination.Path)
					continue
				} else if state.Type == "folder" && destination.Type != "folder" {
					cli.Debug("Skipping folder with invalid type: %s\n", destination.Path)
					continue
				}

				// Check if sync is recommended
				recommended := false
				if source.ModifiedTime > destination.ModifiedTime {
					recommended = true
				}

				// Append synchronizable information
				result = append(result, &Synchronizable{
					Platform:    state.Platform,
					Type:        state.Type,
					Source:      source,
					Destination: destination,
					Recommended: recommended,
				})

				// Ensure that only the first valid result will be processed
				break
			}

			continue
		}
	}

	return result, nil
}

// Sync state of each platform
func SyncState(options *Options) error {

	// Get synchronizable information based on state and options
	synchronizable, err := GetSynchronizables(options)
	if err != nil {
		return err
	}

	// Process each synchronizable item
	for _, item := range synchronizable {

		// Skip items that source not exist
		if !item.Source.Exist {
			continue
		}

		// Fill source and destination information
		source := item.Source.Path
		destination := item.Destination.Path

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
