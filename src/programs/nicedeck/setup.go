package nicedeck

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Perform setup for NiceDeck
func Setup() error {
	if cli.IsLinux() {
		return WriteLinuxDesktopShortcut()
	}
	return nil
}

// Write desktop shortcut for NiceDeck
func WriteLinuxDesktopShortcut() error {

	// Retrieve executable file
	executableFile, err := os.Executable()
	if err != nil {
		return err
	}

	// Icon
	iconFile := fs.ExpandPath("$HOME/.local/share/icons/hicolor/scalable/apps/nicedeck.svg")
	iconContent, err := resourcesContent.ReadFile("resources/nicedeck.svg")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(iconFile), 0774)
	if err != nil {
		return err
	}

	err = os.WriteFile(iconFile, iconContent, 0666)
	if err != nil {
		return err
	}

	// Desktop shortcut
	desktopShortcutFile := fs.ExpandPath("$HOME/.local/share/applications/nicedeck.desktop")
	desktopShortcutContent, err := resourcesContent.ReadFile("resources/nicedeck.desktop")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(desktopShortcutFile), 0700)
	if err != nil {
		return err
	}

	// Match executable with real location
	desktopShortcutContent = bytes.ReplaceAll(
		desktopShortcutContent,
		[]byte("Exec=nicedeck"),
		[]byte("Exec="+executableFile),
	)

	err = os.WriteFile(desktopShortcutFile, desktopShortcutContent, 0644)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Desktop shortcut created at: %s\n", desktopShortcutFile)

	return nil
}
