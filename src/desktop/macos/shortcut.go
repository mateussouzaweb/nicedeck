package macos

import (
	"fmt"
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Retrieve desktop entry shortcut path
func GetShortcutPath(shortcut *shortcuts.Shortcut) string {

	if slices.Contains(shortcut.Tags, "Gaming") {
		return fs.ExpandPath(fmt.Sprintf(
			"$HOME/Applications/Gaming/%s.app",
			shortcut.Name,
		))
	}

	return fs.ExpandPath(fmt.Sprintf(
		"$HOME/Applications/%s.app",
		shortcut.Name,
	))
}

// Create a desktop entry from shortcut data
func CreateShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Make link to main executable
	err := fs.MakeSymlink(shortcut.Executable, destination)
	if err != nil {
		return err
	}

	return nil
}

// Remove desktop entry shortcut
func RemoveShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	err := fs.RemoveSymlink(destination)
	if err != nil {
		return err
	}

	return nil
}
