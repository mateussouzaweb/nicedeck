package steam

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

type Config struct {
	ArtworksPath            string                `json:"artworksPath"`
	ControllerTemplatesPath string                `json:"controllerTemplatesPath"`
	DebugFile               string                `json:"debugFile"`
	IsFlatpak               bool                  `json:"isFlatpak"`
	ShortcutsFile           string                `json:"shortcutsFile"`
	Shortcuts               []*shortcuts.Shortcut `json:"shortcuts"`
	SteamPath               string                `json:"steamPath"`
}

var _config *Config

// Check if Steam installation was done via flatpak
func SteamIsFlatpak() (bool, error) {

	// App can be installed on system or user
	systemFile := os.ExpandEnv("$HOME/.local/share/flatpak/exports/bin/com.valvesoftware.Steam")
	userFile := "/var/lib/flatpak/exports/bin/com.valvesoftware.Steam"

	// Checks what possible file exist
	for _, file := range []string{systemFile, userFile} {
		exist, err := fs.FileExist(file)
		if err != nil {
			return false, err
		} else if exist {
			return true, nil
		}
	}

	return false, nil
}

// Retrieve the full Steam path with given additional path
func GetPath(path string) (string, error) {

	// Fill possible locations
	paths := []string{
		os.ExpandEnv("$HOME/.steam/steam"),
		os.ExpandEnv("$HOME/.local/share/Steam"),
		os.ExpandEnv("$HOME/.var/app/com.valvesoftware.Steam/.steam/steam"),
		os.ExpandEnv("$HOME/snap/steam/common/.local/share/Steam"),
	}

	// Checks what directory path is available
	usePath := ""
	for _, possiblePath := range paths {
		exist, err := fs.DirectoryExist(possiblePath)
		if err != nil {
			return "", err
		} else if exist {
			usePath = filepath.Join(possiblePath, path)
			break
		}
	}

	// Return error if not detected
	if usePath == "" {
		return "", fmt.Errorf("could not detect the Steam installation path")
	}

	// Try to detect the path
	found, err := filepath.Glob(usePath)
	if err != nil {
		return "", err
	}

	if len(found) == 0 {
		return "", fmt.Errorf("could not found the Steam installation path: %s", usePath)
	}

	// Will return only the first result
	return found[0], nil
}

// Load Steam data to runtime config
func Load() error {

	// Retrieve Steam base path
	steamPath, err := GetPath("")
	if err != nil {
		return err
	}

	// Make sure Steam on flatpak has the necessary permission
	// We need this to run flatpak-spawn command to comunicate with others flatpak apps
	isFlatpak, err := SteamIsFlatpak()
	if err != nil {
		return err
	} else if isFlatpak {
		override := "flatpak override --user --talk-name=org.freedesktop.Flatpak com.valvesoftware.Steam"
		err = cli.Command(override).Run()
		if err != nil {
			return err
		}
	}

	// Retrieve user config path
	userConfig, err := GetPath("userdata/*/config")
	if err != nil {
		return err
	}

	// Retrieve controller templates path
	controllerTemplatesPath, err := GetPath("controller_base/templates")
	if err != nil {
		return err
	}

	// Set runtime configs
	_config = &Config{}
	_config.IsFlatpak = isFlatpak
	_config.SteamPath = steamPath
	_config.ControllerTemplatesPath = controllerTemplatesPath
	_config.ArtworksPath = userConfig + "/grid"
	_config.DebugFile = userConfig + "/niceconfig.json"
	_config.ShortcutsFile = userConfig + "/shortcuts.vdf"

	// Load shortcuts
	shortcutsList, err := shortcuts.LoadFromFile(_config.ShortcutsFile)
	if err != nil {
		return err
	}

	_config.Shortcuts = shortcutsList

	return nil
}

// Save runtime state to files
func Save() error {

	var err error

	// Write state to debug file with JSON copy
	// Save JSON copy for debugging
	jsonContent, err := json.MarshalIndent(_config, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON content to file
	err = os.WriteFile(_config.DebugFile, jsonContent, 0666)
	if err != nil {
		return err
	}

	// Sort list of shortcuts (again)
	_config.Shortcuts, err = shortcuts.SortShortcuts(_config.Shortcuts)
	if err != nil {
		return err
	}

	// Save shortcuts
	err = shortcuts.SaveToFile(_config.Shortcuts, _config.ShortcutsFile)
	if err != nil {
		return err
	}

	// Write controller templates
	err = controller.WriteTemplates(_config.ControllerTemplatesPath)
	if err != nil {
		return err
	}

	return nil
}

// Retrieve runtime shortcuts
func GetShortcuts() []*shortcuts.Shortcut {
	return _config.Shortcuts
}

// Add program to the Steam shortcuts library
func AddToShortcuts(shortcut *shortcuts.Shortcut) error {

	var err error

	// Check if Steam was installed via flatpak
	// If yes, then we need to append the flatpak-spawn wrapper
	// This is necessary to have access to host commands
	if _config.IsFlatpak {
		shortcut.Exe = "/usr/bin/flatpak-spawn --host " + shortcut.Exe
	}

	// Determine appId and artworks path
	shortcut.AppID = shortcuts.GenerateShortcutID(shortcut.Exe, shortcut.AppName)
	artworksPath := _config.ArtworksPath

	// Icon: ${APPID}_icon.ico || ${APPID}_icon.png
	if strings.HasSuffix(shortcut.IconURL, ".png") {
		shortcut.Icon = fmt.Sprintf("%s/%v_icon.png", artworksPath, shortcut.AppID)
	} else {
		shortcut.Icon = fmt.Sprintf("%s/%v_icon.ico", artworksPath, shortcut.AppID)
	}

	// Logo: ${APPID}_logo.png || ${APPID}_logo.jpg
	if strings.HasSuffix(shortcut.LogoURL, ".png") {
		shortcut.Logo = fmt.Sprintf("%s/%v_logo.png", artworksPath, shortcut.AppID)
	} else {
		shortcut.Logo = fmt.Sprintf("%s/%v_logo.jpg", artworksPath, shortcut.AppID)
	}

	// Cover: ${APPID}p.png || ${APPID}p.jpg
	if strings.HasSuffix(shortcut.CoverURL, ".png") {
		shortcut.Cover = fmt.Sprintf("%s/%vp.png", artworksPath, shortcut.AppID)
	} else {
		shortcut.Cover = fmt.Sprintf("%s/%vp.jpg", artworksPath, shortcut.AppID)
	}

	// Banner: ${APPID}.png || ${APPID}.jpg
	if strings.HasSuffix(shortcut.BannerURL, ".png") {
		shortcut.Banner = fmt.Sprintf("%s/%v.png", artworksPath, shortcut.AppID)
	} else {
		shortcut.Banner = fmt.Sprintf("%s/%v.jpg", artworksPath, shortcut.AppID)
	}

	// Hero: ${APPID}_hero.png || ${APPID}_hero.jpg
	if strings.HasSuffix(shortcut.HeroURL, ".png") {
		shortcut.Hero = fmt.Sprintf("%s/%v_hero.png", artworksPath, shortcut.AppID)
	} else {
		shortcut.Hero = fmt.Sprintf("%s/%v_hero.jpg", artworksPath, shortcut.AppID)
	}

	// Create list of images to download
	images := map[string]string{}
	images[shortcut.IconURL] = shortcut.Icon
	images[shortcut.LogoURL] = shortcut.Logo
	images[shortcut.CoverURL] = shortcut.Cover
	images[shortcut.BannerURL] = shortcut.Banner
	images[shortcut.HeroURL] = shortcut.Hero

	// Download available shortcut images
	for url, destinationFile := range images {
		if url == "" || destinationFile == "" {
			continue
		}
		err = scraper.DownloadFile(url, destinationFile)
		if err != nil {
			return err
		}
	}

	// Add to shortcuts list
	_config.Shortcuts, err = shortcuts.AddShortcut(_config.Shortcuts, shortcut)
	if err != nil {
		return err
	}

	return nil
}

// Remove program from the Steam shortcuts library
func RemoveFromShortcuts(shortcut *shortcuts.Shortcut) error {

	var err error

	_config.Shortcuts, err = shortcuts.RemoveShortcut(_config.Shortcuts, shortcut)
	if err != nil {
		return err
	}

	return nil
}
