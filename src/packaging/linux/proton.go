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
	AppID       string            `json:"appId"`
	AppName     string            `json:"appName"`
	Installer   string            `json:"installer"`
	Uninstaller string            `json:"uninstaller"`
	Launcher    string            `json:"launcher"`
	Arguments   []string          `json:"arguments"`
	Source      *packaging.Source `json:"source"`
}

// Return package runtime
func (p *Proton) Runtime() string {
	return "proton"
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

	// Prefix is not customizable for now
	implementation := "native"
	version := "Proton - Experimental"

	// Native runtime, such as Proton - Experimental
	if implementation == "native" {
		runtime := p.SteamClientPath()
		runtime = filepath.Join(runtime, "steamapps", "common")
		runtime = filepath.Join(runtime, version, "proton")
		return runtime
	}

	// Custom runtime, such as Proton-GE
	runtime := p.SteamClientPath()
	runtime = filepath.Join(runtime, "compatibilitytools.d")
	runtime = filepath.Join(runtime, version, "proton")
	return runtime
}

// Retrieve proton data path
func (p *Proton) ProtonPath() string {
	return fs.ExpandPath("$GAMES/Proton")
}

// Retrieve proton data path
func (p *Proton) DrivePath() string {
	return filepath.Join(p.ProtonPath(), "pfx", "drive_c")
}

// Retrieve real path for given path
func (p *Proton) RealPath(path string) string {
	return strings.Replace(path, "C:", p.DrivePath(), 1)
}

// Retrieve virtual path for given path
func (p *Proton) VirtualPath(path string) string {
	return strings.Replace(path, p.DrivePath(), "C:", 1)
}

// Install package
func (p *Proton) Install() error {

	// Get relevant paths
	mainPath := p.ProtonPath()
	drivePath := p.DrivePath()

	// Download install from source
	if p.Source != nil {
		p.Source.Destination = p.RealPath(p.Installer)
		err := p.Source.Download(p)
		if err != nil {
			return err
		}
	}

	// Gather information
	steamClientPath := p.SteamClientPath()
	steamRuntime := p.SteamRuntime()
	protonRuntime := p.ProtonRuntime()

	// Create run executable script
	// Will be used to launch applications
	// Write a script to avoid NiceDeck direct dependency
	runFile := filepath.Join(mainPath, "run.sh")
	runScript := fmt.Sprintf(strings.Join([]string{
		`#!/bin/bash`,
		``,
		`# Variables for execution`,
		`export STEAM_COMPAT_CLIENT_INSTALL_PATH=$(realpath "%s")`,
		`export STEAM_COMPAT_DATA_PATH=$(realpath "%s")`,
		`export STEAM_RUNTIME=$(realpath "%s")`,
		`export PROTON_RUNTIME=$(realpath "%s")`,
		`export DRIVE_PATH=$(realpath "%s")`,
		`export LOG_PATH=$(realpath "%s/run.log")`,
		``,
		`# Replace driver path`,
		`set -- "${1/C:/$DRIVE_PATH}" "${@:2}"`,
		``,
		`# Run command`,
		`echo "$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$@" >> "$LOG_PATH"`,
		`"$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$@" >> "$LOG_PATH" 2>&1`}, "\n"),
		steamClientPath,
		mainPath,
		steamRuntime,
		protonRuntime,
		drivePath,
		mainPath,
	)

	err := fs.WriteFile(runFile, runScript)
	if err != nil {
		return err
	}

	err = os.Chmod(runFile, 0775)
	if err != nil {
		return err
	}

	cli.Debug("Running install for %s\n", p.AppName)

	// Run install script
	arguments := []string{cli.Quote(p.Installer)}
	arguments = append(arguments, p.Arguments...)

	err = cli.RunProcess(runFile, arguments)
	if err != nil {
		return err
	}

	return nil
}

// Remove package
func (p *Proton) Remove() error {

	cli.Debug("Running uninstall for %s\n", p.AppName)
	runFile := p.Executable()

	// Remove package by perform the uninstall command
	arguments := []string{cli.Quote(p.Uninstaller)}
	err := cli.RunProcess(runFile, arguments)
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

	launcher := p.RealPath(p.Launcher)
	exist, err := fs.FileExist(launcher)
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
	mainPath := p.ProtonPath()
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
	arguments := []string{cli.Quote(p.Launcher)}
	arguments = append(arguments, p.Arguments...)

	shortcut.ShortcutPath = p.Alias()
	shortcut.LaunchOptions = strings.Join(arguments, " ")

	// Write the desktop shortcut
	err := CreateDesktopShortcut(shortcut)
	if err != nil {
		return err
	}

	return nil
}
