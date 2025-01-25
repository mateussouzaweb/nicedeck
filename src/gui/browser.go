package gui

import "github.com/mateussouzaweb/nicedeck/src/cli"

// Open GUI with selected display mode
func Open(displayMode string, url string) error {

	// Headless mode
	if displayMode == "headless" {
		cli.Printf(cli.ColorWarn, "Running in headless mode...\n")
		cli.Printf(cli.ColorWarn, "Please open the following link in the navigator to use the app: %s\n", url)
		return nil
	}

	// Browser mode
	return cli.Open(url)
}
