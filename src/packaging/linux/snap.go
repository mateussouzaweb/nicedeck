package linux

import (
	"fmt"
	"path/filepath"
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

// Return if package is available
func (s *Snap) Available() bool {
	return cli.IsLinux()
}

// Return package runtime
func (s *Snap) Runtime() string {
	return "snap"
}

// Install program
func (s *Snap) Install() error {

	// Install with CLI command
	script := fmt.Sprintf(
		`sudo snap install %s --channel=%s`,
		s.AppID,
		s.Channel,
	)

	err := cli.Run(script)
	if err != nil {
		return err
	}

	return nil
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

// Run installed program
func (s *Snap) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`exec %s %s`,
		s.Executable(),
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (s *Snap) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for snap application
	shortcutDir := fs.NormalizePath("/var/lib/snapd/desktop/applications")
	shortcutName := fmt.Sprintf("%s_%s.desktop", s.AppID, s.AppID)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)

	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = strings.Join(s.Arguments, " ")

	return nil
}
