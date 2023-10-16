package roms

import (
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Parse and process ROMs for all emulators
func ProcessROMs(includePath string) error {

	// Detect ROMs with parser in all folders
	parsed, err := ParseROMs()
	if err != nil {
		return err
	}

	// Fill list of ROMs to process, based on given include path
	// Also fill detected list including all system ROMs
	process := []*ROM{}
	detected := []string{}

	for _, rom := range parsed {

		// Add to the list of detected ROMs
		// We use the ROM relative path because this info can be found in the shortcut
		detected = append(detected, rom.RelativePath)

		// Add to the list of ROMs to process if match include path condition
		if includePath == "" || strings.Contains(rom.Path, includePath) {
			process = append(process, rom)
		}

	}

	// Print initial process information
	total := len(process)
	cli.Printf(cli.ColorNotice, "%d ROMs detected to process\n", total)
	cli.Printf(cli.ColorNotice, "This could take some time, please be patient...\n")

	// Process each ROM to add or update
	for index, rom := range process {

		cli.Printf(cli.ColorNotice, "Processing ROM [%d/%d]: %s\n", index+1, total, rom.RelativePath)

		// Scrape additional ROM information
		scrape, err := ScrapeROM(rom)
		if err != nil {
			return err
		}

		// Skip if scrape not found anything...
		if scrape.Name == "" {
			cli.Printf(cli.ColorWarn, "Could not detect ROM information. Skipping...\n")
			continue
		}

		// Add to Steam
		err = steam.AddToShortcuts(&shortcuts.Shortcut{
			AppName:       scrape.Name,
			Exe:           rom.LaunchCommand,
			StartDir:      "/var/lib/flatpak/exports/bin/", // Same as main flatpak
			ShortcutPath:  "",
			LaunchOptions: "",
			IconURL:       scrape.IconURL,
			LogoURL:       scrape.LogoURL,
			CoverURL:      scrape.CoverURL,
			BannerURL:     scrape.BannerURL,
			HeroURL:       scrape.HeroURL,
			Tags:          []string{rom.Console, "ROM"},
		})

		if err != nil {
			return err
		}

	}

	cli.Printf(cli.ColorNotice, "Scrapping finished.\n")
	cli.Printf(cli.ColorNotice, "Removing not detect ROMs...\n")

	// Remove ROM shortcuts that was not detected in the current run
	for _, shortcut := range steam.GetShortcuts() {

		// Check if shortcut is managed ROM
		if !slices.Contains(shortcut.Tags, "ROM") {
			continue
		}

		// Check if ROM is on the list of detected ROMs
		found := false
		for _, detectedROM := range detected {
			if strings.Contains(shortcut.Exe, detectedROM) {
				found = true
				break
			}
		}

		// Remove when not found
		if !found {
			cli.Printf(cli.ColorWarn, "Removing not detected ROM: %s\n", shortcut.AppName)
			err = steam.RemoveFromShortcuts(shortcut)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
