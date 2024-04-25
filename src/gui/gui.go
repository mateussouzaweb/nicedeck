package gui

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui/browser"
	"github.com/mateussouzaweb/nicedeck/src/gui/gtk"
	"github.com/mateussouzaweb/nicedeck/src/gui/headless"
	"github.com/mateussouzaweb/nicedeck/src/gui/qt"
)

// Open UI with best available GUI mode
func Open(mode string, url string, version string) error {

	if mode == "gtk" {
		return gtk.Open(url, version)
	} else if mode == "qt" {
		return qt.Open(url, version)
	} else if mode == "browser" {
		return browser.Open(url)
	} else if mode != "headless" {
		cli.Printf(cli.ColorWarn, "Unknown GUI launch mode: %s\n", mode)
		cli.Printf(cli.ColorWarn, "Falling back to headless mode...\n")
	}

	return headless.Open(url)
}
