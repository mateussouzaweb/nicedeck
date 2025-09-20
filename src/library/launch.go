package library

import (
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Launch shortcut
func Launch(shortcut *shortcuts.Shortcut) error {

	var err error
	cli.Printf(cli.ColorSuccess, "Launching: %s\n", shortcut.Name)

	// Launch program based on running system
	program := packaging.Best(&linux.Binary{
		AppID:     shortcut.ID,
		AppBin:    shortcut.Executable,
		Arguments: packaging.NoArguments(),
	}, &macos.Application{
		AppID:     shortcut.ID,
		AppName:   shortcut.Executable,
		Arguments: packaging.NoArguments(),
	}, &windows.Executable{
		AppID:     shortcut.ID,
		AppExe:    shortcut.Executable,
		Arguments: packaging.NoArguments(),
	})

	// Split launch options into parameters and variables
	// Implementation match Steam launch options format
	launchOptions := shortcut.LaunchOptions
	launchVariables := ""

	if strings.Contains(launchOptions, "%command%") {
		launchSplit := strings.Split(launchOptions, "%command%")
		if launchSplit[0] != "" {
			launchVariables = strings.Trim(launchSplit[0], " ")
			launchOptions = strings.Trim(launchSplit[1], " ")
		} else {
			launchOptions = strings.Trim(launchSplit[1], " ")
		}
	}

	// Set environment variables when available
	if launchVariables != "" {
		for _, variable := range strings.Split(launchVariables, " ") {
			variableSplit := strings.Split(variable, "=")
			err := cli.SetEnv(variableSplit[0], variableSplit[1], true)
			if err != nil {
				return err
			}
		}
	}

	cli.Debug("Start directory: %s\n", shortcut.StartDirectory)
	cli.Debug("Executable: %s\n", shortcut.Executable)
	cli.Debug("Launch variables: %s\n", launchVariables)
	cli.Debug("Launch options: %s\n", launchOptions)

	// Launch the program
	if launchOptions != "" {
		err = program.Run([]string{launchOptions})
	} else {
		err = program.Run([]string{})
	}

	return err
}
