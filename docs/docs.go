package docs

import (
	"embed"
	"strings"
)

//go:embed *.md
var files embed.FS

// Retrieve file content
func GetContent(file string, clean bool) (string, error) {

	bytes, err := files.ReadFile(file)
	content := string(bytes)

	// Clean basic markdown syntax
	if clean {
		content = strings.ReplaceAll(content, "# ", "")
		content = strings.ReplaceAll(content, "\\", "")
		content = strings.ReplaceAll(content, "**", "")
	}

	return string(content), err
}
