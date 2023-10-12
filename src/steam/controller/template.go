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

	controllerFile := destinationPath + "/controller_neptune_nicedeck.vdf"
	controllerConfig, err := resourcesContent.ReadFile("resources/controller.vdf")
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
