package ui

import (
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Open UI with best available mode
// We use browser to avoid the need of having to write code for UI application
// Is not perfect, but is good enough for what we need
func Open(address string) error {

	// When there no display, cannot open
	if os.Getenv("DISPLAY") == "" {
		cli.Printf(cli.ColorWarn, "Could not detect display, skipping auto open...\n")
		cli.Printf(cli.ColorWarn, "Please open the link in the navigator: %s\n", address)
		return nil
	}

	// Using Chrome
	exist, err := fs.FileExist("/var/lib/flatpak/exports/bin/com.google.Chrome")
	if err != nil {
		return err
	} else if exist {
		return cli.Command(fmt.Sprintf(
			`flatpak run com.google.Chrome --app=%s;`,
			address,
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
