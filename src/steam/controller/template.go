package controller

import (
	"embed"
	"os"
)

//go:embed resources/*
var resourcesContent embed.FS

// Save controller template on Steam
func SaveTemplate(destinationFile string) error {

	controllerConfig, err := resourcesContent.ReadFile("resources/controller.vdf")
	if err != nil {
		return err
	}

	err = os.WriteFile(destinationFile, controllerConfig, 0666)
	if err != nil {
		return err
	}

	return nil
}
