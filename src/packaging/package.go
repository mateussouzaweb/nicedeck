package packaging

import "github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"

// Package interface
type Package interface {
	Available() bool
	Runtime() string
	Install() error
	Installed() (bool, error)
	Executable() string
	Run(args []string) error
	OnShortcut(shortcut *shortcuts.Shortcut) error
}

// Find best package match based on availability with installed prioritization
func Best(args ...Package) Package {

	var available []Package

	for _, item := range args {
		if item.Available() {
			available = append(available, item)
		}
	}

	if len(available) == 0 {
		return &Missing{}
	}

	for _, item := range available {
		if installed, _ := item.Installed(); installed {
			return item
		}
	}

	return available[0]
}

// Retrieve first installed package
func Installed(args ...Package) Package {

	for _, item := range args {
		if !item.Available() {
			continue
		}
		if installed, _ := item.Installed(); installed {
			return item
		}
	}

	return &Missing{}
}
