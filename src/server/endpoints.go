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
	"syscall"
	"time"

	"github.com/mateussouzaweb/nicedeck/frontend"
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/install"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/nicedeck"
	"github.com/mateussouzaweb/nicedeck/src/roms"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
)

var gridFS fs.FS
var staticFS fs.FS
var gridHandler http.Handler
var staticHandler http.Handler

// Load library result
type LoadLibraryResult struct {
	Status           string `json:"status"`
	Error            string `json:"error"`
	SteamPath        string `json:"steamPath"`
	UserConfigPath   string `json:"userConfigPath"`
	UserArtworksPath string `json:"userArtworksPath"`
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
	artworksPath := library.GetConfig().UserArtworksPath
	gridFS = os.DirFS(artworksPath)
	gridHandler = http.FileServer(http.FS(gridFS))

	// Print loaded data
	result.Status = "OK"
	result.SteamPath = library.GetConfig().SteamPath
	result.UserArtworksPath = library.GetConfig().UserArtworksPath
	result.UserConfigPath = library.GetConfig().UserConfigPath

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

// List shortcuts action
func listPrograms(context *Context) error {
	data := install.GetPrograms()
	return context.Status(http.StatusOK).JSON(data)
}

// List platforms action
func listPlatforms(context *Context) error {
	data := roms.GetPlatforms(&roms.Options{})
	return context.Status(http.StatusOK).JSON(data)
}

// List shortcuts action
func listShortcuts(context *Context) error {
	shortcuts := library.GetShortcuts()
	return context.Status(http.StatusOK).JSON(shortcuts)
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

	// Launch the shortcut
	cli.Printf(cli.ColorSuccess, "Launching: %v\n", shortcut.AppName)
	script := fmt.Sprintf(`cd %s; %s %s`, shortcut.StartDir, shortcut.Exe, shortcut.LaunchOptions)
	err = cli.Command(script).Start()

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

// Run setup data
type RunSetupData struct {
	UseSymlink  bool   `json:"useSymlink"`
	StoragePath string `json:"storagePath"`
}

// Run setup result
type RunSetupResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Run setup action (to install all programs)
func runSetup(context *Context) error {

	result := RunSetupResult{}

	// Bind data
	data := RunSetupData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Run setup by making sure has required structure
	err = install.Structure(data.UseSymlink, data.StoragePath)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Perform desktop install
	err = nicedeck.WriteDesktopShortcut()
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Run install data
type RunInstallData struct {
	Programs []string `json:"programs"`
}

// Run install result
type RunInstallResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Run install action (for specific programs only)
func runInstall(context *Context) error {

	result := RunInstallResult{}

	// Bind data
	data := RunInstallData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Install programs in the list
	for _, program := range data.Programs {
		err := install.Install(program)
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

// Process ROMs data
type ProcessROMsData struct {
	Platforms   []string `json:"platforms"`
	Preferences []string `json:"preferences"`
	Rebuild     bool     `json:"rebuild"`
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
	options := roms.ToOptions(data.Platforms, data.Preferences, data.Rebuild)
	err = roms.Process(options)
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

	// Call xdg-open to use favorite browser
	script := fmt.Sprintf(`xdg-open %s`, data.Link)
	err = cli.Command(script).Run()
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Setup server endpoints
func Setup(version string) error {

	// Load static FS
	staticFS = frontend.GetStaticFS()
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
			fmt.Printf("[%s] %s - %sms\n", context.Request.Method, context.Request.RequestURI, elapsed)

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
	noColor := os.Getenv("NO_COLOR")

	Add("POST", "/api/console/capture", func(context *Context) error {
		cli.Output(&buffer)
		os.Setenv("NO_COLOR", "1")
		return context.Status(http.StatusOK).String("OK")
	})

	Add("POST", "/api/console/release", func(context *Context) error {
		os.Setenv("NO_COLOR", noColor)
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
	Add("POST", "/api/setup", runSetup)
	Add("POST", "/api/install", runInstall)
	Add("POST", "/api/roms", processROMs)
	Add("POST", "/api/link/open", openLink)

	// Capture shutdown request
	Add("POST", "/app/shutdown", func(context *Context) error {
		syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		return nil
	})

	// Grid image request
	Add("GET", "/grid/image/(.*)", func(context *Context) error {

		// Prevent cache headers
		context.Header("Cache-Control", "no-cache, no-store, must-revalidate;")
		context.Header("Pragma", "no-cache")
		context.Header("Expires", "0")
		context.Header("X-Content-Type-Options", "nosniff")

		// Check if requested file exist
		filename := strings.TrimPrefix(context.URI, "/grid/image/")
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

	// Static file request
	// Open fs and http handle for static content
	Add("GET", "/(.*)", func(context *Context) error {
		staticHandler.ServeHTTP(context.Response, context.Request)
		return nil
	})

	return nil
}
