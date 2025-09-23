package fs

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Extract given source .zip file into destination
// Also consider the expected file path to ignore parent directories
func ExtractZip(source string, destination string, expected string) error {

	cli.Debug("Extracting %s to %s\n", source, destination)

	// Open zip file
	archive, err := zip.OpenReader(source)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, archive.Close())
	}()

	// Detects the undesired path based on the expected file path
	undesired := ""
	for _, file := range archive.Reader.File {
		destinationPath := filepath.Join(destination, file.Name)
		if strings.HasSuffix(destinationPath, expected) {
			undesired = strings.TrimSuffix(destinationPath, expected)
			undesired = strings.TrimPrefix(undesired, destination)
			break
		}
	}

	// Process each item of the archive
	// Files outside of the archive will not be removed
	for _, file := range archive.Reader.File {

		// Read directory or file content
		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		defer func() {
			errors.Join(err, fileReader.Close())
		}()

		// Get content information
		destinationPath := filepath.Join(destination, file.Name)
		destinationPath = strings.Replace(destinationPath, undesired, "", 1)

		// If is a directory, just ensure that the folder exists
		isDirectory := file.FileInfo().IsDir()
		if isDirectory {
			err = os.MkdirAll(destinationPath, file.Mode())
			if err != nil {
				return err
			}

			continue
		}

		// If is a regular file, make sure that the parent folder exists
		err = os.MkdirAll(filepath.Dir(destinationPath), 0755)
		if err != nil {
			return err
		}

		// Now create or replace the target file
		flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		fileWriter, err := os.OpenFile(destinationPath, flags, file.Mode())
		if err != nil {
			return err
		}

		defer func() {
			errors.Join(err, fileWriter.Close())
		}()

		// Copy file content to target
		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			return err
		}

	}

	return nil
}

// Extract given source .tar.gz file content into destination
// Also consider the expected file path to ignore parent directories
func ExtractTarGz(source string, destination string, expected string) error {

	cli.Debug("Extracting %s to %s\n", source, destination)

	// Open the archive file
	archive, err := os.Open(source)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, archive.Close())
	}()

	// Create the gzip reader for archive
	gzipReader, err := gzip.NewReader(archive)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, gzipReader.Close())
	}()

	// Detects the undesired path based on the expected file path
	undesired := ""
	tarReader := tar.NewReader(gzipReader)
	for {

		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		} else if err != nil {
			return err
		} else if header == nil {
			continue
		}

		destinationPath := filepath.Join(destination, header.Name)
		if strings.HasSuffix(destinationPath, expected) {
			undesired = strings.TrimSuffix(destinationPath, expected)
			undesired = strings.TrimPrefix(undesired, destination)
			break
		}
	}

	// Process each item of the archive
	// Files outside of the archive will not be removed
	tarReader = tar.NewReader(gzipReader)
	for {

		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		} else if err != nil {
			return err
		} else if header == nil {
			continue
		}

		// Get content information
		destinationPath := filepath.Join(destination, header.Name)
		destinationPath = strings.Replace(destinationPath, undesired, "", 1)
		isDirectory := header.Typeflag == tar.TypeDir
		isRegularFile := header.Typeflag == tar.TypeReg
		fileMode := os.FileMode(header.Mode)

		// Unknown content type, just ignore
		if !isDirectory && !isRegularFile {
			continue
		}

		// If is a directory, just ensure that the folder exists
		if isDirectory {
			err = os.MkdirAll(destinationPath, fileMode)
			if err != nil {
				return err
			}

			continue
		}

		// If is a regular file, make sure that the parent folder exists
		err = os.MkdirAll(filepath.Dir(destinationPath), 0755)
		if err != nil {
			return err
		}

		// Now create or replace the target file
		flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		fileWriter, err := os.OpenFile(destinationPath, flags, fileMode)
		if err != nil {
			return err
		}

		defer func() {
			errors.Join(err, fileWriter.Close())
		}()

		// Copy file content to target
		_, err = io.Copy(fileWriter, tarReader)
		if err != nil {
			return err
		}

	}

	return nil
}
