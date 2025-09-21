package platforms

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/platforms/console"
	"github.com/mateussouzaweb/nicedeck/src/platforms/native"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Scrape data from shortcut and return if was found
func ScrapeShortcut(shortcut *shortcuts.Shortcut) (bool, error) {

	// Scrape additional ROM information
	scrape, err := scraper.ScrapeFromName(shortcut.Name)
	if err != nil {
		return false, err
	}

	// Skip if scrape not found anything...
	if scrape.Name == "" {
		return false, nil
	}

	// Determine best name and images for the shortcut
	name := scrape.Name
	iconURL := ""
	logoURL := ""
	coverURL := ""
	bannerURL := ""
	heroURL := ""

	if len(scrape.IconURLs) > 0 {
		iconURL = scrape.IconURLs[0]
	}
	if len(scrape.LogoURLs) > 0 {
		logoURL = scrape.LogoURLs[0]
	}
	if len(scrape.CoverURLs) > 0 {
		coverURL = scrape.CoverURLs[0]
	}
	if len(scrape.BannerURLs) > 0 {
		bannerURL = scrape.BannerURLs[0]
	}
	if len(scrape.HeroURLs) > 0 {
		heroURL = scrape.HeroURLs[0]
	}

	shortcut.Name = name
	shortcut.IconURL = iconURL
	shortcut.LogoURL = logoURL
	shortcut.CoverURL = coverURL
	shortcut.BannerURL = bannerURL
	shortcut.HeroURL = heroURL

	return true, nil
}

// Parse and process shortcut with given path
func ProcessShortcut(path string, options *Options) (*shortcuts.Shortcut, error) {

	includeNative := true
	includeConsole := true
	shortcut := &shortcuts.Shortcut{}
	nameFormat := "${NAME}"

	// Make sure path is unquoted
	path = cli.Unquote(path)

	// Try to parse a native ROM
	if shortcut.ID == "" && includeNative {

		theOptions := native.ToOptions(options.Platforms, options.Preferences)
		rom, err := native.ParseROM(path, theOptions)
		if err != nil {
			return shortcut, err
		}

		// When valid, create a new shortcut from native ROM
		if rom.Executable != "" {
			shortcutID := shortcuts.GenerateID(rom.Name, rom.Executable)
			shortcut = &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        rom.Program,
				Name:           rom.Name,
				Description:    rom.Description,
				StartDirectory: cli.Quote(rom.StartDirectory),
				Executable:     cli.Quote(rom.Executable),
				LaunchOptions:  rom.LaunchOptions,
				ShortcutPath:   "",
				RelativePath:   rom.Executable,
				IconURL:        "",
				LogoURL:        "",
				CoverURL:       "",
				BannerURL:      "",
				HeroURL:        "",
				Tags:           []string{rom.Runtime},
			}
		}

	}

	// Try to parse a console ROM
	if shortcut.ID == "" && includeConsole {

		theOptions := console.ToOptions(options.Platforms, options.Preferences)
		rom, err := console.ParseROM(path, theOptions)
		if err != nil {
			return shortcut, err
		}

		// When valid, create a new shortcut from console ROM
		if rom.Executable != "" {
			nameFormat = fmt.Sprintf("${NAME} [%s]", rom.Platform)
			startDirectory := filepath.Dir(rom.Executable)
			shortcutID := shortcuts.GenerateID(rom.Name, rom.Executable)
			shortcut = &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        rom.Program,
				Name:           rom.Name,
				Description:    rom.Description,
				StartDirectory: cli.Quote(startDirectory),
				Executable:     cli.Quote(rom.Executable),
				LaunchOptions:  rom.LaunchOptions,
				ShortcutPath:   "",
				RelativePath:   rom.RelativePath,
				IconURL:        "",
				LogoURL:        "",
				CoverURL:       "",
				BannerURL:      "",
				HeroURL:        "",
				Tags:           []string{"Gaming", "ROM", rom.Platform},
			}
		}

	}

	// Scrape additional shortcut information
	if shortcut.ID != "" {

		scraped, err := ScrapeShortcut(shortcut)
		if err != nil {
			return shortcut, err
		} else if !scraped {
			cli.Printf(cli.ColorWarn, "Could not detect shortcut information. Using available data...\n")
		} else if scraped {
			shortcut.Name = strings.Replace(nameFormat, "${NAME}", shortcut.Name, 1)
		}

		return shortcut, nil
	}

	err := fmt.Errorf("could not determine the file shortcut")
	return shortcut, err
}

// Parse and process shortcuts for given platforms
func ProcessShortcuts(options *Options) error {

	theOptions := console.ToOptions(options.Platforms, options.Preferences)

	// Determine if should include ROMs even if scraper was not able to detect it
	optionalScraper := slices.Contains(theOptions.Preferences, "optional-scraper")

	// First, find all existing ROMs path
	// We read the current list of ROMs from the library
	existing := []string{}
	for _, shortcut := range library.Shortcuts.All() {
		if slices.Contains(shortcut.Tags, "ROM") {
			existing = append(existing, shortcut.RelativePath)
		}
	}

	// Detect available ROMs with parser in all folders / systems
	parsed, err := console.ParseROMs(theOptions)
	if err != nil {
		return err
	}

	// Filter ROMs to avoid unnecessary processing
	filtered := console.FilterROMs(parsed, existing, theOptions)
	total := len(filtered)

	// Process new ROMs into shortcuts
	// Skip if not found anything
	if total == 0 {
		cli.Printf(cli.ColorNotice, "No new ROMs to process.\n")
	} else {

		// Print initial process information
		cli.Printf(cli.ColorNotice, "%d new ROMs detected to process.\n", total)
		cli.Printf(cli.ColorNotice, "This could take some time, please be patient...\n")

		// Process each ROM to add or update
		for index, rom := range filtered {

			cli.Printf(cli.ColorNotice, "Processing ROM [%d/%d]: %s\n", index+1, total, rom.RelativePath)

			// Create shortcut information
			startDirectory := filepath.Dir(rom.Executable)
			shortcutID := shortcuts.GenerateID(rom.Name, rom.Executable)
			shortcut := &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        rom.Program,
				Name:           rom.Name,
				Description:    rom.Description,
				StartDirectory: cli.Quote(startDirectory),
				Executable:     cli.Quote(rom.Executable),
				LaunchOptions:  rom.LaunchOptions,
				ShortcutPath:   "",
				RelativePath:   rom.RelativePath,
				IconURL:        "",
				LogoURL:        "",
				CoverURL:       "",
				BannerURL:      "",
				HeroURL:        "",
				Tags:           []string{"Gaming", "ROM", rom.Platform},
			}

			// Scrape additional ROM information
			scraped, err := ScrapeShortcut(shortcut)
			if err != nil {
				return err
			}

			// Skip if scrape not found anything...
			if !optionalScraper && !scraped {
				cli.Printf(cli.ColorWarn, "Could not detect shortcut information. Skipping...\n")
				continue
			} else if optionalScraper && !scraped {
				cli.Printf(cli.ColorWarn, "Could not detect shortcut information. Using available data...\n")
			} else if scraped {
				nameFormat := fmt.Sprintf("${NAME} [%s]", rom.Platform)
				shortcut.Name = strings.Replace(nameFormat, "${NAME}", shortcut.Name, 1)
			}

			// Avoid duplicates by checking on existing library
			// This process also allow switching emulators on existing shortcut
			for _, existing := range library.Shortcuts.All() {

				// Check if shortcut is managed ROM
				if !slices.Contains(existing.Tags, "ROM") {
					continue
				}

				// Check for matching ROM relative path
				if existing.RelativePath != shortcut.RelativePath {
					continue
				}

				// When detected, merge and update existing shortcut
				cli.Debug("Updating existing shortcut: %s\n", existing.ID)
				existing.Merge(shortcut)
				shortcut = existing
				break
			}

			// Add or update into shortcuts library
			err = library.Shortcuts.Set(shortcut, false)
			if err != nil {
				return err
			}

		}

		// Print finished process information
		cli.Printf(cli.ColorNotice, "All ROMs processed.\n")

	}

	// Clean shortcuts for not found ROMs
	cli.Printf(cli.ColorNotice, "Checking for removed ROMs.\n")
	for _, shortcut := range library.Shortcuts.All() {

		// Check if shortcut is managed ROM
		if !slices.Contains(shortcut.Tags, "ROM") {
			continue
		}

		// Check if the ROM of the shortcut is on the list of parsed ROMs
		found := false
		for _, rom := range parsed {
			if shortcut.RelativePath == rom.RelativePath {
				found = true
				break
			}
		}
		if found {
			continue
		}

		// If not found, remove the shortcuts
		cli.Printf(cli.ColorNotice, "Removing shortcut for not detected ROM: %s\n", shortcut.RelativePath)
		err := library.Shortcuts.Remove(shortcut)
		if err != nil {
			return err
		}

	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	return nil
}
