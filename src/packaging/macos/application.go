package macos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Application struct
type Application struct {
	AppID     string            `json:"appId"`
	AppName   string            `json:"appName"`
	Arguments []string          `json:"arguments"`
	Source    *packaging.Source `json:"source"`
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

	// Download from source
	if a.Source != nil {
		err := a.Source.Download(a)
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
	return cli.RunProcess(a.Executable(), args)
}

// Fill shortcut additional details
func (a *Application) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcutDir := fs.ExpandPath("$HOME/Applications")
	shortcutName := fmt.Sprintf("%s/%s.app", shortcut.Tags[0], shortcut.AppName)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)
	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = strings.Join(a.Arguments, " ")

	// Write the application shortcut
	err := os.MkdirAll(filepath.Dir(shortcut.ShortcutPath), 0755)
	if err != nil {
		return err
	}

	err = fs.RemoveFile(shortcut.ShortcutPath)
	if err != nil {
		return err
	}

	err = os.Symlink(shortcut.Exe, shortcut.ShortcutPath)
	if err != nil {
		return err
	}

	return nil
}
