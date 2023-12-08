package roms

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Filter ROMs that match given requirements and return the list to process
func FilterROMs(roms []*ROM, options *Options) []*ROM {

	var existing []*ROM
	var toProcess []*ROM

	// Read current list of ROMs in the library shortcuts
	for _, shortcut := range library.GetShortcuts() {

		// Check if shortcut is managed ROM
		if !slices.Contains(shortcut.Tags, "ROM") {
			continue
		}

		// Check if ROM of the shortcut is on the list of ROMs
		for _, rom := range roms {
			if shortcut.RelativePath == rom.RelativePath {
				existing = append(existing, rom)
				break
			}
		}

	}

	// Fill the list of ROMs to process
	for _, rom := range roms {

		addToList := false

		// Add to the list if ROM matches platform condition
		if len(options.Platforms) == 0 || slices.Contains(options.Platforms, rom.Platform) {
			addToList = true
		}

		// When is not rebuilding, include only new detected ROMs
		if !options.Rebuild {
			for _, item := range existing {
				if item.RelativePath == rom.RelativePath {
					addToList = false
					break
				}
			}
		}

		// Finally, if valid, add to the list of ROMs to process
		if addToList {
			toProcess = append(toProcess, rom)
		}

	}

	return toProcess
}

// Process ROMs to scrape data and add to shortcuts list
func ProcessROMs(parsed []*ROM, options *Options) (int, error) {

	// Filter list to know what ROMs process
	process := FilterROMs(parsed, options)
	total := len(process)

	// Skip if not found anything
	if total == 0 {
		cli.Printf(cli.ColorNotice, "No new ROMs to process.\n")
		return total, nil
	}

	// Print initial process information
	cli.Printf(cli.ColorNotice, "%d new ROMs detected to process.\n", total)
	cli.Printf(cli.ColorNotice, "This could take some time, please be patient...\n")

	// Process each ROM to add or update
	for index, rom := range process {

		cli.Printf(cli.ColorNotice, "Processing ROM [%d/%d]: %s\n", index+1, total, rom.RelativePath)

		// Scrape additional ROM information
		scrape, err := scraper.ScrapeFromName(rom.Name)
		if err != nil {
			return total, err
		}

		// Skip if scrape not found anything...
		if scrape.Name == "" {
			cli.Printf(cli.ColorWarn, "Could not detect ROM information. Skipping...\n")
			continue
		}

		// Determine best name and images for the shortcut
		appName := scrape.Name + " [" + rom.Platform + "]"
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

		// Add to shortcuts library
		err = library.AddToShortcuts(&shortcuts.Shortcut{
			AppName:       appName,
			Exe:           "/var/lib/flatpak/exports/bin/" + rom.Emulator,
			StartDir:      "/var/lib/flatpak/exports/bin/", // Same as main flatpak
			ShortcutPath:  "",
			LaunchOptions: rom.LaunchOptions,
			IconURL:       iconURL,
			LogoURL:       logoURL,
			CoverURL:      coverURL,
			BannerURL:     bannerURL,
			HeroURL:       heroURL,
			Platform:      rom.Platform,
			RelativePath:  rom.RelativePath,
			Tags:          []string{"Gaming", "ROM"},
		})

		if err != nil {
			return total, err
		}

	}

	cli.Printf(cli.ColorNotice, "All ROMs processed.\n")
	return total, nil
}

// Clean shortcuts for not found ROMs
func CleanShortcuts(parsed []*ROM) (int, error) {

	var toRemove []*shortcuts.Shortcut

	// Read current list of ROMs in the library shortcuts
	for _, shortcut := range library.GetShortcuts() {

		// Check if shortcut is managed ROM
		if !slices.Contains(shortcut.Tags, "ROM") {
			continue
		}

		// Check if ROM of the shortcut is on the list of ROMs
		found := false
		for _, rom := range parsed {
			if shortcut.RelativePath == rom.RelativePath {
				found = true
				break
			}
		}

		// If not found, put on list of shortcuts to remove
		if !found {
			toRemove = append(toRemove, shortcut)
		}

	}

	// Validate if has shortcuts to remove
	total := len(toRemove)
	if total == 0 {
		return total, nil
	}

	// Print message when there are ROMs to remove
	cli.Printf(cli.ColorNotice, "Found removed ROMs. Cleaning shortcuts...\n")

	// Remove ROM shortcuts that was not detected in the list of parsed ROMs
	for _, shortcut := range toRemove {

		cli.Printf(cli.ColorNotice, "Removing shortcut for not detected ROM: %s\n", shortcut.RelativePath)
		err := library.RemoveFromShortcuts(shortcut)
		if err != nil {
			return total, err
		}

	}

	return total, nil
}

// Parse and process ROMs for given platforms
func Process(options *Options) error {

	// Detect available ROMs with parser in all folders / systems
	parsed, err := ParseROMs(options)
	if err != nil {
		return err
	}

	// Process new ROMs
	processed, err := ProcessROMs(parsed, options)
	if err != nil {
		return err
	}

	// Remove shortcuts for inexisting ROMs
	removed, err := CleanShortcuts(parsed)
	if err != nil {
		return err
	}

	if processed > 0 || removed > 0 {
		cli.Printf(cli.ColorSuccess, "Process finished!\n")
		cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")
	}

	return nil
}
