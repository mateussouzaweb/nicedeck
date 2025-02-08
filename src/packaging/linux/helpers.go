package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Write a desktop shortcut and return the file path
func WriteDesktopShortcut(appID string, shortcut *shortcuts.Shortcut) (string, error) {

	desktopShortcutName := fmt.Sprintf("%s.desktop", appID)
	desktopShortcutLocation := fs.ExpandPath("$HOME/.local/share/applications")
	desktopShortcutFile := filepath.Join(desktopShortcutLocation, desktopShortcutName)
	iconFile := appID

	// Icon by default follows XDG icon resource name
	// If possible, we download PNG icon from shortcut
	if strings.HasSuffix(shortcut.IconURL, ".png") {
		iconName := fmt.Sprintf("%s.png", appID)
		iconPath := fs.ExpandPath("$HOME/.local/share/icons")
		iconFile = filepath.Join(iconPath, iconName)

		err := fs.DownloadFile(shortcut.IconURL, iconFile, false)
		if err != nil {
			return desktopShortcutFile, err
		}
	}

	// Create and write desktop shortcut
	desktopShortcutContent := os.ExpandEnv(fmt.Sprintf(""+
		"[Desktop Entry]\n"+
		"Type=Application\n"+
		"Name=%s\n"+
		"Icon=%s\n"+
		"Exec=%s\n"+
		"Terminal=false\n"+
		"Categories=%s;",
		shortcut.AppName,
		iconFile,
		shortcut.Exe,
		shortcut.Tags[0],
	))

	err := os.MkdirAll(desktopShortcutLocation, 0700)
	if err != nil {
		return desktopShortcutFile, err
	}

	err = os.WriteFile(desktopShortcutFile, []byte(desktopShortcutContent), 0644)
	if err != nil {
		return desktopShortcutFile, err
	}

	return desktopShortcutFile, nil
}
