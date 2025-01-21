package controller

import (
	"embed"
	"os"
	"path/filepath"
)

//go:embed resources/*
var resourcesContent embed.FS

// Write controller templates on Steam at given destination path
func WriteTemplates(destinationPath string) error {

	controllerFile := filepath.Join(destinationPath, "controller_neptune_nicedeck.vdf")
	controllerConfigSource := filepath.Join("resources", "controller.vdf")
	controllerConfig, err := resourcesContent.ReadFile(controllerConfigSource)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(controllerFile), 0774)
	if err != nil {
		return err
	}

	err = os.WriteFile(controllerFile, controllerConfig, 0666)
	if err != nil {
		return err
	}

	return nil
}
