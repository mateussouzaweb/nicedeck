package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/desktop/macos"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Builds the MacOS .app bundle from already compiled binaries
func build() error {

	// Get project folder
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Define paths for easier access
	binDir := filepath.Join(workingDir, "bin")
	compiledAmd64Path := filepath.Join(binDir, "nicedeck-macos-amd64")
	compiledArm64Path := filepath.Join(binDir, "nicedeck-macos-arm64")
	targetZipPath := filepath.Join(binDir, "nicedeck-macos-universal.zip")

	assetsDir := filepath.Join(workingDir, "packaging", "macos")
	iconSourcePath := filepath.Join(assetsDir, "icon.icns")
	launcherSourcePath := filepath.Join(assetsDir, "launcher.sh")

	appBundleDir := filepath.Join(workingDir, "bin", "NiceDeck.app")
	launcherPath := filepath.Join(appBundleDir, "Contents", "MacOS", "nicedeck")
	launcherAmd64Path := filepath.Join(appBundleDir, "Contents", "MacOS", "nicedeck-amd64")
	launcherArm64Path := filepath.Join(appBundleDir, "Contents", "MacOS", "nicedeck-arm64")

	// Create temporary .app bundle structure
	// Bundle will contains both amd64 and arm64 binaries to simplify the build process
	// Avoid Apple's Universal Binary format because it can be compiled only on MacOS systems
	bundle := &macos.Bundle{
		AppName:          "NiceDeck",
		BundleID:         "com.mateussouzaweb.nicedeck",
		Launcher:         filepath.Base(launcherPath),
		IconPath:         iconSourcePath,
		WorkingDirectory: "",
		Executable:       "",
		Arguments:        []string{},
		Environment:      []string{},
	}

	err = macos.WriteBundle(appBundleDir, bundle)
	if err != nil {
		return err
	}

	// Copy compiled binaries into app bundle
	err = fs.CopyFile(compiledAmd64Path, launcherAmd64Path, true)
	if err != nil {
		return err
	}
	err = fs.CopyFile(compiledArm64Path, launcherArm64Path, true)
	if err != nil {
		return err
	}

	// Write launcher script that selects the correct binary at runtime
	err = fs.CopyFile(launcherSourcePath, launcherPath, true)
	if err != nil {
		return err
	}

	// Make sure launcher scripts are executable
	err = os.Chmod(launcherAmd64Path, 0755)
	if err != nil {
		return err
	}
	err = os.Chmod(launcherArm64Path, 0755)
	if err != nil {
		return err
	}
	err = os.Chmod(launcherPath, 0755)
	if err != nil {
		return err
	}

	// Compress the .app bundle into a .zip file
	compressCommand := cli.Command(fmt.Sprintf(
		`cd %s && zip -r %s %s > /dev/null; rm -rf %s`,
		binDir,
		filepath.Base(targetZipPath),
		filepath.Base(appBundleDir),
		filepath.Base(appBundleDir),
	))

	err = cli.Run(compressCommand)
	if err != nil {
		return err
	}

	return nil
}

// MacOS build process
func main() {

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Run build process
	err := build()
	if err != nil {
		exitCode = 1
		cli.Printf(cli.ColorFatal, "MacOS packaging failed: %s\n", err.Error())
	}

}
