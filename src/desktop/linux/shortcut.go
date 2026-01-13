package linux

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Retrieve desktop entry shortcut path
func GetShortcutPath(shortcut *shortcuts.Shortcut) string {
	return fs.ExpandPath(fmt.Sprintf(
		"$SHARE/applications/%s.desktop",
		fs.NormalizeFilename(shortcut.Name),
	))
}

// Create a desktop entry from shortcut data
func CreateShortcut(shortcut *shortcuts.Shortcut, destination string, overwriteAssets bool) error {

	// Prepare execution context
	context := shortcuts.PrepareContext(shortcut)

	// Icon by default follows XDG icon resource name
	iconFile := filepath.Base(context.Executable)
	iconFile = strings.Replace(iconFile, filepath.Ext(iconFile), "", 1)

	// If available, we use current PNG icon from shortcut
	if strings.HasSuffix(shortcut.IconPath, ".png") {
		iconFile = shortcut.IconPath
	}

	// Map categories into closest values from XDG desktop entry spec
	categories := []string{}
	replaces := map[string]string{
		"Gaming":    "Game",
		"ROM":       "Game",
		"Emulator":  "Game",
		"Proton":    "Game",
		"Utilities": "Utility",
		"Streaming": "Network",
	}

	for _, category := range shortcut.Tags {
		if value, ok := replaces[category]; ok {
			category = value
		}
		if !slices.Contains(categories, category) {
			categories = append(categories, category)
		}
	}

	// Create desktop entry shortcut
	executable := strings.TrimSpace(fmt.Sprintf(
		"%s %s %s",
		strings.Join(context.Environment, " "),
		context.Executable,
		strings.Join(context.Arguments, " "),
	))

	entry := &DesktopEntry{
		HasTerminal: true,
		Terminal:    false,
		Type:        "Application",
		Name:        shortcut.Name,
		Comment:     shortcut.Description,
		Icon:        iconFile,
		Path:        context.WorkingDirectory,
		TryExec:     context.Executable,
		Exec:        executable,
		Categories:  categories,
	}

	// Write desktop shortcut
	err := WriteDesktopFile(destination, entry)
	if err != nil {
		return err
	}

	return nil
}

// Remove desktop entry shortcut
func RemoveShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Remove desktop entry file
	err := fs.RemoveFile(destination)
	if err != nil {
		return err
	}

	return nil
}
