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
func (a *AppImage) Install(shortcut *shortcuts.Shortcut) error {

	// Run before install callback
	// Used to dynamic fetch the app download URL
	err := a.BeforeInstall(a)
	if err != nil {
		return err
	}

	// Download application
	executable := a.Executable()
	err = fs.DownloadFile(a.AppURL, executable, true)
	if err != nil {
		return err
	}

	// Make sure is executable
	err = os.Chmod(executable, 0775)
	if err != nil {
		return err
	}

	// Run after install callback
	// Used to do things like write desktop shortcut or settings
	err = a.AfterInstall(a)
	if err != nil {
		return err
	}

	// Fill shortcut information
	shortcutDir := fs.ExpandPath("$HOME/.local/share/applications")
	shortcutName := fmt.Sprintf("%s.desktop", a.AppID)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)

	shortcut.Exe = executable
	shortcut.StartDir = filepath.Dir(executable)
	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = ""

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
	return fs.ExpandPath(fmt.Sprintf(
		`$APPLICATIONS/%s`,
		a.AppName,
	))
}

// Run installed program
func (a *AppImage) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`exec %s %s`,
		a.Executable(),
		strings.Join(args, " "),
	))
}
