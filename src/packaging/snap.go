package packaging

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

// Install program
func (s *Snap) Install(shortcut *shortcuts.Shortcut) error {

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

	// Fill shortcut information for snap application
	executable := s.Executable()
	startDir := filepath.Dir(executable)
	shortcutDir := fs.NormalizePath("/var/lib/snapd/desktop/applications")
	shortcutName := fmt.Sprintf("%s_%s.desktop", s.AppID, s.AppID)
	shortcutPath := filepath.Join(shortcutDir, shortcutName)

	shortcut.StartDir = startDir
	shortcut.Exe = executable
	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = ""

	// Append shortcut launch arguments
	if len(s.Arguments) > 0 {
		shortcut.LaunchOptions = strings.Join(s.Arguments, " ")
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
