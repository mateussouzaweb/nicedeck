package build

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Represents a command to be executed in the build process
type Command struct {
	Context  *Context
	Callback func() error
}

// Run command callback
func (c *Command) Run() error {
	if c.Context == nil {
		c.Context = Env()
	}

	err := c.Context.Run(c.Callback)
	if err != nil {
		return fmt.Errorf("command error: %w", err)
	}

	return nil
}

// Cmd creates a new build command with the provided script.
func Cmd(script string, args ...any) *Command {
	cmd := cli.Command(fmt.Sprintf(script, args...))
	return &Command{
		Callback: func() error {
			return cmd.Run()
		},
	}
}
