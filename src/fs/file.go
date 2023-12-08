package fs

import (
	"errors"
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

	defer response.Body.Close()

	// Make sure file is created and writable
	err = os.MkdirAll(filepath.Dir(destination), 0774)
	if err != nil {
		return err
	}

	file, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer file.Close()

	// Write HTTP response body to destination file
	_, err = file.ReadFrom(response.Body)
	if err != nil {
		return err
	}

	return nil
}
