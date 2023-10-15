package steam

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

// Check if file already exist
func FileExist(file string) (bool, error) {

	_, err := os.Stat(file)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return false, err
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return true, nil
}

// Download file from URL into destination
func DownloadFile(url string, destinationFile string) error {

	// Check if file exists and if true, skip download process
	exist, err := FileExist(destinationFile)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	// Retrieve file from HTTP
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// Make sure file is created and writable
	err = os.MkdirAll(filepath.Dir(destinationFile), 0774)
	if err != nil {
		return err
	}

	file, err := os.Create(destinationFile)
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
