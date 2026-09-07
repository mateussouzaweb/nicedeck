package launcher

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Parse game details based on the given reference
func Parse(reference string, options *Options) (*Game, error) {

	result := &Game{}

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
func Find(options *Options) ([]*Game, error) {

	var results []*Game

	cli.Printf(cli.ColorNotice, "Checking for games from launchers.\n")

	// Scan each platform for games
	for _, platform := range options.Platforms {
		games, err := platform.Scan()
		if err != nil {
			return nil, err
		}

		// Add game tag and append to final results
		for _, game := range games {
			game.Tags = append([]string{"Game"}, game.Tags...)
			results = append(results, game)
		}
	}

	return results, nil
}

// Filter games that match given requirements and return the list to process
func Filter(list []*Game, existing []string, options *Options) []*Game {

	var toProcess []*Game

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
