package linux

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
	Namespace string   `json:"namespace"`
	AppID     string   `json:"appId"`
	Overrides []string `json:"overrides"`
	Arguments []string `json:"arguments"`
}

// Return if package is available
func (f *Flatpak) Available() bool {
	return cli.IsLinux()
}

// Return package runtime
func (f *Flatpak) Runtime() string {
	return "flatpak"
}

// Retrieve runtime directory based on namespace level
func (f *Flatpak) RuntimeDir() string {
	if f.Namespace == "user" {
		return fs.ExpandPath("$HOME/.local/share/flatpak")
	} else {
		return fs.NormalizePath("/var/lib/flatpak")
	}
}

// Apply program overrides
func (f *Flatpak) ApplyOverrides() error {
	if len(f.Overrides) == 0 {
		return nil
	}

	for _, override := range f.Overrides {
		script := fmt.Sprintf(`flatpak override --user %s %s`, override, f.AppID)
		err := cli.Run(script)
		if err != nil {
			return err
		}
	}

	return nil
}

// Install program
func (f *Flatpak) Install() error {

	// Install with CLI command
	script := fmt.Sprintf(
		`flatpak install --or-update --assumeyes --noninteractive --%s flathub %s`,
		f.Namespace,
		f.AppID,
	)

	err := cli.Run(script)
	if err != nil {
		return err
	}

	// Apply flatpak overrides
	err = f.ApplyOverrides()
	if err != nil {
		return err
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
	return fs.NormalizePath(fmt.Sprintf(
		`%s/exports/bin/%s`,
		f.RuntimeDir(),
		f.AppID,
	))
}

// Run installed program
func (f *Flatpak) Run(args []string) error {
	return cli.Start(fmt.Sprintf(
		`flatpak run %s %s`,
		f.AppID,
		strings.Join(args, " "),
	))
}

// Fill shortcut additional details
func (f *Flatpak) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for flatpak application
	shortcutDir := fs.NormalizePath("exports/share/applications")
	shortcutName := fmt.Sprintf("%s.desktop", f.AppID)
	shortcutPath := filepath.Join(f.RuntimeDir(), shortcutDir, shortcutName)

	shortcut.ShortcutPath = shortcutPath
	shortcut.LaunchOptions = strings.Join(f.Arguments, " ")

	return nil
}
