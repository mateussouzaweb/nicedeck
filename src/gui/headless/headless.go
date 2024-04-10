package headless

import "github.com/mateussouzaweb/nicedeck/src/cli"

// Open UI in headless mode - it won't open anything :D
func Open(address string) error {

	// Display information message
	cli.Printf(cli.ColorWarn, "Running in headless mode...\n")
	cli.Printf(cli.ColorWarn, "Please open the following link in the navigator to use the app: %s\n", address)

	// Create a never ending blocking channel
	keep := make(chan bool, 1)
	<-keep

	return nil
}
