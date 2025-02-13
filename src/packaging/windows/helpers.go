package windows

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Run a program with given set of arguments
func RunProcess(executable string, args []string) error {
	if len(args) > 0 {
		return cli.Start(fmt.Sprintf(``+
			`$Arguments = '%s';`+
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait -ArgumentList $Arguments`,
			strings.Join(args, " "),
			filepath.Dir(executable),
			executable,
		))
	}

	return cli.Start(fmt.Sprintf(
		`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait`,
		filepath.Dir(executable),
		executable,
	))
}

// Create a system shortcut
func CreateSystemShortcut(shortcut *shortcuts.Shortcut) error {
	return cli.Run(fmt.Sprintf(``+
		`$WshShell = New-Object -COMObject WScript.Shell;`+
		`$Shortcut = $WshShell.CreateShortcut("%s");`+
		`$Shortcut.WorkingDirectory = "%s";`+
		`$Shortcut.TargetPath = "%s";`+
		`$Shortcut.Arguments = "%s";`+
		`$Shortcut.Save()`,
		shortcut.ShortcutPath,
		shortcut.StartDir,
		shortcut.Exe,
		strconv.Quote(shortcut.LaunchOptions),
	))
}
