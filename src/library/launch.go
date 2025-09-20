package library

import (
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Launch shortcut
func Launch(shortcut *shortcuts.Shortcut) error {

	cli.Printf(cli.ColorSuccess, "Launching: %s\n", shortcut.Name)
	context := &cli.Context{
		WorkingDirectory: shortcut.StartDirectory,
		Executable:       shortcut.Executable,
		Arguments:        []string{shortcut.LaunchOptions},
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
	}

	cli.Debug("Working directory: %s\n", context.WorkingDirectory)
	cli.Debug("Executable: %s\n", context.Executable)
	cli.Debug("Arguments: %s\n", strings.Join(context.Arguments, " "))
	cli.Debug("Environment: %s\n", strings.Join(context.Environment, " "))

	return context.Run()
}
