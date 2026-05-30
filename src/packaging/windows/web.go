package windows

import (
	"fmt"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Web struct
type Web struct {
	AppID     string               `json:"appId"`
	AppURL    string               `json:"appURL"`
	Wrapper   *packaging.Program   `json:"wrapper"`
	Arguments *packaging.Arguments `json:"arguments"`
}

// Return package runtime
func (w *Web) Runtime() string {
	return "web"
}

// Return if package is available
func (w *Web) Available() bool {
	return cli.IsWindows() && w.Wrapper.Package.Available()
}

// Return package flag path
func (w *Web) Flag() string {
	baseFlags := filepath.Join(fs.ExpandPath("$CONFIG"), ".browser")
	appFlag := fmt.Sprintf("%s-%s.flag", w.Wrapper.ID, w.AppID)
	flag := filepath.Join(baseFlags, appFlag)
	return flag
}

// Install package
func (w *Web) Install() error {

	// Make sure wrapper is installed before installing the web app
	wrapperInstalled, err := w.Wrapper.Package.Installed()
	if err != nil {
		return err
	}

	if !wrapperInstalled {
		err = w.Wrapper.Package.Install()
		if err != nil {
			return err
		}
	}

	// Write flag file to indicate that the web app is installed
	err = fs.WriteFile(w.Flag(), w.AppID)
	if err != nil {
		return err
	}

	return nil
}

// Remove package
func (w *Web) Remove() error {

	// Remove flag file if it exists
	err := fs.RemoveFile(w.Flag())
	if err != nil {
		return err
	}

	return nil
}

// Installed verification
func (w *Web) Installed() (bool, error) {

	// Check if flag file exist
	exist, err := fs.FileExist(w.Flag())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (w *Web) Executable() string {
	return w.Wrapper.Package.Executable()
}

// Return executable alias file path
func (w *Web) Alias() string {
	return ""
}

// Return executable arguments
func (w *Web) Args() []string {
	args := w.Wrapper.Package.Args()
	args = append(args, w.Arguments.Shortcut...)
	args = append(args, fmt.Sprintf("--app=%s", w.AppURL))
	return args
}
