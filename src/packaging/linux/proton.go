package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Proton struct
type Proton struct {
	Implementation   string            `json:"implementation"`
	Version          string            `json:"version"`
	AppName          string            `json:"appName"`
	AppID            string            `json:"appId"`
	InstallExe       string            `json:"installExe"`
	InstallArguments []string          `json:"installArguments"`
	RunExe           string            `json:"runExe"`
	RunArguments     []string          `json:"runArguments"`
	LaunchExe        string            `json:"launchExe"`
	LaunchArguments  []string          `json:"launchArguments"`
	Source           *packaging.Source `json:"source"`
}

// Return package runtime
func (p *Proton) Runtime() string {
	return "steam"
}

// Return if package is available
func (p *Proton) Available() bool {
	return cli.IsLinux()
}

// Return Steam client path
func (p *Proton) SteamClientPath() string {
	return fs.ExpandPath("$HOME/.steam/steam")
}

// Retrieve the steam runtime path
func (p *Proton) SteamRuntime() string {
	runtime := p.SteamClientPath()
	runtime = filepath.Join(runtime, "ubuntu12_32", "steam-runtime", "run.sh")
	return runtime
}

// Retrieve the proton runtime path
func (p *Proton) ProtonRuntime() string {

	// Native runtime, such as Proton - Experimental
	if p.Implementation == "native" {
		runtime := p.SteamClientPath()
		runtime = filepath.Join(runtime, "steamapps", "common")
		runtime = filepath.Join(runtime, p.Version, "proton")
		return runtime
	}

	// Custom runtime, such as Proton-GE
	runtime := p.SteamClientPath()
	runtime = filepath.Join(runtime, "compatibilitytools.d")
	runtime = filepath.Join(runtime, p.Version, "proton")
	return runtime
}

// Retrieve package main path
func (p *Proton) Path() string {
	path := fs.ExpandPath("$GAMES/Proton")
	path = filepath.Join(path, p.AppName)
	return path
}

// Install package
func (p *Proton) Install() error {

	// Get relevant paths
	mainPath := p.Path()
	dataPath := filepath.Join(mainPath, "pfx", "drive_c")
	installPath := filepath.Join(dataPath, p.InstallExe)
	runPath := filepath.Join(dataPath, p.RunExe)
	launchPath := filepath.Join(dataPath, p.LaunchExe)

	// Download install from source
	if p.Source != nil {
		p.Source.Destination = installPath
		err := p.Source.Download(p)
		if err != nil {
			return err
		}
	}

	// Gather information
	steamClientPath := p.SteamClientPath()
	steamRuntime := p.SteamRuntime()
	protonRuntime := p.ProtonRuntime()
	scriptBase := fmt.Sprintf(``+
		`#!/bin/bash`+"\n"+
		`export STEAM_COMPAT_CLIENT_INSTALL_PATH=$(realpath "%s")`+"\n"+
		`export STEAM_COMPAT_DATA_PATH=$(realpath "%s")`+"\n"+
		`export STEAM_RUNTIME=$(realpath "%s")`+"\n"+
		`export PROTON_RUNTIME=$(realpath "%s")`+"\n"+
		`export APP_PATH=$(realpath "%s")`+"\n",
		steamClientPath,
		mainPath,
		steamRuntime,
		protonRuntime,
		mainPath,
	)

	// Create install executable script
	// Write a script to avoid NiceDeck direct dependency
	installFile := filepath.Join(mainPath, "install.sh")
	installScript := fmt.Sprintf(`%s`+
		`export COMMAND=$(realpath "%s")`+"\n"+
		`"$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$COMMAND" %s "$@"`,
		scriptBase,
		installPath,
		strings.Join(p.InstallArguments, " "),
	)

	err := fs.WriteFile(installFile, installScript)
	if err != nil {
		return err
	}

	err = os.Chmod(installFile, 0775)
	if err != nil {
		return err
	}

	// Create run executable script
	// Will be used to direct launch the main application
	// Write a script to avoid NiceDeck direct dependency
	runFile := filepath.Join(mainPath, "run.sh")
	runScript := fmt.Sprintf(`%s`+
		`export COMMAND=$(realpath "%s")`+"\n"+
		`"$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$COMMAND" %s "$@"`,
		scriptBase,
		runPath,
		strings.Join(p.RunArguments, " "),
	)

	err = fs.WriteFile(runFile, runScript)
	if err != nil {
		return err
	}

	err = os.Chmod(runFile, 0775)
	if err != nil {
		return err
	}

	// Create launch executable script
	// Will be used to direct launch games
	// Write a script to avoid NiceDeck direct dependency
	launchFile := filepath.Join(mainPath, "launch.sh")
	launchScript := fmt.Sprintf(`%s`+
		`export COMMAND=$(realpath "%s")`+"\n"+
		`"$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$COMMAND" %s "$@"`,
		scriptBase,
		launchPath,
		strings.Join(p.LaunchArguments, " "),
	)

	err = fs.WriteFile(launchFile, launchScript)
	if err != nil {
		return err
	}

	err = os.Chmod(launchFile, 0775)
	if err != nil {
		return err
	}

	// Run install script
	err = cli.RunProcess(installFile, []string{})
	if err != nil {
		return err
	}

	return nil
}

// Remove package
func (p *Proton) Remove() error {

	// Remove anything inside package folder
	// Because package is located in its own folder
	mainPath := p.Path()
	err := fs.RemoveDirectory(mainPath)
	if err != nil {
		return err
	}

	// Remove alias file
	err = fs.RemoveFile(p.Alias())
	if err != nil {
		return err
	}

	return nil
}

// Installed verification
func (p *Proton) Installed() (bool, error) {

	mainPath := p.Path()
	dataPath := filepath.Join(mainPath, "pfx", "drive_c")
	executablePath := filepath.Join(dataPath, p.RunExe)

	exist, err := fs.FileExist(executablePath)
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
// In Proton implementations, this return the run script file
func (p *Proton) Executable() string {
	mainPath := p.Path()
	runFile := filepath.Join(mainPath, "run.sh")
	return runFile
}

// Return executable alias file path
func (p *Proton) Alias() string {
	return fs.ExpandPath(fmt.Sprintf(
		"$SHARE/applications/%s.desktop",
		p.AppID,
	))
}

// Run installed package
func (p *Proton) Run(args []string) error {
	return cli.RunProcess(p.Executable(), args)
}

// Fill shortcut additional details
func (p *Proton) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for Proton application
	shortcut.ShortcutPath = p.Alias()
	shortcut.LaunchOptions = strings.Join(p.RunArguments, " ")

	// Write the desktop shortcut
	// err := CreateDesktopShortcut(shortcut)
	// if err != nil {
	// 	return err
	// }

	return nil
}
