package windows

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// WinGet struct
type WinGet struct {
	AppID     string               `json:"appId"`
	AppExe    string               `json:"appExe"`
	Arguments *packaging.Arguments `json:"arguments"`
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
		`winget install --accept-package-agreements --accept-source-agreements --disable-interactivity --exact --id %s %s`,
		w.AppID,
		strings.Join(w.Arguments.Install, " "),
	))
}

// Remove package
func (w *WinGet) Remove() error {
	return cli.Run(fmt.Sprintf(
		`winget uninstall --disable-interactivity --exact --id %s %s`,
		w.AppID,
		strings.Join(w.Arguments.Remove, " "),
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
func (w *WinGet) Run(arguments []string) error {
	arguments = append(w.Arguments.Run, arguments...)
	return cli.RunProcess(w.Executable(), arguments)
}

// Fill shortcut additional details
func (w *WinGet) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(w.Arguments.Shortcut, " ")
	return nil
}
