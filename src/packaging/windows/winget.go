package windows

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// WinGet struct
type WinGet struct {
	AppID     string   `json:"appId"`
	AppExe    string   `json:"appExe"`
	Arguments []string `json:"arguments"`
}

// Return package runtime
func (w *WinGet) Runtime() string {
	return "native"
}

// Return if package is available
func (w *WinGet) Available() bool {
	return cli.IsWindows()
}

// Install package
func (w *WinGet) Install() error {
	return cli.Run(fmt.Sprintf(
		`winget install --accept-package-agreements --accept-source-agreements --disable-interactivity --exact --id %s`,
		w.AppID,
	))
}

// Remove package
func (w *WinGet) Remove() error {
	return cli.Run(fmt.Sprintf(
		`winget uninstall --disable-interactivity --exact --id %s`,
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

// Return executable alias file path
func (w *WinGet) Alias() string {
	return w.Executable()
}

// Run installed package
func (w *WinGet) Run(args []string) error {
	return cli.RunProcess(w.Executable(), args)
}

// Fill shortcut additional details
func (w *WinGet) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(w.Arguments, " ")
	return nil
}
