package linux

import (
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Binary struct
type Binary struct {
	AppID     string               `json:"appId"`
	AppBin    string               `json:"appBin"`
	Arguments *packaging.Arguments `json:"arguments"`
	Source    *packaging.Source    `json:"source"`
}

// Return package runtime
func (b *Binary) Runtime() string {
	return "native"
}

// Return if package is available
func (b *Binary) Available() bool {
	return cli.IsLinux()
}

// Install package
func (b *Binary) Install() error {

	// Download from source
	if b.Source != nil {
		err := b.Source.Download(b)
		if err != nil {
			return err
		}

		// Make sure is executable
		if installed, _ := b.Installed(); installed {
			executable := b.Executable()
			err := os.Chmod(executable, 0775)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Remove package
func (b *Binary) Remove() error {

	// Remove executable parent folder
	// Because package is located in its own folder
	err := fs.RemoveDirectory(filepath.Dir(b.Executable()))
	if err != nil {
		return err
	}

	return nil
}

// Installed verification
func (b *Binary) Installed() (bool, error) {
	exist, err := fs.FileExist(b.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (b *Binary) Executable() string {
	return fs.ExpandPath(b.AppBin)
}

// Return executable alias file path
func (b *Binary) Alias() string {
	return ""
}

// Return executable arguments
func (b *Binary) Args() []string {
	return b.Arguments.Shortcut
}
