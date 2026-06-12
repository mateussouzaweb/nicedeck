package fs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Normalize filename to remove undesired characters
func NormalizeFilename(name string) string {

	// Windows does not allow certain characters in file names
	// Linux and MacOS are more permissive, but we keep consistency across platforms
	name = strings.ReplaceAll(name, "<", "")
	name = strings.ReplaceAll(name, ">", "")
	name = strings.ReplaceAll(name, "\\", "")
	name = strings.ReplaceAll(name, "/", "")
	name = strings.ReplaceAll(name, "|", "")
	name = strings.ReplaceAll(name, "?", "")
	name = strings.ReplaceAll(name, "*", "")
	name = strings.ReplaceAll(name, "#", "")
	name = strings.ReplaceAll(name, "~", "")
	name = strings.ReplaceAll(name, "\"", "")

	// These characters are replaced with a hyphen
	name = strings.ReplaceAll(name, ":", " - ")

	// Finally, remove duplicate spaces and trim
	name = strings.Join(strings.Fields(name), " ")
	name = strings.TrimSpace(name)
	name = strings.TrimRight(name, "-")
	name = strings.TrimRight(name, ".")

	return name
}

// Normalize a path with correct system separator
func NormalizePath(path string) string {
	separator := string(os.PathSeparator)
	path = strings.ReplaceAll(path, "/", separator)
	path = strings.ReplaceAll(path, "\\", separator)
	return path
}

// Expand path and return the normalized value
// Accepts and resolves environment variables and glob patterns
func ExpandPath(path string) string {
	path = os.ExpandEnv(path)
	path = NormalizePath(path)

	// When path contains glob patterns, try to find the file that matches the pattern and return it
	// This is useful when the folder or file name is not fixed, but the pattern is fixed
	if strings.ContainsAny(path, "*?[") {
		matches, err := filepath.Glob(path)
		if err != nil || len(matches) == 0 {
			return path
		}

		// Pick the lexically last match, which tends to correspond to the
		// highest version number for typical "vX.Y" style naming
		sort.Strings(matches)
		return matches[len(matches)-1]
	}

	return path
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
