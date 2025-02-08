package windows

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// WinGet struct
type WinGet struct {
	AppID     string   `json:"appId"`
	AppExe    string   `json:"appExe"`
	Arguments []string `json:"arguments"`
}

// Return if package is available
func (w *WinGet) Available() bool {
	return cli.IsWindows()
}

// Return package runtime
func (w *WinGet) Runtime() string {
	return "native"
}

// Install program
func (w *WinGet) Install() error {
	return cli.Run(fmt.Sprintf(
		`winget install --accept-package-agreements --accept-source-agreements --disable-interactivity --exact --id %s`,
		w.AppID,
	))
}

// Installed verification
func (w *WinGet) Installed() (bool, error) {
	exist, err := fs.FileExist(w.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (w *WinGet) Executable() string {
	return fs.ExpandPath(w.AppExe)
}

// Run installed program
func (w *WinGet) Run(args []string) error {
	if len(args) > 0 {
		return cli.Start(fmt.Sprintf(``+
			`$arguments = '%s';`+
			`Start-Process -FilePath "%s" -PassThru -Wait -ArgumentList $arguments`,
			strings.Join(args, " "),
			w.Executable(),
		))
	}

	return cli.Start(fmt.Sprintf(
		`Start-Process -FilePath "%s" -PassThru -Wait`,
		w.Executable(),
	))
}

// Fill shortcut additional details
func (w *WinGet) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(w.Arguments, " ")
	return nil
}
