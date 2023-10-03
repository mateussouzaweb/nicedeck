package fs

import "os"

// Return if directory exist on given path
func DirectoryExist(path string) bool {

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
