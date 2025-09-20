package windows

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Executable struct
type Executable struct {
	AppID     string               `json:"appId"`
	AppExe    string               `json:"appExe"`
	AppAlias  string               `json:"appAlias"`
	Arguments *packaging.Arguments `json:"arguments"`
	Source    *packaging.Source    `json:"source"`
}

// Return package runtime
func (e *Executable) Runtime() string {
	return "native"
}

// Return if package is available
func (e *Executable) Available() bool {
	return cli.IsWindows()
}

// Install package
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

// Remove package
func (e *Executable) Remove() error {

	// Remove executable parent folder
	// Because package is located in its own folder
	err := fs.RemoveDirectory(filepath.Dir(e.Executable()))
	if err != nil {
		return err
	}

	// Remove alias file
	err = fs.RemoveFile(e.Alias())
	if err != nil {
		return err
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

// Return executable alias file path
func (e *Executable) Alias() string {
	return fs.ExpandPath(e.AppAlias)
}

// Fill shortcut additional details
func (e *Executable) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcut.ShortcutPath = e.Alias()
	shortcut.LaunchOptions = strings.Join(e.Arguments.Shortcut, " ")

	// Write system alias on shortcut location
	err := os.MkdirAll(filepath.Dir(shortcut.ShortcutPath), 0755)
	if err != nil {
		return err
	}

	script := fmt.Sprintf(``+
		`$WshShell = New-Object -COMObject WScript.Shell;`+
		`$Shortcut = $WshShell.CreateShortcut("%s");`+
		`$Shortcut.WorkingDirectory = "%s";`+
		`$Shortcut.TargetPath = "%s";`+
		`$Shortcut.Arguments = "%s";`+
		`$Shortcut.Save()`,
		shortcut.ShortcutPath,
		strings.ReplaceAll(shortcut.StartDirectory, `"`, ``),
		strings.ReplaceAll(shortcut.Executable, `"`, ``),
		strings.ReplaceAll(shortcut.LaunchOptions, `"`, `\"`),
	)

	command := cli.Command(script)
	err = cli.Run(command)
	if err != nil {
		return err
	}

	return nil
}
