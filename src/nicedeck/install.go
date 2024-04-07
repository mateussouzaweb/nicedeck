package nicedeck

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Desktop install for NiceDeck
func DesktopInstall() error {

	// Check if executable exist first
	// If not exist, ignore self installer
	executable := os.ExpandEnv("$HOME/Applications/NiceDeck")
	exist, err := fs.FileExist(executable)
	if err != nil {
		return err
	} else if !exist {
		return nil
	}

	// Check if desktop shortcut exist or create one
	desktopShortcut := os.ExpandEnv("$HOME/.local/share/applications/com.mateussouzaweb.NiceDeck.desktop")
	exist, err = fs.FileExist(desktopShortcut)
	if err != nil {
		return err
	} else if !exist {
		return WriteDesktopShortcut()
	}

	return nil
}

// Write desktop shortcut for NiceDeck
func WriteDesktopShortcut() error {

	// Replace special variables
	replaceVars := func(content []byte) []byte {
		return bytes.ReplaceAll(content, []byte("$HOME"), []byte(os.Getenv("HOME")))
	}

	// Icon
	iconFile := os.ExpandEnv("$HOME/.local/share/nicedeck/nicedeck.svg")
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

	err = os.WriteFile(desktopShortcutFile, replaceVars(desktopShortcutContent), 0774)
	if err != nil {
		return err
	}

	return nil
}
