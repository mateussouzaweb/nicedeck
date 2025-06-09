package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

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
	"github.com/mateussouzaweb/nicedeck/src/steam"
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
	shortcuts := library.GetShortcuts()
	for _, shortcut := range shortcuts {
		cli.Printf(cli.ColorDefault, "%v - %s\n", shortcut.AppID, shortcut.AppName)
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

	// Retrieve appID
	reference := context.Arg("--id", "")
	if reference == "" {
		return fmt.Errorf("shortcut appID is required")
	}

	// Convert value to Uint
	referenceID, err := strconv.ParseUint(reference, 10, 64)
	if err != nil {
		return err
	}

	// Init user library
	err = library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// Find shortcut reference
	shortcut := library.GetShortcut(uint(referenceID))
	if shortcut.AppID == 0 {
		return fmt.Errorf("could not found shortcut with appID: %v", referenceID)
	}

	// Launch program based on running system
	appID := fmt.Sprintf("%v", shortcut.AppID)
	executable := steam.CleanExec(shortcut.Exe)
	program := packaging.Best(&linux.Binary{
		AppID:  appID,
		AppBin: executable,
	}, &macos.Application{
		AppID:   appID,
		AppName: executable,
	}, &windows.Executable{
		AppID:  appID,
		AppExe: executable,
	})

	// Launch the shortcut
	cli.Printf(cli.ColorSuccess, "Launching: %s\n", shortcut.AppName)
	if shortcut.LaunchOptions != "" {
		err = program.Run([]string{shortcut.LaunchOptions})
	} else {
		err = program.Run([]string{})
	}

	return err
}

// Modify shortcut
func modifyShortcut(context Context) error {

	// Retrieve appID
	reference := context.Arg("--id", "")
	if reference == "" {
		return fmt.Errorf("shortcut appID is required")
	}

	// Convert value to Uint
	referenceID, err := strconv.ParseUint(reference, 10, 64)
	if err != nil {
		return err
	}

	// Init user library
	err = library.Init(context.Version)
	if err != nil {
		return err
	}

	// Load user library
	err = library.Load()
	if err != nil {
		return err
	}

	// Find shortcut reference
	shortcut := library.GetShortcut(uint(referenceID))
	if shortcut.AppID == 0 {
		return fmt.Errorf("could not found shortcut with appID: %v", referenceID)
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, library.Save())
	}()

	// Retrieve action and data
	update := context.Flag("--update", false)
	delete := context.Flag("--delete", false)
	appName := context.Arg("--app-name", shortcut.AppName)
	startDir := context.Arg("--start-dir", shortcut.StartDir)
	exe := context.Arg("--exe", shortcut.Exe)
	launchOptions := context.Arg("--launch-options", shortcut.LaunchOptions)
	iconURL := context.Arg("--icon-url", shortcut.IconURL)
	logoURL := context.Arg("--logo-url", shortcut.LogoURL)
	coverURL := context.Arg("--cover-url", shortcut.CoverURL)
	bannerURL := context.Arg("--banner-url", shortcut.BannerURL)
	heroURL := context.Arg("--hero-url", shortcut.HeroURL)

	// Update shortcut
	if update {
		shortcut.AppName = appName
		shortcut.StartDir = startDir
		shortcut.Exe = exe
		shortcut.LaunchOptions = launchOptions
		shortcut.IconURL = iconURL
		shortcut.LogoURL = logoURL
		shortcut.CoverURL = coverURL
		shortcut.BannerURL = bannerURL
		shortcut.HeroURL = heroURL

		err := library.AddToShortcuts(shortcut, true)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %v updated!\n", shortcut.AppID)
	}

	// Delete shortcut
	if delete {
		err := library.RemoveFromShortcuts(shortcut)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %v removed!\n", shortcut.AppID)
	}

	return nil
}

// Install programs
func installPrograms(context Context) error {

	// Retrieve list of programs to install
	list := context.Multiple("--programs", ",")
	if len(list) == 0 {
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

	// Run Steam setup by making sure has required settings
	err = steam.Setup()
	if err != nil {
		return err
	}

	// Install programs in the list
	for _, program := range list {
		err := programs.Install(program)
		if err != nil {
			return err
		}
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// Remove programs
func removePrograms(context Context) error {

	// Retrieve list of programs to remove
	list := context.Multiple("--programs", ",")
	if len(list) == 0 {
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

	// Run Steam setup by making sure has required settings
	err = steam.Setup()
	if err != nil {
		return err
	}

	// Remove programs in the list
	for _, program := range list {
		err := programs.Remove(program)
		if err != nil {
			return err
		}
	}

	cli.Printf(cli.ColorSuccess, "Remove process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// Sync state
func syncState(context Context) error {

	// Retrieve command details
	preferences := context.Multiple("--preferences", ",")
	include := context.Multiple("--platforms", ",")

	if len(include) == 0 {
		return fmt.Errorf("platform list is required")
	}

	// Process synchronization
	options := platforms.ToOptions(include, preferences)
	err := platforms.SyncState(options)
	if err != nil {
		return err
	}

	return nil
}

// Process ROMs
func processROMs(context Context) error {

	// Retrieve command details
	preferences := context.Multiple("--preferences", ",")
	include := context.Multiple("--platforms", ",")

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
