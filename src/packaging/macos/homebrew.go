package macos

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Homebrew struct
type Homebrew struct {
	AppID     string               `json:"appId"`
	AppName   string               `json:"appName"`
	Arguments *packaging.Arguments `json:"arguments"`
}

// Return package runtime
func (h *Homebrew) Runtime() string {
	return "native"
}

// Return if package is available
func (h *Homebrew) Available() bool {
	return cli.IsMacOS()
}

// Install package
func (h *Homebrew) Install() error {
	script := fmt.Sprintf(
		`brew install --cask %s %s`,
		h.AppID,
		strings.Join(h.Arguments.Install, " "),
	)

	command := cli.Command(script)
	return cli.Run(command)
}

// Remove package
func (h *Homebrew) Remove() error {
	script := fmt.Sprintf(
		`brew uninstall --cask %s %s`,
		h.AppID,
		strings.Join(h.Arguments.Remove, " "),
	)

	command := cli.Command(script)
	return cli.Run(command)
}

// Installed verification
func (h *Homebrew) Installed() (bool, error) {

	// MacOs.app files are considered a directory
	if filepath.Ext(h.Executable()) == ".app" {
		exist, err := fs.DirectoryExist(h.Executable())
		if err != nil {
			return false, err
		} else if exist {
			return true, nil
		}
	}

	exist, err := fs.FileExist(h.Executable())
	if err != nil {
		return false, err
	} else if exist {
		return true, nil
	}

	return false, nil
}

// Return executable file path
func (h *Homebrew) Executable() string {
	return fs.NormalizePath(fmt.Sprintf(
		`/Applications/%s`,
		h.AppName,
	))
}

// Return executable alias file path
func (h *Homebrew) Alias() string {
	return h.Executable()
}

// Return executable arguments
func (h *Homebrew) Args() []string {
	return h.Arguments.Shortcut
}
