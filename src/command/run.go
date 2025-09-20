package command

import (
	"fmt"
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Run command based on given arguments
func Run(version string, args []string, done chan bool) error {

	// Retrieve args and command to run
	command := cli.Arg(args, "0", "server")

	// Remove command index from arguments
	index := slices.Index(args, command)
	if index != -1 {
		args = slices.Delete(args, index, index+1)
	}

	// Create context
	context := Context{
		Args:    args,
		Version: version,
		Done:    done,
	}

	// Process required command
	var err error
	switch command {
	case "version":
		err = printVersion(context)
	case "help":
		err = printHelp(context)
	case "programs":
		err = listPrograms(context)
	case "platforms":
		err = listPlatforms(context)
	case "shortcuts":
		err = listShortcuts(context)
	case "scrape":
		err = scrapeData(context)
	case "launch":
		err = launchShortcut(context)
	case "create":
		err = createShortcut(context)
	case "add":
		err = addShortcut(context)
	case "modify":
		err = modifyShortcut(context)
	case "install":
		err = installPrograms(context)
	case "remove":
		err = removePrograms(context)
	case "backup-state":
		err = backupState(context)
	case "restore-state":
		err = restoreState(context)
	case "process-roms":
		err = processROMs(context)
	case "server":
		err = runServer(context)
	default:
		err = fmt.Errorf("could not find command: %s", command)
	}

	return err
}
