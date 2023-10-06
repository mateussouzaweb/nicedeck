package emulationstation

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"
)

//go:embed resources/*
var resourcesContent embed.FS

func WriteConfigs() error {

	// Replace special variables
	replaceVars := func(content []byte) []byte {
		return bytes.ReplaceAll(content, []byte("$HOME"), []byte(os.Getenv("HOME")))
	}

	// Settings
	settingsFile := os.ExpandEnv("$HOME/.emulationstation/es_settings.xml")
	settingsConfig, err := resourcesContent.ReadFile("resources/es_settings.xml")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(settingsFile), 0774)
	if err != nil {
		return err
	}

	err = os.WriteFile(settingsFile, replaceVars(settingsConfig), 0666)
	if err != nil {
		return err
	}

	// Systems
	systemsFile := os.ExpandEnv("$HOME/.emulationstation/custom_systems/es_systems.xml")
	systemsConfig, err := resourcesContent.ReadFile("resources/es_systems.xml")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(systemsFile), 0774)
	if err != nil {
		return err
	}

	err = os.WriteFile(systemsFile, replaceVars(systemsConfig), 0666)
	if err != nil {
		return err
	}

	// Icon
	iconFile := os.ExpandEnv("$HOME/.emulationstation/icon.png")
	iconContent, err := resourcesContent.ReadFile("resources/icon.png")
	if err != nil {
		return err
	}

	err = os.WriteFile(iconFile, iconContent, 0666)
	if err != nil {
		return err
	}

	// Desktop shortcut
	desktopShortcutFile := os.ExpandEnv("$HOME/.local/share/applications/emulationstation-de.desktop")
	desktopShortcutContent, err := resourcesContent.ReadFile("resources/emulationstation-de.desktop")
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
