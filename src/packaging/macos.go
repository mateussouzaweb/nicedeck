package packaging

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// MacOS struct
type MacOS struct {
	AppID   string `json:"appId"`
	AppName string `json:"appName"`
}

// Return if package is available
func (m *MacOS) Available() bool {
	return cli.IsMacOS()
}

// Install program
func (m *MacOS) Install(shortcut *shortcuts.Shortcut) error {
	return fmt.Errorf("cannot perform package installations")
}

// Installed verification
func (m *MacOS) Installed() (bool, error) {
	exist, err := fs.FileExist(m.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (m *MacOS) Executable() string {
	return fs.NormalizePath(fmt.Sprintf(
		`/Applications/%s`,
		m.AppName,
	))
}

// Run installed program
func (m *MacOS) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`open -n %s --args %s`,
		m.Executable(),
		strings.Join(args, " "),
	))
}
