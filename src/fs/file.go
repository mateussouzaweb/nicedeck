package fs

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

// Remove file from system if exist
func RemoveFile(path string) error {

	exist, err := FileExist(path)
	if err != nil {
		return err
	} else if exist {
		return os.Remove(path)
	}

	return nil
}

// Copy file from given source path into destination path
func CopyFile(source string, destination string) error {

	var err error

	// Open source file
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, sourceFile.Close())
	}()

	// Ensure destination path exist
	err = os.MkdirAll(filepath.Dir(destination), 0755)
	if err != nil {
		return err
	}

	// Open destination file
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, destinationFile.Close())
	}()

	// Copy content from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// Write data to file
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	// Get permissions from source path
	stat, err := os.Stat(source)
	if err != nil {
		return err
	}

	// Apply copied permissions to file
	err = os.Chmod(destination, stat.Mode())
	if err != nil {
		return err
	}

	return err
}

// Move file to another location
func MoveFile(source string, destination string) error {

	// Check if file exist
	exist, err := FileExist(source)
	if err != nil {
		return err
	} else if !exist {
		return nil
	}

	// Ensure that destination folder exists
	err = os.MkdirAll(filepath.Dir(destination), 0774)
	if err != nil {
		return err
	}

	// Move file using rename command
	err = os.Rename(source, destination)
	if err != nil {
		return err
	}

	return nil
}

// Download file from URL into destination
func DownloadFile(url string, destination string, overwriteExisting bool) error {

	// Check if file exists and skip download if already exist
	if !overwriteExisting {
		exist, err := FileExist(destination)
		if err != nil {
			return err
		} else if exist {
			return nil
		}
	}

	// Retrieve file from HTTP
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, response.Body.Close())
	}()

	// Ensure that destination folder exists
	err = os.MkdirAll(filepath.Dir(destination), 0774)
	if err != nil {
		return err
	}

	// Make sure file is created and writable
	file, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, file.Close())
	}()

	// Write HTTP response body to destination file
	_, err = file.ReadFrom(response.Body)
	if err != nil {
		return err
	}

	return nil
}
