package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mateussouzaweb/nicedeck/docs"
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/platforms"
	"github.com/mateussouzaweb/nicedeck/src/programs"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/server"
)

// Print application version
func printVersion(context Context) error {
	cli.Printf(cli.ColorDefault, "%s\n", context.Version)
	return nil
}

// Print application help
func printHelp(_ Context) error {

	help, err := docs.Read("HELP.md")
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorDefault, "%s\n", help)
	return nil
}

// List available programs
func listPrograms(_ Context) error {

	list, err := programs.GetPrograms()
	if err != nil {
		return err
	}

	for _, program := range list {
		cli.Printf(cli.ColorDefault, "%s - %s\n", program.ID, program.Name)
	}

	return nil
}

// List available platforms
func listPlatforms(_ Context) error {

	options := platforms.Options{}
	list, err := platforms.GetPlatforms(&options)
	if err != nil {
		return err
	}

	for _, platform := range list {
		cli.Printf(cli.ColorDefault, "%s - %s\n", platform.Name, platform.Console)
	}

	return nil
}

// List user shortcuts
func listShortcuts(context Context) error {

	// Init user library
	err := library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// List available shortcuts
	shortcuts := library.Shortcuts.All()
	for _, shortcut := range shortcuts {
		cli.Printf(cli.ColorDefault, "%s - %s\n", shortcut.ID, shortcut.Name)
	}

	return nil
}

// Scrape data
func scrapeData(context Context) error {

	// Retrieve search terms
	search := context.Arg("--search", "")
	if search == "" {
		return fmt.Errorf("search terms is required")
	}

	// Scrape term data
	data, err := scraper.ScrapeFromName(search)
	if err != nil {
		return err
	}

	// Print result
	result, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorDefault, "%s\n", string(result))
	return nil
}

// Launch shortcut
func launchShortcut(context Context) error {

	// Retrieve ID
	referenceID := context.Arg("--id", "")
	if referenceID == "" {
		return fmt.Errorf("shortcut ID is required")
	}

	// Init user library
	err := library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// Find shortcut reference
	shortcut := library.Shortcuts.Get(referenceID)
	if shortcut.ID == "" {
		return fmt.Errorf("could not found shortcut with ID: %s", referenceID)
	}

	// Launch program based on running system
	program := packaging.Best(&linux.Binary{
		AppID:  shortcut.ID,
		AppBin: shortcut.Executable,
	}, &macos.Application{
		AppID:   shortcut.ID,
		AppName: shortcut.Executable,
	}, &windows.Executable{
		AppID:  shortcut.ID,
		AppExe: shortcut.Executable,
	})

	// Launch the shortcut
	cli.Printf(cli.ColorSuccess, "Launching: %s\n", shortcut.Name)
	if shortcut.LaunchOptions != "" {
		err = program.Run([]string{shortcut.LaunchOptions})
	} else {
		err = program.Run([]string{})
	}

	return err
}

// Modify shortcut
func modifyShortcut(context Context) error {

	// Retrieve ID
	referenceID := context.Arg("--id", "")
	if referenceID == "" {
		return fmt.Errorf("shortcut ID is required")
	}

	// Init user library
	err := library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// Find shortcut reference
	shortcut := library.Shortcuts.Get(referenceID)
	if shortcut.ID == "" {
		return fmt.Errorf("could not found shortcut with ID: %s", referenceID)
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, library.Save())
	}()

	// Retrieve action and data
	update := context.Flag("--update", false)
	delete := context.Flag("--delete", false)
	platform := context.Arg("--platform", shortcut.Platform)
	program := context.Arg("--program", shortcut.Program)
	layer := context.Arg("--layer", shortcut.Layer)
	theType := context.Arg("--type", shortcut.Type)
	name := context.Arg("--name", shortcut.Name)
	description := context.Arg("--description", shortcut.Description)
	startDirectory := context.Arg("--start-directory", shortcut.StartDirectory)
	executable := context.Arg("--executable", shortcut.Executable)
	launchOptions := context.Arg("--launch-options", shortcut.LaunchOptions)
	iconURL := context.Arg("--icon-url", shortcut.IconURL)
	logoURL := context.Arg("--logo-url", shortcut.LogoURL)
	coverURL := context.Arg("--cover-url", shortcut.CoverURL)
	bannerURL := context.Arg("--banner-url", shortcut.BannerURL)
	heroURL := context.Arg("--hero-url", shortcut.HeroURL)

	// Update shortcut
	if update {
		shortcut.Platform = platform
		shortcut.Program = program
		shortcut.Layer = layer
		shortcut.Type = theType
		shortcut.Name = name
		shortcut.Description = description
		shortcut.StartDirectory = startDirectory
		shortcut.Executable = executable
		shortcut.LaunchOptions = launchOptions
		shortcut.IconURL = iconURL
		shortcut.LogoURL = logoURL
		shortcut.CoverURL = coverURL
		shortcut.BannerURL = bannerURL
		shortcut.HeroURL = heroURL

		err := library.Shortcuts.Update(shortcut, true)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %s updated!\n", shortcut.ID)
	}

	// Delete shortcut
	if delete {
		err := library.Shortcuts.Remove(shortcut)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %s removed!\n", shortcut.ID)
	}

	return nil
}

// Install programs
func installPrograms(context Context) error {

	// Retrieve command details
	include := context.Multiple("--programs", ",")
	preferences := context.Multiple("--preferences", ",")

	if len(include) == 0 {
		return fmt.Errorf("programs list is required")
	}

	// Init user library
	err := library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, library.Save())
	}()

	// Install programs in the list
	options := programs.ToOptions(include, preferences)
	err = programs.Install(options)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// Remove programs
func removePrograms(context Context) error {

	// Retrieve command details
	include := context.Multiple("--programs", ",")
	preferences := context.Multiple("--preferences", ",")

	if len(include) == 0 {
		return fmt.Errorf("programs list is required")
	}

	// Init user library
	err := library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, library.Save())
	}()

	// Remove programs in the list
	options := programs.ToOptions(include, preferences)
	err = programs.Remove(options)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Remove process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// Backup state
func backupState(context Context) error {

	// Retrieve command details
	preferences := context.Multiple("--preferences", ",")
	include := context.Multiple("--platforms", ",")

	if len(include) == 0 {
		return fmt.Errorf("platform list is required")
	}

	// Process synchronization
	options := platforms.ToOptions(include, preferences)
	err := platforms.SyncState("backup", options)
	if err != nil {
		return err
	}

	return nil
}

// Restore state
func restoreState(context Context) error {

	// Retrieve command details
	include := context.Multiple("--platforms", ",")
	preferences := context.Multiple("--preferences", ",")

	if len(include) == 0 {
		return fmt.Errorf("platform list is required")
	}

	// Process synchronization
	options := platforms.ToOptions(include, preferences)
	err := platforms.SyncState("restore", options)
	if err != nil {
		return err
	}

	return nil
}

// Process ROMs
func processROMs(context Context) error {

	// Retrieve command details
	include := context.Multiple("--platforms", ",")
	preferences := context.Multiple("--preferences", ",")

	if len(include) == 0 {
		return fmt.Errorf("platform list is required")
	}

	// Init user library
	err := library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, library.Save())
	}()

	// Process ROMs to add/update/remove
	options := platforms.ToOptions(include, preferences)
	err = platforms.Process(options)
	if err != nil {
		return err
	}

	return nil
}

// Run server
func runServer(context Context) error {

	var err error

	// Retrieve server options
	displayMode := context.Arg("--gui", "")
	developmentMode := context.Flag("--dev", false)
	listenAddress := context.Arg("--address", "127.0.0.1:14935")
	targetURL := "http://" + listenAddress

	// Open UI with target URL when ready
	ready := make(chan bool, 1)
	go func() {
		<-ready

		// Headless mode
		if displayMode == "headless" {
			cli.Printf(cli.ColorWarn, "Running in headless mode...\n")
			cli.Printf(cli.ColorWarn, "Please open the following link in the navigator to use the app: %s\n", targetURL)
			return
		}

		// Browser mode
		errors.Join(err, cli.Open(targetURL))
	}()

	// Init server
	go func() {
		errors.Join(err, server.Init(
			context.Version,
			developmentMode,
			listenAddress,
			ready,
			context.Done,
		))
	}()

	context.Wait()
	return err
}
