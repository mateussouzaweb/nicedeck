package server

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mateussouzaweb/nicedeck/frontend"
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/platforms"
	"github.com/mateussouzaweb/nicedeck/src/programs"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

var gridFS fs.FS
var staticFS fs.FS
var gridHandler http.Handler
var staticHandler http.Handler

// Load library result
type LoadLibraryResult struct {
	Status       string `json:"status"`
	Error        string `json:"error"`
	SteamRuntime string `json:"steamRuntime"`
	SteamPath    string `json:"steamPath"`
	ConfigPath   string `json:"configPath"`
	ArtworksPath string `json:"artworksPath"`
}

// Load library action
func loadLibrary(context *Context) error {

	result := LoadLibraryResult{}

	// Load user library
	err := library.Load()
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Create FS with loaded artworks path
	config := library.GetConfig()
	gridFS = os.DirFS(config.ArtworksPath)
	gridHandler = http.FileServer(http.FS(gridFS))

	// Print loaded data
	result.Status = "OK"
	result.SteamRuntime = config.SteamRuntime
	result.SteamPath = config.SteamPath
	result.ConfigPath = config.ConfigPath
	result.ArtworksPath = config.ArtworksPath

	return context.Status(200).JSON(result)
}

// Save library result
type SaveLibraryResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Save library action
func saveLibrary(context *Context) error {

	result := SaveLibraryResult{}

	// Save user library
	err := library.Save()
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// List platforms result
type ListProgramsResult struct {
	Status string              `json:"status"`
	Error  string              `json:"error"`
	Data   []*programs.Program `json:"data"`
}

// List programs action
func listPrograms(context *Context) error {

	result := ListProgramsResult{}

	// Retrieve programs
	data, err := programs.GetPrograms()
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	result.Data = data
	return context.Status(http.StatusOK).JSON(result)
}

// List platforms result
type ListPlatformsResult struct {
	Status string                `json:"status"`
	Error  string                `json:"error"`
	Data   []*platforms.Platform `json:"data"`
}

// List platforms action
func listPlatforms(context *Context) error {

	result := ListPlatformsResult{}

	// Retrieve platforms
	data, err := platforms.GetPlatforms(&platforms.Options{})
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	result.Data = data
	return context.Status(http.StatusOK).JSON(result)
}

// List shortcuts result
type ListShortcutsResult struct {
	Status string                `json:"status"`
	Error  string                `json:"error"`
	Data   []*shortcuts.Shortcut `json:"data"`
}

// List shortcuts action
func listShortcuts(context *Context) error {
	data := library.GetShortcuts()
	result := ListShortcutsResult{}
	result.Status = "OK"
	result.Data = data
	return context.Status(http.StatusOK).JSON(result)
}

// Launch shortcut data
type LaunchShortcutData struct {
	AppID uint `json:"appId"`
}

// Launch shortcut result
type LaunchShortcutResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Launch shortcut action
func launchShortcut(context *Context) error {

	result := LaunchShortcutResult{}

	// Bind data
	data := LaunchShortcutData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Find shortcut reference
	shortcut := library.GetShortcut(data.AppID)
	if shortcut.AppID == 0 {
		err := fmt.Errorf("could not found shortcut with appID: %v", data.AppID)
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
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
	cli.Printf(cli.ColorSuccess, "Launching: %v\n", shortcut.AppName)
	if shortcut.LaunchOptions != "" {
		err = program.Run([]string{shortcut.LaunchOptions})
	} else {
		err = program.Run([]string{})
	}

	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Modify shortcut data
type ModifyShortcutData struct {
	Action        string `json:"action"`
	AppID         uint   `json:"appId"`
	AppName       string `json:"appName"`
	StartDir      string `json:"startDir"`
	Exe           string `json:"exe"`
	LaunchOptions string `json:"launchOptions"`
	IconURL       string `json:"iconUrl"`
	LogoURL       string `json:"logoUrl"`
	CoverURL      string `json:"coverUrl"`
	BannerURL     string `json:"bannerUrl"`
	HeroURL       string `json:"heroUrl"`
}

// Modify shortcut result
type ModifyShortcutResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Modify shortcut action
func modifyShortcut(context *Context) error {

	result := ModifyShortcutResult{}

	// Bind data
	data := ModifyShortcutData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Find shortcut reference
	shortcut := library.GetShortcut(data.AppID)
	if shortcut.AppID == 0 {
		err := fmt.Errorf("could not found shortcut with appID: %v", data.AppID)
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Update shortcut
	if data.Action == "update" {
		shortcut.AppName = data.AppName
		shortcut.StartDir = data.StartDir
		shortcut.Exe = data.Exe
		shortcut.LaunchOptions = data.LaunchOptions
		shortcut.IconURL = data.IconURL
		shortcut.LogoURL = data.LogoURL
		shortcut.CoverURL = data.CoverURL
		shortcut.BannerURL = data.BannerURL
		shortcut.HeroURL = data.HeroURL

		err := library.AddToShortcuts(shortcut, true)
		if err != nil {
			result.Status = "ERROR"
			result.Error = err.Error()
			return context.Status(400).JSON(result)
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %v updated!\n", shortcut.AppID)
	}

	// Delete shortcut
	if data.Action == "delete" {
		err := library.RemoveFromShortcuts(shortcut)
		if err != nil {
			result.Status = "ERROR"
			result.Error = err.Error()
			return context.Status(400).JSON(result)
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %v removed!\n", shortcut.AppID)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Install programs data
type InstallProgramsData struct {
	Programs []string `json:"programs"`
}

// Install programs result
type InstallProgramsResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Install programs action
func installPrograms(context *Context) error {

	result := InstallProgramsResult{}

	// Bind data
	data := InstallProgramsData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Run Steam setup by making sure has required settings
	err = steam.Setup()
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Install programs in the list
	for _, program := range data.Programs {
		err := programs.Install(program)
		if err != nil {
			result.Status = "ERROR"
			result.Error = err.Error()
			return context.Status(400).JSON(result)
		}
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Remove programs data
type RemoveProgramsData struct {
	Programs []string `json:"programs"`
}

// Remove programs result
type RemoveProgramsResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Remove programs actions
func removePrograms(context *Context) error {

	result := RemoveProgramsResult{}

	// Bind data
	data := RemoveProgramsData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Remove programs in the list
	for _, program := range data.Programs {
		err := programs.Remove(program)
		if err != nil {
			result.Status = "ERROR"
			result.Error = err.Error()
			return context.Status(400).JSON(result)
		}
	}

	cli.Printf(cli.ColorSuccess, "Remove process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Backup state data
type BackupStateData struct {
	Platforms   []string `json:"platforms"`
	Preferences []string `json:"preferences"`
}

// Backup state result
type BackupStateResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Backup state action
func backupState(context *Context) error {

	result := BackupStateResult{}

	// Bind data
	data := BackupStateData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Process synchronization
	options := platforms.ToOptions(data.Platforms, data.Preferences)
	err = platforms.SyncState("backup", options)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Restore state data
type RestoreStateData struct {
	Platforms   []string `json:"platforms"`
	Preferences []string `json:"preferences"`
}

// Restore state result
type RestoreStateResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Restore state action
func restoreState(context *Context) error {

	result := RestoreStateResult{}

	// Bind data
	data := RestoreStateData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Process synchronization
	options := platforms.ToOptions(data.Platforms, data.Preferences)
	err = platforms.SyncState("restore", options)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Process ROMs data
type ProcessROMsData struct {
	Platforms   []string `json:"platforms"`
	Preferences []string `json:"preferences"`
}

// Process ROMS result
type ProcessROMsResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Process ROMs action (to update library)
func processROMs(context *Context) error {

	result := ProcessROMsResult{}

	// Bind data
	data := ProcessROMsData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Process ROMs to add/update/remove
	options := platforms.ToOptions(data.Platforms, data.Preferences)
	err = platforms.Process(options)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Scrape data result
type ScrapeDataResult struct {
	Status string                `json:"status"`
	Error  string                `json:"error"`
	Result *scraper.ScrapeResult `json:"result"`
}

// Scrape data action
func scrapeData(context *Context) error {

	result := ScrapeDataResult{}

	// Bind data
	term := context.Request.URL.Query().Get("term")

	// Scrape term data
	data, err := scraper.ScrapeFromName(term)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	result.Result = data
	return context.Status(200).JSON(result)
}

// Open link data
type OpenLinkData struct {
	Link string `json:"link"`
}

// Open link result
type OpenLinkResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Open link action
func openLink(context *Context) error {

	result := OpenLinkResult{}

	// Bind data
	data := OpenLinkData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Call open from system favorite browser
	err = cli.Open(data.Link)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Setup server endpoints and shutdown channel
func Setup(version string, developmentMode bool, shutdown chan bool) error {

	// Init user library
	err := library.Init(version)
	if err != nil {
		return err
	}

	// Load static FS
	staticFS = frontend.GetStaticFS(developmentMode)
	staticHandler = http.FileServer(http.FS(staticFS))

	// Access log middleware
	Use(func(next Handler) Handler {
		return func(context *Context) error {

			// Run handle
			start := time.Now()
			err := next(context)
			end := time.Since(start)

			// Print access log
			elapsed := strconv.FormatInt(int64(end/time.Microsecond), 10)
			fmt.Printf(
				"[%s] %s - %d - %sms\n",
				context.Request.Method,
				context.Request.RequestURI,
				context.StatusCode,
				elapsed,
			)

			// Return resulting error
			return err
		}
	})

	// Logger middleware
	Use(func(next Handler) Handler {
		return func(context *Context) error {

			// Run handler
			err := next(context)

			// Print message when there is error
			if err != nil {
				cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			}

			// Return resulting error
			return err
		}
	})

	// Any command in routes should output to buffer
	// This can be read or clear later with endpoint
	var buffer bytes.Buffer
	noColor := cli.GetEnv("NO_COLOR", "")

	Add("POST", "/api/console/capture", func(context *Context) error {
		cli.Output(&buffer)
		cli.SetEnv("NO_COLOR", "1", true)
		return context.Status(http.StatusOK).String("OK")
	})

	Add("POST", "/api/console/release", func(context *Context) error {
		cli.SetEnv("NO_COLOR", noColor, true)
		cli.Output(os.Stdout)
		return context.Status(http.StatusOK).String("OK")
	})

	Add("GET", "/api/console/output", func(context *Context) error {
		return context.Status(http.StatusOK).String(buffer.String())
	})

	Add("POST", "/api/console/clear", func(context *Context) error {
		buffer.Reset()
		return context.Status(http.StatusOK).String("OK")
	})

	// Print version command
	Add("GET", "/api/version", func(context *Context) error {
		return context.Status(http.StatusOK).String(version)
	})

	// Specific routes
	Add("GET", "/api/programs", listPrograms)
	Add("GET", "/api/platforms", listPlatforms)
	Add("GET", "/api/shortcuts", listShortcuts)
	Add("GET", "/api/scrape", scrapeData)
	Add("POST", "/api/library/load", loadLibrary)
	Add("POST", "/api/library/save", saveLibrary)
	Add("POST", "/api/shortcut/launch", launchShortcut)
	Add("POST", "/api/shortcut/modify", modifyShortcut)
	Add("POST", "/api/programs/install", installPrograms)
	Add("POST", "/api/programs/remove", removePrograms)
	Add("POST", "/api/state/backup", backupState)
	Add("POST", "/api/state/restore", restoreState)
	Add("POST", "/api/roms", processROMs)
	Add("POST", "/api/link/open", openLink)

	// Grid image request
	Add("GET", "/grid/image/(.*)", func(context *Context) error {

		// Prevent cache headers
		context.Header("Cache-Control", "no-cache, no-store, must-revalidate;")
		context.Header("Pragma", "no-cache")
		context.Header("Expires", "0")
		context.Header("X-Content-Type-Options", "nosniff")

		// Check if requested file exist
		filename := strings.TrimPrefix(context.URI, "/grid/image/")
		filename = strings.ReplaceAll(filename, "/", string(os.PathSeparator))
		file, err := gridFS.Open(filename)
		if err == nil {
			defer file.Close()
		}

		// Reply with default image when not found
		if err != nil && errors.Is(err, fs.ErrNotExist) {
			context.Request.URL.Path = "/img/default/cover.png"
			staticHandler.ServeHTTP(context.Response, context.Request)
			return nil
		}

		// Server error when are other type of error
		if err != nil {
			return context.Status(http.StatusInternalServerError).Error(err)
		}

		// Server file with handler
		handler := http.StripPrefix("/grid/image/", gridHandler)
		handler.ServeHTTP(context.Response, context.Request)

		return nil
	})

	// 404 handle
	Add("GET", "/404", func(context *Context) error {
		notFound := http.StatusText(http.StatusNotFound)
		context.Status(http.StatusNotFound).String(notFound)
		return nil
	})

	// Static file request
	// Open fs and http handle for static content
	Add("GET", "/(.*)", func(context *Context) error {
		staticHandler.ServeHTTP(context.Response, context.Request)
		return nil
	})

	// Capture shutdown request
	Add("POST", "/app/shutdown", func(context *Context) error {
		shutdown <- true
		return nil
	})

	return nil
}
