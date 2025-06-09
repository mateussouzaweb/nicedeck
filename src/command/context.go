package command

import (
	"strconv"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Context struct
type Context struct {
	Version string
	Args    []string
	Done    chan bool
}

// Retrieve argument with given keys or in given index position
func (c *Context) Arg(keys string, defaultValue string) string {
	value := cli.Arg(c.Args, keys, defaultValue)

	if strings.ContainsAny(value, "\"'`") {
		value, _ = strconv.Unquote(value)
	}

	return value
}

// Retrieve argument with given keys or in given index position
func (c *Context) Flag(keys string, defaultValue bool) bool {
	return cli.Flag(c.Args, keys, defaultValue)
}

// Retrieve multiple values for arguments with given key
func (c *Context) Multiple(key string, separator string) []string {
	return cli.Multiple(c.Args, key, separator)
}

// Wait for application to shutdown
func (c *Context) Wait() {
	<-c.Done
}

// Ask for application to shutdown
func (c *Context) Shutdown() {
	c.Done <- true
}
