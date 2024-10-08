//go:build qt

package gui

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui/browser"
	"github.com/mateussouzaweb/nicedeck/src/gui/headless"
	"github.com/mateussouzaweb/nicedeck/src/gui/qt"
)

// Open GUI with selected display mode
func Open(displayMode string, url string, version string, developmentMode bool) error {
	switch displayMode {
	case "", "qt":
		return qt.Open(url, version, developmentMode)
	case "browser":
		return browser.Open(url, developmentMode)
	case "headless":
		return headless.Open(url)
	default:
		cli.Printf(cli.ColorWarn, "Unknown GUI display mode: %s\n", displayMode)
		cli.Printf(cli.ColorWarn, "Falling back to headless mode...\n")
		return headless.Open(url)
	}
}
