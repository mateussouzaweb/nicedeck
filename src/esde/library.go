package esde

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/esde/settings"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Shortcut alias
type Shortcut = shortcuts.Shortcut

// Library struct
type Library struct {
	BasePath string `json:"basePath"`
}

// String representation of the library
func (l *Library) String() string {
	return "ES-DE"
}

// Load library
func (l *Library) Load() error {

	// Windows portable version uses application path
	if cli.IsWindows() {
		l.BasePath = fs.ExpandPath("$APPLICATIONS/ES-DE/ES-DE")
		return nil
	}

	// Default path for other OS (Linux and MacOS)
	l.BasePath = fs.ExpandPath("$HOME/ES-DE")
	return nil
}

// Save library
func (l *Library) Save() error {

	installed, err := GetPackage().Installed()
	if err != nil {
		return err
	} else if !installed {
		return nil
	}

	// Write settings
	err = settings.WriteSettings(l.BasePath)
	if err != nil {
		return err
	}

	return nil
}

// Export shortcuts to internal format
func (l *Library) Export() []*Shortcut {
	results := make([]*Shortcut, 0)
	return results
}

// Add shortcut to the library
func (l *Library) Add(shortcut *Shortcut) error {
	return nil
}

// Update shortcut on library
func (l *Library) Update(shortcut *Shortcut, overwriteAssets bool) error {
	return nil
}

// Remove shortcut from the library
func (l *Library) Remove(shortcut *Shortcut) error {
	return nil
}
