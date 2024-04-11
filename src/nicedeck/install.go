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

// Check if program is running inside flatpak
func IsRunningInsideFlatpak() bool {
	return os.Getenv("FLATPAK_ID") == "com.mateussouzaweb.NiceDeck"
}

// Write desktop shortcut for NiceDeck
func WriteDesktopShortcut() error {

	// Check if is running via flatpak
	if IsRunningInsideFlatpak() {
		cli.Printf(cli.ColorWarn, "NiceDeck is running via Flatpak\n")
		cli.Printf(cli.ColorWarn, "Cannot install desktop shortcut. Skipping...\n")
	}

	// Retrieve executable file
	executableFile, err := os.Executable()
	if err != nil {
		return err
	}

	// Replace special variables
	replaceVars := func(content []byte) []byte {
		return bytes.ReplaceAll(content, []byte("$HOME"), []byte(os.Getenv("HOME")))
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
	desktopShortcutFile := os.ExpandEnv("$HOME/.local/share/applications/com.mateussouzaweb.NiceDeck.desktop")
	desktopShortcutContent, err := resourcesContent.ReadFile("resources/com.mateussouzaweb.NiceDeck.desktop")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(desktopShortcutFile), 0774)
	if err != nil {
		return err
	}

	// Match executable with real location
	desktopShortcutContent = replaceVars(desktopShortcutContent)
	desktopShortcutContent = bytes.ReplaceAll(
		desktopShortcutContent,
		[]byte("Exec=nicedeck"),
		[]byte("Exec="+executableFile),
	)

	err = os.WriteFile(desktopShortcutFile, replaceVars(desktopShortcutContent), 0774)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Desktop shortcut created at: %s\n", desktopShortcutFile)

	return nil
}
