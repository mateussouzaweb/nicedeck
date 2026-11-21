package controller

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/vdf"
)

type Shortcut = steam.Shortcut

//go:embed resources/*
var resourcesContent embed.FS

// Assigns controller template to given account shortcuts
func AssignController(basePath string, accountId string, shortcuts []*Shortcut) error {

	// No need to process if there is no shortcuts
	if len(shortcuts) == 0 {
		return nil
	}

	// Determine config, config set and controller template
	configPath := filepath.Join(
		basePath, "steamapps", "common", "Steam Controller Configs",
		accountId, "config",
	)

	configSet := "configset_controller_neptune.vdf"
	configSetPath := filepath.Join(configPath, configSet)

	// Steam Deck controller template (neptune)
	// with NiceDeck customizations
	controllerTemplate := "controller_neptune_nicedeck.vdf"
	controllerTemplatePath := filepath.Join(
		basePath, "controller_base", "templates",
		controllerTemplate,
	)

	// Write custom controller templates
	err := fs.CopyEmbedded(
		resourcesContent,
		"resources/controller.vdf",
		controllerTemplatePath,
		false, // Do not expand environment variables
		true,  // Overwrite existing
	)
	if err != nil {
		return fmt.Errorf("could not install Steam controller template: %s", err)
	}

	// Read config set file on config path
	data, err := vdf.ParseText(configSetPath)
	if err != nil {
		return fmt.Errorf("error while reading controller config set: %w", err)
	}

	// Define controller template on shortcuts
	changed := false
	configs := data.Object("controller_config")

	for _, shortcut := range shortcuts {

		// Layouts are keyed by normalized shortcut title
		// Only Steam native games uses AppID
		title := strings.ToLower(shortcut.AppName)
		title = strings.Map(func(r rune) rune {
			if strings.ContainsRune("/\\?%*:|\"<>.", r) {
				return -1
			}
			return r
		}, title)

		if title == "" {
			continue
		}

		// When config already exists, skip it
		// Value is already defined or is user customized value
		if _, exists := configs[title]; exists {
			continue
		}

		cli.Debug("Assigning controller layout to %d\n", shortcut.AppID)
		configs[title] = vdf.Vdf{"template": controllerTemplate}
		changed = true

	}

	// If changed, write content to VDF file
	if changed {
		data["controller_config"] = configs
		err = vdf.WriteText(configSetPath, data)
		if err != nil {
			return err
		}
	}

	return nil
}
