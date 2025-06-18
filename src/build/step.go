package build

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Step represents a single step in the build process
type Step struct {
	ID      string
	Name    string
	Context *Context
	Command *Command
}

// Run step process
func (s *Step) Run() error {
	if s.Context == nil {
		s.Context = Env()
	}

	cli.Printf(cli.ColorNotice, "- Running step: %s (%s)\n", s.ID, s.Name)
	err := s.Context.Run(s.Command.Run)
	if err != nil {
		return fmt.Errorf("step error: %w", err)
	}

	return nil
}
