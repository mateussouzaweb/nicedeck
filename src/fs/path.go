package fs

import (
	"os"
	"strings"
)

// Normalize a path with correct system separator
func NormalizePath(path string) string {
	separator := string(os.PathSeparator)
	path = strings.ReplaceAll(path, "/", separator)
	path = strings.ReplaceAll(path, "\\", separator)
	return path
}

// Expand path and return the normalized value
func ExpandPath(path string) string {
	path = os.ExpandEnv(path)
	return NormalizePath(path)
}
