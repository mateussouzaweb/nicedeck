package fs

import (
	"errors"
	"os"
)

// Check if directory exist at given path
func DirectoryExist(path string) (bool, error) {

	stat, err := os.Stat(path)
	if err == nil {
		return stat.IsDir(), nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	return false, nil
}
