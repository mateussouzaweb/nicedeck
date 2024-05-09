package frontend

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed *
var content embed.FS

// Get static content with local or embedded files
func GetStaticFS(useLocalFS bool) fs.FS {
	if useLocalFS {
		return os.DirFS("./frontend")
	}
	return fs.FS(content)
}
