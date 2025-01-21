package emulationstation

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write settings for EmulationStation
func WriteSettings() error {

	// Settings (write file only if not exist yet)
	settingsFile := fs.ExpandPath("$HOME/ES-DE/settings/es_settings.xml")
	settingsExist, err := fs.FileExist(settingsFile)
	if err != nil {
		return err
	}

	// Create file if not exist
	if !settingsExist {
		settingsConfigSource := fs.NormalizePath("resources/es_settings.xml")
		settingsConfig, err := resourcesContent.ReadFile(settingsConfigSource)
		if err != nil {
			return err
		}

		err = os.MkdirAll(filepath.Dir(settingsFile), 0774)
		if err != nil {
			return err
		}

		settingsConfig = []byte(os.ExpandEnv(string(settingsConfig)))
		err = os.WriteFile(settingsFile, settingsConfig, 0666)
		if err != nil {
			return err
		}
	}

	// Systems
	systemsFile := fs.ExpandPath("$HOME/ES-DE/custom_systems/es_systems.xml")
	systemsConfigSource := fs.NormalizePath("resources/es_systems.xml")
	systemsConfig, err := resourcesContent.ReadFile(systemsConfigSource)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(systemsFile), 0774)
	if err != nil {
		return err
	}

	err = os.WriteFile(systemsFile, systemsConfig, 0666)
	if err != nil {
		return err
	}

	// Find Rules
	findRulesFile := fs.ExpandPath("$HOME/ES-DE/custom_systems/es_find_rules.xml")
	findRulesConfigSource := fs.NormalizePath("resources/es_find_rules.xml")
	findRulesConfig, err := resourcesContent.ReadFile(findRulesConfigSource)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(findRulesFile), 0774)
	if err != nil {
		return err
	}

	err = os.WriteFile(findRulesFile, findRulesConfig, 0666)
	if err != nil {
		return err
	}

	// Icon
	iconFile := fs.ExpandPath("$HOME/ES-DE/icon.png")
	iconContentSource := fs.NormalizePath("resources/icon.png")
	iconContent, err := resourcesContent.ReadFile(iconContentSource)
	if err != nil {
		return err
	}

	err = os.WriteFile(iconFile, iconContent, 0666)
	if err != nil {
		return err
	}

	// Desktop shortcut
	desktopShortcutFile := fs.ExpandPath("$HOME/.local/share/applications/emulationstation-de.desktop")
	desktopShortcutContentSource := fs.NormalizePath("resources/emulationstation-de.desktop")
	desktopShortcutContent, err := resourcesContent.ReadFile(desktopShortcutContentSource)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(desktopShortcutFile), 0700)
	if err != nil {
		return err
	}

	desktopShortcutContent = []byte(os.ExpandEnv(string(desktopShortcutContent)))
	err = os.WriteFile(desktopShortcutFile, desktopShortcutContent, 0644)
	if err != nil {
		return err
	}

	return nil
}
