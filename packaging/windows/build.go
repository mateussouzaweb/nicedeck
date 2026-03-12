package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/build"
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/version"
)

// Builds the Windows embedded resources to .syso file
func buildResourceSyso(workingDir string, architecture string) func() error {
	return func() error {

		// Gather information
		cmdDir := filepath.Join(workingDir, "cmd", "nicedeck")
		assetsDir := filepath.Join(workingDir, "packaging", "windows")

		appRcName := fmt.Sprintf("nicedeck_windows_%s.rc", architecture)
		appRcDestination := filepath.Join(assetsDir, appRcName)
		sysoName := fmt.Sprintf("nicedeck_windows_%s.syso", architecture)
		sysoDestination := filepath.Join(cmdDir, sysoName)

		exeVersion := fmt.Sprintf("%s.0", version.Get())
		fileVersion := strings.ReplaceAll(exeVersion, ".", ",")
		productVersion := strings.ReplaceAll(exeVersion, ".", ",")

		// Read file content
		templatePath := filepath.Join(assetsDir, "template.rc")
		rcProperties, err := os.ReadFile(templatePath)
		if err != nil {
			return err
		}

		// Replace variables in rc script
		replaces := map[string]string{
			"@{COMPANY_NAME}":        "NiceDeck",
			"@{PRODUCT_NAME}":        "NiceDeck",
			"@{PRODUCT_DESCRIPTION}": "NiceDeck application",
			"@{EXE_NAME}":            "nicedeck.exe",
			"@{EXE_VERSION}":         exeVersion,
			"@{FILE_VERSION}":        fileVersion,
			"@{PRODUCT_VERSION}":     productVersion,
		}
		for key, value := range replaces {
			rcProperties = bytes.ReplaceAll(rcProperties, []byte(key), []byte(value))
		}

		err = fs.WriteFile(appRcDestination, string(rcProperties))
		if err != nil {
			return err
		}

		// Compile Windows resources file
		target := "x86_64-w64-mingw32"
		if architecture == "arm64" {
			target = "arm64-w64-mingw32"
		}

		compileCommand := cli.Command(fmt.Sprintf(
			`llvm-windres-18 --target=%s %s -O coff -o %s`,
			target, appRcDestination, sysoDestination,
		))

		err = cli.Run(compileCommand)
		if err != nil {
			return err
		}

		return nil
	}
}

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
		ID:      "build-windows-resources-amd64",
		Name:    "Building Windows Resources AMD64",
		Context: build.Env("GOARCH=amd64"),
		Command: &build.Command{
			Callback: buildResourceSyso(workingDir, "amd64"),
		},
	}).Add(&build.Step{
		ID:      "build-windows-resources-arm64",
		Name:    "Building Windows Resources ARM64",
		Context: build.Env("GOARCH=arm64"),
		Command: &build.Command{
			Callback: buildResourceSyso(workingDir, "arm64"),
		},
	}).Add(&build.Step{
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
	}).Add(&build.Step{
		ID:      "clean-windows-rc-files",
		Name:    "Cleaning Windows RC Files",
		Context: build.Env(),
		Command: build.Cmd(
			"rm -f %s/packaging/windows/nicedeck_windows_*.rc",
			workingDir,
		),
	}).Add(&build.Step{
		ID:      "clean-windows-syso-files",
		Name:    "Cleaning Windows SYSO Files",
		Context: build.Env(),
		Command: build.Cmd(
			"rm -f %s/cmd/nicedeck/nicedeck_windows_*.syso",
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
