package library

import (
	"path/filepath"
	"time"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/desktop"
	"github.com/mateussouzaweb/nicedeck/src/esde"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

// Shortcut alias
type Shortcut = shortcuts.Shortcut

// Diff struct
type Diff struct {
	Added    []*Shortcut
	Removed  []*Shortcut
	Updated  []*Shortcut
	Existing []*Shortcut
}

// Synchronizable interface
type Synchronizable interface {
	String() string
	Load() error
	Save() error
	Export() []*Shortcut
	Add(*Shortcut) error
	Update(*Shortcut, bool) error
	Remove(*Shortcut) error
}

// Global shortcuts library
var Shortcuts = &shortcuts.Library{}
var Steam = &steam.Library{}
var ESDE = &esde.Library{}
var Desktop = &desktop.Library{}

// var EpicGames = &epic.Library{}
// var GOG = &gog.Library{}

// Load library from config path
func Load() error {

	// Normalize path
	configPath := filepath.Join(fs.ExpandPath("$APPLICATIONS"), "NiceDeck")
	shortcutsConfigPath := filepath.Join(configPath, "shortcuts.json")
	steamConfigPath := filepath.Join(configPath, "steam.json")
	esdeConfigPath := filepath.Join(configPath, "esde.json")
	desktopConfigPath := filepath.Join(configPath, "desktop.json")

	// Init shortcuts library
	err := Shortcuts.Init(shortcutsConfigPath)
	if err != nil {
		return err
	}

	// Init Steam library
	err = Steam.Init(steamConfigPath)
	if err != nil {
		return err
	}

	// Init ESDE library
	err = ESDE.Init(esdeConfigPath)
	if err != nil {
		return err
	}

	// Init Desktop library
	err = Desktop.Init(desktopConfigPath)
	if err != nil {
		return err
	}

	// Load shortcuts
	err = Shortcuts.Load()
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

	// Check if there are recent changes to synchronize
	if len(Shortcuts.History) == 0 {
		cli.Debug("No recent changes to synchronize\n")
		return nil
	}

	// Perform synchronization to additional libraries after saving main library
	// This happens only on this context as one-way sync
	libraries := make([]Synchronizable, 0)
	libraries = append(libraries, Steam)
	libraries = append(libraries, ESDE)
	libraries = append(libraries, Desktop)
	// libraries = append(libraries, EpicGames)
	// libraries = append(libraries, GOG)

	for _, library := range libraries {

		cli.Debug("Synchronizing recent changes from %s to library\n", library.String())

		// Load library data
		err := library.Load()
		if err != nil {
			return err
		}

		// Perform synchronization based on recent history only
		for _, history := range Shortcuts.History {
			if history.Action == "added" {
				err := library.Add(history.Result)
				if err != nil {
					return err
				}
			}
			if history.Action == "updated" {
				err := library.Update(history.Result, true)
				if err != nil {
					return err
				}
			}
			if history.Action == "removed" {
				err := library.Remove(history.Original)
				if err != nil {
					return err
				}
			}
		}

		// Save library
		err = library.Save()
		if err != nil {
			return err
		}

	}

	// Clear history after synchronization
	Shortcuts.Reset()

	return nil
}

// Fill missing details when shortcut is provided as ID only
func Fill(list []*Shortcut) []*Shortcut {

	for index, shortcut := range list {
		if shortcut.ID != "" && shortcut.Name == "" {
			found := Shortcuts.Get(shortcut.ID)
			if found.ID == shortcut.ID {
				list[index] = found
			}
		}
	}

	return list
}

// Compare two shortcut lists and return the differences
func Compare(current []*Shortcut, compare []*Shortcut) Diff {

	// Compare with maps for fast lookup O(1)
	currentMap := make(map[string]bool)
	compareMap := make(map[string]bool)

	// Fill lookup tables
	for _, shortcut := range current {
		currentMap[shortcut.ID] = true
	}
	for _, shortcut := range compare {
		compareMap[shortcut.ID] = true
	}

	// Prepare diff slices
	existing := make([]*Shortcut, 0)
	added := make([]*Shortcut, 0)
	updated := make([]*Shortcut, 0)
	removed := make([]*Shortcut, 0)

	// Detect removed and existing entries
	for _, shortcut := range current {
		if _, ok := compareMap[shortcut.ID]; ok {
			existing = append(existing, shortcut)
		} else {
			removed = append(removed, shortcut)
		}
	}

	// Detect added entries
	for _, shortcut := range compare {
		if _, ok := currentMap[shortcut.ID]; !ok {
			added = append(added, shortcut)
		}
	}

	// Helper to get comparable shortcut by ID
	getCompareById := func(ID string) *Shortcut {
		for _, item := range compare {
			if item.ID == ID {
				return item
			}
		}
		return &Shortcut{}
	}

	// Detect update entries
	// Compare based on modification timestamp
	for _, shortcut := range existing {

		// Get comparable entry
		comparable := getCompareById(shortcut.ID)
		if comparable.ID == "" {
			continue
		}

		// Comparable entry must have at least one minute diff between timestamps
		// In such case, entry is considered as newest entry reference
		if shortcut.Timestamp >= comparable.Timestamp {
			continue
		}
		if shortcut.Timestamp+60 >= comparable.Timestamp {
			continue
		}

		cli.Debug(
			"Found update on %s: %s vs %s\n",
			shortcut.ID,
			time.Unix(shortcut.Timestamp, 0).String(),
			time.Unix(comparable.Timestamp, 0).String(),
		)

		// Merge information and mark as updated entry
		shortcut.Merge(comparable)
		updated = append(updated, shortcut)

	}

	if len(added) == 0 && len(removed) == 0 && len(updated) == 0 {
		cli.Debug("Diff: no differences found\n")
	} else {
		cli.Debug("Diff: added: %d / updated: %d / removed: %d\n", len(added), len(updated), len(removed))
	}

	return Diff{
		Added:    added,
		Removed:  removed,
		Updated:  updated,
		Existing: existing,
	}
}

// Sync libraries to add, update or remove entries
func Sync() error {

	libraries := make([]Synchronizable, 0)
	libraries = append(libraries, Steam)
	libraries = append(libraries, ESDE)
	libraries = append(libraries, Desktop)
	// libraries = append(libraries, EpicGames)
	// libraries = append(libraries, GOG)

	// Perform synchronization process to main shortcuts library
	// Please note that shortcuts library is already loaded
	// Load and compare libraries to find differences
	// Then, apply differences to main shortcuts library
	// Process is done gradually for each additional library
	for _, library := range libraries {

		cli.Debug("Synchronizing %s to library\n", library.String())

		// Load library data
		err := library.Load()
		if err != nil {
			return err
		}

		// Export library shortcuts to internal format
		// Avoid processing empty libraries
		exported := Fill(library.Export())
		if len(exported) == 0 {
			continue
		}

		// Compare library with main shortcuts library
		diff := Compare(Shortcuts.All(), exported)

		// Apply differences to main shortcuts library
		for _, shortcut := range diff.Added {
			err := Shortcuts.Add(shortcut)
			if err != nil {
				return err
			}
		}
		for _, shortcut := range diff.Updated {
			err := Shortcuts.Update(shortcut, false)
			if err != nil {
				return err
			}
		}
		for _, shortcut := range diff.Removed {
			err := Shortcuts.Remove(shortcut)
			if err != nil {
				return err
			}
		}

	}

	// Clear history after synchronization
	Shortcuts.Reset()

	// Perform synchronization process to additional libraries
	// At this stage, main library is full synchronized
	// We now find and apply differences to each additional library
	for _, library := range libraries {

		cli.Debug("Synchronizing library to %s\n", library.String())

		// Compare library with main shortcuts library
		current := Fill(library.Export())
		diff := Compare(current, Shortcuts.All())

		// Apply differences to additional library
		for _, shortcut := range diff.Added {
			err := library.Add(shortcut)
			if err != nil {
				return err
			}
		}
		for _, shortcut := range diff.Updated {
			err := library.Update(shortcut, true)
			if err != nil {
				return err
			}
		}
		for _, shortcut := range diff.Removed {
			err := library.Remove(shortcut)
			if err != nil {
				return err
			}
		}

		// Save library
		err := library.Save()
		if err != nil {
			return err
		}

	}

	return nil
}
