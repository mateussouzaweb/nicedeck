package fs

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Retrieve JSON content from URL and put into target
func RetrieveJSON(url string, target any) error {

	cli.Debug("Requesting JSON %s\n", url)

	// Read content from HTTP request
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, res.Body.Close())
	}()

	// Read body content
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Write decoded content into target
	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}

// Read JSON from file content and put into target
func ReadJSON(path string, target any) error {

	// Check if file exist
	exist, err := FileExist(path)
	if err != nil {
		return err
	} else if exist {

		cli.Debug("Reading JSON %s\n", path)

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Write decoded content to target pointer
		err = json.Unmarshal(content, target)
		if err != nil {
			return err
		}
	}

	return nil
}

// Write JSON content from source into target path
func WriteJSON(path string, source any) error {

	cli.Debug("Writing JSON at %s\n", path)

	// Convert source to JSON representation
	content, err := json.MarshalIndent(source, "", "  ")
	if err != nil {
		return err
	}

	// Make sure destination folder path exist
	err = os.MkdirAll(filepath.Dir(path), 0774)
	if err != nil {
		return err
	}

	// Write JSON content to file
	err = os.WriteFile(path, content, 0666)
	if err != nil {
		return err
	}

	return nil
}
