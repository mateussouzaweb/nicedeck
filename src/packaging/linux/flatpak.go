package linux

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Flatpak struct
type Flatpak struct {
	AppID     string               `json:"appId"`
	Namespace string               `json:"namespace"`
	Overrides []string             `json:"overrides"`
	Arguments *packaging.Arguments `json:"arguments"`
}

// Return package runtime
func (f *Flatpak) Runtime() string {
	return "flatpak"
}

// Retrieve runtime directory based on namespace level
func (f *Flatpak) RuntimeDir() string {
	if f.Namespace == "user" {
		return fs.ExpandPath("$SHARE/flatpak")
	} else {
		return fs.NormalizePath("/var/lib/flatpak")
	}
}

// Return if package is available
func (f *Flatpak) Available() bool {
	return cli.IsLinux()
}

// Apply package overrides
func (f *Flatpak) ApplyOverrides() error {
	if len(f.Overrides) == 0 {
		return nil
	}

	for _, override := range f.Overrides {
		script := fmt.Sprintf(`flatpak override --user %s %s`, override, f.AppID)
		command := cli.Command(script)
		err := cli.Run(command)
		if err != nil {
			return err
		}
	}

	return nil
}

// Install package
func (f *Flatpak) Install() error {

	// Install with CLI command
	script := fmt.Sprintf(
		`flatpak install --or-update --assumeyes --noninteractive --%s flathub %s %s`,
		f.Namespace,
		f.AppID,
		strings.Join(f.Arguments.Install, " "),
	)

	command := cli.Command(script)
	err := cli.Run(command)
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

// Remove package
func (f *Flatpak) Remove() error {
	script := fmt.Sprintf(
		`flatpak uninstall --assumeyes --noninteractive --%s %s %s`,
		f.Namespace,
		f.AppID,
		strings.Join(f.Arguments.Remove, " "),
	)

	command := cli.Command(script)
	return cli.Run(command)
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

// Return executable alias file path
func (f *Flatpak) Alias() string {
	return filepath.Join(f.RuntimeDir(), fs.NormalizePath(
		fmt.Sprintf("exports/share/applications/%s.desktop", f.AppID),
	))
}

// Run installed package
func (f *Flatpak) Run(arguments []string) error {
	arguments = append(f.Arguments.Run, arguments...)
	return cli.RunProcess(f.Executable(), arguments)
}

// Fill shortcut additional details
func (f *Flatpak) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for flatpak application
	shortcut.ShortcutPath = f.Alias()
	shortcut.LaunchOptions = strings.Join(f.Arguments.Shortcut, " ")

	return nil
}
