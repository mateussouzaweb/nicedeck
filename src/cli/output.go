package cli

import (
	"fmt"
	"io"
	"os"
)

// Output
var output io.Writer = os.Stdout

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

// Output set the destination for console messages
func Output(writer io.Writer) {
	output = writer
}

// NoColor check if should avoid color output on console messages
func NoColor() bool {
	return os.Getenv("NO_COLOR") != "" || os.Getenv("CLICOLOR") == "0"
}

// Print a message to console output with color level
func Printf(color string, format string, args ...any) {
	if NoColor() {
		fmt.Fprintf(output, format, args...)
	} else {
		fmt.Fprintf(output, color+format+ColorReset, args...)
	}
}

// Print debug message
func Debug(format string, args ...any) {
	if os.Getenv("DEBUG") == "1" {
		Printf(ColorPurple, "[DEBUG] "+format, args...)
	}
}
