package fs

import (
	"crypto/md5"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Copy file from given embedded source path into destination path
// Embedded copy always expand environment variables on destination path
func CopyEmbedded(source embed.FS, path string, destination string, overwriteExisting bool) error {

	// Read embedded content
	embedded, err := source.ReadFile(path)
	if err != nil {
		return err
	}

	// Expand embedded to replace variables
	content := []byte(os.ExpandEnv(string(embedded)))

	// Check if destination file exists
	destinationExist, err := FileExist(destination)
	if err != nil {
		return err
	}

	// Ignore when cannot overwrite existing file
	if destinationExist && !overwriteExisting {
		return nil
	}

	// Check if both file are equals before copy
	// Verification will avoid unnecessary copy to destination files
	if destinationExist {
		currentContent, err := os.ReadFile(destination)
		if err != nil {
			return err
		}

		// Ignore operation if checksum matches
		// Checksum is based on expanded embedded content
		contentChecksum := fmt.Sprintf("%x", md5.Sum(content))
		currentChecksum := fmt.Sprintf("%x", md5.Sum(currentContent))
		if contentChecksum == currentChecksum {
			return nil
		}
	}

	cli.Debug("Copying %s to %s\n", path, destination)

	// Ensure destination path exist
	err = os.MkdirAll(filepath.Dir(destination), 0755)
	if err != nil {
		return err
	}

	// Write data to file
	err = os.WriteFile(destination, content, 0666)
	if err != nil {
		return err
	}

	return nil
}
