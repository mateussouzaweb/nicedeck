package fs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

// Copy directory content to destination
// When content already exists, it will be replaced
func CopyDirectory(source string, destination string) error {

	// Read stat from source path
	stat, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("source path is not a directory")
	}

	// Ensure destination path exist
	err = os.MkdirAll(destination, stat.Mode())
	if err != nil {
		return err
	}

	// Read entries in source path
	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	// Process list of entries, but skip symbolic links
	for _, entry := range entries {

		sourcePath := filepath.Join(source, entry.Name())
		destinationPath := filepath.Join(destination, entry.Name())

		if entry.IsDir() {
			err = CopyDirectory(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		} else if entry.Type().IsRegular() {
			err = CopyFile(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
