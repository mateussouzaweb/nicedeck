package linux

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Snap struct
type Snap struct {
	AppID     string   `json:"appId"`
	AppBin    string   `json:"appBin"`
	Channel   string   `json:"channel"`
	Arguments []string `json:"arguments"`
}

// Return package runtime
func (s *Snap) Runtime() string {
	return "snap"
}

// Return if package is available
func (s *Snap) Available() bool {
	return cli.IsLinux()
}

// Install package
func (s *Snap) Install() error {
	return cli.Run(fmt.Sprintf(
		`sudo snap install %s --channel=%s`,
		s.AppID,
		s.Channel,
	))
}

// Remove package
func (s *Snap) Remove() error {
	return cli.Run(fmt.Sprintf(
		`sudo snap remove %s`,
		s.AppID,
	))
}

// Installed verification
func (s *Snap) Installed() (bool, error) {
	exist, err := fs.FileExist(s.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (s *Snap) Executable() string {
	return fs.NormalizePath(fmt.Sprintf(
		`/snap/bin/%s`,
		s.AppBin,
	))
}

// Return executable alias file path
func (s *Snap) Alias() string {
	return fs.NormalizePath(fmt.Sprintf(
		"/var/lib/snapd/desktop/applications/%s_%s.desktop",
		s.AppID,
		s.AppID,
	))
}

// Run installed package
func (s *Snap) Run(args []string) error {
	return cli.RunProcess(s.Executable(), args)
}

// Fill shortcut additional details
func (s *Snap) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for snap application
	shortcut.ShortcutPath = s.Alias()
	shortcut.LaunchOptions = strings.Join(s.Arguments, " ")

	return nil
}
