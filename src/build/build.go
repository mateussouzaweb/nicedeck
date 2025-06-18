package build

import "fmt"

// Executable interface
type Executable interface {
	Run() error
}

// Build represents a collection of steps to be executed
type Build struct {
	Context *Context
	Steps   []Executable
}

// Add a new step to the build
func (b *Build) Add(steps ...Executable) *Build {
	b.Steps = append(b.Steps, steps...)
	return b
}

// Run all steps in the build process
func (b *Build) Run() error {
	if b.Context == nil {
		b.Context = Env()
	}

	for _, step := range b.Steps {
		err := b.Context.Run(step.Run)
		if err != nil {
			return fmt.Errorf("build error: %w", err)
		}
	}

	return nil
}

// Creates a new build instance with context
func New(context *Context) *Build {
	return &Build{
		Context: context,
		Steps:   make([]Executable, 0),
	}
}
