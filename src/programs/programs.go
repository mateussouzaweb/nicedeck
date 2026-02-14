package programs

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

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
	Flags       []string          `json:"flags"`
	Website     string            `json:"website"`
	IconURL     string            `json:"iconUrl"`
	LogoURL     string            `json:"logoUrl"`
	CoverURL    string            `json:"coverUrl"`
	BannerURL   string            `json:"bannerUrl"`
	HeroURL     string            `json:"heroUrl"`
	Package     packaging.Package `json:"-"`
}

// Retrieve list of available programs to install
func GetPrograms() ([]*Program, error) {

	var programs []*Program
	var available []*Program

	// Retrieve all possible programs
	programs = append(programs, AmazonGames())
	programs = append(programs, Azahar())
	programs = append(programs, BattleNet())
	programs = append(programs, BraveBrowser())
	programs = append(programs, Bottles())
	programs = append(programs, Cemu())
	programs = append(programs, ChiakiNG())
	programs = append(programs, Citron())
	programs = append(programs, Discord())
	programs = append(programs, Dolphin())
	programs = append(programs, DuckStation())
	programs = append(programs, EAApp())
	programs = append(programs, Eden())
	programs = append(programs, EpicGames())
	programs = append(programs, ESDE())
	programs = append(programs, Firefox())
	programs = append(programs, Flycast())
	programs = append(programs, GeForceNow())
	programs = append(programs, GOGGalaxy())
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
	programs = append(programs, Steam())
	programs = append(programs, UbisoftConnect())
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

	// Flag installed programs and runtime
	for _, program := range available {
		runtime := fmt.Sprintf("--%s", program.Package.Runtime())
		program.Flags = append(program.Flags, runtime)

		installed, err := program.Package.Installed()
		if err != nil {
			return available, err
		}
		if installed {
			program.Flags = append(program.Flags, "--installed")
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

		// Make sure required folders exist
		if len(program.Folders) > 0 {
			for _, folder := range program.Folders {
				err := os.MkdirAll(fs.ExpandPath(folder), 0755)
				if err != nil {
					return err
				}
			}
		}

		// Determine if can install package
		canInstallPackage := true
		if slices.Contains(program.Flags, "--browser-shortcut") {
			if slices.Contains(program.Flags, "--installed") {
				cli.Printf(cli.ColorWarn, "Skipping browser installation because it already is installed.\n")
				canInstallPackage = false
			}
		} else if slices.Contains(program.Flags, "--system") {
			cli.Printf(cli.ColorWarn, "Skipping %s installation because it already is provided from system.\n", program.Name)
			canInstallPackage = false
		}

		// Run program installation when possible
		if canInstallPackage {
			cli.Printf(cli.ColorNotice, "Installing %s...\n", program.Name)
			err = program.Package.Install()
			if err != nil {
				return err
			}

			cli.Printf(cli.ColorSuccess, "%s installed!\n", program.Name)
		}

		// Add desktop flag or tag to control automatic shortcut creation
		// Based on formats that requires desktop shortcut creation
		// These packages do not create shortcuts by default
		if program.Package.Alias() == "" {
			program.Tags = append(program.Tags, "Desktop")
		} else if slices.Contains(program.Flags, "--browser-shortcut") {
			program.Tags = append(program.Tags, "Desktop")
		}

		// Retrieve shortcut information
		executable := program.Package.Executable()
		arguments := program.Package.Args()
		startDirectory := filepath.Dir(executable)
		shortcutID := shortcuts.GenerateID(program.Name, executable)

		// Create final shortcut specs
		shortcut := &shortcuts.Shortcut{
			ID:             shortcutID,
			Program:        program.ID,
			Name:           program.Name,
			Description:    program.Description,
			StartDirectory: cli.Quote(startDirectory),
			Executable:     cli.Quote(executable),
			LaunchOptions:  strings.Join(arguments, " "),
			RelativePath:   "",
			IconPath:       program.IconURL,
			LogoPath:       program.LogoURL,
			CoverPath:      program.CoverURL,
			BannerPath:     program.BannerURL,
			HeroPath:       program.HeroURL,
			Tags:           program.Tags,
		}

		// Add to shortcuts list
		err = library.Shortcuts.Set(shortcut, false)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "%s shortcut created!\n", program.Name)

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

		// Remove from shortcuts list
		executable := program.Package.Executable()
		shortcut := library.Shortcuts.Find(program.Name, executable)

		if shortcut.ID != "" {
			err = library.Shortcuts.Remove(shortcut)
			if err != nil {
				return err
			}

			cli.Printf(cli.ColorSuccess, "%s shortcut removed!\n", program.Name)
		}

		// Determine if can remove package
		canRemovePackage := true
		if slices.Contains(program.Flags, "--browser-shortcut") {
			cli.Printf(cli.ColorWarn, "Note: Only the %s shortcut was be removed because it is a browser shortcut.\n", program.Name)
			canRemovePackage = false
		} else if slices.Contains(program.Flags, "--system") {
			cli.Printf(cli.ColorWarn, "Note: %s is provided by the system and cannot removed.\n", program.Name)
			canRemovePackage = false
		} else if slices.Contains(program.Flags, "--nicedeck") {
			cli.Printf(cli.ColorWarn, "Warning: %s cannot be fully removed because it is running.\n", program.Name)
			cli.Printf(cli.ColorWarn, "Please close the program and remove it manually.\n")
			canRemovePackage = false
		}

		// Run program removal when possible
		if canRemovePackage {
			cli.Printf(cli.ColorNotice, "Removing %s...\n", program.Name)
			err = program.Package.Remove()
			if err != nil {
				return err
			}

			cli.Printf(cli.ColorSuccess, "%s removed!\n", program.Name)
		}

	}

	return nil
}
