package steam

import (
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
)

// Check the runtime that Steam installation is running
func GetRuntime() (string, error) {

	// Check flatpak --system install
	flatpakSystemFile := "/var/lib/flatpak/exports/bin/com.valvesoftware.Steam"
	exist, err := fs.FileExist(flatpakSystemFile)
	if err != nil {
		return "", err
	} else if exist {
		return "flatpak", nil
	}

	// Check flatpak --user install
	flatpakUserFile := os.ExpandEnv("$HOME/.local/share/flatpak/exports/bin/com.valvesoftware.Steam")
	exist, err = fs.FileExist(flatpakUserFile)
	if err != nil {
		return "", err
	} else if exist {
		return "flatpak", nil
	}

	return "native", nil
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
	steamRuntime, err := GetRuntime()
	if err != nil {
		return fmt.Errorf("could not determine the Steam runtime: %s", err)
	} else if steamRuntime == "flatpak" {
		script := "flatpak override --user --talk-name=org.freedesktop.Flatpak com.valvesoftware.Steam"
		err = cli.Run(script)
		if err != nil {
			return fmt.Errorf("could not perform Steam setup with flatpak runtime: %s", err)
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
