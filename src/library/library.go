package library

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

// var Desktop *desktop.Library
// var EpicGames *epic.Library
// var GOG *gog.Library
var Shortcuts = &shortcuts.Library{}
var Steam = &steam.Library{}

// Load library from config path
func Load() error {

	// Normalize path
	configPath := "$APPLICATIONS/NiceDeck"
	configPath = fs.ExpandPath(configPath)

	// Load shortcuts
	err := Shortcuts.Load(fmt.Sprintf("%s/shortcuts.json", configPath))
	if err != nil {
		return err
	}

	// Load Steam data
	err = Steam.Load(fmt.Sprintf("%s/steam.json", configPath))
	if err != nil {
		return err
	}

	// Sync Steam shortcuts to internal library
	// Please note that Steam data will have preference
	// Internal library should update data when necessary
	for _, steamShortcut := range Steam.Shortcuts {
		shortcut := Steam.ToInternal(steamShortcut)
		err := Shortcuts.Set(shortcut, false)
		if err != nil {
			return err
		}
	}

	// Desktop.Load()
	// EpicGames.Load()
	// GOG.Load()

	return nil
}

// Save library on config path
func Save() error {

	// Save shortcuts library
	err := Shortcuts.Save()
	if err != nil {
		return err
	}

	// Sync change history with Steam to update shortcuts
	for _, history := range Shortcuts.History {
		err := Steam.Sync(history)
		if err != nil {
			return err
		}
	}

	// Save Steam library
	err = Steam.Save()
	if err != nil {
		return err
	}

	// Desktop.Sync()
	// Desktop.Save()
	// EpicGames.Sync()
	// EpicGames.Save()
	// GOG.Sync()
	// GOG.Save()

	return nil
}
