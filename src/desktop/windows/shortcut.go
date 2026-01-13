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

	name := fs.NormalizeFilename(shortcut.Name)
	categories := shortcut.Tags

	// Specify category resolution to place shortcut accordingly
	accepts := []string{
		"Gaming",
		"Utilities",
	}
	replaces := map[string]string{
		"ROM":       "Gaming",
		"Emulator":  "Gaming",
		"Streaming": "Utilities",
	}

	// Check for accepted categories in shortcut tags
	// When found, place shortcut inside respective folder
	for _, category := range categories {
		if value, ok := replaces[category]; ok {
			category = value
		}
		if slices.Contains(accepts, category) {
			return fs.ExpandPath(fmt.Sprintf(
				"$START_MENU/%s/%s.lnk", category, name,
			))
		}
	}

	// Default location
	return fs.ExpandPath(fmt.Sprintf(
		"$START_MENU/%s.lnk", name,
	))
}

// Create a desktop entry from shortcut data
func CreateShortcut(shortcut *shortcuts.Shortcut, destination string) error {

	// Prepare execution context
	context := shortcuts.PrepareContext(shortcut)

	// Ensure shortcut directory exists
	err := os.MkdirAll(filepath.Dir(destination), 0755)
	if err != nil {
		return err
	}

	// Detect icon path or defaults to executable icon
	// Windows accepts only .ico format for shortcuts
	iconPath := context.Executable
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
		strings.ReplaceAll(context.WorkingDirectory, `"`, ``),
		strings.ReplaceAll(context.Executable, `"`, ``),
		strings.ReplaceAll(strings.Join(context.Arguments, " "), `"`, `""`),
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
