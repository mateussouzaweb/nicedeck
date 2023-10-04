package steam

import (
	"os"
	"path/filepath"
)

// Retrieve the full user data path with given additonal path
func GetUserPath(path string) (string, error) {

	// Retrieve home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Find user data paths
	path = home + "/.local/share/Steam/userdata/*" + path
	directories, err := filepath.Glob(path)
	if err != nil {
		return "", err
	}

	// Will return only the first result
	return directories[0], nil
}
