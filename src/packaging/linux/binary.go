package linux

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Binary struct
type Binary struct {
	AppID     string                         `json:"appId"`
	AppBin    string                         `json:"appBin"`
	Arguments []string                       `json:"arguments"`
	Source    func() (string, string, error) `json:"-"`
}

// Return if package is available
func (b *Binary) Available() bool {
	return cli.IsLinux()
}

// Return package runtime
func (b *Binary) Runtime() string {
	return "native"
}

// Install program
func (b *Binary) Install() error {

	// Skip when cannot install
	if b.Source == nil {
		return nil
	}

	// Retrieve source details
	sourceURL, sourceType, err := b.Source()
	if err != nil {
		return err
	}

	// From ZIP format
	if sourceType == "zip" {

		// Download Zip
		destination := b.Executable()
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
		destination := b.Executable()
		err := fs.DownloadFile(sourceURL, destination, true)
		if err != nil {
			return err
		}
	}

	// Make sure is executable
	if installed, _ := b.Installed(); installed {
		executable := b.Executable()
		err := os.Chmod(executable, 0775)
		if err != nil {
			return err
		}
	}

	return nil
}

// Installed verification
func (b *Binary) Installed() (bool, error) {
	exist, err := fs.FileExist(b.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (b *Binary) Executable() string {
	return fs.ExpandPath(b.AppBin)
}

// Run installed program
func (b *Binary) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`exec %s %s`,
		b.Executable(),
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (b *Binary) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Write the desktop shortcut
	desktopShortcut, err := WriteDesktopShortcut(b.AppID, shortcut)
	if err != nil {
		return err
	}

	// Fill shortcut information for application
	shortcut.ShortcutPath = desktopShortcut
	shortcut.LaunchOptions = strings.Join(b.Arguments, " ")

	return nil
}
