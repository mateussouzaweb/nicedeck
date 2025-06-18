package build

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Context represents a build context with environment variables and arguments
type Context struct {
	Env []string
}

// Set an environment variable in the context
func (c *Context) Set(key string, value string) *Context {
	c.Env = append(c.Env, key+"="+value)
	return c
}

// Init context process
func (c *Context) Init() error {
	for _, item := range c.Env {
		data := strings.Split(item, "=")
		err := cli.SetEnv(data[0], data[1], true)
		return err
	}

	return nil
}

// Complete context process
func (c *Context) Done() error {
	for _, item := range c.Env {
		data := strings.Split(item, "=")
		err := cli.UnsetEnv(data[0])
		return err
	}

	return nil
}

// Executes the context process for a callback function
func (c *Context) Run(callback func() error) error {

	err := c.Init()
	if err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	defer func() {
		doneErr := c.Done()
		if doneErr != nil {
			doneErr = fmt.Errorf("context error: %w", c.Done())
			errors.Join(err, doneErr)
		}
	}()

	if callback != nil {
		return callback()
	}

	return nil
}

// Create a new context with environment variables
func Env(env ...string) *Context {
	return &Context{
		Env: env,
	}
}
