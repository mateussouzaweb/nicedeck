package management

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/platforms"
	"github.com/mateussouzaweb/nicedeck/src/platforms/console"
	"github.com/mateussouzaweb/nicedeck/src/platforms/launcher"
	"github.com/mateussouzaweb/nicedeck/src/platforms/native"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Parse and process platform to shortcut with given path
func PlatformToShortcut(name string, path string, options *platforms.Options) (*shortcuts.Shortcut, error) {

	processConsoles := true
	processLaunchers := true
	processNative := true
	shortcut := &shortcuts.Shortcut{}
	nameFormat := "${NAME}"

	// Make sure path is unquoted
	path = cli.Unquote(path)
	path = fs.NormalizePath(path)

	// Try to parse a console ROM
	if shortcut.ID == "" && processConsoles {

		theOptions, err := console.ToOptions(
			options.Include,
			options.Preferences,
		)
		if err != nil {
			return shortcut, err
		}

		rom, err := console.Parse(path, theOptions)
		if err != nil {
			return shortcut, err
		}

		// When valid, create a new shortcut from console ROM
		if rom.Executable != "" {
			nameFormat = fmt.Sprintf("${NAME} [%s]", rom.Platform)
			shortcutID := shortcuts.GenerateID(rom.Name, rom.Executable)
			shortcut = &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        rom.Program,
				Name:           rom.Name,
				Description:    rom.Description,
				StartDirectory: cli.Quote(rom.StartDirectory),
				Executable:     cli.Quote(rom.Executable),
				LaunchOptions:  rom.LaunchOptions,
				Reference:      rom.Reference,
				IconPath:       "",
				LogoPath:       "",
				CoverPath:      "",
				BannerPath:     "",
				HeroPath:       "",
				Tags:           rom.Tags,
			}
		}

	}

	// Try to parse a launcher game
	if shortcut.ID == "" && processLaunchers {

		theOptions, err := launcher.ToOptions(
			options.Include,
			options.Preferences,
		)
		if err != nil {
			return shortcut, err
		}

		game, err := launcher.Parse(path, theOptions)
		if err != nil {
			return shortcut, err
		}

		// When valid, create a new shortcut from launcher game
		if game.Executable != "" {
			shortcutID := shortcuts.GenerateID(game.Name, game.Executable)
			shortcut = &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        game.Program,
				Name:           game.Name,
				Description:    game.Description,
				StartDirectory: cli.Quote(game.StartDirectory),
				Executable:     cli.Quote(game.Executable),
				LaunchOptions:  game.LaunchOptions,
				Reference:      game.Reference,
				IconPath:       "",
				LogoPath:       "",
				CoverPath:      "",
				BannerPath:     "",
				HeroPath:       "",
				Tags:           game.Tags,
			}
		}

	}

	// Try to parse a native ROM
	if shortcut.ID == "" && processNative {

		theOptions, err := native.ToOptions(
			options.Include,
			options.Preferences,
		)
		if err != nil {
			return shortcut, err
		}

		app, err := native.Parse(path, theOptions)
		if err != nil {
			return shortcut, err
		}

		// When valid, create a new shortcut from native application
		if app.Executable != "" {
			shortcutID := shortcuts.GenerateID(app.Name, app.Executable)
			shortcut = &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        strings.ToLower(app.Runtime),
				Name:           app.Name,
				Description:    app.Description,
				StartDirectory: cli.Quote(app.StartDirectory),
				Executable:     cli.Quote(app.Executable),
				LaunchOptions:  app.LaunchOptions,
				Reference:      app.Reference,
				IconPath:       "",
				LogoPath:       "",
				CoverPath:      "",
				BannerPath:     "",
				HeroPath:       "",
				Tags:           app.Tags,
			}
		}

	}

	// Scrape additional shortcut information
	if shortcut.ID != "" {

		// Force custom specific name for scrapping
		if name != "" {
			shortcut.Name = name
		}

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

// Parse and process shortcuts for given console platforms
func ProcessConsolePlatforms(options *platforms.Options) error {

	// Transform options into console specific options
	theOptions, err := console.ToOptions(
		options.Include, options.Preferences,
	)
	if err != nil {
		return err
	}

	// Determine if should include ROMs even if scraper was not able to detect it
	optionalScraper := slices.Contains(
		theOptions.Preferences,
		"optional-scraper",
	)

	// Return if shortcut is managed from this platform type
	isManaged := func(shortcut *shortcuts.Shortcut) bool {
		return slices.Contains(shortcut.Tags, "ROM")
	}

	// First, find all existing ROMs references
	// We read the current list of ROMs from the library
	existing := []string{}
	for _, shortcut := range GetShortcuts() {
		if isManaged(shortcut) {
			existing = append(existing, shortcut.Reference)
		}
	}

	// Detect available ROMs with parser in all folders / systems
	parsed, err := console.Find(theOptions)
	if err != nil {
		return err
	}

	// Filter ROMs to avoid unnecessary processing
	filtered := console.Filter(parsed, existing, theOptions)
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

			cli.Printf(cli.ColorNotice, "Processing ROM [%d/%d]: %s / %s\n", index+1, total, rom.Platform, rom.Name)

			// Create shortcut information
			shortcutID := shortcuts.GenerateID(rom.Name, rom.Executable)
			shortcut := &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        rom.Program,
				Name:           rom.Name,
				Description:    rom.Description,
				StartDirectory: cli.Quote(rom.StartDirectory),
				Executable:     cli.Quote(rom.Executable),
				LaunchOptions:  rom.LaunchOptions,
				Reference:      rom.Reference,
				IconPath:       "",
				LogoPath:       "",
				CoverPath:      "",
				BannerPath:     "",
				HeroPath:       "",
				Tags:           rom.Tags,
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
			for _, existing := range GetShortcuts() {

				// Check if shortcut is managed ROM
				if !isManaged(existing) {
					continue
				}

				// Check for matching ROM reference
				if existing.Reference != shortcut.Reference {
					continue
				}

				// When detected, merge and update existing shortcut
				cli.Debug("Updating existing shortcut: %s\n", existing.ID)
				existing.Merge(shortcut)
				shortcut = existing
				break
			}

			// Add or update into shortcuts library
			err = SetShortcut(shortcut, false)
			if err != nil {
				return err
			}

		}

		// Print finished process information
		cli.Printf(cli.ColorNotice, "All console ROMs processed.\n")

	}

	// Clean shortcuts for not found ROMs
	cli.Printf(cli.ColorNotice, "Checking for removed console ROMs.\n")
	for _, shortcut := range GetShortcuts() {

		// Check if shortcut is managed ROM
		if !isManaged(shortcut) {
			continue
		}

		// Check if the ROM of the shortcut is on the list of parsed ROMs
		found := false
		for _, rom := range parsed {
			if shortcut.Reference == rom.Reference {
				found = true
				break
			}
		}
		if found {
			continue
		}

		// If not found, remove the shortcuts
		cli.Printf(cli.ColorNotice, "Removing shortcut for not detected console ROM: %s\n", shortcut.Reference)
		err := RemoveShortcut(shortcut)
		if err != nil {
			return err
		}

	}

	return nil
}

// Parse and process shortcuts for given launcher platforms
func ProcessLauncherPlatforms(options *platforms.Options) error {

	// Transform options into launcher specific options
	theOptions, err := launcher.ToOptions(
		options.Include,
		options.Preferences,
	)
	if err != nil {
		return err
	}

	// Determine if should include games even if scraper was not able to detect it
	optionalScraper := slices.Contains(
		theOptions.Preferences,
		"optional-scraper",
	)

	// Return if shortcut is managed from this platform type
	isManaged := func(shortcut *shortcuts.Shortcut) bool {
		return slices.Contains(shortcut.Tags, "Game")
	}

	// First, find all existing games references
	// We read the current list of games from the library
	existing := []string{}
	for _, shortcut := range GetShortcuts() {
		if isManaged(shortcut) {
			existing = append(existing, shortcut.Reference)
		}
	}

	// Detect available games with parser
	parsed, err := launcher.Find(theOptions)
	if err != nil {
		return err
	}

	// Filter games to avoid unnecessary processing
	filtered := launcher.Filter(parsed, existing, theOptions)
	total := len(filtered)

	// Process new games into shortcuts
	// Skip if not found anything
	if total == 0 {
		cli.Printf(cli.ColorNotice, "No new launcher games to process.\n")
	} else {

		// Print initial process information
		cli.Printf(cli.ColorNotice, "%d new launcher games detected to process.\n", total)
		cli.Printf(cli.ColorNotice, "This could take some time, please be patient...\n")

		// Process each ROM to add or update
		for index, game := range filtered {

			cli.Printf(cli.ColorNotice, "Processing game [%d/%d]: %s / %s\n", index+1, total, game.Launcher, game.Name)

			// Create shortcut information
			shortcutID := shortcuts.GenerateID(game.Name, game.Executable)
			shortcut := &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        game.Program,
				Name:           game.Name,
				Description:    game.Description,
				StartDirectory: cli.Quote(game.StartDirectory),
				Executable:     cli.Quote(game.Executable),
				LaunchOptions:  game.LaunchOptions,
				Reference:      game.Reference,
				IconPath:       "",
				LogoPath:       "",
				CoverPath:      "",
				BannerPath:     "",
				HeroPath:       "",
				Tags:           game.Tags,
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
			}

			// Avoid duplicates by checking on existing library
			for _, existing := range GetShortcuts() {

				// Check if shortcut is managed game
				if !isManaged(existing) {
					continue
				}

				// Check for matching game reference
				if existing.Reference != shortcut.Reference {
					continue
				}

				// When detected, merge and update existing shortcut
				cli.Debug("Updating existing shortcut: %s\n", existing.ID)
				existing.Merge(shortcut)
				shortcut = existing
				break
			}

			// Add or update into shortcuts library
			err = SetShortcut(shortcut, false)
			if err != nil {
				return err
			}

		}

		// Print finished process information
		cli.Printf(cli.ColorNotice, "All launcher games processed.\n")

	}

	// Clean shortcuts for not found games
	cli.Printf(cli.ColorNotice, "Checking for removed launcher games.\n")
	for _, shortcut := range GetShortcuts() {

		// Check if shortcut is managed game
		if !isManaged(shortcut) {
			continue
		}

		// Check if the game of the shortcut is on the list of parsed games
		found := false
		for _, game := range parsed {
			if shortcut.Reference == game.Reference {
				found = true
				break
			}
		}
		if found {
			continue
		}

		// If not found, remove the shortcuts
		cli.Printf(cli.ColorNotice, "Removing shortcut for not detected launcher game: %s\n", shortcut.Reference)
		err := RemoveShortcut(shortcut)
		if err != nil {
			return err
		}

	}

	return nil
}

// Parse and process shortcuts for given native platforms
func ProcessNativePlatforms(options *platforms.Options) error {

	// Transform options into native specific options
	theOptions, err := native.ToOptions(
		options.Include,
		options.Preferences,
	)
	if err != nil {
		return err
	}

	// Determine if should include apps even if scraper was not able to detect it
	optionalScraper := slices.Contains(
		theOptions.Preferences,
		"optional-scraper",
	)

	// Return if shortcut is managed from this platform type
	isManaged := func(shortcut *shortcuts.Shortcut) bool {
		return slices.Contains(shortcut.Tags, "App")
	}

	// First, find all existing apps references
	// We read the current list of apps from the library
	existing := []string{}
	for _, shortcut := range GetShortcuts() {
		if isManaged(shortcut) {
			existing = append(existing, shortcut.Reference)
		}
	}

	// Detect available apps with parser
	parsed, err := native.Find(theOptions)
	if err != nil {
		return err
	}

	// Filter apps to avoid unnecessary processing
	filtered := native.Filter(parsed, existing, theOptions)
	total := len(filtered)

	// Process new apps into shortcuts
	// Skip if not found anything
	if total == 0 {
		cli.Printf(cli.ColorNotice, "No new apps to process.\n")
	} else {

		// Print initial process information
		cli.Printf(cli.ColorNotice, "%d new apps detected to process.\n", total)
		cli.Printf(cli.ColorNotice, "This could take some time, please be patient...\n")

		// Process each app to add or update
		for index, app := range filtered {

			cli.Printf(cli.ColorNotice, "Processing app [%d/%d]: %s\n", index+1, total, app.Name)

			// Create shortcut information
			shortcutID := shortcuts.GenerateID(app.Name, app.Executable)
			shortcut := &shortcuts.Shortcut{
				ID:             shortcutID,
				Program:        strings.ToLower(app.Name),
				Name:           app.Name,
				Description:    app.Description,
				StartDirectory: cli.Quote(app.StartDirectory),
				Executable:     cli.Quote(app.Executable),
				LaunchOptions:  app.LaunchOptions,
				Reference:      app.Reference,
				IconPath:       "",
				LogoPath:       "",
				CoverPath:      "",
				BannerPath:     "",
				HeroPath:       "",
				Tags:           app.Tags,
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
			}

			// Avoid duplicates by checking on existing library
			// This process also allow switching emulators on existing shortcut
			for _, existing := range GetShortcuts() {

				// Check if shortcut is managed app
				if !isManaged(existing) {
					continue
				}

				// Check for matching app reference
				if existing.Reference != shortcut.Reference {
					continue
				}

				// When detected, merge and update existing shortcut
				cli.Debug("Updating existing shortcut: %s\n", existing.ID)
				existing.Merge(shortcut)
				shortcut = existing
				break
			}

			// Add or update into shortcuts library
			err = SetShortcut(shortcut, false)
			if err != nil {
				return err
			}

		}

		// Print finished process information
		cli.Printf(cli.ColorNotice, "All apps processed.\n")

	}

	// Clean shortcuts for not found apps
	cli.Printf(cli.ColorNotice, "Checking for removed apps.\n")
	for _, shortcut := range GetShortcuts() {

		// Check if shortcut is managed App
		if !isManaged(shortcut) {
			continue
		}

		// Check if the app of the shortcut is on the list of parsed apps
		found := false
		for _, app := range parsed {
			if shortcut.Reference == app.Reference {
				found = true
				break
			}
		}
		if found {
			continue
		}

		// If not found, remove the shortcuts
		cli.Printf(cli.ColorNotice, "Removing shortcut for not detected app: %s\n", shortcut.Reference)
		err := RemoveShortcut(shortcut)
		if err != nil {
			return err
		}

	}

	return nil
}

// ProcessPlatforms to detect shortcuts
func ProcessPlatforms(options *platforms.Options) error {

	processConsoles := true
	processLaunchers := true
	processNative := true

	if processConsoles {
		err := ProcessConsolePlatforms(options)
		if err != nil {
			return err
		}
	}

	if processLaunchers {
		err := ProcessLauncherPlatforms(options)
		if err != nil {
			return err
		}
	}

	if processNative {
		err := ProcessNativePlatforms(options)
		if err != nil {
			return err
		}
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	return nil
}
