package library

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// var Steam *steam.Library
// var Desktop *desktop.Library
// var EpicGames *epic.Library
// var GOG *gog.Library
var Shortcuts = &shortcuts.Library{}

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

	// Desktop.ReadShortcuts()
	// Steam.LoadProfile()
	// EpicGames.LoadProfile()
	// GOG.LoadProfile()

	return nil
}

// Save library on config path
func Save() error {

	err := Shortcuts.Save()
	if err != nil {
		return err
	}

	// Desktop.ReadShortcuts()
	// Steam.LoadProfile()
	// EpicGames.LoadProfile()
	// GOG.LoadProfile()

	return nil
}
