package shortcuts

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Image struct
type Image struct {
	TargetDirectory string   `json:"targetDirectory"`
	TargetName      string   `json:"targetName"`
	TargetPath      string   `json:"targetPath"`
	Extensions      []string `json:"extensions"`
}

// Retrieve target path for the image with the specified extension
func (i *Image) Path(extension string) string {
	path := fmt.Sprintf("%s/%s%s", i.TargetDirectory, i.TargetName, extension)
	return fs.NormalizePath(path)
}

// Process image to sync image from source
// Also keep only one valid format and remove others
func (i *Image) Process(source string, overwriteExisting bool) error {

	var err error

	// Remove duplicated or unnecessary images on function exit
	// We should keep only the final target path image on exit
	defer func() {
		for _, extension := range i.Extensions {
			file := i.Path(extension)
			if file != i.TargetPath {
				errors.Join(err, fs.RemoveFile(file))
			}
		}
	}()

	// Find the valid extension for the given value
	validExtension := func(value string) string {
		for _, extension := range i.Extensions {
			if strings.HasSuffix(value, extension) {
				return extension
			}
		}

		return ""
	}

	// Determine processing details based on source
	sourceIsURL := strings.HasPrefix(source, "http")
	sourceExtension := validExtension(source)
	sourceTarget := i.Path(sourceExtension)
	sourceEqualsTarget := source == sourceTarget
	sourceValid := sourceExtension != ""

	// When no valid extension found, ignore further processing
	// In such case, target path will be empty, representing no image
	if !sourceValid {
		i.TargetPath = ""
		return nil
	}

	// Check if provided source path image exists
	sourcePathExist := false
	if !sourceIsURL {
		sourcePathExist, err = fs.FileExist(source)
		if err != nil {
			return err
		}
	}

	// If source image exists locally, copy it to the target path
	// This case is valid only when source and target are different
	if sourcePathExist && !sourceEqualsTarget {
		err := fs.CopyFile(source, sourceTarget, overwriteExisting)
		if err != nil {
			return err
		}

		i.TargetPath = sourceTarget
		return nil
	}

	// If source image is URL, download the image
	// This case act as priority when source path is equal to target path
	if sourceIsURL {
		err := fs.DownloadFile(source, sourceTarget, overwriteExisting)
		if err != nil {
			return err
		}

		i.TargetPath = sourceTarget
		return nil
	}

	// If source path exists and is equal to target, set the target path
	// This case act when no download is needed
	if !sourceIsURL && sourcePathExist && sourceEqualsTarget {
		i.TargetPath = sourceTarget
	}

	return nil
}

// Remove all images variations from valid extension
func (i *Image) Remove() error {

	for _, extension := range i.Extensions {
		file := i.Path(extension)
		err := fs.RemoveFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}
