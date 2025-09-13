package settings

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write settings for ES-DE
func WriteSettings(destinationPath string) error {

	// Settings (write file only if not exist yet)
	settingsFile := filepath.Join(destinationPath, "settings", "es_settings.xml")
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

		settingsContent := os.ExpandEnv(string(settingsConfig))
		err = fs.WriteFile(settingsFile, settingsContent)
		if err != nil {
			return err
		}
	}

	// Systems
	systemsFile := filepath.Join(destinationPath, "custom_systems", "es_systems.xml")
	systemsConfig, err := resourcesContent.ReadFile("resources/es_systems.xml")
	if err != nil {
		return err
	}

	systemsContent := os.ExpandEnv(string(systemsConfig))
	err = fs.WriteFile(systemsFile, systemsContent)
	if err != nil {
		return err
	}

	// Find Rules
	findRulesFile := filepath.Join(destinationPath, "custom_systems", "es_find_rules.xml")
	findRulesConfig, err := resourcesContent.ReadFile("resources/es_find_rules.xml")
	if err != nil {
		return err
	}

	findRulesContent := os.ExpandEnv(string(findRulesConfig))
	err = fs.WriteFile(findRulesFile, findRulesContent)
	if err != nil {
		return err
	}

	return nil
}
