package gui

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui/browser"
	"github.com/mateussouzaweb/nicedeck/src/gui/gtk"
	"github.com/mateussouzaweb/nicedeck/src/gui/headless"
)

// Open UI with best available GUI mode
func Open(mode string, url string, version string, developmentMode bool) error {

	// if mode == "qt" {
	// 	return qt.Open(url, version, developmentMode)
	// }

	if mode == "gtk" {
		return gtk.Open(url, version, developmentMode)
	} else if mode == "browser" {
		return browser.Open(url, developmentMode)
	} else if mode != "headless" {
		cli.Printf(cli.ColorWarn, "Unknown GUI launch mode: %s\n", mode)
		cli.Printf(cli.ColorWarn, "Falling back to headless mode...\n")
	}

	return headless.Open(url)
}
