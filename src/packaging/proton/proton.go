package proton

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

//go:embed resources/*
var resourcesContent embed.FS

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

// Retrieve the steam runtime path
func (p *Proton) SteamRuntime() (string, error) {

	runtime, err := steam.GetBasePath()
	if err != nil {
		return "", err
	}

	runtime = filepath.Join(runtime, "ubuntu12_32", "steam-runtime", "run.sh")
	return runtime, nil
}

// Retrieve the proton runtime path
func (p *Proton) ProtonRuntime() (string, error) {

	runtime, err := steam.GetBasePath()
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
	steamPackage := steam.GetPackage()
	if !steamPackage.Available() {
		return fmt.Errorf("requirement error, Steam must be installed")
	}

	// Gather information
	dataPath := p.ProtonPath()
	drivePath := p.DrivePath()
	steamInstallType := steamPackage.Runtime()
	steamClientPath, err := steam.GetBasePath()
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
	// Makes verification to check at the installer file
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

	// Create run executable script to avoid NiceDeck direct dependency
	// Will be used to launch applications
	runFile := filepath.Join(dataPath, "run.sh")
	runScript, err := resourcesContent.ReadFile("resources/run.sh")
	if err != nil {
		return err
	}

	// When Steam is installed with flatpak, make use of flatpak sandbox paths
	if steamInstallType == "flatpak" {
		appID := steamPackage.(*linux.Flatpak).AppID
		appFolder := fmt.Sprintf("/.var/app/%s", appID)
		steamClientPath = strings.Replace(steamClientPath, appFolder, "", 1)
		steamRuntime = strings.Replace(steamRuntime, appFolder, "", 1)
		protonRuntime = strings.Replace(protonRuntime, appFolder, "", 1)
	}

	replaces := map[string]string{
		"${DATA_PATH}":         dataPath,
		"${DRIVE_PATH}":        drivePath,
		"${INSTALL_TYPE}":      steamInstallType,
		"${PROTON_RUNTIME}":    protonRuntime,
		"${STEAM_CLIENT_PATH}": steamClientPath,
		"${STEAM_RUNTIME}":     steamRuntime,
	}
	for key, value := range replaces {
		runScript = bytes.ReplaceAll(runScript, []byte(key), []byte(value))
	}

	err = fs.WriteFile(runFile, string(runScript))
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
	err := linux.CreateDesktopShortcut(shortcut)
	if err != nil {
		return err
	}

	return nil
}
