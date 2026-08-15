package command

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/management"
	"github.com/mateussouzaweb/nicedeck/src/platforms"
	"github.com/mateussouzaweb/nicedeck/src/platforms/state"
	"github.com/mateussouzaweb/nicedeck/src/programs"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/server"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/version"
)

//go:embed resources/*
var resourcesContent embed.FS

// Print application version
func printVersion(_ Context) error {
	cli.Printf(cli.ColorDefault, "%s\n", version.Get())
	return nil
}

// Print application help
func printHelp(_ Context) error {

	help, err := resourcesContent.ReadFile("resources/help")
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorDefault, "%s\n", help)
	return nil
}

// List available programs
func listPrograms(_ Context) error {

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// List available programs
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

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// List console platforms
	consoleList, err := platforms.GetConsoles()
	if err != nil {
		return err
	}

	for _, platform := range consoleList {
		cli.Printf(cli.ColorDefault, "%s - %s\n", platform.Name, platform.Console)
	}

	// List native platforms
	nativeList, err := platforms.GetSystemNative()
	if err != nil {
		return err
	}

	for _, platform := range nativeList {
		cli.Printf(cli.ColorDefault, "%s - %s\n", platform.Name, platform.Runtime)
	}

	return nil
}

// List user shortcuts
func listShortcuts(_ Context) error {

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// List available shortcuts
	shortcuts := management.GetShortcuts()
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

	// Create scrape options
	options := scraper.ToOptions(
		search,
		context.Flag("--icon", false),
		context.Flag("--logo", false),
		context.Flag("--cover", false),
		context.Flag("--banner", false),
		context.Flag("--hero", false),
	)

	// Scrape term data
	data, err := management.ScrapeData(options)
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

// Sync library
func syncLibrary(_ Context) error {

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, management.SaveLibrary())
	}()

	// Sync library
	err = management.SyncLibrary()
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Libraries synchronized!\n")

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
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Find shortcut reference
	shortcut := management.GetShortcut(referenceID)
	if shortcut.ID == "" {
		return fmt.Errorf("could not found shortcut with ID: %s", referenceID)
	}

	// Launch the shortcut
	return management.LaunchShortcut(shortcut)
}

// Parse and create a new shortcut from path
func createShortcut(context Context) error {

	// Retrieve path
	path := context.Arg("--path", "")
	if path == "" {
		return fmt.Errorf("file path is required")
	}

	// Retrieve name
	name := context.Arg("--name", "")

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, management.SaveLibrary())
	}()

	// Process shortcut for path
	include := []string{}
	preferences := []string{}
	options := platforms.ToOptions(include, preferences)
	shortcut, err := management.ProcessPlatformShortcut(
		name,
		path,
		options,
	)

	if err != nil {
		return err
	} else if shortcut.ID == "" {
		return fmt.Errorf("could not determine the shortcut")
	}

	// Add shortcut
	err = management.SetShortcut(shortcut, true)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Shortcut %s created!\n", shortcut.ID)

	return nil
}

// Add shortcut
func addShortcut(context Context) error {

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, management.SaveLibrary())
	}()

	// Retrieve data
	ID := context.Arg("--id", "")
	program := context.Arg("--program", "")
	name := context.Arg("--name", "")
	description := context.Arg("--description", "")
	startDirectory := context.Arg("--start-directory", "")
	executable := context.Arg("--executable", "")
	launchOptions := context.Arg("--launch-options", "")
	relativePath := context.Arg("--relative-path", "")
	iconPath := context.Arg("--icon-path", "")
	logoPath := context.Arg("--logo-path", "")
	coverPath := context.Arg("--cover-path", "")
	bannerPath := context.Arg("--banner-path", "")
	heroPath := context.Arg("--hero-path", "")
	tags := context.Arg("--tags", "")

	if ID == "" {
		ID = shortcuts.GenerateID(name, executable)
	}
	if startDirectory == "" {
		startDirectory = filepath.Dir(executable)
	}

	// Add shortcut
	shortcut := &shortcuts.Shortcut{
		ID:             ID,
		Program:        program,
		Name:           name,
		Description:    description,
		StartDirectory: cli.Quote(startDirectory),
		Executable:     cli.Quote(executable),
		LaunchOptions:  launchOptions,
		RelativePath:   relativePath,
		IconPath:       iconPath,
		LogoPath:       logoPath,
		CoverPath:      coverPath,
		BannerPath:     bannerPath,
		HeroPath:       heroPath,
		Tags:           strings.Split(tags, ","),
	}

	err = management.SetShortcut(shortcut, true)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Shortcut %s added!\n", shortcut.ID)

	return nil
}

// Modify shortcut
func modifyShortcut(context Context) error {

	// Retrieve ID
	referenceID := context.Arg("--id", "")
	if referenceID == "" {
		return fmt.Errorf("shortcut ID is required")
	}

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Find shortcut reference
	shortcut := management.GetShortcut(referenceID)
	if shortcut.ID == "" {
		return fmt.Errorf("could not found shortcut with ID: %s", referenceID)
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, management.SaveLibrary())
	}()

	// Retrieve action and data
	update := context.Flag("--update", false)
	delete := context.Flag("--delete", false)
	program := context.Arg("--program", shortcut.Program)
	name := context.Arg("--name", shortcut.Name)
	description := context.Arg("--description", shortcut.Description)
	startDirectory := context.Arg("--start-directory", shortcut.StartDirectory)
	executable := context.Arg("--executable", shortcut.Executable)
	launchOptions := context.Arg("--launch-options", shortcut.LaunchOptions)
	relativePath := context.Arg("--relative-path", shortcut.RelativePath)
	iconPath := context.Arg("--icon-path", shortcut.IconPath)
	logoPath := context.Arg("--logo-path", shortcut.LogoPath)
	coverPath := context.Arg("--cover-path", shortcut.CoverPath)
	bannerPath := context.Arg("--banner-path", shortcut.BannerPath)
	heroPath := context.Arg("--hero-path", shortcut.HeroPath)
	tags := context.Arg("--tags", strings.Join(shortcut.Tags, ","))

	// Update shortcut
	if update {
		shortcut.Program = program
		shortcut.Name = name
		shortcut.Description = description
		shortcut.StartDirectory = cli.Quote(startDirectory)
		shortcut.Executable = cli.Quote(executable)
		shortcut.LaunchOptions = launchOptions
		shortcut.RelativePath = relativePath
		shortcut.IconPath = iconPath
		shortcut.LogoPath = logoPath
		shortcut.CoverPath = coverPath
		shortcut.BannerPath = bannerPath
		shortcut.HeroPath = heroPath
		shortcut.Tags = strings.Split(tags, ",")

		err := management.UpdateShortcut(shortcut, true)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %s updated!\n", shortcut.ID)
	}

	// Delete shortcut
	if delete {
		err := management.RemoveShortcut(shortcut)
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
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, management.SaveLibrary())
	}()

	// Install programs in the list
	options := programs.ToOptions(include, preferences)
	err = management.InstallPrograms(options)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")

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
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, management.SaveLibrary())
	}()

	// Remove programs in the list
	options := programs.ToOptions(include, preferences)
	err = management.RemovePrograms(options)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Remove process finished!\n")

	return nil
}

// List state
func listState(context Context) error {

	// Retrieve command details
	action := context.Arg("--action", "")
	include := context.Multiple("--platforms", ",")
	preferences := context.Multiple("--preferences", ",")

	if action == "" {
		return fmt.Errorf("action is required")
	}
	if len(include) == 0 {
		return fmt.Errorf("platform list is required")
	}

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Retrieve synchronizable results
	options := state.ToOptions(action, include, preferences)
	result, err := management.GetState(options)
	if err != nil {
		return err
	}

	// Print results
	transform := func(condition bool, value string, alternative string) string {
		if condition {
			return value
		}
		return alternative
	}

	for _, item := range result {
		entryResult := fmt.Sprintf(``+
			`Platform: %s (%s)`+"\n"+
			`Recommended: %s`+"\n"+
			`Source: %s`+"\n"+
			`Size: %d bytes%s | Modified: %d`+"\n"+
			`Destination: %s`+"\n"+
			`Size: %d bytes%s | Modified: %d`+"\n",
			item.Platform,
			item.Type,
			transform(item.Recommended, "YES", "NO"),
			item.Source.Path,
			int(item.Source.Size),
			transform(item.Source.Exist, "", "(no exist)"),
			int(item.Source.ModifiedTime),
			item.Destination.Path,
			int(item.Destination.Size),
			transform(item.Destination.Exist, "", "(no exist)"),
			int(item.Destination.ModifiedTime),
		)

		cli.Printf(cli.ColorDefault, "%s\n", entryResult)
	}

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

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Process synchronization
	options := state.ToOptions("backup", include, preferences)
	err = management.SyncState(options)
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

	// Init user library
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Process synchronization
	options := state.ToOptions("restore", include, preferences)
	err = management.SyncState(options)
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
	err := management.InitLibrary()
	if err != nil {
		return err
	}

	// Load user library
	err = management.LoadLibrary()
	if err != nil {
		return err
	}

	// Make sure to save library on finish
	defer func() {
		errors.Join(err, management.SaveLibrary())
	}()

	// Process ROMs to add/update/remove
	options := platforms.ToOptions(include, preferences)
	err = management.ProcessPlatformShortcuts(options)
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
			developmentMode,
			listenAddress,
			ready,
			context.Done,
		))
	}()

	context.Wait()
	return err
}
