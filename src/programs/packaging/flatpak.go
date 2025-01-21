package packaging

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Flatpak struct
type Flatpak struct {
	AppID     string   `json:"appId"`
	Overrides []string `json:"overrides"`
	Arguments []string `json:"arguments"`
}

// Return if package is available
func (f *Flatpak) Available() bool {
	return cli.IsLinux()
}

// Install program with flatpak
func (f *Flatpak) Install(shortcut *shortcuts.Shortcut) error {

	// Install with CLI command
	script := fmt.Sprintf(
		"flatpak install --or-update --assumeyes --noninteractive --system flathub %s",
		f.AppID,
	)

	err := cli.Run(script)
	if err != nil {
		return err
	}

	// Apply flatpak overrides
	if len(f.Overrides) > 0 {
		for _, override := range f.Overrides {
			script := fmt.Sprintf("flatpak override --user %s %s", override, f.AppID)
			err := cli.Run(script)
			if err != nil {
				return err
			}
		}
	}

	// Fill shortcut information for flatpak app
	executable := f.Executable()
	startDir := filepath.Dir(executable)
	shortcutDir := "/var/lib/flatpak/exports/share/applications"

	shortcut.StartDir = startDir
	shortcut.Exe = executable
	shortcut.ShortcutPath = shortcutDir + "/" + f.AppID + ".desktop"
	shortcut.LaunchOptions = ""

	// Append shortcut launch arguments
	if len(f.Arguments) > 0 {
		shortcut.LaunchOptions = strings.Join(f.Arguments, " ")
	}

	return nil
}

// Installed verification
func (f *Flatpak) Installed() (bool, error) {
	exist, err := fs.FileExist(f.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (f *Flatpak) Executable() string {
	return fmt.Sprintf(
		`/var/lib/flatpak/exports/bin/%s`,
		f.AppID,
	)
}

// Run installed program
func (f *Flatpak) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`flatpak run %s %s`,
		f.AppID,
		strings.Join(args, " "),
	))
}
