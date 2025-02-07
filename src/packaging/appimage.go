package packaging

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

type AppImageCallback func(a *AppImage) error

// AppImage struct
type AppImage struct {
	AppID         string           `json:"appId"`
	AppName       string           `json:"appName"`
	AppURL        string           `json:"appUrl"`
	Arguments     []string         `json:"arguments"`
	BeforeInstall AppImageCallback `json:"-"`
	AfterInstall  AppImageCallback `json:"-"`
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

	// Run before install callback
	// Used to dynamic fetch the app download URL
	if a.BeforeInstall != nil {
		err := a.BeforeInstall(a)
		if err != nil {
			return err
		}
	}

	// Download application when possible
	if a.AppURL != "" {
		executable := a.Executable()
		err := fs.DownloadFile(a.AppURL, executable, true)
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

	// Run after install callback
	// Used to do things like write desktop shortcut or settings
	if a.AfterInstall != nil {
		err := a.AfterInstall(a)
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
	return cli.Start(fmt.Sprintf(
		`exec %s %s`,
		a.Executable(),
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (a *AppImage) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcutDir := fs.ExpandPath("$HOME/.local/share/applications")
	shortcutName := fmt.Sprintf("%s.desktop", a.AppID)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)

	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = strings.Join(a.Arguments, " ")

	return nil
}
