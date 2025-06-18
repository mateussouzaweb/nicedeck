package build

import (
	"fmt"
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
