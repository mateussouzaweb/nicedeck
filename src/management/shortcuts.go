package management

import (
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Retrieve all shortcuts in the library
func GetShortcuts() []*shortcuts.Shortcut {
	return library.Shortcuts.All()
}

// Retrieve shortcut in the library with given ID
func GetShortcut(ID string) *shortcuts.Shortcut {
	return library.Shortcuts.Get(ID)
}

// Find shortcut with given name and executable combination
func FindShortcut(name string, executable string) *shortcuts.Shortcut {
	return library.Shortcuts.Find(name, executable)
}

// Launch shortcut
func LaunchShortcut(shortcut *shortcuts.Shortcut) error {
	return library.Shortcuts.Launch(shortcut)
}

// Set shortcut into library by adding or updating it
func SetShortcut(shortcut *shortcuts.Shortcut, overwriteAssets bool) error {
	return library.Shortcuts.Set(shortcut, overwriteAssets)
}

// Update shortcut on library
func UpdateShortcut(shortcut *shortcuts.Shortcut, overwriteAssets bool) error {
	return library.Shortcuts.Update(shortcut, overwriteAssets)
}

// Remove shortcut from the library
func RemoveShortcut(shortcut *shortcuts.Shortcut) error {
	return library.Shortcuts.Remove(shortcut)
}

// Retrieve path for shortcut images
func GetShortcutsImagesPath() string {
	return library.Shortcuts.ImagesPath
}
