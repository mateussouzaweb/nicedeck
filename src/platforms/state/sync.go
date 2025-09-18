package state

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Sync state of each platform
func SyncState(action string, options *Options) error {

	// Default action is copy from source to destination path as backup method
	// However, user can choose to restore state with optional preference
	restoreState := action == "restore"

	// Process each state
	for _, state := range GetStates(options) {

		// Check if should process this platform
		if !slices.Contains(options.Platforms, state.Platform) {
			continue
		}

		// Source are in multiple locations due to multiple runtimes and operating systems
		// To ensure compatibility, we process just the first valid location for source
		for _, source := range state.Source {

			// Fill source and destination information
			source := fs.ExpandPath(source)
			destination := fs.ExpandPath(state.Path)

			// When using restore method, invert path information
			if restoreState {
				source = fs.ExpandPath(state.Path)
				destination = fs.ExpandPath(source)
			}

			// Process file or folder state
			if state.Type == "file" {

				// Ensure file exist
				exist, err := fs.FileExist(source)
				if err != nil {
					return err
				} else if !exist {
					cli.Debug("Skipping file not detected: %s\n", source)
					continue
				}

				// Copy file
				cli.Printf(cli.ColorNotice, "Synchronizing file from %s to %s...\n", source, destination)
				err = fs.CopyFile(source, destination, true)
				if err != nil {
					return err
				}

			} else if state.Type == "folder" {

				// Ensure folder exist
				exist, err := fs.DirectoryExist(source)
				if err != nil {
					return err
				} else if !exist {
					cli.Debug("Skipping folder not detected: %s\n", source)
					continue
				}

				// Recursive copy content
				cli.Printf(cli.ColorNotice, "Synchronizing folder from %s to %s...\n", source, destination)
				err = fs.CopyDirectory(source, destination)
				if err != nil {
					return err
				}

			}

			// Ensure that only the first valid result will be processed
			break
		}
	}

	cli.Printf(cli.ColorNotice, "State synchronized.\n")
	return nil
}
