package packaging

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Homebrew struct
type Homebrew struct {
	AppID   string `json:"appId"`
	AppName string `json:"appName"`
}

// Return if package is available
func (h *Homebrew) Available() bool {
	return cli.IsMacOS()
}

// Install program
func (h *Homebrew) Install(shortcut *shortcuts.Shortcut) error {
	return cli.Run(fmt.Sprintf(
		`brew install --cask %s`,
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

// Run installed program
func (h *Homebrew) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`open -n %s --args %s`,
		h.Executable(),
		strings.Join(args, " "),
	))
}
