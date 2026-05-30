package fs

import (
	"errors"
	"os"
	"path/filepath"
)

type Info struct {
	Path         string
	Type         string
	Exist        bool
	Size         int64
	ModifiedTime int64
}

// Returns information about a file or folder path
func GetInfo(path string) (*Info, error) {

	info := &Info{
		Path:         ExpandPath(path),
		Type:         "none",
		Exist:        false,
		Size:         0,
		ModifiedTime: 0,
	}

	// Check for path existence and type
	stat, err := os.Stat(info.Path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return info, err
		}
	}
	if errors.Is(err, os.ErrNotExist) {
		return info, nil
	}

	// Set existence, size and modified time
	info.Exist = true
	info.Size = stat.Size()
	info.ModifiedTime = stat.ModTime().Unix()

	// Determine type when path exist
	if stat.IsDir() {
		info.Type = "folder"
	} else if stat.Mode().IsRegular() {
		info.Type = "file"
	} else {
		info.Type = "link"
	}

	// When path is a directory, we need to calculate the modified time of the directory
	// Return the most recent modified time for all files available in the directory
	if stat.IsDir() {

		// Note: walkDir does not follow symbolic links
		filepath.WalkDir(info.Path, func(filePath string, dir os.DirEntry, err error) error {

			// Stop in case of errors
			if err != nil {
				return err
			}

			// Ignore directories
			if dir.IsDir() {
				return nil
			}

			// Get file info
			fileInfo, err := dir.Info()
			if err != nil {
				return err
			}

			// Update modified time if file is more recent than current modified time
			if fileInfo.ModTime().Unix() > info.ModifiedTime {
				info.ModifiedTime = fileInfo.ModTime().Unix()
			}

			return nil
		})

	}

	return info, nil
}
