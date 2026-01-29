package server

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mateussouzaweb/nicedeck/frontend"
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/platforms"
	"github.com/mateussouzaweb/nicedeck/src/platforms/console"
	"github.com/mateussouzaweb/nicedeck/src/platforms/native"
	"github.com/mateussouzaweb/nicedeck/src/platforms/state"
	"github.com/mateussouzaweb/nicedeck/src/programs"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

var gridFS fs.FS
var staticFS fs.FS
var gridHandler http.Handler
var staticHandler http.Handler

// Library data result
type LibraryData struct {
	Timestamp  int64  `json:"timestamp"`
	ImagesPath string `json:"imagesPath"`
}

// Load library result
type LoadLibraryResult struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   LibraryData `json:"data"`
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

	// Create FS with loaded images path
	imagesPath := library.Shortcuts.ImagesPath
	gridFS = os.DirFS(imagesPath)
	gridHandler = http.FileServer(http.FS(gridFS))

	data := LibraryData{
		Timestamp:  time.Now().Unix(),
		ImagesPath: library.Shortcuts.ImagesPath,
	}

	// Print loaded data
	result.Status = "OK"
	result.Data = data

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

// Sync library result
type SyncLibraryResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Sync library action
func syncLibrary(context *Context) error {

	result := SyncLibraryResult{}

	// Sync user library
	err := library.Sync()
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
	Status  string              `json:"status"`
	Error   string              `json:"error"`
	Console []*console.Platform `json:"console"`
	Native  []*native.Platform  `json:"native"`
}

// List platforms action
func listPlatforms(context *Context) error {

	result := ListPlatformsResult{}

	// Retrieve console platforms
	consoleOptions := &console.Options{}
	consoleList, err := console.GetPlatforms(consoleOptions)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Retrieve native platforms
	nativeOptions := &native.Options{}
	nativeList, err := native.GetPlatforms(nativeOptions)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	result.Console = consoleList
	result.Native = nativeList
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
	data := library.Shortcuts.All()
	result := ListShortcutsResult{}
	result.Status = "OK"
	result.Data = data
	return context.Status(http.StatusOK).JSON(result)
}

// Launch shortcut data
type LaunchShortcutData struct {
	ID string `json:"id"`
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
	shortcut := library.Shortcuts.Get(data.ID)
	if shortcut.ID == "" {
		err := fmt.Errorf("could not found shortcut with ID: %s", data.ID)
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Launch the shortcut
	err = library.Shortcuts.Launch(shortcut)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Create shortcut data
type CreateShortcutData struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Create shortcut result
type CreateShortcutResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Create shortcut action
func createShortcut(context *Context) error {

	result := CreateShortcutResult{}

	// Bind data
	data := CreateShortcutData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Validate requirements
	if data.Path == "" {
		err := fmt.Errorf("file path is required")
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Process shortcut for path
	options := &platforms.Options{}
	shortcut, err := platforms.ProcessShortcut(
		data.Name,
		data.Path,
		options,
	)

	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	} else if shortcut.ID == "" {
		err := fmt.Errorf("could not determine the shortcut")
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Add shortcut
	err = library.Shortcuts.Set(shortcut, true)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	cli.Printf(cli.ColorSuccess, "Shortcut %s created!\n", shortcut.ID)

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Add shortcut data
type AddShortcutData struct {
	ID             string   `json:"id"`
	Program        string   `json:"program"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	StartDirectory string   `json:"startDirectory"`
	Executable     string   `json:"executable"`
	LaunchOptions  string   `json:"launchOptions"`
	RelativePath   string   `json:"relativePath"`
	IconURL        string   `json:"iconUrl"`
	LogoURL        string   `json:"logoUrl"`
	CoverURL       string   `json:"coverUrl"`
	BannerURL      string   `json:"bannerUrl"`
	HeroURL        string   `json:"heroUrl"`
	Tags           []string `json:"tags"`
}

// Add shortcut result
type AddShortcutResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Add shortcut action
func addShortcut(context *Context) error {

	result := AddShortcutResult{}

	// Bind data
	data := AddShortcutData{}
	err := context.Bind(&data)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	if data.ID == "" {
		data.ID = shortcuts.GenerateID(data.Name, data.Executable)
	}
	if data.StartDirectory == "" {
		data.StartDirectory = filepath.Dir(data.Executable)
	}

	// Add shortcut
	shortcut := &shortcuts.Shortcut{
		ID:             data.ID,
		Program:        data.Program,
		Name:           data.Name,
		Description:    data.Description,
		StartDirectory: cli.Quote(data.StartDirectory),
		Executable:     cli.Quote(data.Executable),
		LaunchOptions:  data.LaunchOptions,
		RelativePath:   data.RelativePath,
		IconURL:        data.IconURL,
		LogoURL:        data.LogoURL,
		CoverURL:       data.CoverURL,
		BannerURL:      data.BannerURL,
		HeroURL:        data.HeroURL,
		Tags:           data.Tags,
	}

	err = library.Shortcuts.Set(shortcut, true)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	cli.Printf(cli.ColorSuccess, "Shortcut %s added!\n", shortcut.ID)

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Modify shortcut data
type ModifyShortcutData struct {
	Action         string   `json:"action"`
	ID             string   `json:"id"`
	Program        string   `json:"program"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	StartDirectory string   `json:"startDirectory"`
	Executable     string   `json:"executable"`
	LaunchOptions  string   `json:"launchOptions"`
	RelativePath   string   `json:"relativePath"`
	IconURL        string   `json:"iconUrl"`
	LogoURL        string   `json:"logoUrl"`
	CoverURL       string   `json:"coverUrl"`
	BannerURL      string   `json:"bannerUrl"`
	HeroURL        string   `json:"heroUrl"`
	Tags           []string `json:"tags"`
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
	shortcut := library.Shortcuts.Get(data.ID)
	if shortcut.ID == "" {
		err := fmt.Errorf("could not found shortcut with ID: %s", data.ID)
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	// Update shortcut
	if data.Action == "update" {
		shortcut.Program = data.Program
		shortcut.Name = data.Name
		shortcut.Description = data.Description
		shortcut.StartDirectory = cli.Quote(data.StartDirectory)
		shortcut.Executable = cli.Quote(data.Executable)
		shortcut.LaunchOptions = data.LaunchOptions
		shortcut.RelativePath = data.RelativePath
		shortcut.IconURL = data.IconURL
		shortcut.LogoURL = data.LogoURL
		shortcut.CoverURL = data.CoverURL
		shortcut.BannerURL = data.BannerURL
		shortcut.HeroURL = data.HeroURL
		shortcut.Tags = data.Tags

		// Empty images to download again
		shortcut.IconPath = ""
		shortcut.LogoPath = ""
		shortcut.CoverPath = ""
		shortcut.BannerPath = ""
		shortcut.HeroPath = ""

		err := library.Shortcuts.Update(shortcut, true)
		if err != nil {
			result.Status = "ERROR"
			result.Error = err.Error()
			return context.Status(400).JSON(result)
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %s updated!\n", shortcut.ID)
	}

	// Delete shortcut
	if data.Action == "delete" {
		err := library.Shortcuts.Remove(shortcut)
		if err != nil {
			result.Status = "ERROR"
			result.Error = err.Error()
			return context.Status(400).JSON(result)
		}

		cli.Printf(cli.ColorSuccess, "Shortcut %s removed!\n", shortcut.ID)
	}

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Install programs data
type InstallProgramsData struct {
	Programs    []string `json:"programs"`
	Preferences []string `json:"preferences"`
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

	// Install programs in the list
	options := programs.ToOptions(data.Programs, data.Preferences)
	err = programs.Install(options)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")

	result.Status = "OK"
	return context.Status(200).JSON(result)
}

// Remove programs data
type RemoveProgramsData struct {
	Programs    []string `json:"programs"`
	Preferences []string `json:"preferences"`
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
	options := programs.ToOptions(data.Programs, data.Preferences)
	err = programs.Remove(options)
	if err != nil {
		result.Status = "ERROR"
		result.Error = err.Error()
		return context.Status(400).JSON(result)
	}

	cli.Printf(cli.ColorSuccess, "Remove process finished!\n")

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
	options := state.ToOptions(data.Platforms, data.Preferences)
	err = state.SyncState("backup", options)
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
	options := state.ToOptions(data.Platforms, data.Preferences)
	err = state.SyncState("restore", options)
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
	err = platforms.ProcessShortcuts(options)
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
	Add("POST", "/api/library/sync", syncLibrary)
	Add("POST", "/api/shortcut/launch", launchShortcut)
	Add("POST", "/api/shortcut/create", createShortcut)
	Add("POST", "/api/shortcut/add", addShortcut)
	Add("POST", "/api/shortcut/modify", modifyShortcut)
	Add("POST", "/api/programs/install", installPrograms)
	Add("POST", "/api/programs/remove", removePrograms)
	Add("POST", "/api/state/backup", backupState)
	Add("POST", "/api/state/restore", restoreState)
	Add("POST", "/api/roms", processROMs)
	Add("POST", "/api/link/open", openLink)

	// Grid image request
	Add("GET", "/grid/image/(.*)", func(context *Context) error {

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
