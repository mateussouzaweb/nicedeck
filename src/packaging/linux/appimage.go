package linux

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

// AppImage struct
type AppImage struct {
	AppID     string            `json:"appId"`
	AppName   string            `json:"appName"`
	Arguments []string          `json:"arguments"`
	Source    *packaging.Source `json:"source"`
}

// Return if package is available
func (a *AppImage) Available() bool {
	return cli.IsLinux()
}

// Return package runtime
func (a *AppImage) Runtime() string {
	return "appimage"
}

// Install program
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

// Run installed program
func (a *AppImage) Run(args []string) error {
	return cli.RunProcess(a.Executable(), args)
}

// Fill shortcut additional details
func (a *AppImage) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcutDir := fs.ExpandPath("$SHARE/applications")
	shortcutName := fmt.Sprintf("%s.desktop", a.AppID)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)
	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = strings.Join(a.Arguments, " ")

	// Write the desktop shortcut
	err := CreateDesktopShortcut(shortcut)
	if err != nil {
		return err
	}

	return nil
}
