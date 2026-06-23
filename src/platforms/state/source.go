package state

import "github.com/mateussouzaweb/nicedeck/src/cli"

// Source struct
type Source struct {
	Linux   []string `json:"linux"`
	MacOS   []string `json:"macos"`
	Windows []string `json:"windows"`
}

// Return paths based on current system
func (s *Source) Paths() []string {
	if cli.IsLinux() {
		return s.Linux
	} else if cli.IsMacOS() {
		return s.MacOS
	} else if cli.IsWindows() {
		return s.Windows
	}

	return []string{}
}
