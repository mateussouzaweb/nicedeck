package library

import (
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/esde"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

// var Desktop *desktop.Library
// var EpicGames *epic.Library
// var GOG *gog.Library
var Shortcuts = &shortcuts.Library{}
var Steam = &steam.Library{}
var ESDE = &esde.Library{}

// Load library from config path
func Load() error {

	// Normalize path
	configPath := filepath.Join(fs.ExpandPath("$APPLICATIONS"), "NiceDeck")
	shortcutsConfigPath := filepath.Join(configPath, "shortcuts.json")
	steamConfigPath := filepath.Join(configPath, "steam.json")

	// Load shortcuts
	err := Shortcuts.Load(shortcutsConfigPath)
	if err != nil {
		return err
	}

	// Load Steam data
	err = Steam.Load(steamConfigPath)
	if err != nil {
		return err
	}

	// Load ES-DE data
	err = ESDE.Load()
	if err != nil {
		return err
	}

	// Desktop.Load()
	// EpicGames.Load()
	// GOG.Load()

	// Sync libraries
	err = Sync()
	if err != nil {
		return err
	}

	return nil
}

// Save library on config path
func Save() error {

	// Sync libraries
	err := Sync()
	if err != nil {
		return err
	}

	// Save shortcuts library
	err = Shortcuts.Save()
	if err != nil {
		return err
	}

	// Save Steam library
	err = Steam.Save()
	if err != nil {
		return err
	}

	// Save ES-DE library
	err = ESDE.Save()
	if err != nil {
		return err
	}

	// Desktop.Save()
	// EpicGames.Save()
	// GOG.Save()

	return nil
}

// Sync libraries to add, update or remove entries
func Sync() error {

	// Sync Steam library
	err := Steam.Sync(Shortcuts)
	if err != nil {
		return err
	}

	return nil
}
