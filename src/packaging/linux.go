package packaging

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Linux struct
type Linux struct {
	AppID     string   `json:"appId"`
	AppBin    string   `json:"appBin"`
	Arguments []string `json:"arguments"`
}

// Return if package is available
func (l *Linux) Available() bool {
	return cli.IsLinux()
}

// Return package runtime
func (l *Linux) Runtime() string {
	return "native"
}

// Install program
func (l *Linux) Install() error {

	cli.Printf(cli.ColorWarn, "Warning: Unable to install Linux native packages.\n")
	cli.Printf(cli.ColorWarn, "Please make sure to manually download and install the program.\n")
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", l.Executable())

	// Make sure is executable
	if installed, _ := l.Installed(); installed {
		executable := l.Executable()
		err := os.Chmod(executable, 0775)
		if err != nil {
			return err
		}
	}

	return nil
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

// Fill shortcut additional details
func (l *Linux) OnShortcut(shortcut *shortcuts.Shortcut) error {
	shortcut.LaunchOptions = strings.Join(l.Arguments, " ")
	return nil
}
