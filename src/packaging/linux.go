package packaging

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Linux struct
type Linux struct {
	AppID  string `json:"appId"`
	AppBin string `json:"appBin"`
}

// Return if package is available
func (l *Linux) Available() bool {
	return cli.IsLinux()
}

// Install program
func (l *Linux) Install(shortcut *shortcuts.Shortcut) error {
	return fmt.Errorf("cannot perform package installations")
}

// Installed verification
func (l *Linux) Installed() (bool, error) {
	exist, err := fs.FileExist(l.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (l *Linux) Executable() string {
	return fs.ExpandPath(l.AppBin)
}

// Run installed program
func (l *Linux) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`exec %s %s`,
		l.Executable(),
		strings.Join(args, " "),
	))
}
