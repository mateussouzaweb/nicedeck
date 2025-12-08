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

	// Steam Deck controller template
	err := fs.CopyEmbedded(
		resourcesContent,
		"resources/controller.vdf",
		filepath.Join(destinationPath, "controller_neptune_nicedeck.vdf"),
		true,
	)

	if err != nil {
		return err
	}

	return nil
}
