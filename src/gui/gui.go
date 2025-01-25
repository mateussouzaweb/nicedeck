package gui

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui/qt"
)

// Open GUI with selected display mode
func Open(displayMode string, url string, version string, developmentMode bool) error {

	// In QT mode, open the QT application wrapper
	if displayMode == "qt" {
		return qt.Open(url, version, developmentMode)
	}

	// In browser mode, open the link on browser
	// In headless mode, display information message
	switch displayMode {
	case "browser":
		cli.Open(url)
	case "headless", "":
		cli.Printf(cli.ColorWarn, "Running in headless mode...\n")
		cli.Printf(cli.ColorWarn, "Please open the following link in the navigator to use the app: %s\n", url)
	}

	// Create a never ending blocking channel to keep application running
	keep := make(chan bool, 1)
	<-keep

	return nil
}
