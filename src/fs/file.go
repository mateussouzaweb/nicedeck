package fs

import (
	"errors"
	"os"
)

// Check if file exist at given path
func FileExist(path string) (bool, error) {

	stat, err := os.Stat(path)
	if err == nil {
		return !stat.IsDir(), nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	return false, nil
}
