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

// Install program
func (m *Missing) Install(shortcut *shortcuts.Shortcut) error {
	return fmt.Errorf("cannot be installed, package is missing")
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
