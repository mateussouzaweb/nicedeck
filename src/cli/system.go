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
		command := Command(script)
		return Run(command)
	} else if IsMacOS() {
		script := fmt.Sprintf(`open %s`, file)
		command := Command(script)
		return Run(command)
	} else if IsWindows() {
		script := fmt.Sprintf(`Start-Process "%s"`, file)
		command := Command(script)
		return Run(command)
	}

	return nil
}

// Run a program with given set of arguments
func RunProcess(executable string, args []string) error {

	script := ""
	arguments := strings.Join(args, " ")
	workingDirectory := filepath.Dir(executable)

	if IsLinux() {
		script = fmt.Sprintf(
			`cd "%s" && exec "%s" %s`,
			workingDirectory,
			executable,
			arguments,
		)
	} else if IsMacOS() {
		script = fmt.Sprintf(
			`cd "%s" && open -n "%s" --args %s`,
			workingDirectory,
			executable,
			arguments,
		)
	} else if IsWindows() && len(args) > 0 {
		script = fmt.Sprintf(``+
			`$Arguments = '%s';`+
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait -ArgumentList $Arguments`,
			arguments,
			workingDirectory,
			executable,
		)
	} else if IsWindows() {
		script = fmt.Sprintf(
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait`,
			workingDirectory,
			executable,
		)
	}

	if script != "" {
		command := Command(script)
		command.Dir = workingDirectory
		return Start(command)
	}

	return nil
}
