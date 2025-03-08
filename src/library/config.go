package library

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Config struct
type Config struct {
	SteamRuntime  string                `json:"steamRuntime"`
	SteamPath     string                `json:"steamPath"`
	ConfigPath    string                `json:"configPath"`
	ArtworksPath  string                `json:"artworksPath"`
	StateFile     string                `json:"stateFile"`
	ShortcutsFile string                `json:"shortcutsFile"`
	Shortcuts     []*shortcuts.Shortcut `json:"shortcuts"`
}

var _config Config

// Load data to runtime config
func Load() error {

	var err error

	// Check how Steam is running
	steamRuntime, err := steam.GetRuntime()
	if err != nil {
		return fmt.Errorf("could not determine Steam runtime: %s", err)
	}

	// Retrieve Steam base path
	steamPath, err := steam.GetPath()
	if err != nil {
		return fmt.Errorf("could not detect Steam installation: %s", err)
	}

	// Retrieve Steam user config path
	configPath, err := steam.GetConfigPath()
	if err != nil {
		return fmt.Errorf("could not detect Steam user config path: %s", err)
	}

	// Set default runtime configs
	_config = Config{
		SteamRuntime:  steamRuntime,
		SteamPath:     steamPath,
		ConfigPath:    configPath,
		ArtworksPath:  filepath.Join(configPath, "grid"),
		StateFile:     filepath.Join(configPath, "niceconfig.json"),
		ShortcutsFile: filepath.Join(configPath, "shortcuts.vdf"),
	}

	// Show message based on Steam detection
	if _config.SteamPath == "" {
		cli.Printf(cli.ColorWarn, "Steam installation was not detected.\n")
		cli.Printf(cli.ColorWarn, "Please make sure to install and login into Steam first.\n")
	}

	// Load config file if exist
	exist, err := fs.FileExist(_config.StateFile)
	if err != nil {
		return err
	} else if exist {

		// Read config file content
		content, err := os.ReadFile(_config.StateFile)
		if err != nil {
			return err
		}

		// Fill config information from file content when available
		// This file contains the extended state for shortcuts
		err = json.Unmarshal(content, &_config)
		if err != nil {
			return err
		}

	}

	// Load shortcuts from file
	shortcutsList, err := shortcuts.LoadFromFile(_config.ShortcutsFile)
	if err != nil {
		return err
	}

	// When already exist a list of shortcuts from file, we should merge data
	// The preferencial content always are from the .vdf file
	// This file can possible be updated by other services or Steam UI
	if len(shortcutsList) > 0 {
		_config.Shortcuts = shortcuts.MergeShortcuts(
			shortcutsList,
			_config.Shortcuts,
			func(target *shortcuts.Shortcut, source *shortcuts.Shortcut) {

				// Merge relevant data
				target.IconURL = source.IconURL
				target.Logo = source.Logo
				target.LogoURL = source.LogoURL
				target.Cover = source.Cover
				target.CoverURL = source.CoverURL
				target.Banner = source.Banner
				target.BannerURL = source.BannerURL
				target.Hero = source.Hero
				target.HeroURL = source.HeroURL
				target.Description = source.Description
				target.Platform = source.Platform
				target.RelativePath = source.RelativePath

				// Merge tags to not lose current ones
				for _, tag := range source.Tags {
					if !slices.Contains(target.Tags, tag) {
						target.Tags = append(target.Tags, tag)
					}
				}

			},
			false,
		)
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

	// Make sure config folder path exist
	err = os.MkdirAll(_config.ConfigPath, 0774)
	if err != nil {
		return err
	}

	// Write JSON content to config file
	err = os.WriteFile(_config.StateFile, jsonContent, 0666)
	if err != nil {
		return err
	}

	// Save shortcuts file
	err = shortcuts.SaveToFile(_config.Shortcuts, _config.ShortcutsFile)
	if err != nil {
		return err
	}

	return nil
}

// Retrieve runtime config
func GetConfig() *Config {
	return &_config
}

// Retrieve runtime shortcuts
func GetShortcuts() []*shortcuts.Shortcut {

	if _config.Shortcuts == nil {
		_config.Shortcuts = make([]*shortcuts.Shortcut, 0)
	}

	return _config.Shortcuts
}

// Retrieve runtime shortcut with given appID
func GetShortcut(appID uint) *shortcuts.Shortcut {
	return shortcuts.GetShortcut(_config.Shortcuts, appID)
}

// Find shortcut with given executable and appName combination
func FindShortcut(appExe string, appName string) *shortcuts.Shortcut {
	executable := steam.EnsureExec(_config.SteamRuntime, appExe)
	appID := shortcuts.GenerateShortcutID(executable, appName)
	return GetShortcut(appID)
}

// Ensure that shortcut has the correct settings
func EnsureShortcut(shortcut *shortcuts.Shortcut) error {

	var err error

	// Ensure executable is correct for steam
	shortcut.Exe = steam.EnsureExec(_config.SteamRuntime, shortcut.Exe)

	// Determine appID and artworks path
	shortcut.AppID = shortcuts.GenerateShortcutID(shortcut.Exe, shortcut.AppName)
	artworksPath := _config.ArtworksPath
	remove := []string{}

	// Logo: ${APPID}_logo.png
	logoPng := fmt.Sprintf("%s/%v_logo.png", artworksPath, shortcut.AppID)
	logoPng = fs.NormalizePath(logoPng)

	if shortcut.LogoURL != "" {
		shortcut.Logo = logoPng
	} else {
		shortcut.Logo = ""
		remove = append(remove, logoPng)
	}

	// Icon: ${APPID}_icon.ico || ${APPID}_icon.png
	iconPng := fmt.Sprintf("%s/%v_icon.png", artworksPath, shortcut.AppID)
	iconPng = fs.NormalizePath(iconPng)

	iconIco := fmt.Sprintf("%s/%v_icon.ico", artworksPath, shortcut.AppID)
	iconIco = fs.NormalizePath(iconIco)

	if strings.HasSuffix(shortcut.IconURL, ".png") {
		shortcut.Icon = iconPng
		remove = append(remove, iconIco)
	} else if shortcut.IconURL != "" {
		shortcut.Icon = iconIco
		remove = append(remove, iconPng)
	} else {
		shortcut.Icon = ""
		remove = append(remove, iconPng)
		remove = append(remove, iconIco)
	}

	// Cover: ${APPID}p.png || ${APPID}p.jpg
	coverPng := fmt.Sprintf("%s/%vp.png", artworksPath, shortcut.AppID)
	coverPng = fs.NormalizePath(coverPng)

	coverJpg := fmt.Sprintf("%s/%vp.jpg", artworksPath, shortcut.AppID)
	coverJpg = fs.NormalizePath(coverJpg)

	if strings.HasSuffix(shortcut.CoverURL, ".png") {
		shortcut.Cover = coverPng
		remove = append(remove, coverJpg)
	} else if shortcut.CoverURL != "" {
		shortcut.Cover = coverJpg
		remove = append(remove, coverPng)
	} else {
		shortcut.Cover = ""
		remove = append(remove, coverPng)
		remove = append(remove, coverJpg)
	}

	// Banner: ${APPID}.png || ${APPID}.jpg
	bannerPng := fmt.Sprintf("%s/%v.png", artworksPath, shortcut.AppID)
	bannerPng = fs.NormalizePath(bannerPng)

	bannerJpg := fmt.Sprintf("%s/%v.jpg", artworksPath, shortcut.AppID)
	bannerJpg = fs.NormalizePath(bannerJpg)

	if strings.HasSuffix(shortcut.BannerURL, ".png") {
		shortcut.Banner = bannerPng
		remove = append(remove, bannerJpg)
	} else if shortcut.BannerURL != "" {
		shortcut.Banner = bannerJpg
		remove = append(remove, bannerPng)
	} else {
		shortcut.Banner = ""
		remove = append(remove, bannerPng)
		remove = append(remove, bannerJpg)
	}

	// Hero: ${APPID}_hero.png || ${APPID}_hero.jpg
	heroPng := fmt.Sprintf("%s/%v_hero.png", artworksPath, shortcut.AppID)
	heroPng = fs.NormalizePath(heroPng)

	heroJpg := fmt.Sprintf("%s/%v_hero.jpg", artworksPath, shortcut.AppID)
	heroJpg = fs.NormalizePath(heroJpg)

	if strings.HasSuffix(shortcut.HeroURL, ".png") {
		shortcut.Hero = heroPng
		remove = append(remove, heroJpg)
	} else if shortcut.HeroURL != "" {
		shortcut.Hero = heroJpg
		remove = append(remove, heroPng)
	} else {
		shortcut.Hero = ""
		remove = append(remove, heroPng)
		remove = append(remove, heroJpg)
	}

	// Remove duplicated or unnecessary images
	for _, file := range remove {
		err = fs.RemoveFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}

// Add program to the shortcuts list
func AddToShortcuts(shortcut *shortcuts.Shortcut, overwriteArtworks bool) error {

	var err error

	// Make sure shortcut settings is correct
	err = EnsureShortcut(shortcut)
	if err != nil {
		return err
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
		err := fs.DownloadFile(url, destinationFile, overwriteArtworks)
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

	// Remove all images of the shortcut
	images := []string{
		shortcut.Icon,
		shortcut.Logo,
		shortcut.Cover,
		shortcut.Banner,
		shortcut.Hero,
	}

	for _, file := range images {
		if file == "" {
			continue
		}
		err = fs.RemoveFile(file)
		if err != nil {
			return err
		}
	}

	// Remove the shortcut from list
	_config.Shortcuts, err = shortcuts.RemoveShortcut(_config.Shortcuts, shortcut)
	if err != nil {
		return err
	}

	return nil
}
