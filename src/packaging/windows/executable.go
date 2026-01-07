package windows

import (
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Executable struct
type Executable struct {
	AppID     string               `json:"appId"`
	AppExe    string               `json:"appExe"`
	Arguments *packaging.Arguments `json:"arguments"`
	Source    *packaging.Source    `json:"source"`
}

// Return package runtime
func (e *Executable) Runtime() string {
	return "native"
}

// Return if package is available
func (e *Executable) Available() bool {
	return cli.IsWindows()
}

// Install package
func (e *Executable) Install() error {

	// Download from source
	if e.Source != nil {
		err := e.Source.Download(e)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove package
func (e *Executable) Remove() error {

	// Remove executable parent folder
	// Because package is located in its own folder
	err := fs.RemoveDirectory(filepath.Dir(e.Executable()))
	if err != nil {
		return err
	}

	return nil
}

// Installed verification
func (e *Executable) Installed() (bool, error) {
	exist, err := fs.FileExist(e.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (e *Executable) Executable() string {
	return fs.ExpandPath(e.AppExe)
}

// Return executable alias file path
func (e *Executable) Alias() string {
	return ""
}

// Return executable arguments
func (e *Executable) Args() []string {
	return e.Arguments.Shortcut
}
