package fs

import (
	"os"
	"strings"
)

// Normalize a path with correct system separator
func NormalizePath(path string) string {
	path = strings.ReplaceAll(path, "/", string(os.PathSeparator))
	return path
}

// Expand path and return the normalized value
func ExpandPath(path string) string {
	path = os.ExpandEnv(path)
	return NormalizePath(path)
}
