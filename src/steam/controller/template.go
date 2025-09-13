package controller

import (
	"embed"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write controller templates on Steam at given destination path
func WriteTemplates(destinationPath string) error {

	controllerFile := filepath.Join(destinationPath, "controller_neptune_nicedeck.vdf")
	controllerConfig, err := resourcesContent.ReadFile("resources/controller.vdf")
	if err != nil {
		return err
	}

	err = fs.WriteFile(controllerFile, string(controllerConfig))
	if err != nil {
		return err
	}

	return nil
}
