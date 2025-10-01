package steam

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
)

// Retrieve Steam package
func GetPackage() packaging.Package {
	return packaging.Installed(&linux.Flatpak{
		Namespace: "system",
		AppID:     "com.valvesoftware.Steam",
		Overrides: []string{
			fs.ExpandPath("--filesystem=$GAMES"),
			"--talk-name=org.freedesktop.Flatpak",
		},
		Arguments: packaging.NoArguments(),
	}, &linux.Flatpak{
		Namespace: "user",
		AppID:     "com.valvesoftware.Steam",
		Overrides: []string{
			fs.ExpandPath("--filesystem=$GAMES"),
			"--talk-name=org.freedesktop.Flatpak",
		},
		Arguments: packaging.NoArguments(),
	}, &linux.Snap{
		AppID:     "steam",
		AppBin:    "steam",
		Arguments: packaging.NoArguments(),
	}, &linux.Binary{
		AppID:     "steam",
		AppBin:    "/usr/bin/steam",
		Arguments: packaging.NoArguments(),
	}, &macos.Homebrew{
		AppID:     "steam",
		AppName:   "Steam.app",
		Arguments: packaging.NoArguments(),
	}, &windows.WinGet{
		AppID:     "Valve.Steam",
		AppExe:    "$PROGRAMS_X86/Steam/Steam.exe",
		Arguments: packaging.NoArguments(),
	})
}

// Retrieve the absolute Steam path
func GetBasePath() (string, error) {

	// Fill possible locations
	paths := []string{
		fs.ExpandPath("$VAR/com.valvesoftware.Steam/.steam/steam"),
		fs.ExpandPath("$HOME/snap/steam/common/.local/share/Steam"),
		fs.ExpandPath("$HOME/.steam/steam"),
		fs.ExpandPath("$SHARE/Steam"),
		fs.ExpandPath("$CONFIG/Steam"),
		fs.ExpandPath("$PROGRAMS_X86/Steam"),
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

	// Retrieve Steam base path
	steamPath, err := GetBasePath()
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
		return "", nil
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
