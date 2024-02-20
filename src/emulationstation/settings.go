package emulationstation

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write settings for EmulationStation
func WriteSettings() error {

	// Replace special variables
	replaceVars := func(content []byte) []byte {
		return bytes.ReplaceAll(content, []byte("$HOME"), []byte(os.Getenv("HOME")))
	}

	// Settings (write file only if not exist yet)
	settingsFile := os.ExpandEnv("$HOME/ES-DE/settings/es_settings.xml")
	settingsExist, err := fs.FileExist(settingsFile)
	if err != nil {
		return err
	}

	// Create file if not exist
	if !settingsExist {
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
	}

	// Systems
	systemsFile := os.ExpandEnv("$HOME/ES-DE/custom_systems/es_systems.xml")
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

	// Find Rules
	findRulesFile := os.ExpandEnv("$HOME/ES-DE/custom_systems/es_find_rules.xml")
	findRulesConfig, err := resourcesContent.ReadFile("resources/es_find_rules.xml")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(findRulesFile), 0774)
	if err != nil {
		return err
	}

	err = os.WriteFile(findRulesFile, replaceVars(findRulesConfig), 0666)
	if err != nil {
		return err
	}

	// Icon
	iconFile := os.ExpandEnv("$HOME/ES-DE/icon.png")
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
