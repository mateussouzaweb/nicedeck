package windows

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Retrieve desktop entry shortcut path
func GetShortcutPath(shortcut *shortcuts.Shortcut) string {

	if slices.Contains(shortcut.Tags, "Gaming") {
		return fs.ExpandPath(fmt.Sprintf(
			"$START_MENU/Gaming/%s.lnk",
			shortcut.Name,
		))
	}

	return fs.ExpandPath(fmt.Sprintf(
		"$START_MENU/%s.lnk",
		shortcut.Name,
	))
}

// Create a desktop entry from shortcut data
func CreateShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Ensure shortcut directory exists
	err := os.MkdirAll(filepath.Dir(destination), 0755)
	if err != nil {
		return err
	}

	// Write system shortcut on location using PowerShell
	script := fmt.Sprintf(``+
		`$WshShell = New-Object -COMObject WScript.Shell;`+
		`$Shortcut = $WshShell.CreateShortcut("%s");`+
		`$Shortcut.WorkingDirectory = "%s";`+
		`$Shortcut.TargetPath = "%s";`+
		`$Shortcut.Arguments = "%s";`+
		`$Shortcut.Save()`,
		destination,
		strings.ReplaceAll(shortcut.StartDirectory, `"`, ``),
		strings.ReplaceAll(shortcut.Executable, `"`, ``),
		strings.ReplaceAll(shortcut.LaunchOptions, `"`, `\"`),
	)

	command := cli.Command(script)
	err = cli.Run(command)
	if err != nil {
		return err
	}

	return nil
}

// Remove desktop entry shortcut
func RemoveShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	err := fs.RemoveFile(destination)
	if err != nil {
		return err
	}

	return nil
}
