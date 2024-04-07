package gui

import (
	"os"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui/webview"
)

// Open UI with best available GUI mode
func Open(address string, width int, height int) error {

	// When there no display, cannot open
	if os.Getenv("DISPLAY") == "" {

		// Display information message
		cli.Printf(cli.ColorWarn, "Could not detect display, skipping auto open...\n")
		cli.Printf(cli.ColorWarn, "Please open the following link in the navigator to use the app: %s\n", address)

		// Create a never ending blocking channel
		keep := make(chan bool, 1)
		<-keep

		return nil
	}

	// return browser.Open(address, width, height)
	return webview.Open(address, width, height)
}
