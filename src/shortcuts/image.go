package shortcuts

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Image struct
type Image struct {
	SourcePath      string   `json:"sourcePath"`
	SourceURL       string   `json:"sourceUrl"`
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

// Process image to keep only one format and remove others
func (i *Image) Process(overwriteExisting bool) error {

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

	// Determine source path and URL valid extension
	sourceURLExtension := validExtension(i.SourceURL)
	sourcePathExtension := validExtension(i.SourcePath)

	// When no valid extension found, ignore further processing
	sourceURLValid := sourceURLExtension != ""
	sourcePathValid := sourcePathExtension != ""
	if !sourceURLValid && !sourcePathValid {
		return nil
	}

	// Determine target details based valid extension
	sourceURLTarget := i.Path(sourceURLExtension)
	sourcePathTarget := i.Path(sourcePathExtension)
	sourcePathEqualsTarget := i.SourcePath == sourcePathTarget

	// Check if provided source path image exists
	sourcePathExist := false
	if i.SourcePath != "" {
		sourcePathExist, err = fs.FileExist(i.SourcePath)
		if err != nil {
			return err
		}
	}

	// If source image exists locally, copy it to the target path
	// This case is valid only when source and target are different
	if sourcePathValid && sourcePathExist && !sourcePathEqualsTarget {
		i.TargetPath = sourcePathTarget
		err := fs.CopyFile(i.SourcePath, i.TargetPath, overwriteExisting)
		if err != nil {
			return err
		}

		return nil
	}

	// If source image URL exists, download the image
	// This case act as priority when source path is equal to target path
	if sourceURLValid {
		i.TargetPath = sourceURLTarget
		err := fs.DownloadFile(i.SourceURL, i.TargetPath, overwriteExisting)
		if err != nil {
			return err
		}

		return nil
	}

	// If source path exists and is equal to target, set the target path
	// This case act when no download is needed
	if sourcePathValid && sourcePathExist && sourcePathEqualsTarget {
		i.TargetPath = sourcePathTarget
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
