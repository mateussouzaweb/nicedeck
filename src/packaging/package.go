package packaging

import "github.com/mateussouzaweb/nicedeck/src/shortcuts"

// Package interface
type Package interface {
	Runtime() string
	Available() bool
	Install() error
	Remove() error
	Installed() (bool, error)
	Executable() string
	Alias() string
	Run(arguments []string) error
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
