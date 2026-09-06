package launcher

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/platforms/launcher/game"
)

// Parse game details based on the given reference
func Parse(reference string, options *Options) (*game.Game, error) {

	result := &game.Game{}

	// Find games in launchers and return the list of detected games
	games, err := Find(options)
	if err != nil {
		return result, err
	}

	// Check if the given reference matches any of the detected ROMs
	for _, item := range games {
		if item.Reference == reference {
			return item, nil
		}
		if item.Name == reference {
			return item, nil
		}
	}

	return result, nil
}

// Find games in launchers and return the list of detected games
func Find(options *Options) ([]*game.Game, error) {

	var results []*game.Game

	cli.Printf(cli.ColorNotice, "Checking for games from launchers.\n")

	for _, platform := range options.Platforms {

		// Skip not included platforms to gain performance
		if !slices.Contains(options.Include, platform.Name) {
			continue
		}

		// Scan the platform for games
		games, err := platform.Scan()
		if err != nil {
			return nil, err
		}

		// Append the detected games to the results list
		results = append(results, games...)

	}

	return results, nil
}

// Filter games that match given requirements and return the list to process
func Filter(list []*game.Game, existing []string, options *Options) []*game.Game {

	var toProcess []*game.Game

	// Rebuild option will include every ROM in the platform
	rebuild := slices.Contains(options.Preferences, "rebuild")

	// Fill the list of ROMs to process
	for _, item := range list {
		addToList := false

		// Add to the list if ROM matches platform condition
		if len(options.Platforms) == 0 {
			addToList = true
		} else if slices.Contains(options.Include, item.Platform) {
			addToList = true
		} else {
			cli.Debug("No platform match. Game skipped to the process list: %s\n", item.Reference)
		}

		// When is not rebuilding, include only new detected ROMs
		if !rebuild {
			if slices.Contains(existing, item.Reference) {
				cli.Debug("Existing. Game skipped to the process list: %s\n", item.Reference)
				addToList = false
			}
		}

		// Finally, if valid, add to the list of ROMs to process
		if addToList {
			cli.Debug("Game added to the process list: %s\n", item.Reference)
			toProcess = append(toProcess, item)
		}
	}

	return toProcess
}
