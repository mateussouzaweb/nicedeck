package gui

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Open UI with best available browser mode
// We use browser to avoid the need of having to write code for UI application
// Is not perfect, but is good enough for what we need
func OpenWithBrowser(address string, width int, height int) error {

	// When there no display, cannot open
	if os.Getenv("DISPLAY") == "" {

		// Display information message
		cli.Printf(cli.ColorWarn, "Could not detect display, skipping auto open...\n")
		cli.Printf(cli.ColorWarn, "Please open the link in the navigator: %s\n", address)

		// Create a never ending blocking channel
		keep := make(chan bool, 1)
		<-keep

		return nil
	}

	// Chrome like args
	chromeArgs := []string{
		fmt.Sprintf("--app=%s", address),
		fmt.Sprintf("--window-size=%d,%d", width, height),
		"--window-position=center",
		"--bwsi",
		"--allow-insecure-localhost",
		"--disable-background-mode",
		"--disable-background-networking",
		"--disable-background-timer-throttling",
		"--disable-backgrounding-occluded-windows",
		"--disable-breakpad",
		"--disable-client-side-phishing-detection",
		"--disable-component-extensions-with-background-pages",
		"--disable-component-update",
		"--disable-default-apps",
		"--disable-dev-shm-usage",
		"--disable-domain-reliability",
		"--disable-extensions",
		"--disable-features=site-per-process",
		"--disable-hang-monitor",
		"--disable-infobars",
		"--disable-ipc-flooding-protection",
		"--disable-notifications",
		"--disable-plugins",
		"--disable-plugins-discovery",
		"--disable-popup-blocking",
		"--disable-prompt-on-repost",
		"--disable-renderer-backgrounding",
		"--disable-sync",
		"--disable-sync-preferences",
		"--disable-translate",
		"--disable-windows10-custom-titlebar",
		"--inprivate",
		"--incognito",
		"--metrics-recording-only",
		"--no-default-browser-check",
		"--no-first-run",
		"--no-service-autorun",
		"--no-pings",
		"--safebrowsing-disable-auto-update",
		"--safe-mode",
		"--password-store=basic",
		"--remote-allow-origins=*",
		"--remote-debugging-port=0",
		"--use-mock-keychain",
	}

	// Using Google Chrome
	exist, err := fs.FileExist("/var/lib/flatpak/exports/bin/com.google.Chrome")
	if err != nil {
		return err
	} else if exist {
		return RunProcess(fmt.Sprintf(
			`flatpak run com.google.Chrome %s`,
			strings.Join(chromeArgs, " "),
		))
	}

	// Using Microsoft Edge
	exist, err = fs.FileExist("/var/lib/flatpak/exports/bin/com.microsoft.Edge")
	if err != nil {
		return err
	} else if exist {
		return RunProcess(fmt.Sprintf(
			`flatpak run com.microsoft.Edge %s`,
			strings.Join(chromeArgs, " "),
		))
	}

	// Using Brave Browser
	exist, err = fs.FileExist("/var/lib/flatpak/exports/bin/com.brave.Browser")
	if err != nil {
		return err
	} else if exist {
		return RunProcess(fmt.Sprintf(
			`flatpak run com.brave.Browser %s`,
			strings.Join(chromeArgs, " "),
		))
	}

	// Using Firefox
	exist, err = fs.FileExist("/var/lib/flatpak/exports/bin/org.mozilla.firefox")
	if err != nil {
		return err
	} else if exist {
		return RunProcess(fmt.Sprintf(
			`flatpak run org.mozilla.firefox --kiosk %s;`,
			address,
		))
	}

	// Fallback to XDG open
	err = cli.Command(fmt.Sprintf(
		`xdg-open %s;`,
		address,
	)).Run()
	if err != nil {
		return err
	}

	// XDG open do not keep the process running
	// So we need to create a never ending blocking channel
	keep := make(chan bool, 1)
	<-keep

	return nil
}

// Run process with blocking channel
func RunProcess(script string) error {

	// Start the command
	command := cli.Command(script)
	err := command.Start()
	if err != nil {
		return err
	}

	// Waiting until it closes and report back to main channel
	finished := make(chan bool, 1)

	go func() {
		err = command.Wait()
		finished <- true
	}()

	<-finished
	return err
}
