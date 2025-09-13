package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/desktop"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Create a desktop shortcut
func CreateDesktopShortcut(shortcut *shortcuts.Shortcut) error {

	// Icon by default follows XDG icon resource name
	// If possible, we download PNG icon from shortcut
	iconFile := filepath.Base(shortcut.ShortcutPath)
	iconFile = strings.Replace(iconFile, ".desktop", "", 1)

	if strings.HasSuffix(shortcut.IconURL, ".png") {
		iconName := fmt.Sprintf("%s.png", iconFile)
		iconPath := fs.ExpandPath("$SHARE/icons")
		iconFile = filepath.Join(iconPath, iconName)

		err := fs.DownloadFile(shortcut.IconURL, iconFile, false)
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
	entry := &desktop.DesktopEntry{
		Terminal:   false,
		Type:       "Application",
		Name:       shortcut.Name,
		Comment:    shortcut.Description,
		Icon:       iconFile,
		Exec:       executable,
		Categories: categories,
	}

	// Create and write desktop shortcut
	err := desktop.WriteDesktopFile(shortcut.ShortcutPath, entry)
	if err != nil {
		return err
	}

	return nil
}
