package shortcuts

import (
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Prepare execution context from shortcut data
func PrepareContext(shortcut *Shortcut) *cli.Context {

	context := &cli.Context{
		WorkingDirectory: shortcut.StartDirectory,
		Executable:       shortcut.Executable,
		Arguments:        []string{},
		Environment:      []string{},
	}

	// Split launch options into parameters and environment variables
	// Implementation match Steam launch options format
	if strings.Contains(shortcut.LaunchOptions, "%command%") {
		split := strings.Split(shortcut.LaunchOptions, "%command%")
		arguments := strings.Trim(split[1], " ")
		context.Arguments = []string{arguments}

		if split[0] != "" {
			environment := strings.Trim(split[0], " ")
			context.Environment = strings.Split(environment, " ")
		}
	} else if shortcut.LaunchOptions != "" {
		context.Arguments = []string{shortcut.LaunchOptions}
	}

	// Detect if executable is a URL link to switch executable to URL opener
	// In that case, additional arguments are not accepted by URLs
	isURL := strings.Contains(context.Executable, "://")
	if isURL && (cli.IsLinux() || cli.IsMacOS()) {
		context.Arguments = []string{context.Executable}

		if cli.IsLinux() {
			context.Executable = "xdg-open"
		} else {
			context.Executable = "open"
		}
	}

	// On MacOS, enforce the open command with executable as argument
	// Behavior shares application command preparation with desktop app bundles
	if !isURL && cli.IsMacOS() {
		arguments := []string{"-n", context.Executable, "--args"}
		context.Arguments = append(arguments, context.Arguments...)
		context.Executable = "open"
	}

	return context
}
