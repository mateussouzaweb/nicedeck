package fs

import (
	"errors"
	"os"
	"path/filepath"
)

// Check if given path is a symbolic link
func IsSymlink(path string) (bool, error) {

	// os.Lstat does not follow symbolic links
	info, err := os.Lstat(path)
	if err == nil {
		return info.Mode()&os.ModeSymlink == os.ModeSymlink, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	return false, nil
}

// Remove an existing symlink
func RemoveSymlink(link string) error {

	// If already exists, then remove it first
	existing, err := IsSymlink(link)
	if err != nil {
		return err
	} else if existing {
		err := os.Remove(link)
		if err != nil {
			return err
		}
	}

	return nil
}

// Make symlink from target
func MakeSymlink(target string, link string) error {

	// If already exists, then remove it first
	err := RemoveSymlink(link)
	if err != nil {
		return err
	}

	// Make sure parent path exists
	err = os.MkdirAll(filepath.Dir(link), 0755)
	if err != nil {
		return err
	}

	// Make link to main executable
	err = os.Symlink(target, link)
	if err != nil {
		return err
	}

	return nil
}
