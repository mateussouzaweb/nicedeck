package steam

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
)

// Check the runtime that Steam installation is running
func GetRuntime() (string, error) {

	// Check from environment variable
	fromEnv := cli.GetEnv("STEAM_RUNTIME", "")
	if fromEnv != "" {
		return fromEnv, nil
	}

	// Check flatpak --system install
	flatpakSystemFile := fs.NormalizePath("/var/lib/flatpak/exports/bin/com.valvesoftware.Steam")
	exist, err := fs.FileExist(flatpakSystemFile)
	if err != nil {
		return "", err
	} else if exist {
		return "flatpak", nil
	}

	// Check flatpak --user install
	flatpakUserFile := fs.ExpandPath("$HOME/.local/share/flatpak/exports/bin/com.valvesoftware.Steam")
	exist, err = fs.FileExist(flatpakUserFile)
	if err != nil {
		return "", err
	} else if exist {
		return "flatpak", nil
	}

	// Check from snap install
	snapFile := fs.NormalizePath("/snap/bin/steam")
	exist, err = fs.FileExist(snapFile)
	if err != nil {
		return "", err
	} else if exist {
		return "snap", nil
	}

	return "native", nil
}

// Retrieve the absolute Steam path
func GetPath() (string, error) {

	// Check from environment variable
	fromEnv := cli.GetEnv("STEAM_PATH", "")
	if fromEnv != "" {
		return fs.ExpandPath(fromEnv), nil
	}

	// Fill possible locations
	paths := []string{
		fs.ExpandPath("$HOME/.steam/steam"),
		fs.ExpandPath("$HOME/.local/share/Steam"),
		fs.ExpandPath("$HOME/.var/app/com.valvesoftware.Steam/.steam/steam"),
		fs.ExpandPath("$HOME/snap/steam/common/.local/share/Steam"),
		fs.ExpandPath("$HOME/Library/Application Support/Steam"),
		fs.ExpandPath("$HOMEDRIVE\\Program Files (x86)\\Steam"),
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

// Retrieve Steam user config path
func GetConfigPath() (string, error) {

	// Check from environment variable
	fromEnv := cli.GetEnv("STEAM_USER_CONFIG_PATH", "")
	if fromEnv != "" {
		return fs.ExpandPath(fromEnv), nil
	}

	// Retrieve Steam base path
	steamPath, err := GetPath()
	if err != nil {
		return "", fmt.Errorf("could not detect Steam installation: %s", err)
	}

	// Steam can contains more than one user
	// At this time, we manage only the first user
	globPath := fs.NormalizePath(steamPath + "/userdata/*/config")
	configPaths, err := filepath.Glob(globPath)
	if err != nil {
		return "", fmt.Errorf("could not detect Steam user configuration: %s", err)
	}

	// Make sure zero config is ignored (this is not a valid user)
	if len(configPaths) > 0 {
		invalidPath := fs.NormalizePath("/0/config")
		if strings.Contains(configPaths[0], invalidPath) {
			configPaths = configPaths[1:]
		}
	}

	// Check if results was found
	if len(configPaths) == 0 {
		return "", fmt.Errorf("no users detected, please make sure to login into Steam first")
	}

	return configPaths[0], nil
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
	controllerTemplatesPaths := filepath.Join(steamPath, "controller_base", "templates")
	err = controller.WriteTemplates(controllerTemplatesPaths)
	if err != nil {
		return fmt.Errorf("could not perform Steam controller setup: %s", err)
	}

	return nil
}
