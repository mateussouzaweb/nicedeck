package cli

import (
	"io/fs"
	"os"
)

// Return if directory exist on given path
func ExistDirectory(path string) bool {

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// Return if file exist on given path
func ExistFile(path string) bool {

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// Read file and retrieve content from it
func ReadFile(path string) ([]byte, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	return content, nil
}

// Write content on file
func WriteFile(file string, content []byte, perm fs.FileMode) error {

	err := os.WriteFile(file, content, perm)
	if err != nil {
		return err
	}

	return nil
}
