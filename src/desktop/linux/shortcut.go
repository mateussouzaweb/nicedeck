package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Retrieve desktop entry shortcut path
func GetShortcutPath(shortcut *shortcuts.Shortcut) string {

	// Transform name into slug format
	pattern := regexp.MustCompile("[^a-z0-9]+")
	slug := strings.ToLower(shortcut.Name)
	slug = strings.Trim(pattern.ReplaceAllString(slug, "-"), "-")

	return fs.ExpandPath(fmt.Sprintf(
		"$SHARE/applications/%s.desktop",
		slug,
	))
}

// Create a desktop entry from shortcut data
func CreateShortcut(shortcut *shortcuts.Shortcut, destination string, overwriteAssets bool) error {

	// Icon by default follows XDG icon resource name
	iconFile := filepath.Base(destination)
	iconFile = strings.Replace(iconFile, ".desktop", "", 1)

	// If possible, we copy PNG icon from shortcut
	if strings.HasSuffix(shortcut.IconPath, ".png") {
		iconName := fmt.Sprintf("%s.png", iconFile)
		iconPath := fs.ExpandPath("$SHARE/icons")
		iconFile = filepath.Join(iconPath, iconName)

		err := fs.CopyFile(shortcut.IconPath, iconFile, overwriteAssets)
		if err != nil {
			return err
		}
	}

	// Map categories into closest values from desktop menu spec
	categories := shortcut.Tags
	replaces := map[string]string{
		"Gaming":    "Game",
		"Utilities": "Utility",
		"Streaming": "Network",
	}

	for index, category := range categories {
		if value, ok := replaces[category]; ok {
			categories[index] = value
		}
	}

	// Create and write desktop entry shortcut
	executable := fmt.Sprintf(
		"%s %s",
		os.ExpandEnv(shortcut.Executable),
		os.ExpandEnv(shortcut.LaunchOptions),
	)

	entry := &DesktopEntry{
		Terminal:   false,
		Type:       "Application",
		Name:       shortcut.Name,
		Comment:    shortcut.Description,
		Icon:       iconFile,
		Exec:       executable,
		Categories: categories,
	}

	// Create and write desktop shortcut
	err := WriteDesktopFile(destination, entry)
	if err != nil {
		return err
	}

	return nil
}

// Remove desktop entry shortcut
func RemoveShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Remove possible stored icon file
	iconPath := fs.ExpandPath("$SHARE/icons")
	iconFile := filepath.Base(destination)
	iconFile = strings.Replace(iconFile, ".desktop", "", 1)
	iconFile = fmt.Sprintf("%s.png", iconFile)
	iconFile = filepath.Join(iconPath, iconFile)

	err := fs.RemoveFile(iconFile)
	if err != nil {
		return err
	}

	// Remove desktop entry file
	err = fs.RemoveFile(destination)
	if err != nil {
		return err
	}

	return nil
}
