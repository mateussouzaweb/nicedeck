package frontend

import (
	"embed"
	"io/fs"
)

//go:embed *
var content embed.FS

// Get static content
func GetStaticFS() fs.FS {
	return fs.FS(content)
}
