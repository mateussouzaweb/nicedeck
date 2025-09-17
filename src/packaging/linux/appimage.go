package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
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
	}

	// Make sure is executable
	if installed, _ := a.Installed(); installed {
		executable := a.Executable()
		err := os.Chmod(executable, 0775)
		if err != nil {
			return err
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

	// Remove alias file
	err = fs.RemoveFile(a.Alias())
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
	return fs.ExpandPath(fmt.Sprintf(
		"$SHARE/applications/%s.desktop",
		a.AppID,
	))
}

// Run installed package
func (a *AppImage) Run(arguments []string) error {
	arguments = append(a.Arguments.Run, arguments...)
	return cli.RunProcess(a.Executable(), arguments)
}

// Fill shortcut additional details
func (a *AppImage) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcut.ShortcutPath = a.Alias()
	shortcut.LaunchOptions = strings.Join(a.Arguments.Shortcut, " ")

	// Write the desktop shortcut
	err := CreateDesktopShortcut(shortcut)
	if err != nil {
		return err
	}

	return nil
}
