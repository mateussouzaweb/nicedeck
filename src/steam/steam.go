package steam

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
)

// Retrieve Steam package
func GetPackage() packaging.Package {
	return packaging.Installed(&linux.Flatpak{
		Namespace: "system",
		AppID:     "com.valvesoftware.Steam",
		Overrides: []string{"--talk-name=org.freedesktop.Flatpak"},
	}, &linux.Flatpak{
		Namespace: "user",
		AppID:     "com.valvesoftware.Steam",
		Overrides: []string{"--talk-name=org.freedesktop.Flatpak"},
	}, &linux.Snap{
		AppID:  "steam",
		AppBin: "steam",
	}, &linux.Binary{
		AppID:  "steam",
		AppBin: "/usr/share/bin/steam",
	}, &macos.Homebrew{
		AppID:   "steam",
		AppName: "Steam.app",
	}, &windows.WinGet{
		AppID:  "Valve.Steam",
		AppExe: "$PROGRAMS_X86\\Steam\\Steam.exe",
	})
}

// Check the runtime that Steam installation is running
func GetRuntime() (string, error) {

	// Check from environment variable
	fromEnv := cli.GetEnv("STEAM_RUNTIME", "")
	if fromEnv != "" {
		return fromEnv, nil
	}

	// Check from packaging program
	program := GetPackage()
	return program.Runtime(), nil
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
		fs.ExpandPath("$VAR/com.valvesoftware.Steam/.steam/steam"),
		fs.ExpandPath("$HOME/snap/steam/common/.local/share/Steam"),
		fs.ExpandPath("$HOME/.steam/steam"),
		fs.ExpandPath("$SHARE/Steam"),
		fs.ExpandPath("$CONFIG/Steam"),
		fs.ExpandPath("$PROGRAMS_X86\\Steam"),
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
	fromEnv := cli.GetEnv("STEAM_CONFIG_PATH", "")
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

// Ensure executable by adding special wrappers when necessary
func EnsureExec(runtime string, exec string) string {
	if runtime == "flatpak" {
		return "/usr/bin/flatpak-spawn --host " + CleanExec(exec)
	}
	return CleanExec(exec)
}

// Clean executable by removing special wrappers
func CleanExec(exec string) string {
	exec = strings.Replace(exec, "/usr/bin/flatpak-spawn --host", "", 1)
	exec = strings.Trim(exec, " ")
	return exec
}

// Perform Steam setup
func Setup() error {

	// Retrieve Steam base path
	// Skip if Steam installation was not found
	steamPath, err := GetPath()
	if err != nil {
		return fmt.Errorf("could not detect Steam installation: %s", err)
	} else if steamPath == "" {
		cli.Printf(cli.ColorWarn, "Steam not detected, skipping Steam setup process...\n")
		return nil
	}

	// Write controller templates
	controllerTemplatesPaths := filepath.Join(steamPath, "controller_base", "templates")
	err = controller.WriteTemplates(controllerTemplatesPaths)
	if err != nil {
		return fmt.Errorf("could not perform Steam controller setup: %s", err)
	}

	// Make sure Steam on flatpak has the necessary permission
	if _, ok := GetPackage().(*linux.Flatpak); ok {
		err := GetPackage().(*linux.Flatpak).ApplyOverrides()
		if err != nil {
			return fmt.Errorf("could not perform Steam runtime setup: %s", err)
		}
	}

	return nil
}
