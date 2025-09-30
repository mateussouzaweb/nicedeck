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
	AppID       string               `json:"appId"`
	AppName     string               `json:"appName"`
	Installer   string               `json:"installer"`
	Uninstaller string               `json:"uninstaller"`
	Launcher    string               `json:"launcher"`
	Arguments   *packaging.Arguments `json:"arguments"`
	Source      *packaging.Source    `json:"source"`
}

// Return package runtime
func (p *Proton) Runtime() string {
	return "proton"
}

// Return if package is available
func (p *Proton) Available() bool {
	return cli.IsLinux()
}

// Steam installed verification
func (p *Proton) SteamInstalled() bool {

	steamPackage := packaging.Installed(&Flatpak{
		Namespace: "system",
		AppID:     "com.valvesoftware.Steam",
		Overrides: []string{"--talk-name=org.freedesktop.Flatpak"},
		Arguments: packaging.NoArguments(),
	}, &Flatpak{
		Namespace: "user",
		AppID:     "com.valvesoftware.Steam",
		Overrides: []string{"--talk-name=org.freedesktop.Flatpak"},
		Arguments: packaging.NoArguments(),
	}, &Snap{
		AppID:     "steam",
		AppBin:    "steam",
		Arguments: packaging.NoArguments(),
	}, &Binary{
		AppID:     "steam",
		AppBin:    "/usr/bin/steam",
		Arguments: packaging.NoArguments(),
	})

	return steamPackage.Available()
}

// Return Steam client data path
func (p *Proton) SteamPath() (string, error) {

	// Fill possible locations
	paths := []string{
		fs.ExpandPath("$VAR/com.valvesoftware.Steam/.steam/steam"),
		fs.ExpandPath("$HOME/snap/steam/common/.local/share/Steam"),
		fs.ExpandPath("$HOME/.steam/steam"),
		fs.ExpandPath("$SHARE/Steam"),
		fs.ExpandPath("$CONFIG/Steam"),
		fs.ExpandPath("$PROGRAMS_X86/Steam"),
	}

	// Checks what directory path is available
	for _, possiblePath := range paths {
		exist, err := fs.DirectoryExist(possiblePath)
		if err != nil {
			return "", err
		} else if exist {
			return possiblePath, nil
		}
	}

	return "", nil
}

// Retrieve the steam runtime path
func (p *Proton) SteamRuntime() (string, error) {

	runtime, err := p.SteamPath()
	if err != nil {
		return "", err
	}

	runtime = filepath.Join(runtime, "ubuntu12_32", "steam-runtime", "run.sh")
	return runtime, nil
}

// Retrieve the proton runtime path
func (p *Proton) ProtonRuntime() (string, error) {

	runtime, err := p.SteamPath()
	if err != nil {
		return "", err
	}

	// Prefix is not customizable for now
	implementation := "native"
	version := "Proton - Experimental"

	// Native runtime, such as Proton - Experimental
	if implementation == "native" {
		runtime = filepath.Join(runtime, "steamapps", "common")
		runtime = filepath.Join(runtime, version, "proton")
		return runtime, nil
	}

	// Custom runtime, such as Proton-GE
	runtime = filepath.Join(runtime, "compatibilitytools.d")
	runtime = filepath.Join(runtime, version, "proton")
	return runtime, nil
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

	// Make sure Steam is installed
	if !p.SteamInstalled() {
		return fmt.Errorf("requirement error, Steam must be installed")
	}

	// Gather information
	mainPath := p.ProtonPath()
	drivePath := p.DrivePath()
	steamClientPath, err := p.SteamPath()
	if err != nil {
		return err
	}

	steamRuntime, err := p.SteamRuntime()
	if err != nil {
		return err
	}

	protonRuntime, err := p.ProtonRuntime()
	if err != nil {
		return err
	}

	// Make sure that Proton is installed
	// When missing, request Proton installation from Steam URL handler
	protonInstalled, err := fs.FileExist(protonRuntime)
	if err != nil {
		return err
	} else if !protonInstalled {
		defer cli.Open("steam://install/1493710") // Proton Experimental
		return fmt.Errorf("proton install missing, please install proton first")
	}

	// Download install from source
	// Also makes verification to check at the installer file
	if p.Source != nil {
		originalLauncher := p.Launcher
		p.Launcher = p.Installer
		defer func() {
			p.Launcher = originalLauncher
		}()

		p.Source.Destination = p.RealPath(p.Installer)
		err := p.Source.Download(p)
		if err != nil {
			return err
		}
	}

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
		``,
		`# Replace C: with driver path`,
		`set -- "${1/C:/$DRIVE_PATH}" "${@:2}"`,
		``,
		`# Go to target executable path`,
		`# This step is required for some games`,
		`cd "$(dirname "$1")"`,
		``,
		`# Run command`,
		`exec "$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$@" 2>&1`}, "\n"),
		steamClientPath,
		mainPath,
		steamRuntime,
		protonRuntime,
		drivePath,
	)

	err = fs.WriteFile(runFile, runScript)
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
	arguments = append(arguments, p.Arguments.Install...)
	directory := filepath.Dir(p.RealPath(p.Installer))

	context := &cli.Context{
		WorkingDirectory: directory,
		Executable:       runFile,
		Arguments:        arguments,
		Environment:      []string{},
	}

	err = context.Run()
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
	arguments = append(arguments, p.Arguments.Remove...)
	directory := filepath.Dir(p.RealPath(p.Uninstaller))

	context := &cli.Context{
		WorkingDirectory: directory,
		Executable:       runFile,
		Arguments:        arguments,
		Environment:      []string{},
	}

	err := context.Run()
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

// Fill shortcut additional details
func (p *Proton) OnShortcut(shortcut *shortcuts.Shortcut) error {

	// Fill shortcut information for Proton application
	arguments := []string{cli.Quote(p.Launcher)}
	arguments = append(arguments, p.Arguments.Shortcut...)

	shortcut.ShortcutPath = p.Alias()
	shortcut.LaunchOptions = strings.Join(arguments, " ")

	// Write the desktop shortcut
	err := CreateDesktopShortcut(shortcut)
	if err != nil {
		return err
	}

	return nil
}
