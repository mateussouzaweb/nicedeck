package main

import (
	"os"

	"github.com/mateussouzaweb/nicedeck/src/build"
	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Linux build process
func main() {

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Get project folder
	workingDir, err := os.Getwd()
	if err != nil {
		exitCode = 1
		cli.Printf(cli.ColorFatal, "Linux packaging failed: %s\n", err.Error())
		return
	}

	// Build steps
	steps := build.New(build.Env(
		"GOOS=linux",
	)).Add(&build.Step{
		ID:      "build-linux-amd64",
		Name:    "Building Linux AMD64",
		Context: build.Env("GOARCH=amd64"),
		Command: build.Cmd(
			"go build -buildvcs=false -o %s/bin/nicedeck-linux-amd64 %s/cmd/nicedeck",
			workingDir,
			workingDir,
		),
	}).Add(&build.Step{
		ID:      "build-linux-arm64",
		Name:    "Building Linux ARM64",
		Context: build.Env("GOARCH=arm64"),
		Command: build.Cmd(
			"go build -buildvcs=false -o %s/bin/nicedeck-linux-arm64 %s/cmd/nicedeck",
			workingDir,
			workingDir,
		),
	})

	// Run build process
	err = steps.Run()
	if err != nil {
		exitCode = 1
		cli.Printf(cli.ColorFatal, "Linux packaging failed: %s\n", err.Error())
	}

}
