package linux

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

// Binary struct
type Binary struct {
	AppID     string            `json:"appId"`
	AppBin    string            `json:"appBin"`
	Arguments []string          `json:"arguments"`
	Source    *packaging.Source `json:"source"`
}

// Return if package is available
func (b *Binary) Available() bool {
	return cli.IsLinux()
}

// Return package runtime
func (b *Binary) Runtime() string {
	return "native"
}

// Install program
func (b *Binary) Install() error {

	// Download from source
	if b.Source != nil {
		err := b.Source.Download(b)
		if err != nil {
			return err
		}
	}

	// Make sure is executable
	if installed, _ := b.Installed(); installed {
		executable := b.Executable()
		err := os.Chmod(executable, 0775)
		if err != nil {
			return err
		}
	}

	return nil
}

// Installed verification
func (b *Binary) Installed() (bool, error) {
	exist, err := fs.FileExist(b.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (b *Binary) Executable() string {
	return fs.ExpandPath(b.AppBin)
}

// Run installed program
func (b *Binary) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`cd "%s" && exec "%s" %s`,
		filepath.Dir(b.Executable()),
		b.Executable(),
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (b *Binary) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for application
	shortcutDir := fs.ExpandPath("$HOME/.local/share/applications")
	shortcutName := fmt.Sprintf("%s.desktop", b.AppID)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)

	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = strings.Join(b.Arguments, " ")

	// Write the desktop shortcut
	err := WriteDesktopShortcut(b.AppID, shortcutPath, shortcut)
	if err != nil {
		return err
	}

	return nil
}
