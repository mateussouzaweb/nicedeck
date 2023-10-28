package library

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Config struct
type Config struct {
	IsFlatpak               bool                  `json:"isFlatpak"`
	SteamPath               string                `json:"steamPath"`
	UserConfigPath          string                `json:"userConfigPath"`
	ControllerTemplatesPath string                `json:"controllerTemplatesPath"`
	Shortcuts               []*shortcuts.Shortcut `json:"shortcuts"`
}

var _config *Config

// Load data to runtime config
func Load() error {

	// Retrieve Steam base path
	steamPath, err := steam.GetPath("")
	if err != nil {
		return err
	}

	// Make sure Steam on flatpak has the necessary permission
	// We need this to run flatpak-spawn command to comunicate with others flatpak apps
	isFlatpak, err := steam.IsFlatpak()
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
	userConfigPath, err := steam.GetPath("userdata/*/config")
	if err != nil {
		return err
	}

	// Retrieve controller templates path
	controllerTemplatesPath, err := steam.GetPath("controller_base/templates")
	if err != nil {
		return err
	}

	// Set runtime configs
	_config = &Config{}
	_config.IsFlatpak = isFlatpak
	_config.SteamPath = steamPath
	_config.UserConfigPath = userConfigPath
	_config.ControllerTemplatesPath = controllerTemplatesPath

	// Load config file if exist
	exist, err := fs.FileExist(_config.UserConfigPath + "/niceconfig.json")
	if err != nil {
		return err
	} else if exist {

		// Read config file content
		content, err := os.ReadFile(_config.UserConfigPath + "/niceconfig.json")
		if err != nil {
			return err
		}

		// Fill config information from file content when available
		// This file contains the extended state for shortcuts
		err = json.Unmarshal(content, _config)
		if err != nil {
			return err
		}

	}

	// Load shortcuts from file
	shortcutsList, err := shortcuts.LoadFromFile(_config.UserConfigPath + "/shortcuts.vdf")
	if err != nil {
		return err
	}

	// When already exist a list of shortcuts from Steam, we should merge data
	// The preferencial content always are from the .vdf file
	// This file can possible be updated by other services or Steam UI
	if len(shortcutsList) > 0 {
		_config.Shortcuts = shortcuts.MergeShortcuts(_config.Shortcuts, shortcutsList)
	}

	return nil
}

// Save runtime state to files
func Save() error {

	var err error

	// Sort list of shortcuts (again)
	_config.Shortcuts, err = shortcuts.SortShortcuts(_config.Shortcuts)
	if err != nil {
		return err
	}

	// Convert config state to JSON representation
	jsonContent, err := json.MarshalIndent(_config, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON content to config file
	err = os.WriteFile(_config.UserConfigPath+"/niceconfig.json", jsonContent, 0666)
	if err != nil {
		return err
	}

	// Save shortcuts file
	err = shortcuts.SaveToFile(_config.Shortcuts, _config.UserConfigPath+"/shortcuts.vdf")
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

// Add program to the shortcuts list
func AddToShortcuts(shortcut *shortcuts.Shortcut) error {

	var err error

	// Check if Steam was installed via flatpak
	// If yes, then we need to append the flatpak-spawn wrapper
	// This is necessary to have access to host commands
	if _config.IsFlatpak {
		shortcut.Exe = "/usr/bin/flatpak-spawn --host " + shortcut.Exe
	}

	// Determine appID and artworks path
	shortcut.AppID = shortcuts.GenerateShortcutID(shortcut.Exe, shortcut.AppName)
	artworksPath := _config.UserConfigPath + "/grid"

	// Logo: ${APPID}_logo.png
	shortcut.Logo = fmt.Sprintf("%s/%v_logo.png", artworksPath, shortcut.AppID)

	// Icon: ${APPID}_icon.ico || ${APPID}_icon.png
	if strings.HasSuffix(shortcut.IconURL, ".png") {
		shortcut.Icon = fmt.Sprintf("%s/%v_icon.png", artworksPath, shortcut.AppID)
	} else {
		shortcut.Icon = fmt.Sprintf("%s/%v_icon.ico", artworksPath, shortcut.AppID)
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

// Remove program from the shortcuts list
func RemoveFromShortcuts(shortcut *shortcuts.Shortcut) error {

	var err error

	_config.Shortcuts, err = shortcuts.RemoveShortcut(_config.Shortcuts, shortcut)
	if err != nil {
		return err
	}

	return nil
}
