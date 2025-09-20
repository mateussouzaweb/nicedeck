package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Process context struct
type Context struct {
	WorkingDirectory string
	Executable       string
	Arguments        []string
	Environment      []string
}

// Run a program with based on their context
func (c *Context) Run() error {

	if c.WorkingDirectory == "" {
		c.WorkingDirectory = filepath.Dir(c.Executable)
	}

	arguments := strings.Join(c.Arguments, " ")
	script := ""

	if IsLinux() {
		script = fmt.Sprintf(
			`cd "%s" && exec "%s" %s`,
			c.WorkingDirectory,
			c.Executable,
			arguments,
		)
	} else if IsMacOS() {
		script = fmt.Sprintf(
			`cd "%s" && open -n "%s" --args %s`,
			c.WorkingDirectory,
			c.Executable,
			arguments,
		)
	} else if IsWindows() && len(c.Arguments) > 0 {
		script = fmt.Sprintf(``+
			`$Arguments = '%s';`+
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait -ArgumentList $Arguments`,
			arguments,
			c.WorkingDirectory,
			c.Executable,
		)
	} else if IsWindows() {
		script = fmt.Sprintf(
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait`,
			c.WorkingDirectory,
			c.Executable,
		)
	}

	if script != "" {
		command := Command(script)
		command.Dir = c.WorkingDirectory
		command.Env = append(os.Environ(), c.Environment...)
		return Start(command)
	}

	return nil
}
