package main

import (
	"os"

	"github.com/mateussouzaweb/nicedeck/src/build"
	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Windows build process
func main() {

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Get project folder
	workingDir, err := os.Getwd()
	if err != nil {
		exitCode = 1
		cli.Printf(cli.ColorFatal, "Windows packaging failed: %s\n", err.Error())
		return
	}

	// Build steps
	steps := build.New(build.Env(
		"GOOS=windows",
	)).Add(&build.Step{
		ID:      "build-windows-cli-amd64",
		Name:    "Building Windows CLI AMD64",
		Context: build.Env("GOARCH=amd64"),
		Command: build.Cmd(
			"go build -buildvcs=false -o %s/bin/nicedeck-windows-cli-amd64.exe %s/cmd/nicedeck",
			workingDir,
			workingDir,
		),
	}).Add(&build.Step{
		ID:      "build-windows-cli-arm64",
		Name:    "Building Windows CLI ARM64",
		Context: build.Env("GOARCH=arm64"),
		Command: build.Cmd(
			"go build -buildvcs=false -o %s/bin/nicedeck-windows-cli-arm64.exe %s/cmd/nicedeck",
			workingDir,
			workingDir,
		),
	}).Add(&build.Step{
		ID:      "build-windows-amd64",
		Name:    "Building Windows AMD64",
		Context: build.Env("GOARCH=amd64"),
		Command: build.Cmd(
			"go build -buildvcs=false -ldflags=\"-H windowsgui\" -o %s/bin/nicedeck-windows-amd64.exe %s/cmd/nicedeck",
			workingDir,
			workingDir,
		),
	}).Add(&build.Step{
		ID:      "build-windows-arm64",
		Name:    "Building Windows ARM64",
		Context: build.Env("GOARCH=arm64"),
		Command: build.Cmd(
			"go build -buildvcs=false -ldflags=\"-H windowsgui\" -o %s/bin/nicedeck-windows-arm64.exe %s/cmd/nicedeck",
			workingDir,
			workingDir,
		),
	})

	// Run build process
	err = steps.Run()
	if err != nil {
		exitCode = 1
		cli.Printf(cli.ColorFatal, "Windows packaging failed: %s\n", err.Error())
	}

}
