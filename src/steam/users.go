package steam

import (
	"os"
	"path/filepath"
)

// Retrieve the user data path
func GetUserDataPath() (string, error) {

	// Retrieve home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Find user data paths
	path := home + "/.local/share/Steam/userdata/*"
	directories, err := filepath.Glob(path)
	if err != nil {
		return "", err
	}

	// Will return only the first result
	return directories[0], nil
}
