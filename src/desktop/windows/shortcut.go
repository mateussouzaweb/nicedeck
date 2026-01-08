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

	// Remove invalid characters from name
	// Windows does not allow certain characters in file names
	name := shortcut.Name
	name = strings.ReplaceAll(name, "<", "")
	name = strings.ReplaceAll(name, ">", "")
	name = strings.ReplaceAll(name, "\\", "")
	name = strings.ReplaceAll(name, "/", "")
	name = strings.ReplaceAll(name, "|", "")
	name = strings.ReplaceAll(name, ":", "")
	name = strings.ReplaceAll(name, "?", "")
	name = strings.ReplaceAll(name, "*", "")
	name = strings.ReplaceAll(name, "\"", "")

	// Check for specific categories to place shortcut accordingly
	categories := []string{"Gaming", "Utilities"}
	for _, category := range categories {
		if !slices.Contains(shortcut.Tags, category) {
			continue
		}

		return fs.ExpandPath(fmt.Sprintf(
			"$START_MENU/%s/%s.lnk", category, name,
		))
	}

	// Default location
	return fs.ExpandPath(fmt.Sprintf(
		"$START_MENU/%s.lnk", name,
	))
}

// Create a desktop entry from shortcut data
func CreateShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Ensure shortcut directory exists
	err := os.MkdirAll(filepath.Dir(destination), 0755)
	if err != nil {
		return err
	}

	// Detect icon path or defaults to executable icon
	// Windows accepts only .ico format for shortcuts
	iconPath := shortcut.Executable
	if strings.HasSuffix(shortcut.IconPath, ".ico") {
		iconPath = shortcut.IconPath
	}

	// Write system shortcut on location using PowerShell
	script := fmt.Sprintf(``+
		`$WshShell = New-Object -COMObject WScript.Shell;`+
		`$Shortcut = $WshShell.CreateShortcut("%s");`+
		`$Shortcut.WorkingDirectory = "%s";`+
		`$Shortcut.TargetPath = "%s";`+
		`$Shortcut.Arguments = "%s";`+
		`$Shortcut.IconLocation = "%s,0";`+
		`$Shortcut.Save()`,
		destination,
		strings.ReplaceAll(shortcut.StartDirectory, `"`, ``),
		strings.ReplaceAll(shortcut.Executable, `"`, ``),
		strings.ReplaceAll(shortcut.LaunchOptions, `"`, `""`),
		strings.ReplaceAll(iconPath, `"`, ``),
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
