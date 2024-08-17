package platforms

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// State struct
type State struct {
	Platform    string `json:"platform"`
	Emulator    string `json:"emulator"`
	Type        string `json:"type"`
	Source      string `json:"path"`
	Destination string `json:"destination"`
}

// Retrieve save state of each platform
func GetStates(options *Options) []*State {

	states := []*State{}

	states = append(states, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "folder",
		Source:      "$HOME/.config/Ryujinx/bis/user/save",
		Destination: "$HOME/Games/State/Ryujinx/bis/user/save",
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "folder",
		Source:      "$HOME/.config/Ryujinx/bis/user/saveMeta",
		Destination: "$HOME/Games/State/Ryujinx/bis/user/saveMeta",
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "file",
		Source:      "$HOME/.config/Ryujinx/system/Profiles.json",
		Destination: "$HOME/Games/State/Ryujinx/system/Profiles.json",
	})

	return states
}

// Sync state of each platform
func SyncState(options *Options) error {

	// Default action is copy from source to destination as backup method
	// However, user can choose to restore state with optional preference
	restoreState := slices.Contains(options.Preferences, "restore-state")

	// Process each state
	for _, state := range GetStates(options) {

		// Check if should process this platform
		if !slices.Contains(options.Platforms, state.Platform) {
			continue
		}

		// Check if should process this emulator
		// if len(options.Emulators) > 0 {
		// 	if !slices.Contains(options.Emulators, state.Emulator) {
		// 		continue
		// 	}
		// }

		// Fill source and destination information
		source := state.Source
		destination := state.Destination

		if restoreState {
			source = state.Destination
			destination = state.Source
		}

		// Process file or folder state
		if state.Type == "file" {

			// Ensure file exist
			exist, err := fs.FileExist(source)
			if err != nil {
				return err
			} else if !exist {
				cli.Printf(cli.ColorNotice, "Skipping file not detected: %s\n", source)
				continue
			}

			// Copy file
			cli.Printf(cli.ColorNotice, "Synchronizing file from %s to %s...\n", source, destination)
			err = fs.CopyFile(source, destination)
			if err != nil {
				return err
			}

		} else if state.Type == "folder" {

			// Ensure folder exist
			exist, err := fs.DirectoryExist(source)
			if err != nil {
				return err
			} else if !exist {
				cli.Printf(cli.ColorNotice, "Skipping folder not detected: %s\n", source)
				continue
			}

			// Recursive copy content
			cli.Printf(cli.ColorNotice, "Synchronizing folder from %s to %s...\n", source, destination)
			err = fs.CopyDirectory(source, destination)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
