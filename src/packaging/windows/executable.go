package windows

import (
	"fmt"
	"path/filepath"
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
	return RunProcess(e.Executable(), args)
}

// Fill shortcut additional details
func (e *Executable) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcutDir := fs.ExpandPath("$APPDATA\\Microsoft\\Windows\\Start Menu\\Programs")
	shortcutName := fmt.Sprintf("%s.lnk", shortcut.AppName)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)
	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = strings.Join(e.Arguments, " ")

	// Write system shortcut on start menu
	err := cli.Run(fmt.Sprintf(``+
		`$WshShell = New-Object -COMObject WScript.Shell;`+
		`$Shortcut = $WshShell.CreateShortcut("%s");`+
		`$Shortcut.WorkingDirectory = "%s";`+
		`$Shortcut.TargetPath = "%s";`+
		`$Shortcut.Arguments = "%s";`+
		`$Shortcut.Save()`,
		shortcut.ShortcutPath,
		shortcut.StartDir,
		shortcut.Exe,
		strings.ReplaceAll(shortcut.LaunchOptions, `"`, `\"`),
	))

	if err != nil {
		return err
	}

	return nil
}
