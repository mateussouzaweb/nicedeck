package linux

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

// Binary struct
type Binary struct {
	AppID     string               `json:"appId"`
	AppBin    string               `json:"appBin"`
	Arguments *packaging.Arguments `json:"arguments"`
	Source    *packaging.Source    `json:"source"`
}

// Return package runtime
func (b *Binary) Runtime() string {
	return "native"
}

// Return if package is available
func (b *Binary) Available() bool {
	return cli.IsLinux()
}

// Install package
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

// Remove package
func (b *Binary) Remove() error {

	// Remove executable parent folder
	// Because package is located in its own folder
	err := fs.RemoveDirectory(filepath.Dir(b.Executable()))
	if err != nil {
		return err
	}

	// Remove alias file
	err = fs.RemoveFile(b.Alias())
	if err != nil {
		return err
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

// Return executable alias file path
func (b *Binary) Alias() string {
	return fs.ExpandPath(fmt.Sprintf(
		"$SHARE/applications/%s.desktop",
		b.AppID,
	))
}

// Fill shortcut additional details
func (b *Binary) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for binary application
	shortcut.ShortcutPath = b.Alias()
	shortcut.LaunchOptions = strings.Join(b.Arguments.Shortcut, " ")

	// Write the desktop shortcut
	err := CreateDesktopShortcut(shortcut)
	if err != nil {
		return err
	}

	return nil
}
