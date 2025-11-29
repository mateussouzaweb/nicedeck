package fs

import "os"

// Read file modification time in UTC unix timestamp
func ModificationTime(path string) (int64, error) {

	// Check if file exist
	exist, err := FileExist(path)
	if err != nil {
		return 0, err
	} else if !exist {
		return 0, nil
	}

	// Read file info
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	modificationTime := fileInfo.ModTime().UTC().Unix()
	return modificationTime, nil
}
