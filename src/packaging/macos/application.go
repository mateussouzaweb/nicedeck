package macos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Application struct
type Application struct {
	AppID     string            `json:"appId"`
	AppName   string            `json:"appName"`
	AppAlias  string            `json:"appAlias"`
	Arguments []string          `json:"arguments"`
	Source    *packaging.Source `json:"source"`
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

	// Add package to quarantine
	if installed, _ := a.Installed(); installed {
		script := fmt.Sprintf(`xattr -r -d com.apple.quarantine %s`, a.Executable())
		err := cli.Run(script)
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

	// Remove alias file
	err = fs.RemoveFile(a.Alias())
	if err != nil {
		return err
	}

	return nil
}

// Installed verification
func (a *Application) Installed() (bool, error) {
	exist, err := fs.FileExist(a.Executable())
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
	return fs.ExpandPath(a.AppAlias)
}

// Run installed package
func (a *Application) Run(args []string) error {
	return cli.RunProcess(a.Executable(), args)
}

// Fill shortcut additional details
func (a *Application) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcut.ShortcutPath = a.Alias()
	shortcut.LaunchOptions = strings.Join(a.Arguments, " ")

	// Write the application shortcut
	err := os.MkdirAll(filepath.Dir(shortcut.ShortcutPath), 0755)
	if err != nil {
		return err
	}

	err = fs.RemoveFile(shortcut.ShortcutPath)
	if err != nil {
		return err
	}

	err = os.Symlink(shortcut.Exe, shortcut.ShortcutPath)
	if err != nil {
		return err
	}

	return nil
}
