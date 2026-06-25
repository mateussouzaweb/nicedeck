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
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

//go:embed resources/*
var resourcesContent embed.FS

// Proton struct
type Proton struct {
	AppID       string               `json:"appId"`
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

// Retrieve the proton path for proton runtime
func (p *Proton) ProtonPath() (string, error) {

	path, err := steam.GetBasePath()
	if err != nil {
		return "", err
	}

	// Prefix is not customizable for now
	implementation := "native"
	version := "Proton - Experimental"

	// Native runtime, such as Proton - Experimental
	if implementation == "native" {
		path = filepath.Join(path, "steamapps", "common", version)
		return path, nil
	}

	// Custom runtime, such as Proton-GE
	path = filepath.Join(path, "compatibilitytools.d", version)
	return path, nil
}

// Retrieve the proton runtime path
func (p *Proton) ProtonRuntime() (string, error) {

	path, err := p.ProtonPath()
	if err != nil {
		return "", err
	}

	runtime := filepath.Join(path, "proton")
	return runtime, nil
}

// Retrieve the wine binary path
func (p *Proton) WineBinary() (string, error) {

	path, err := p.ProtonPath()
	if err != nil {
		return "", err
	}

	binary := filepath.Join(path, "files", "bin", "wine")
	return binary, nil
}

// Retrieve proton data path
func (p *Proton) DataPath() string {
	return fs.ExpandPath("$GAMES/Proton")
}

// Retrieve proton wine path
func (p *Proton) WinePath() string {
	return filepath.Join(p.DataPath(), "pfx")
}

// Retrieve proton drive path
func (p *Proton) DrivePath() string {
	return filepath.Join(p.DataPath(), "pfx", "drive_c")
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
		return fmt.Errorf("requirement error, Steam is not available")
	}

	installed, err := steamPackage.Installed()
	if err != nil {
		return err
	} else if !installed {
		return fmt.Errorf("requirement error, Steam must be installed")
	}

	// Gather information
	dataPath := p.DataPath()
	winePath := p.WinePath()
	drivePath := p.DrivePath()

	steamPath, err := steam.GetBasePath()
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

	wineBinary, err := p.WineBinary()
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

	// Create run executable script to avoid NiceDeck direct dependency
	// Will be used to launch applications
	runFile := filepath.Join(dataPath, "run.sh")
	runScript, err := resourcesContent.ReadFile("resources/run.sh")
	if err != nil {
		return err
	}

	// Retrieve runtime for Steam in Flatpak mode
	steamInstallType := steamPackage.Runtime()
	steamFlatpakID := "com.valvesoftware.Steam"

	// Replace variables in run script
	replaces := map[string]string{
		"@{DATA_PATH}":      dataPath,
		"@{DRIVE_PATH}":     drivePath,
		"@{WINE_PATH}":      winePath,
		"@{FLATPAK_ID}":     steamFlatpakID,
		"@{INSTALL_TYPE}":   steamInstallType,
		"@{PROTON_RUNTIME}": protonRuntime,
		"@{STEAM_PATH}":     steamPath,
		"@{STEAM_RUNTIME}":  steamRuntime,
		"@{WINE_BINARY}":    wineBinary,
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

	// Copy extra executable for custom modifications
	// File is created only if not exist yet
	err = fs.CopyEmbedded(
		resourcesContent,
		"resources/extra.sh",
		filepath.Join(dataPath, "extra.sh"),
		false, // Do not expand environment variables
		false, // Do not overwrite existing
	)
	if err != nil {
		return err
	}

	// Download program from source
	// When no need for installer, download from source
	// When installer is needed, download from source and makes verification to check at the installer file
	if p.Source != nil && p.Installer != "" {
		defer func(launcher string) {
			p.Launcher = launcher
			p.Source.Destination = ""
		}(p.Launcher)

		p.Launcher = p.Installer
		p.Source.Destination = p.RealPath(p.Installer)
		err := p.Source.Download(p)
		if err != nil {
			return err
		}

	} else if p.Source != nil {
		p.Source.Destination = p.RealPath(p.Launcher)
		err := p.Source.Download(p)
		if err != nil {
			return err
		}
	}

	// When installer is needed, run installer
	if p.Installer != "" {
		cli.Debug("Running install for %s\n", p.AppID)

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
	}

	return nil
}

// Remove package
func (p *Proton) Remove() error {

	// Remove package by perform the uninstall command
	if p.Uninstaller != "" {
		cli.Debug("Running uninstall for %s\n", p.AppID)

		runFile := p.Executable()
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
	}

	// Remove launcher parent folder
	// Because package is located in its own folder
	directory := filepath.Dir(fs.ExpandPath(p.Launcher))
	err := fs.RemoveDirectory(directory)
	if err != nil {
		return err
	}

	// Remove installer file
	// Because installer is placed in another location
	if p.Installer != "" {
		err := fs.RemoveFile(fs.ExpandPath(p.Installer))
		if err != nil {
			return err
		}
	}

	return nil
}

// Installed verification
func (p *Proton) Installed() (bool, error) {

	launcher := fs.ExpandPath(p.RealPath(p.Launcher))
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
	dataPath := p.DataPath()
	runFile := filepath.Join(dataPath, "run.sh")
	return runFile
}

// Return executable alias file path
func (p *Proton) Alias() string {
	return ""
}

// Return executable arguments
func (p *Proton) Args() []string {
	arguments := []string{cli.Quote(p.Launcher)}
	arguments = append(arguments, p.Arguments.Shortcut...)
	return arguments
}
