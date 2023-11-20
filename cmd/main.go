package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/install"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/roms"
	"github.com/mateussouzaweb/nicedeck/src/server"
	"github.com/mateussouzaweb/nicedeck/src/ui"
)

var version = "Version 0.0.12"
var address = "127.0.0.1:14935"

// Print version command
func printVersion(context *server.Context) error {
	return context.Status(http.StatusOK).String(version)
}

// List shortcuts command
func listPrograms(context *server.Context) error {
	data := install.GetPrograms()
	return context.Status(http.StatusOK).JSON(data)
}

// List platforms command
func listPlatforms(context *server.Context) error {
	data := roms.GetPlatforms(&roms.Options{})
	return context.Status(http.StatusOK).JSON(data)
}

// List shortcuts command
func listShortcuts(context *server.Context) error {

	// Load library
	err := library.Load()
	if err != nil {
		return err
	}

	// List detected shortcuts
	shortcuts := library.GetShortcuts()
	return context.Status(http.StatusOK).JSON(shortcuts)
}

// Run setup command (to install all programs)
func runSetup(context *server.Context) error {

	// Parse form data
	err := context.Request.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		return err
	}

	installOnMicroSD := context.Request.FormValue("microsd_install") == "Y"
	microSDPath := context.Request.FormValue("microsd_path")

	// Load library
	err = library.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := library.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}()

	// Run setup by making sure has required structure
	err = install.Structure(installOnMicroSD, microSDPath)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Setup completed!\n")
	return context.Status(200).String("OK")
}

// Run install command (for specific programs only)
func runInstall(context *server.Context) error {

	// Parse form data
	err := context.Request.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		return err
	}

	programs := []string{}
	for _, value := range context.Request.Form["programs[]"] {
		programs = append(programs, strings.ReplaceAll(value, " ", ""))
	}

	// Load library
	err = library.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := library.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}()

	// Install programs in the list
	for _, program := range programs {
		err := install.Install(program)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return context.Status(200).String("OK")
}

// Process ROMs command (to update library)
func processROMs(context *server.Context) error {

	// Parse form data
	err := context.Request.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		return err
	}

	platforms := []string{}
	for _, value := range context.Request.Form["platforms[]"] {
		platforms = append(platforms, strings.ReplaceAll(value, " ", ""))
	}

	preferences := []string{}
	for _, value := range context.Request.Form["preferences[]"] {
		preferences = append(preferences, strings.ReplaceAll(value, " ", ""))
	}

	rebuild := context.Request.FormValue("rebuild") == "Y"

	// Load library
	err = library.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := library.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}()

	// Process ROMs to add/update/remove
	options := roms.ToOptions(platforms, preferences, rebuild)
	err = roms.Process(options)
	if err != nil {
		return err
	}

	return context.Status(200).String("OK")
}

// Main command
func main() {

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Graceful shutdown support
	exit := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Capture exit (CTRL-C)
	go func() {
		<-exit
		done <- true
	}()

	// Run the program server
	go func() {

		// Access log middleware
		server.Use(func(next server.Handler) server.Handler {
			return func(context *server.Context) error {

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
		var buffer bytes.Buffer

		server.Use(func(next server.Handler) server.Handler {
			noColor := os.Getenv("NO_COLOR")
			return func(context *server.Context) error {
				// Set logger to buffer
				cli.Output(&buffer)
				os.Setenv("NO_COLOR", "1")

				// Run handler
				err := next(context)
				if err != nil {
					cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
				}

				// Restore logger to stdout
				os.Setenv("NO_COLOR", noColor)
				cli.Output(os.Stdout)

				// Return resulting error
				return err
			}
		})

		// Any command in routes should output to buffer
		// This can be read or clear later with endpoint
		server.Add("GET", "/api/console", func(context *server.Context) error {
			return context.Status(http.StatusOK).String(buffer.String())
		})
		server.Add("GET", "/api/clear", func(context *server.Context) error {
			buffer.Reset()
			return context.Status(http.StatusOK).String("OK")
		})

		// Specific routes
		server.Add("GET", "/api/version", printVersion)
		server.Add("GET", "/api/programs", listPrograms)
		server.Add("GET", "/api/platforms", listPlatforms)
		server.Add("GET", "/api/shortcuts", listShortcuts)
		server.Add("POST", "/api/setup", runSetup)
		server.Add("POST", "/api/install", runInstall)
		server.Add("POST", "/api/roms", processROMs)

		// Capture shutdown request
		server.Add("POST", "/app/shutdown", func(context *server.Context) error {
			done <- true
			return nil
		})

		err := server.Start(address)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
		}

		done <- true
	}()

	// Open server website page in browser
	// We should wait for the serve goes up first
	go func() {
		time.Sleep(1 * time.Second)
		err := ui.Open("http://" + address)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}
	}()

	<-done
}
