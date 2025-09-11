package fs

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
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
		cli.Debug("Removing %s\n", path)
		return os.Remove(path)
	}

	return nil
}

// Move file to another location
func MoveFile(source string, destination string) error {

	// Check if file exist
	exist, err := FileExist(source)
	if err != nil {
		return err
	} else if !exist {
		return nil
	}

	cli.Debug("Moving %s to %s\n", source, destination)

	// Ensure that destination folder exists
	err = os.MkdirAll(filepath.Dir(destination), 0774)
	if err != nil {
		return err
	}

	// Move file using rename command
	err = os.Rename(source, destination)
	if err != nil {
		return err
	}

	return nil
}

// Copy file from given source path into destination path
func CopyFile(source string, destination string, overwriteExisting bool) error {

	// Skip copy if source and destination path are equals
	if source == destination {
		return nil
	}

	// Check if destination file exists
	destinationExist, err := FileExist(destination)
	if err != nil {
		return err
	}

	// Skip copy if already exist
	if !overwriteExisting && destinationExist {
		return nil
	}

	// Retrieve stat for source file
	sourceStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	// Check if both file are equals before copy
	// Verification will avoid unnecessary copy of large files
	if destinationExist {
		destinationStat, err := os.Stat(destination)
		if err != nil {
			return err
		}

		// Use simple file mode and size verification
		// sameFile := os.SameFile(sourceStat, destinationStat)
		sameFile := sourceStat.Mode() == destinationStat.Mode()
		sameFile = sameFile && sourceStat.Size() == destinationStat.Size()
		sameFile = sameFile && sourceStat.ModTime().Equal(destinationStat.ModTime())
		if sameFile {
			return nil
		}
	}

	cli.Debug("Coping %s to %s\n", source, destination)

	// Open source file
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, sourceFile.Close())
	}()

	// Ensure destination path exist
	err = os.MkdirAll(filepath.Dir(destination), 0755)
	if err != nil {
		return err
	}

	// Open destination file
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, destinationFile.Close())
	}()

	// Copy content from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// Write data to file
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	// Apply copied permissions to file
	err = os.Chmod(destination, sourceStat.Mode())
	if err != nil {
		return err
	}

	// Apply access and modification times to file
	err = os.Chtimes(destination, sourceStat.ModTime(), sourceStat.ModTime())
	if err != nil {
		return err
	}

	return err
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

	cli.Debug("Downloading %s to %s\n", url, destination)

	// Retrieve file from HTTP
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, response.Body.Close())
	}()

	// Ensure that destination folder exists
	err = os.MkdirAll(filepath.Dir(destination), 0774)
	if err != nil {
		return err
	}

	// Make sure file is created and writable
	file, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, file.Close())
	}()

	// Write HTTP response body to destination file
	_, err = file.ReadFrom(response.Body)
	if err != nil {
		return err
	}

	return nil
}

// Write content into target path
func WriteFile(path string, content string) error {

	cli.Debug("Writing file at %s\n", path)

	// Make sure destination folder path exist
	err := os.MkdirAll(filepath.Dir(path), 0774)
	if err != nil {
		return err
	}

	// Write content to file
	err = os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err
	}

	return nil
}
