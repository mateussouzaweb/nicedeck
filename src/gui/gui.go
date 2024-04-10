package gui

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui/browser"
	"github.com/mateussouzaweb/nicedeck/src/gui/headless"
	"github.com/mateussouzaweb/nicedeck/src/gui/webview"
)

// Open UI with best available GUI mode
func Open(mode string, address string, width int, height int) error {

	if mode == "webview" {
		return webview.Open(address, width, height)
	} else if mode == "browser" {
		return browser.Open(address, width, height)
	} else if mode != "headless" {
		cli.Printf(cli.ColorWarn, "Unknown GUI launch mode: %s", mode)
		cli.Printf(cli.ColorWarn, "Falling back to headless mode...")
	}

	return headless.Open(address)
}
