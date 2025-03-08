package macos

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Homebrew struct
type Homebrew struct {
	AppID     string   `json:"appId"`
	AppName   string   `json:"appName"`
	Arguments []string `json:"arguments"`
}

// Return package runtime
func (h *Homebrew) Runtime() string {
	return "native"
}

// Return if package is available
func (h *Homebrew) Available() bool {
	return cli.IsMacOS()
}

// Install package
func (h *Homebrew) Install() error {
	return cli.Run(fmt.Sprintf(
		`brew install --cask %s`,
		h.AppID,
	))
}

// Remove package
func (h *Homebrew) Remove() error {
	return cli.Run(fmt.Sprintf(
		`brew uninstall --cask %s`,
		h.AppID,
	))
}

// Installed verification
func (h *Homebrew) Installed() (bool, error) {
	exist, err := fs.FileExist(h.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (h *Homebrew) Executable() string {
	return fs.NormalizePath(fmt.Sprintf(
		`/Applications/%s`,
		h.AppName,
	))
}

// Return executable alias file path
func (h *Homebrew) Alias() string {
	return h.Executable()
}

// Run installed package
func (h *Homebrew) Run(args []string) error {
	return cli.RunProcess(h.Executable(), args)
}

// Fill shortcut additional details
func (h *Homebrew) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(h.Arguments, " ")
	return nil
}
