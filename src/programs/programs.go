package programs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Program struct
type Program struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Category    string            `json:"category"`
	Tags        []string          `json:"tags"`
	Folders     []string          `json:"folders"`
	Website     string            `json:"website"`
	IconURL     string            `json:"iconUrl"`
	LogoURL     string            `json:"logoUrl"`
	CoverURL    string            `json:"coverUrl"`
	BannerURL   string            `json:"bannerUrl"`
	HeroURL     string            `json:"heroUrl"`
	Package     packaging.Package `json:"-"`
	OnInstall   func() error      `json:"-"`
	OnRemove    func() error      `json:"-"`
}

// Retrieve list of available programs to install
func GetPrograms() ([]*Program, error) {

	var programs []*Program
	var available []*Program

	// Retrieve all possible programs
	programs = append(programs, Azahar())
	programs = append(programs, BraveBrowser())
	programs = append(programs, Bottles())
	programs = append(programs, Cemu())
	programs = append(programs, ChiakiNG())
	programs = append(programs, Citron())
	programs = append(programs, Dolphin())
	programs = append(programs, DuckStation())
	programs = append(programs, Eden())
	programs = append(programs, ESDE())
	programs = append(programs, Firefox())
	programs = append(programs, Flycast())
	programs = append(programs, GeForceNow())
	programs = append(programs, GoogleChrome())
	programs = append(programs, HeroicGamesLauncher())
	programs = append(programs, Lutris())
	programs = append(programs, MelonDS())
	programs = append(programs, MGBA())
	programs = append(programs, MicrosoftEdge())
	programs = append(programs, MoonlightGameStreaming())
	programs = append(programs, NiceDeck())
	programs = append(programs, PCSX2())
	programs = append(programs, PPSSPP())
	programs = append(programs, ProtonPlus())
	programs = append(programs, Redream())
	programs = append(programs, RPCS3())
	programs = append(programs, Ryujinx())
	programs = append(programs, ShadPS4())
	programs = append(programs, Simple64())
	programs = append(programs, Vita3K())
	programs = append(programs, XboxCloudGaming())
	programs = append(programs, Xemu())
	programs = append(programs, Xenia())

	// Filter to return only available programs
	for _, program := range programs {
		if program.Package.Available() {
			available = append(available, program)
		}
	}

	return available, nil
}

// Retrieve program with given ID
func GetProgramByID(id string) (*Program, error) {

	programs, err := GetPrograms()
	notFound := &Program{
		Package: &packaging.Missing{},
	}

	if err != nil {
		return notFound, err
	}

	for _, program := range programs {
		if id == program.ID {
			return program, nil
		}
	}

	return notFound, nil
}

// Install program with given ID
func Install(options *Options) error {

	for _, id := range options.Programs {

		program, err := GetProgramByID(id)
		if err != nil {
			return err
		}

		// Program not found
		if program.ID == "" {
			return fmt.Errorf("program not found: %s", id)
		}

		// Program not available
		if !program.Package.Available() {
			return fmt.Errorf("program is not available to install: %s", id)
		}

		// Print step message
		cli.Printf(cli.ColorNotice, "Installing %s...\n", program.Name)

		// Make sure required folders exist
		if len(program.Folders) > 0 {
			for _, folder := range program.Folders {
				err := os.MkdirAll(fs.ExpandPath(folder), 0755)
				if err != nil {
					return err
				}
			}
		}

		// Run program installation
		err = program.Package.Install()
		if err != nil {
			return err
		}

		// Perform additional steps after installation
		if program.OnInstall != nil {
			err = program.OnInstall()
			if err != nil {
				return err
			}
		}

		// Fill basic shortcut information
		executable := program.Package.Executable()
		startDirectory := filepath.Dir(executable)
		shortcut := &shortcuts.Shortcut{
			Platform:       "PC",
			Program:        program.ID,
			Layer:          "native",
			Type:           "program",
			Name:           program.Name,
			Description:    program.Description,
			StartDirectory: startDirectory,
			Executable:     executable,
			LaunchOptions:  "",
			ShortcutPath:   "",
			RelativePath:   "",
			IconURL:        program.IconURL,
			LogoURL:        program.LogoURL,
			CoverURL:       program.CoverURL,
			BannerURL:      program.BannerURL,
			HeroURL:        program.HeroURL,
			Tags:           program.Tags,
		}

		// Fill additional shortcut information from package
		err = program.Package.OnShortcut(shortcut)
		if err != nil {
			return err
		}

		// Add to shortcuts list
		err = library.Shortcuts.Set(shortcut, false)
		if err != nil {
			return err
		}

		// Print success message
		cli.Printf(cli.ColorSuccess, "%s installed!\n", program.Name)

	}

	return nil
}

// Remove program with given options
func Remove(options *Options) error {

	for _, id := range options.Programs {

		program, err := GetProgramByID(id)
		if err != nil {
			return err
		}

		// Program not found
		if program.ID == "" {
			return fmt.Errorf("program not found: %s", id)
		}

		// Program not available
		if !program.Package.Available() {
			return fmt.Errorf("program is not available to remove: %s", id)
		}

		// Print step message
		cli.Printf(cli.ColorNotice, "Removing %s...\n", program.Name)

		// Run program removal
		err = program.Package.Remove()
		if err != nil {
			return err
		}

		// Perform additional steps after remove
		if program.OnRemove != nil {
			err = program.OnRemove()
			if err != nil {
				return err
			}
		}

		// Remove from shortcuts list
		executable := program.Package.Executable()
		shortcut := library.Shortcuts.Find(executable, program.Name)

		if shortcut.ID != "" {
			err = library.Shortcuts.Remove(shortcut)
			if err != nil {
				return err
			}
		}

		// Print success message
		cli.Printf(cli.ColorSuccess, "%s removed!\n", program.Name)

	}

	return nil
}
