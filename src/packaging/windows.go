package packaging

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Windows struct
type Windows struct {
	AppID     string   `json:"appId"`
	AppExe    string   `json:"appExe"`
	Arguments []string `json:"arguments"`
}

// Return if package is available
func (w *Windows) Available() bool {
	return cli.IsWindows()
}

// Return package runtime
func (w *Windows) Runtime() string {
	return "native"
}

// Install program
func (w *Windows) Install() error {

	cli.Printf(cli.ColorWarn, "Warning: Unable to install Windows native packages.")
	cli.Printf(cli.ColorWarn, "Warning: Please make sure to manually download and install the program.")
	cli.Printf(cli.ColorWarn, "Warning: Expected executable: %s", w.Executable())

	return nil
}

// Installed verification
func (w *Windows) Installed() (bool, error) {
	exist, err := fs.FileExist(w.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (w *Windows) Executable() string {
	return fs.ExpandPath(w.AppExe)
}

// Run installed program
func (w *Windows) Run(args []string) error {
	if len(args) > 0 {
		return cli.Start(fmt.Sprintf(
			`Start-Process -FilePath "%s" -ArgumentList "%s" -PassThru -Wait`,
			w.Executable(),
			strings.Join(args, " "),
		))
	}

	return cli.Start(fmt.Sprintf(
		`Start-Process -FilePath "%s" -PassThru -Wait`,
		w.Executable(),
	))
}

// Fill shortcut additional details
func (w *Windows) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(w.Arguments, " ")
	return nil
}
