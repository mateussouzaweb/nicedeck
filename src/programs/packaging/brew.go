package packaging

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Brew struct
type Brew struct {
	AppID   string `json:"appId"`
	AppName string `json:"appName"`
}

// Return if package is available
func (b *Brew) Available() bool {
	return cli.IsMacOS()
}

// Install program with brew
func (b *Brew) Install(shortcut *shortcuts.Shortcut) error {
	return cli.Run(fmt.Sprintf(
		`brew install --cask %s`,
		b.AppID,
	))
}

// Installed verification
func (b *Brew) Installed() (bool, error) {
	exist, err := fs.FileExist(b.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (b *Brew) Executable() string {
	return fs.NormalizePath(fmt.Sprintf(
		`/Applications/%s.app`,
		b.AppName,
	))
}

// Run installed program
func (b *Brew) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`open -n %s --args %s`,
		b.Executable(),
		strings.Join(args, " "),
	))
}
