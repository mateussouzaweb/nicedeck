package emulationstation

import (
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Package struct
type Package struct {
	AppID       string `json:"appId"`
	Format      string `json:"format"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Launcher    string `json:"launcher"`
}

// Install program with archive
func (p *Package) Install(shortcut *shortcuts.Shortcut) error {

	// Get latest available version
	latest, err := GetLatestRelease()
	if err != nil {
		return err
	}

	// Download application
	executable := p.Executable()
	err = fs.DownloadFile(latest, executable, true)
	if err != nil {
		return err
	}

	// Make sure is executable
	err = os.Chmod(executable, 0775)
	if err != nil {
		return err
	}

	// Write settings
	err = WriteSettings()
	if err != nil {
		return err
	}

	// Fill shortcut information
	desktopShortcut := os.ExpandEnv("$HOME/.local/share/applications/emulationstation-de.desktop")
	shortcut.Exe = executable
	shortcut.StartDir = filepath.Dir(executable)
	shortcut.ShortcutPath = desktopShortcut
	shortcut.LaunchOptions = ""

	return nil
}

// Installed verification
func (p *Package) Installed() (bool, error) {

	executable := p.Executable()
	exist, err := fs.FileExist(executable)
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (p *Package) Executable() string {
	return os.ExpandEnv("$APPLICATIONS/EmulationStation-DE.AppImage")
}
