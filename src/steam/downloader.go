package steam

import (
	"net/http"
	"os"
	"path/filepath"
)

// Download file from URL into destination
func DownloadFile(url string, destinationFile string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = os.MkdirAll(filepath.Dir(destinationFile), 0774)
	if err != nil {
		return err
	}

	file, err := os.Create(destinationFile)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.ReadFrom(response.Body)
	if err != nil {
		return err
	}

	return nil
}
