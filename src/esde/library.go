package esde

import (
	"github.com/mateussouzaweb/nicedeck/src/esde/settings"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Library struct
type Library struct {
	BasePath string `json:"basePath"`
}

// Load library
func (l *Library) Load() error {
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
