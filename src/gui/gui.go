package gui

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui/browser"
	"github.com/mateussouzaweb/nicedeck/src/gui/headless"
	"github.com/mateussouzaweb/nicedeck/src/gui/webview"
)

// Open UI with best available GUI mode
func Open(mode string, url string, version string) error {

	if mode == "webview" {
		return webview.Open(url, version)
	} else if mode == "browser" {
		return browser.Open(url)
	} else if mode != "headless" {
		cli.Printf(cli.ColorWarn, "Unknown GUI launch mode: %s", mode)
		cli.Printf(cli.ColorWarn, "Falling back to headless mode...")
	}

	return headless.Open(url)
}
