package packaging

import "github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"

// Package interface
type Package interface {
	Available() bool
	Runtime() string
	Install(shortcut *shortcuts.Shortcut) error
	Installed() (bool, error)
	Executable() string
	Run(args []string) error
}

// Retrieve first available package
func Available(args ...Package) Package {

	for _, item := range args {
		if item.Available() {
			return item
		}
	}

	return &Missing{}
}
