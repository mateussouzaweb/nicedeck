package browser

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// Open UI in browser mode with best available browser
func Open(url string, developmentMode bool) error {

	// Using Google Chrome, Microsoft Edge or Brave Browser
	chromiumPrograms := []string{
		"google-chrome",
		"microsoft-edge",
		"brave-browser",
	}

	// Args for chromium programs
	chromiumArgs := []string{
		fmt.Sprintf("--app=%s", url),
		fmt.Sprintf("--window-size=%d,%d", 1280, 800),
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
		"--use-mock-keychain",
	}

	// Extra development flags
	if developmentMode {
		chromiumArgs = append(
			chromiumArgs,
			"--remote-allow-origins=*",
			"--remote-debugging-port=0",
		)
	}

	// Run the first available browser
	for _, programID := range chromiumPrograms {
		program, err := programs.GetProgramByID(programID)
		if err != nil {
			return err
		}

		if program.Package.Available() {
			installed, err := program.Package.Installed()
			if err != nil {
				return err
			} else if installed {
				return program.Package.Run(chromiumArgs)
			}
		}
	}

	// Using Firefox
	program, err := programs.GetProgramByID("firefox")
	if err != nil {
		return err
	}

	if program.Package.Available() {
		installed, err := program.Package.Installed()
		if err != nil {
			return err
		} else if installed {
			return program.Package.Run([]string{
				"--kiosk",
				url,
			})
		}
	}

	// Fallback to system open command
	err = cli.Open(url)
	if err != nil {
		return err
	}

	// Open may not keep the process running
	// So we need to create a never ending blocking channel
	keep := make(chan bool, 1)
	<-keep

	return nil
}
