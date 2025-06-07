package docs

import (
	"embed"
)

//go:embed *
var fs embed.FS

// GetFS retrieve file system with embedded files
func GetFS() embed.FS {
	return fs
}

// Read template and return its content
func Read(file string) (string, error) {
	content, err := GetFS().ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
