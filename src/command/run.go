package command

import (
	"fmt"
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Run command based on given arguments
func Run(args []string, done chan bool) error {

	// Retrieve args and command to run
	command := cli.Arg(args, "0", "server")

	// Remove command index from arguments
	index := slices.Index(args, command)
	if index != -1 {
		args = slices.Delete(args, index, index+1)
	}

	// Create context
	context := Context{
		Args: args,
		Done: done,
	}

	// Process required command
	var err error
	switch command {
	case "version":
		err = printVersion(context)
	case "help":
		err = printHelp(context)
	case "server":
		err = runServer(context)
	case "scrape":
		err = scrapeData(context)
	case "sync":
		err = syncLibrary(context)
	case "list-programs":
		err = listPrograms(context)
	case "install-programs":
		err = installPrograms(context)
	case "remove-programs":
		err = removePrograms(context)
	case "list-platforms":
		err = listPlatforms(context)
	case "process-platforms":
		err = processPlatforms(context)
	case "list-state":
		err = listState(context)
	case "backup-state":
		err = backupState(context)
	case "restore-state":
		err = restoreState(context)
	case "list-shortcuts":
		err = listShortcuts(context)
	case "launch-shortcut":
		err = launchShortcut(context)
	case "create-shortcut":
		err = createShortcut(context)
	case "add-shortcut":
		err = addShortcut(context)
	case "modify-shortcut":
		err = modifyShortcut(context)
	default:
		err = fmt.Errorf("could not find command: %s", command)
	}

	return err
}
