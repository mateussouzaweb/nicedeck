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

	// Open the archive file
	archive, err := os.Open(source)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, archive.Close())
	}()

	// Get file info to determine size for zip.NewReader
	archiveInfo, err := archive.Stat()
	if err != nil {
		return err
	}

	// Open zip file to see its content
	archiveReader, err := zip.NewReader(archive, archiveInfo.Size())
	if err != nil {
		return err
	}

	// Detects the undesired path based on the expected file path
	undesired := ""
	for _, file := range archiveReader.File {
		destinationPath := NormalizePath(file.Name)
		destinationPath = filepath.Join(destination, destinationPath)

		if strings.HasSuffix(destinationPath, expected) {
			undesired = strings.TrimSuffix(destinationPath, expected)
			undesired = strings.TrimPrefix(undesired, destination)
			break
		}
	}

	// Process each item of the archive
	// Files outside of the archive will not be removed
	for _, file := range archiveReader.File {

		// Get content information
		destinationPath := NormalizePath(file.Name)
		destinationPath = filepath.Join(destination, destinationPath)
		destinationPath = strings.Replace(destinationPath, undesired, "", 1)
		isDirectory := file.FileInfo().IsDir()

		// Make sure that path is not a directory
		separator := string(os.PathSeparator)
		if strings.HasSuffix(NormalizePath(file.Name), separator) {
			isDirectory = true
		}

		// If is a directory, just ensure that the folder exists
		if isDirectory {
			err = os.MkdirAll(destinationPath, 0755)
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

		// Read file content
		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		defer func() {
			errors.Join(err, fileReader.Close())
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

		destinationPath := NormalizePath(header.Name)
		destinationPath = filepath.Join(destination, destinationPath)

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
		destinationPath := NormalizePath(header.Name)
		destinationPath = filepath.Join(destination, destinationPath)
		destinationPath = strings.Replace(destinationPath, undesired, "", 1)
		isDirectory := header.Typeflag == tar.TypeDir
		isRegularFile := header.Typeflag == tar.TypeReg

		// Unknown content type, just ignore
		if !isDirectory && !isRegularFile {
			continue
		}

		// If is a directory, just ensure that the folder exists
		if isDirectory {
			err = os.MkdirAll(destinationPath, 0755)
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
		fileMode := os.FileMode(header.Mode)
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
