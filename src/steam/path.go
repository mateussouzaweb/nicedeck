package steam

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Retrieve the full Steam path with given additional path
func GetPath(path string) (string, error) {

	// Fill possible locations
	paths := []string{
		os.ExpandEnv("$HOME/.steam/steam"),
		os.ExpandEnv("$HOME/.local/share/Steam"),
		os.ExpandEnv("$HOME/.var/app/com.valvesoftware.Steam/.steam/steam"),
		os.ExpandEnv("$HOME/snap/steam/common/.local/share/Steam"),
	}

	// Check what path is available
	usePath := ""
	for _, possiblePath := range paths {
		stat, err := os.Stat(possiblePath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
		if stat.IsDir() {
			usePath = possiblePath + path
			break
		}
	}

	// Return error if not detected
	if usePath == "" {
		return "", fmt.Errorf("could not detect the steam installation path")
	}

	// Try to detect the path
	found, err := filepath.Glob(usePath)
	if err != nil {
		return "", err
	}

	if len(found) == 0 {
		return "", fmt.Errorf("could not found the steam installation path: %s", usePath)
	}

	// Will return only the first result
	return found[0], nil
}
