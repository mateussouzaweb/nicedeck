package emulationstation

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Perform setup for EmulationStation
func Setup() error {
	if cli.IsLinux() {
		return WriteLinuxSettings()
	}
	return nil
}

// Write settings for EmulationStation
func WriteLinuxSettings() error {

	// Settings (write file only if not exist yet)
	settingsFile := fs.ExpandPath("$HOME/ES-DE/settings/es_settings.xml")
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

		settingsConfig = []byte(os.ExpandEnv(string(settingsConfig)))
		err = os.WriteFile(settingsFile, settingsConfig, 0666)
		if err != nil {
			return err
		}
	}

	// Systems
	systemsFile := fs.ExpandPath("$HOME/ES-DE/custom_systems/es_systems.xml")
	systemsConfig, err := resourcesContent.ReadFile("resources/es_systems.xml")
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
	findRulesConfig, err := resourcesContent.ReadFile("resources/es_find_rules.xml")
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
	iconContent, err := resourcesContent.ReadFile("resources/icon.png")
	if err != nil {
		return err
	}

	err = os.WriteFile(iconFile, iconContent, 0666)
	if err != nil {
		return err
	}

	// Desktop shortcut
	desktopShortcutFile := fs.ExpandPath("$HOME/.local/share/applications/emulationstation-de.desktop")
	desktopShortcutContent, err := resourcesContent.ReadFile("resources/emulationstation-de.desktop")
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
