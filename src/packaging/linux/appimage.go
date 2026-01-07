package linux

import (
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// AppImage struct
type AppImage struct {
	AppID     string               `json:"appId"`
	AppName   string               `json:"appName"`
	Arguments *packaging.Arguments `json:"arguments"`
	Source    *packaging.Source    `json:"source"`
}

// Return package runtime
func (a *AppImage) Runtime() string {
	return "appimage"
}

// Return if package is available
func (a *AppImage) Available() bool {
	return cli.IsLinux()
}

// Install package
func (a *AppImage) Install() error {

	// Download from source
	if a.Source != nil {
		err := a.Source.Download(a)
		if err != nil {
			return err
		}

		// Make sure is executable
		if installed, _ := a.Installed(); installed {
			executable := a.Executable()
			err := os.Chmod(executable, 0775)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Remove package
func (a *AppImage) Remove() error {

	// Remove executable parent folder
	// Because package is located in its own folder
	err := fs.RemoveDirectory(filepath.Dir(a.Executable()))
	if err != nil {
		return err
	}

	return nil
}

// Installed verification
func (a *AppImage) Installed() (bool, error) {
	exist, err := fs.FileExist(a.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (a *AppImage) Executable() string {
	return fs.ExpandPath(a.AppName)
}

// Return executable alias file path
func (a *AppImage) Alias() string {
	return ""
}

// Return executable arguments
func (a *AppImage) Args() []string {
	return a.Arguments.Shortcut
}
