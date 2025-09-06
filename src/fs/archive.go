package fs

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Unzip given source ZIP file into destination
func Unzip(source, destination string) error {

	cli.Debug("Unzipping %s to %s\n", source, destination)

	// Open zip file
	archive, err := zip.OpenReader(source)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, archive.Close())
	}()

	// Process each item of the archive
	// When destination already exists
	// Files outside of the archive will not be removed
	for _, file := range archive.Reader.File {
		path := filepath.Join(destination, file.Name)

		// Read directory or file content
		reader, err := file.Open()
		if err != nil {
			return err
		}

		defer func() {
			errors.Join(err, reader.Close())
		}()

		// If is a directory, just ensure that the folder exists
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}

			continue
		}

		// When is file, first remove the existing file
		err = RemoveFile(path)
		if err != nil {
			return err
		}

		// Now create the target file again
		writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		defer func() {
			errors.Join(err, writer.Close())
		}()

		// Copy file content to target
		_, err = io.Copy(writer, reader)
		if err != nil {
			return err
		}
	}

	return nil
}
