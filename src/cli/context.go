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

	// Fallback to executable path when empty
	if c.WorkingDirectory == "" {
		c.WorkingDirectory = filepath.Dir(c.Executable)
	}

	// Make sure data in unquoted
	c.WorkingDirectory = Unquote(c.WorkingDirectory)
	c.Executable = Unquote(c.Executable)

	// Create script with arguments
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

	// Print debug data
	Debug("Working directory: %s\n", c.WorkingDirectory)
	Debug("Executable: %s\n", c.Executable)
	Debug("Arguments: %s\n", strings.Join(c.Arguments, " "))
	Debug("Environment: %s\n", strings.Join(c.Environment, " "))

	// Run the script
	if script != "" {
		command := Command(script)
		command.Dir = c.WorkingDirectory
		command.Env = append(os.Environ(), c.Environment...)
		return Start(command)
	}

	return nil
}
