package esde

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write settings for ES-DE
func WriteSettings() error {

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

	systemsConfig = []byte(os.ExpandEnv(string(systemsConfig)))
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

	findRulesConfig = []byte(os.ExpandEnv(string(findRulesConfig)))
	err = os.WriteFile(findRulesFile, findRulesConfig, 0666)
	if err != nil {
		return err
	}

	return nil
}
