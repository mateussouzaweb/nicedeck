package cli

import (
	"fmt"
	"os"
)

// Colors
var (
	ColorDefault = ""
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorPurple  = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorGray    = "\033[37m"
	ColorWhite   = "\033[97m"
	ColorNotice  = ColorCyan
	ColorWarn    = ColorYellow
	ColorFatal   = ColorRed
	ColorSuccess = ColorGreen
)

// NoColor check if should avoid color output on console
func NoColor() bool {
	return os.Getenv("NO_COLOR") != "" || os.Getenv("CLICOLOR") == "0"
}

// Print displays a info to standard output
func Printf(color string, format string, args ...any) {
	if NoColor() {
		fmt.Printf(format, args...)
	} else {
		fmt.Printf(color+format+ColorReset, args...)
	}
}
