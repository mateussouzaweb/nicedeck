package steam

import (
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
)

// Check if Steam installation was done via flatpak
func IsFlatpak() (bool, error) {

	// App can be installed on system or user
	systemFile := "/var/lib/flatpak/exports/bin/com.valvesoftware.Steam"
	userFile := os.ExpandEnv("$HOME/.local/share/flatpak/exports/bin/com.valvesoftware.Steam")

	// Checks what possible file exist
	for _, file := range []string{systemFile, userFile} {
		exist, err := fs.FileExist(file)
		if err != nil {
			return false, err
		} else if exist {
			return true, nil
		}
	}

	return false, nil
}

// Retrieve the absolute Steam path
func GetPath() (string, error) {

	// Fill possible locations
	paths := []string{
		os.ExpandEnv("$HOME/.steam/steam"),
		os.ExpandEnv("$HOME/.local/share/Steam"),
		os.ExpandEnv("$HOME/.var/app/com.valvesoftware.Steam/.steam/steam"),
		os.ExpandEnv("$HOME/snap/steam/common/.local/share/Steam"),
	}

	// Checks what directory path is available
	for _, possiblePath := range paths {
		exist, err := fs.DirectoryExist(possiblePath)
		if err != nil {
			return "", err
		} else if exist {
			return possiblePath, nil
		}
	}

	return "", nil
}

// Perform Steam setup
func Setup() error {

	// Retrieve Steam base path
	steamPath, err := GetPath()
	if err != nil {
		return fmt.Errorf("could not detect Steam installation: %s", err)
	}

	// Skip if Steam installation was not found
	if steamPath == "" {
		cli.Printf(cli.ColorWarn, "Steam not detected, skipping Steam setup process...\n")
		return nil
	}

	// Make sure Steam on flatpak has the necessary permission
	// We need this to run flatpak-spawn command to communicate with others flatpak apps
	isFlatpak, err := IsFlatpak()
	if err != nil {
		return fmt.Errorf("could not determine if Steam is from Flatpak: %s", err)
	} else if isFlatpak {
		script := "flatpak override --user --talk-name=org.freedesktop.Flatpak com.valvesoftware.Steam"
		err = cli.Run(script)
		if err != nil {
			return fmt.Errorf("could not perform Steam setup for Flatpak: %s", err)
		}
	}

	// Write controller templates
	controllerTemplatesPaths := steamPath + "/controller_base/templates"
	err = controller.WriteTemplates(controllerTemplatesPaths)
	if err != nil {
		return fmt.Errorf("could not perform Steam controller setup: %s", err)
	}

	return nil
}
