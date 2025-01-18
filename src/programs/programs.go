package programs

import (
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Package interface
type Package interface {
	Available() bool
	Install(shortcut *shortcuts.Shortcut) error
	Installed() (bool, error)
	Executable() string
	Run(args []string) error
}

// Program struct
type Program struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Category        string   `json:"category"`
	Tags            []string `json:"tags"`
	RequiredFolders []string `json:"requiredFolders"`
	IconURL         string   `json:"iconUrl"`
	LogoURL         string   `json:"logoUrl"`
	CoverURL        string   `json:"coverUrl"`
	BannerURL       string   `json:"bannerUrl"`
	HeroURL         string   `json:"heroUrl"`
	Package         Package  `json:"-"`
}

// Retrieve list of available programs to install
func GetPrograms() ([]*Program, error) {

	var programs []*Program

	programs = append(programs, BraveBrowser())
	programs = append(programs, Bottles())
	programs = append(programs, Cemu())
	programs = append(programs, Chiaki())
	programs = append(programs, Citra())
	programs = append(programs, Dolphin())
	programs = append(programs, DuckStation())
	programs = append(programs, EmulationStationDE())
	programs = append(programs, Firefox())
	programs = append(programs, Flycast())
	programs = append(programs, GeForceNow())
	programs = append(programs, GoogleChrome())
	programs = append(programs, HeroicGamesLauncher())
	programs = append(programs, Lime3DS())
	programs = append(programs, Lutris())
	programs = append(programs, MelonDS())
	programs = append(programs, MGBA())
	programs = append(programs, MicrosoftEdge())
	programs = append(programs, MoonlightGameStreaming())
	programs = append(programs, PCSX2())
	programs = append(programs, PPSSPP())
	programs = append(programs, ProtonPlus())
	programs = append(programs, RPCS3())
	programs = append(programs, Ryujinx())
	programs = append(programs, Simple64())
	programs = append(programs, XboxCloudGaming())
	programs = append(programs, Xemu())
	programs = append(programs, Yuzu())

	return programs, nil
}

// Retrieve program with given ID
func GetProgramByID(id string) (*Program, error) {

	programs, err := GetPrograms()
	if err != nil {
		return &Program{}, err
	}

	for _, program := range programs {
		if id == program.ID {
			return program, nil
		}
	}

	return &Program{}, nil
}

// Install program with given ID
func Install(id string) error {

	program, err := GetProgramByID(id)
	if err != nil {
		return err
	}

	// Program not found
	if program.ID == "" {
		return fmt.Errorf("program not found: %s", id)
	}

	// Program not available
	if program.Package.Available() {
		return fmt.Errorf("program is not available to install: %s", id)
	}

	// Print step message
	cli.Printf(cli.ColorNotice, "Installing %s...\n", program.Name)

	// Make sure required folders exist
	if len(program.RequiredFolders) > 0 {
		for _, folder := range program.RequiredFolders {
			err := os.MkdirAll(os.ExpandEnv(folder), 0755)
			if err != nil {
				return err
			}
		}
	}

	// Fill basic shortcut information
	shortcut := &shortcuts.Shortcut{
		AppName:   program.Name,
		Tags:      program.Tags,
		IconURL:   program.IconURL,
		LogoURL:   program.LogoURL,
		CoverURL:  program.CoverURL,
		BannerURL: program.BannerURL,
		HeroURL:   program.HeroURL,
	}

	// Run program installation with shortcut
	err = program.Package.Install(shortcut)
	if err != nil {
		return err
	}

	// Add to shortcuts list
	err = library.AddToShortcuts(shortcut, false)
	if err != nil {
		return err
	}

	// Print success message
	cli.Printf(cli.ColorSuccess, "%s installed!\n", program.Name)
	return nil
}
