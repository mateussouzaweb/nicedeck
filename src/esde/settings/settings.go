package settings

import (
	"embed"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write settings for ES-DE
func WriteSettings(destinationPath string) error {

	// Settings
	err := fs.CopyEmbedded(
		resourcesContent,
		"resources/es_settings.xml",
		filepath.Join(destinationPath, "settings", "es_settings.xml"),
		false, // Do not overwrite existing
	)

	if err != nil {
		return err
	}

	// Systems
	err = fs.CopyEmbedded(
		resourcesContent,
		"resources/es_systems.xml",
		filepath.Join(destinationPath, "custom_systems", "es_systems.xml"),
		true,
	)

	if err != nil {
		return err
	}

	// Find Rules
	err = fs.CopyEmbedded(
		resourcesContent,
		"resources/es_find_rules.xml",
		filepath.Join(destinationPath, "custom_systems", "es_find_rules.xml"),
		true,
	)

	if err != nil {
		return err
	}

	return nil
}
