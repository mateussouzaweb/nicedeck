package scraper

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Download file from URL into destination
func DownloadFile(url string, destinationFile string) error {

	// Check if file exists and skip download if already exist
	exist, err := fs.FileExist(destinationFile)
	if err != nil {
		return err
	} else if exist {
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
