package windows

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Executable struct
type Executable struct {
	AppID     string            `json:"appId"`
	AppExe    string            `json:"appExe"`
	Arguments []string          `json:"arguments"`
	Source    *packaging.Source `json:"source"`
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

	// Download from source
	if e.Source != nil {
		err := e.Source.Download(e)
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
