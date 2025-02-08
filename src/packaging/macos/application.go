package macos

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Application struct
type Application struct {
	AppID     string                         `json:"appId"`
	AppName   string                         `json:"appName"`
	Arguments []string                       `json:"arguments"`
	Source    func() (string, string, error) `json:"-"`
}

// Return if package is available
func (a *Application) Available() bool {
	return cli.IsMacOS()
}

// Return package runtime
func (a *Application) Runtime() string {
	return "native"
}

// Install program
func (a *Application) Install() error {

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

	// From DMG format
	if sourceType == "dmg" {

		// Download DMG
		destination := a.Executable()
		dmgFile := strings.Replace(destination, ".app", ".dmg", 1)
		err = fs.DownloadFile(sourceURL, dmgFile, true)
		if err != nil {
			return err
		}

		// Extract application from DMG
		appName := filepath.Base(destination)
		script := fmt.Sprintf(
			`hdiutil attach %s
			cp /Volumes/%s %s
			hdiutil detach /Volumes/%s`,
			dmgFile,
			appName,
			destination,
			appName,
		)

		err = cli.Run(script)
		if err != nil {
			return err
		}

		// Remove DMG file
		err = fs.RemoveFile(dmgFile)
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

	// Add program to quarantine
	if installed, _ := a.Installed(); installed {
		script := fmt.Sprintf(`xattr -r -d com.apple.quarantine %s`, a.Executable())
		err := cli.Run(script)
		if err != nil {
			return err
		}
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

// Run installed program
func (a *Application) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`open -n %s --args %s`,
		a.Executable(),
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (a *Application) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(a.Arguments, " ")
	return nil
}
