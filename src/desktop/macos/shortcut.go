package macos

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Retrieve desktop entry shortcut path
func GetShortcutPath(shortcut *shortcuts.Shortcut) string {
	return fs.ExpandPath(fmt.Sprintf(
		"$HOME/Applications/%s.app",
		fs.NormalizeFilename(shortcut.Name),
	))
}

// Create a desktop entry from shortcut data
func CreateShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Prepare execution context and bundle data
	context := shortcuts.PrepareContext(shortcut)
	bundle := &Bundle{
		AppName:          shortcut.Name,
		BundleID:         fmt.Sprintf("com.nicedeck.%s", shortcut.ID),
		Launcher:         "launcher",
		IconPath:         "",
		WorkingDirectory: context.WorkingDirectory,
		Executable:       context.Executable,
		Arguments:        context.Arguments,
		Environment:      context.Environment,
	}

	// If available, we use current PNG icon from shortcut
	if strings.HasSuffix(shortcut.IconPath, ".png") {
		bundle.IconPath = shortcut.IconPath
	}

	// Write .app bundle to destination
	err := WriteBundle(destination, bundle)
	if err != nil {
		return err
	}

	return nil
}

// Remove desktop entry shortcut
func RemoveShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Remove .app bundle directory
	err := fs.RemoveDirectory(destination)
	if err != nil {
		return err
	}

	return nil
}
