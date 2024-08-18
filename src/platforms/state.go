package platforms

import (
	"os"
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
		Destination: "$STATE/Ryujinx/bis/user/save",
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "folder",
		Source:      "$HOME/.config/Ryujinx/bis/user/saveMeta",
		Destination: "$STATE/Ryujinx/bis/user/saveMeta",
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "file",
		Source:      "$HOME/.config/Ryujinx/system/Profiles.json",
		Destination: "$STATE/Ryujinx/system/Profiles.json",
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

		// Fill source and destination information
		source := os.ExpandEnv(state.Source)
		destination := os.ExpandEnv(state.Destination)

		if restoreState {
			source = os.ExpandEnv(state.Destination)
			destination = os.ExpandEnv(state.Source)
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
