package nicedeck

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write desktop shortcut for NiceDeck
func WriteDesktopShortcut() error {

	// Check if is running via flatpak
	if cli.GetEnv("FLATPAK_ID", "") != "" {
		cli.Printf(cli.ColorWarn, "NiceDeck is running via Flatpak\n")
		cli.Printf(cli.ColorWarn, "Cannot install desktop shortcut. Skipping...\n")
		return nil
	}

	// Retrieve executable file
	executableFile, err := os.Executable()
	if err != nil {
		return err
	}

	// Icon
	iconFile := os.ExpandEnv("$HOME/.local/share/icons/hicolor/scalable/apps/nicedeck.svg")
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
	desktopShortcutFile := os.ExpandEnv("$HOME/.local/share/applications/nicedeck.desktop")
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
