package steam

import (
	"os"
	"path/filepath"
)

// Retrieve the full steam path with given additional path
func GetPath(path string) (string, error) {

	path = os.ExpandEnv("$HOME/.local/share/Steam/" + path)
	directories, err := filepath.Glob(path)
	if err != nil {
		return "", err
	}

	// Will return only the first result
	return directories[0], nil
}
