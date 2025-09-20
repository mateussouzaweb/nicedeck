package cli

import (
	"fmt"
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
