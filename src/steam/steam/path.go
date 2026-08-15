package steam

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Retrieve the absolute Steam path
func GetBasePath() (string, error) {

	// Fill possible locations
	paths := []string{
		fs.ExpandPath("$VAR/com.valvesoftware.Steam/.steam/steam"),
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
