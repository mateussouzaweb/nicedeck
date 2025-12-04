package shortcuts

import (
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

	toRemove := []string{}

	// Determine the target filename and the extension
	for _, extension := range i.Extensions {
		if strings.HasSuffix(i.SourcePath, extension) {
			i.TargetPath = i.Path(extension)
		} else if strings.HasSuffix(i.SourceURL, extension) {
			i.TargetPath = i.Path(extension)
		} else {
			toRemove = append(toRemove, i.Path(extension))
		}
	}

	// Remove duplicated or unnecessary images
	for _, file := range toRemove {
		err := fs.RemoveFile(file)
		if err != nil {
			return err
		}
	}

	// Ignore image processing when no target path is detected
	if i.TargetPath == "" {
		return nil
	}

	// If source image exists locally, copy it to the target path
	if i.SourcePath != "" {
		exist, err := fs.FileExist(i.SourcePath)
		if err != nil {
			return err
		} else if exist && i.SourcePath != i.TargetPath {
			err := fs.CopyFile(i.SourcePath, i.TargetPath, overwriteExisting)
			if err != nil {
				return err
			}

			return nil
		}
	}

	// If source image URL exists, download the image
	if i.SourceURL != "" {
		err := fs.DownloadFile(i.SourceURL, i.TargetPath, overwriteExisting)
		if err != nil {
			return err
		}
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
