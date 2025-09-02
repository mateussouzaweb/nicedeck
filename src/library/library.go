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

	// Sync additional libraries
	err = Sync()
	if err != nil {
		return err
	}

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

	// Save ES-DE data
	err = ESDE.Save()
	if err != nil {
		return err
	}

	// Desktop.Sync()
	// Desktop.Save()
	// EpicGames.Sync()
	// EpicGames.Save()
	// GOG.Sync()
	// GOG.Save()

	// Clean history of changes
	Shortcuts.History = Shortcuts.History[:0]

	return nil
}

// Sync additional libraries into internal library to add or update entries
func Sync() error {

	// Steam shortcuts
	for _, steamShortcut := range Steam.Shortcuts {
		shortcut := Steam.ToInternal(steamShortcut)
		existing := Shortcuts.Get(shortcut.ID)

		if existing.ID != "" {
			existing.Merge(shortcut)
			shortcut = existing
		}

		err := Shortcuts.Set(shortcut, false)
		if err != nil {
			return err
		}
	}

	return nil
}
