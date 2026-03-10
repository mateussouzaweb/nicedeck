package main

import (
	"os"

	"github.com/mateussouzaweb/nicedeck/src/build"
	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Main build process
func main() {

	// Main build process that combines all platforms
	all := build.New(build.Env())
	all.Add(&build.Step{
		ID:      "clean-previous-builds",
		Name:    "Clean up previous builds",
		Command: build.Cmd("[ -d bin/ ] && rm -r bin/ || true"),
	}).Add(&build.Step{
		ID:      "create-bin-directory",
		Name:    "Create bin directory",
		Command: build.Cmd("mkdir -p bin/"),
	}).Add(&build.Step{
		ID:      "build-linux",
		Name:    "Build Linux",
		Context: build.Env(),
		Command: build.Cmd("go run packaging/linux/build.go"),
	}).Add(&build.Step{
		ID:      "build-macos",
		Name:    "Build MacOS",
		Context: build.Env(),
		Command: build.Cmd("go run packaging/macos/build.go"),
	}).Add(&build.Step{
		ID:      "build-windows",
		Name:    "Build Windows",
		Context: build.Env(),
		Command: build.Cmd("go run packaging/windows/build.go"),
	}).Add(&build.Step{
		ID:      "set-permissions",
		Name:    "Make everything executable",
		Command: build.Cmd("chmod +x bin/*"),
	})

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Run the entire build process
	err := all.Run()
	if err != nil {
		exitCode = 1
		cli.Printf(cli.ColorFatal, "Build failed: %s\n", err.Error())
	}

}
