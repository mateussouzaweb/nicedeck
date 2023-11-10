package steam

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Check if Steam installation was done via flatpak
func IsFlatpak() (bool, error) {

	// App can be installed on system or user
	systemFile := os.ExpandEnv("$HOME/.local/share/flatpak/exports/bin/com.valvesoftware.Steam")
	userFile := "/var/lib/flatpak/exports/bin/com.valvesoftware.Steam"

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

// Retrieve the full Steam paths with given additional pattern
func GetPaths(pattern string) ([]string, error) {

	var found = []string{}

	// Fill possible locations
	paths := []string{
		os.ExpandEnv("$HOME/.steam/steam"),
		os.ExpandEnv("$HOME/.local/share/Steam"),
		os.ExpandEnv("$HOME/.var/app/com.valvesoftware.Steam/.steam/steam"),
		os.ExpandEnv("$HOME/snap/steam/common/.local/share/Steam"),
	}

	// Checks what directory path is available
	usePath := ""
	for _, possiblePath := range paths {
		exist, err := fs.DirectoryExist(possiblePath)
		if err != nil {
			return found, err
		} else if exist {
			usePath = filepath.Join(possiblePath, pattern)
			break
		}
	}

	// Return error if not detected
	if usePath == "" {
		return found, fmt.Errorf("could not detect the Steam installation path")
	}

	// Try to detect the path
	found, err := filepath.Glob(usePath)
	if err != nil {
		return found, err
	}

	if len(found) == 0 {
		return found, fmt.Errorf("could not found the Steam installation path: %s", usePath)
	}

	// Will return only the first result
	return found, nil
}
