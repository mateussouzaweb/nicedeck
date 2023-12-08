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
		cli.Printf(cli.ColorWarn, "Could not detect display, skipping auto open...\n")
		cli.Printf(cli.ColorWarn, "Please open the link in the navigator: %s\n", address)
		return nil
	}

	// Using Chrome
	chromeArgs := []string{
		fmt.Sprintf("--app=%s", address),
		fmt.Sprintf("--window-size=%d,%d", width, height),
		"--window-position=center",
		"--bwsi",
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
		"--disable-popup-blocking",
		"--disable-prompt-on-repost",
		"--disable-renderer-backgrounding",
		"--disable-sync",
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
		"--password-store=basic",
		"--remote-debugging-port=0",
		"--use-mock-keychain",
	}

	exist, err := fs.FileExist("/var/lib/flatpak/exports/bin/com.google.Chrome")
	if err != nil {
		return err
	} else if exist {
		return cli.Command(fmt.Sprintf(
			`flatpak run com.google.Chrome %s`,
			strings.Join(chromeArgs, " "),
		)).Run()
	}

	// Using Firefox
	exist, err = fs.FileExist("/var/lib/flatpak/exports/bin/org.mozilla.firefox")
	if err != nil {
		return err
	} else if exist {
		return cli.Command(fmt.Sprintf(
			`flatpak run org.mozilla.firefox --kiosk %s;`,
			address,
		)).Run()
	}

	// Fallback to XDG Open
	return cli.Command(fmt.Sprintf(
		`xdg-open %s;`,
		address,
	)).Run()
}
