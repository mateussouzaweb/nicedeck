package packaging

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Missing struct
type Missing struct{}

// Return if package is available
func (m *Missing) Available() bool {
	return false
}

// Return package runtime
func (m *Missing) Runtime() string {
	return "none"
}

// Install program
func (m *Missing) Install() error {
	return fmt.Errorf("cannot perform package installations")
}

// Installed verification
func (m *Missing) Installed() (bool, error) {
	return false, nil
}

// Return executable file path
func (m *Missing) Executable() string {
	return ""
}

// Run installed program
func (m *Missing) Run(args []string) error {
	return nil
}

// Fill shortcut additional details
func (m *Missing) OnShortcut(shortcut *shortcuts.Shortcut) error {
	return nil
}
