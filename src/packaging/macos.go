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
	AppID     string   `json:"appId"`
	AppName   string   `json:"appName"`
	Arguments []string `json:"arguments"`
}

// Return if package is available
func (m *MacOS) Available() bool {
	return cli.IsMacOS()
}

// Return package runtime
func (m *MacOS) Runtime() string {
	return "native"
}

// Install program
func (m *MacOS) Install() error {
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
	return fs.ExpandPath(m.AppName)
}

// Run installed program
func (m *MacOS) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`open -n %s --args %s`,
		m.Executable(),
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (m *MacOS) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(m.Arguments, " ")
	return nil
}
