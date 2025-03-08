package packaging

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Missing struct
type Missing struct{}

// Return package runtime
func (m *Missing) Runtime() string {
	return "none"
}

// Return if package is available
func (m *Missing) Available() bool {
	return false
}

// Install package
func (m *Missing) Install() error {
	return fmt.Errorf("cannot perform package installations")
}

// Remove package
func (m *Missing) Remove() error {
	return nil
}

// Installed verification
func (m *Missing) Installed() (bool, error) {
	return false, nil
}

// Return executable file path
func (m *Missing) Executable() string {
	return ""
}

// Return executable alias file path
func (m *Missing) Alias() string {
	return ""
}

// Run installed package
func (m *Missing) Run(args []string) error {
	return nil
}

// Fill shortcut additional details
func (m *Missing) OnShortcut(shortcut *shortcuts.Shortcut) error {
	return nil
}
