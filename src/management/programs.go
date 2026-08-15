package management

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/programs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Install programs with given options
func InstallPrograms(options *programs.Options) error {

	for _, id := range options.Programs {

		program, err := programs.GetProgramByID(id)
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
		if slices.Contains(program.Flags, "--system") {
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

		cli.Printf(cli.ColorNotice, "Creating shortcut for %s...\n", program.Name)

		// Add desktop flag or tag to control automatic shortcut creation
		// Based on formats that requires desktop shortcut creation
		// These packages do not create shortcuts by default
		if program.Package.Alias() == "" {
			program.Tags = append(program.Tags, "Desktop")
		} else if slices.Contains(program.Flags, "--web") {
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
		err = SetShortcut(shortcut, false)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "%s shortcut created!\n", program.Name)

	}

	return nil
}

// Remove programs with given options
func RemovePrograms(options *programs.Options) error {

	for _, id := range options.Programs {

		program, err := programs.GetProgramByID(id)
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
		shortcut := FindShortcut(program.Name, executable)

		if shortcut.ID != "" {
			cli.Printf(cli.ColorNotice, "Removing shortcut for %s...\n", program.Name)
			err = RemoveShortcut(shortcut)
			if err != nil {
				return err
			}

			cli.Printf(cli.ColorSuccess, "%s shortcut removed!\n", program.Name)
		}

		// Determine if can remove package
		canRemovePackage := true
		if slices.Contains(program.Flags, "--system") {
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
