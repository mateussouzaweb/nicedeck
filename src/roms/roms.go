package roms

import (
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Parse and process ROMs for all emulators
func ProcessROMs() error {

	// Detect ROMs with parser
	parsed, err := ParseROMs()
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorNotice, "%d ROMs detected\n", len(parsed))
	cli.Printf(cli.ColorNotice, "This process could take some time, please be patient...\n")

	detected := []string{}

	// Process each ROM to add or update
	for _, rom := range parsed {

		cli.Printf(cli.ColorNotice, "Processing ROM: %s\n", rom.RelativePath)

		// Add to the list of detected ROMs
		detected = append(detected, rom.RelativePath)

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
		for _, detectRom := range detected {
			// We use the ROM relative path because this info can be found in the shortcut.Exe
			if strings.Contains(shortcut.Exe, detectRom) {
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
