package macos

import (
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Application struct
type Application struct {
	AppID     string               `json:"appId"`
	AppName   string               `json:"appName"`
	Arguments *packaging.Arguments `json:"arguments"`
	Source    *packaging.Source    `json:"source"`
}

// Return package runtime
func (a *Application) Runtime() string {
	return "native"
}

// Return if package is available
func (a *Application) Available() bool {
	return cli.IsMacOS()
}

// Install package
func (a *Application) Install() error {

	// Download from source
	if a.Source != nil {
		err := a.Source.Download(a)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove package
func (a *Application) Remove() error {

	// Remove executable parent folder
	// Because package is located in its own folder
	err := fs.RemoveDirectory(filepath.Dir(a.Executable()))
	if err != nil {
		return err
	}

	return nil
}

// Installed verification
func (a *Application) Installed() (bool, error) {

	// AppName.app files are considered a directory
	exist, err := fs.DirectoryExist(a.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (a *Application) Executable() string {
	return fs.ExpandPath(a.AppName)
}

// Return executable alias file path
func (a *Application) Alias() string {
	return ""
}

// Return executable arguments
func (a *Application) Args() []string {
	return a.Arguments.Shortcut
}
