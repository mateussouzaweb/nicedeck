package cli

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Open file or URL in configured system app
func Open(file string) error {

	if IsLinux() {
		script := fmt.Sprintf(`xdg-open %s`, file)
		return Run(script)
	} else if IsMacOS() {
		script := fmt.Sprintf(`open %s`, file)
		return Run(script)
	} else if IsWindows() {
		script := fmt.Sprintf(`Start-Process "%s"`, file)
		return Run(script)
	}

	return nil
}

// Run a program with given set of arguments
func RunProcess(executable string, args []string) error {

	workingDirectory := filepath.Dir(executable)
	arguments := strings.Join(args, " ")

	if IsLinux() {
		return Start(fmt.Sprintf(
			`cd "%s" && exec "%s" %s`,
			workingDirectory,
			executable,
			arguments,
		))
	} else if IsMacOS() {
		return Start(fmt.Sprintf(
			`cd "%s" && open -n "%s" --args %s`,
			workingDirectory,
			executable,
			arguments,
		))
	} else if IsWindows() && len(args) > 0 {
		return Start(fmt.Sprintf(``+
			`$Arguments = '%s';`+
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait -ArgumentList $Arguments`,
			arguments,
			workingDirectory,
			executable,
		))
	} else if IsWindows() {
		return Start(fmt.Sprintf(
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait`,
			workingDirectory,
			executable,
		))
	}

	return nil
}
