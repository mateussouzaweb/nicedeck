package linux

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// AppImage struct
type AppImage struct {
	AppID     string                         `json:"appId"`
	AppName   string                         `json:"appName"`
	Arguments []string                       `json:"arguments"`
	Source    func() (string, string, error) `json:"-"`
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

	// Skip when cannot install
	if a.Source == nil {
		return nil
	}

	// Retrieve source details
	sourceURL, sourceType, err := a.Source()
	if err != nil {
		return err
	}

	// From ZIP format
	if sourceType == "zip" {

		// Download Zip
		destination := a.Executable()
		zipFile := fmt.Sprintf("%s.zip", destination)
		err := fs.DownloadFile(sourceURL, zipFile, true)
		if err != nil {
			return err
		}

		// Extract ZIP
		err = fs.Unzip(zipFile, destination)
		if err != nil {
			return err
		}

		// Remove ZIP file
		err = fs.RemoveFile(zipFile)
		if err != nil {
			return err
		}

	}

	// From direct file
	if sourceType == "file" {
		destination := a.Executable()
		err := fs.DownloadFile(sourceURL, destination, true)
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
	return cli.Start(fmt.Sprintf(
		`exec %s %s`,
		a.Executable(),
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (a *AppImage) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Write the desktop shortcut
	desktopShortcut, err := WriteDesktopShortcut(a.AppID, shortcut)
	if err != nil {
		return err
	}

	// Fill shortcut information for application
	shortcut.ShortcutPath = desktopShortcut
	shortcut.LaunchOptions = strings.Join(a.Arguments, " ")

	return nil
}
