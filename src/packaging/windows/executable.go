package windows

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Executable struct
type Executable struct {
	AppID     string                         `json:"appId"`
	AppExe    string                         `json:"appExe"`
	Arguments []string                       `json:"arguments"`
	Source    func() (string, string, error) `json:"-"`
}

// Return if package is available
func (e *Executable) Available() bool {
	return cli.IsWindows()
}

// Return package runtime
func (e *Executable) Runtime() string {
	return "native"
}

// Install program
func (e *Executable) Install() error {

	// Skip when cannot install
	if e.Source == nil {
		return nil
	}

	// Retrieve source details
	sourceURL, sourceType, err := e.Source()
	if err != nil {
		return err
	}

	// From ZIP format
	if sourceType == "zip" {

		// Download Zip
		destination := e.Executable()
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
		destination := e.Executable()
		err := fs.DownloadFile(sourceURL, destination, true)
		if err != nil {
			return err
		}
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
	return fs.ExpandPath(e.AppExe)
}

// Run installed program
func (e *Executable) Run(args []string) error {
	if len(args) > 0 {
		return cli.Start(fmt.Sprintf(
			`Start-Process -FilePath "%s" -ArgumentList "%s" -PassThru -Wait`,
			e.Executable(),
			strings.Join(args, " "),
		))
	}

	return cli.Start(fmt.Sprintf(
		`Start-Process -FilePath "%s" -PassThru -Wait`,
		e.Executable(),
	))
}

// Fill shortcut additional details
func (e *Executable) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(e.Arguments, " ")
	return nil
}
