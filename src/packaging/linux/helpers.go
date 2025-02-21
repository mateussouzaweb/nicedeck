package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
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
	categories := strings.Join(shortcut.Tags, ";")
	categories = strings.ReplaceAll(categories, "Gaming", "Game")
	categories = strings.ReplaceAll(categories, "Utilities", "Utility")
	categories = strings.ReplaceAll(categories, "Streaming", "Network")

	// Create and write desktop shortcut
	desktopShortcutContent := os.ExpandEnv(fmt.Sprintf(""+
		"[Desktop Entry]\n"+
		"Type=Application\n"+
		"Name=%s\n"+
		"Comment=%s\n"+
		"Icon=%s\n"+
		"Exec=%s %s\n"+
		"Terminal=false\n"+
		"Categories=%s;",
		shortcut.AppName,
		shortcut.Description,
		iconFile,
		shortcut.Exe,
		shortcut.LaunchOptions,
		categories,
	))

	err := os.MkdirAll(filepath.Dir(shortcut.ShortcutPath), 0700)
	if err != nil {
		return err
	}

	err = os.WriteFile(shortcut.ShortcutPath, []byte(desktopShortcutContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
