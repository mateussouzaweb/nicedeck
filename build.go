package main

import (
	"os"

	"github.com/mateussouzaweb/nicedeck/src/build"
	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Cmd creates a new build command with the provided script
func Cmd(script string) *build.Command {
	cmd := cli.Command(script)
	return &build.Command{
		Callback: func() error {
			return cmd.Run()
		},
	}
}

// Env create a new build context with given environment variables
func Env(env ...string) *build.Context {
	return build.Env(env...)
}

func main() {

	// Linux build process
	linux := build.New(Env(
		"GOOS=linux",
	)).Add(&build.Step{
		ID:      "build-linux-x64",
		Name:    "Building Linux X64",
		Context: Env("GOARCH=amd64"),
		Command: Cmd("go build -o bin/nicedeck-linux-amd64 cmd/main.go"),
	}).Add(&build.Step{
		ID:      "build-linux-arm64",
		Name:    "Building Linux ARM64",
		Context: Env("GOARCH=arm64"),
		Command: Cmd("go build -o bin/nicedeck-linux-arm64 cmd/main.go"),
	})

	// MacOS build process
	macos := build.New(Env(
		"GOOS=darwin",
	)).Add(&build.Step{
		ID:      "build-macos-x64",
		Name:    "Building MacOS Intel",
		Context: Env("GOARCH=amd64"),
		Command: Cmd("go build -o bin/nicedeck-macos-amd64 cmd/main.go"),
	}).Add(&build.Step{
		ID:      "build-macos-arm64",
		Name:    "Building MacOS ARM64",
		Context: Env("GOARCH=arm64"),
		Command: Cmd("go build -o bin/nicedeck-macos-arm64 cmd/main.go"),
	})

	// Windows build process
	windows := build.New(Env(
		"GOOS=windows",
	)).Add(&build.Step{
		ID:      "build-windows-cli-amd64",
		Name:    "Building Windows CLI AMD64",
		Context: Env("GOARCH=amd64"),
		Command: Cmd("go build -o bin/nicedeck-windows-cli-amd64.exe cmd/main.go"),
	}).Add(&build.Step{
		ID:      "build-windows-cli-arm64",
		Name:    "Building Windows CLI ARM64",
		Context: Env("GOARCH=arm64"),
		Command: Cmd("go build -o bin/nicedeck-windows-cli-arm64.exe cmd/main.go"),
	}).Add(&build.Step{
		ID:      "build-windows-amd64",
		Name:    "Building Windows AMD64",
		Context: Env("GOARCH=amd64"),
		Command: Cmd("go build -ldflags=\"-H windowsgui\" -o bin/nicedeck-windows-amd64.exe cmd/main.go"),
	}).Add(&build.Step{
		ID:      "build-windows-arm64",
		Name:    "Building Windows ARM64",
		Context: Env("GOARCH=arm64"),
		Command: Cmd("go build -ldflags=\"-H windowsgui\" -o bin/nicedeck-windows-arm64.exe cmd/main.go"),
	})

	// Main build process that combines all platforms
	all := build.New(Env())
	all.Add(&build.Step{
		ID:      "clean",
		Name:    "Clean up previous builds",
		Command: Cmd("[ -d bin/ ] && rm -r bin/ || true"),
	}).Add(&build.Step{
		ID:      "create-bin-directory",
		Name:    "Create bin directory",
		Command: Cmd("mkdir -p bin/"),
	}).Add(
		linux,
		macos,
		windows,
	).Add(&build.Step{
		ID:      "set-permissions",
		Name:    "Make everything executable",
		Command: Cmd("chmod +x bin/*"),
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
