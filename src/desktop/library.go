package desktop

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/desktop/linux"
	"github.com/mateussouzaweb/nicedeck/src/desktop/macos"
	"github.com/mateussouzaweb/nicedeck/src/desktop/windows"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Shortcut alias
type Shortcut = shortcuts.Shortcut

// Library struct
type Library struct {
	DatabasePath string            `json:"databasePath"`
	References   map[string]string `json:"references"`
}

// String representation of the library
func (l *Library) String() string {
	return "Desktop"
}

// Init library
func (l *Library) Init(databasePath string) error {
	l.DatabasePath = databasePath
	return nil
}

// Load library
func (l *Library) Load() error {

	// Reset and fill basic information
	l.References = make(map[string]string, 0)

	// Read database file content
	err := fs.ReadJSON(l.DatabasePath, &l)
	if err != nil {
		return err
	}

	return nil
}

// Save library
func (l *Library) Save() error {

	// Save database state to file
	err := fs.WriteJSON(l.DatabasePath, l)
	if err != nil {
		return err
	}

	return nil
}

// Export shortcuts to internal format
func (l *Library) Export() []*Shortcut {
	results := make([]*Shortcut, 0)
	for id := range l.References {
		results = append(results, &Shortcut{ID: id})
	}

	return results
}

// Add shortcut to the library
func (l *Library) Add(shortcut *Shortcut) error {

	// Skip if not tagged for desktop but keep reference
	if !slices.Contains(shortcut.Tags, "Desktop") {
		l.References[shortcut.ID] = ""
		return nil
	}

	// Update shortcut if already present
	if l.References[shortcut.ID] != "" {
		return l.Update(shortcut, true)
	}

	cli.Debug("Adding shortcut to desktop: %s\n", shortcut.ID)

	// Write system shortcut
	var err error
	var path string

	if cli.IsLinux() {
		path = linux.GetShortcutPath(shortcut)
		err = linux.CreateShortcut(shortcut, path, false)
	} else if cli.IsMacOS() {
		path = macos.GetShortcutPath(shortcut)
		err = macos.CreateShortcut(shortcut, path)
	} else if cli.IsWindows() {
		path = windows.GetShortcutPath(shortcut)
		err = windows.CreateShortcut(shortcut, path)
	}

	if err != nil {
		return err
	} else if path != "" {
		l.References[shortcut.ID] = path
	}

	return nil
}

// Update shortcut on library
func (l *Library) Update(shortcut *Shortcut, overwriteAssets bool) error {

	// Remove shortcut if not tagged for desktop
	if !slices.Contains(shortcut.Tags, "Desktop") {
		return l.Remove(shortcut)
	}

	// Add shortcut if not present yet
	if l.References[shortcut.ID] == "" {
		return l.Add(shortcut)
	}

	// Retrieve desired shortcut path
	var path string
	if cli.IsLinux() {
		path = linux.GetShortcutPath(shortcut)
	} else if cli.IsMacOS() {
		path = macos.GetShortcutPath(shortcut)
	} else if cli.IsWindows() {
		path = windows.GetShortcutPath(shortcut)
	}

	// Replace existing shortcut if path as been changed
	if l.References[shortcut.ID] != path {
		cli.Debug("Replacing shortcut on desktop: %s\n", shortcut.ID)
		err := l.Remove(shortcut)
		if err != nil {
			return err
		}

		return l.Add(shortcut)
	}

	// Update reference
	cli.Debug("Updating shortcut on desktop: %s\n", shortcut.ID)
	l.References[shortcut.ID] = path

	// Write system shortcut
	var err error
	if cli.IsLinux() {
		err = linux.CreateShortcut(shortcut, path, true)
	} else if cli.IsMacOS() {
		err = macos.CreateShortcut(shortcut, path)
	} else if cli.IsWindows() {
		err = windows.CreateShortcut(shortcut, path)
	}
	if err != nil {
		return err
	}

	return nil
}

// Remove shortcut from the library
func (l *Library) Remove(shortcut *Shortcut) error {

	// Skip when empty reference
	if l.References[shortcut.ID] == "" {
		delete(l.References, shortcut.ID)
		return nil
	}

	cli.Debug("Removing shortcut from desktop: %s\n", shortcut.ID)

	// Remove system shortcut
	var err error
	path := l.References[shortcut.ID]

	if cli.IsLinux() {
		err = linux.RemoveShortcut(shortcut, path)
	} else if cli.IsMacOS() {
		err = macos.RemoveShortcut(shortcut, path)
	} else if cli.IsWindows() {
		err = windows.RemoveShortcut(shortcut, path)
	}
	if err != nil {
		return err
	}

	// Remove reference
	delete(l.References, shortcut.ID)
	return nil
}
