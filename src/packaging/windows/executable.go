package windows

import (
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Executable struct
type Executable struct {
	AppID       string               `json:"appId"`
	Installer   string               `json:"installer"`
	Uninstaller string               `json:"uninstaller"`
	Launcher    string               `json:"launcher"`
	Arguments   *packaging.Arguments `json:"arguments"`
	Source      *packaging.Source    `json:"source"`
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

	// Download program from source
	// When no need for installer, download from source
	// When installer is needed, download from source and makes verification to check at the installer file
	if e.Source != nil && e.Installer != "" {
		originalLauncher := e.Launcher
		e.Launcher = e.Installer
		defer func() {
			e.Launcher = originalLauncher
		}()

		err := e.Source.Download(e)
		if err != nil {
			return err
		}

	} else if e.Source != nil {
		err := e.Source.Download(e)
		if err != nil {
			return err
		}
	}

	// When installer is needed, run installer
	if e.Installer != "" {
		cli.Debug("Running install for %s\n", e.AppID)

		// Run install script
		executable := e.Installer
		arguments := e.Arguments.Install
		directory := filepath.Dir(executable)

		context := &cli.Context{
			WorkingDirectory: directory,
			Executable:       executable,
			Arguments:        arguments,
			Environment:      []string{},
		}

		err := context.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove package
func (e *Executable) Remove() error {

	// When uninstaller is needed, run uninstaller
	if e.Uninstaller != "" {
		cli.Debug("Running uninstall for %s\n", e.AppID)

		// Run uninstall script
		executable := e.Uninstaller
		arguments := e.Arguments.Remove
		directory := filepath.Dir(executable)

		context := &cli.Context{
			WorkingDirectory: directory,
			Executable:       executable,
			Arguments:        arguments,
			Environment:      []string{},
		}

		err := context.Run()
		if err != nil {
			return err
		}
	}

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
	return fs.ExpandPath(e.Launcher)
}

// Return executable alias file path
func (e *Executable) Alias() string {
	return ""
}

// Return executable arguments
func (e *Executable) Args() []string {
	return e.Arguments.Shortcut
}
