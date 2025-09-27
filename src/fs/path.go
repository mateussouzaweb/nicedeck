package fs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

// Find real path for expected file or directory
func FindPath(base string, expected string) (string, error) {

	found := ""
	err := filepath.WalkDir(base,
		func(realPath string, dir os.DirEntry, err error) error {

			// Stop in case of errors
			if err != nil {
				return err
			}

			// When file or directory matches
			if strings.HasSuffix(realPath, expected) {
				found = realPath
				return fs.SkipAll
			}

			return nil
		},
	)

	if err == fs.SkipAll {
		return found, nil
	}
	if found == "" {
		err = fmt.Errorf("expected file not found: %s", expected)
	}

	return found, err
}
