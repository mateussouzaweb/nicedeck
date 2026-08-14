package fs

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Read XML from file content and put into target
func ReadXML(path string, target any) error {

	// Check if file exist
	exist, err := FileExist(path)
	if err != nil {
		return err
	} else if exist {

		cli.Debug("Reading XML %s\n", path)

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Write decoded content to target pointer
		err = xml.Unmarshal(content, target)
		if err != nil {
			return err
		}
	}

	return nil
}

// Write XML content from source into target path
func WriteXML(path string, source any) error {

	cli.Debug("Writing XML at %s\n", path)

	// Convert source to XML representation
	content, err := xml.MarshalIndent(source, "", "  ")
	if err != nil {
		return err
	}

	// Make sure destination folder path exist
	err = os.MkdirAll(filepath.Dir(path), 0774)
	if err != nil {
		return err
	}

	// Write XML content to file
	output := append([]byte(xml.Header), content...)
	err = os.WriteFile(path, output, 0666)
	if err != nil {
		return err
	}

	return nil
}
