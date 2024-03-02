package library

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Config struct
type Config struct {
	IsFlatpak               bool                  `json:"isFlatpak"`
	SteamPath               string                `json:"steamPath"`
	UserConfigPath          string                `json:"userConfigPath"`
	UserArtworksPath        string                `json:"userArtworksPath"`
	ControllerTemplatesPath string                `json:"controllerTemplatesPath"`
	Shortcuts               []*shortcuts.Shortcut `json:"shortcuts"`
}

var _config Config

// Load data to runtime config
func Load() error {

	// Retrieve Steam base path
	steamPaths, err := steam.GetPaths("")
	if err != nil {
		return fmt.Errorf("could not detect Steam - please make sure to install Steam first: %s", err)
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
	userConfigPaths, err := steam.GetPaths("userdata/*/config")
	if err != nil {
		return fmt.Errorf("could not detect Steam user configuration - please make sure to login into Steam first: %s", err)
	}

	// Make sure zero config is ignored (this is not a valid user)
	if strings.Contains(userConfigPaths[0], "/0/config") {
		userConfigPaths = userConfigPaths[1:]
	}

	// Retrieve controller templates path
	controllerTemplatesPaths, err := steam.GetPaths("controller_base/templates")
	if err != nil {
		return err
	}

	// Set runtime configs
	_config = Config{}
	_config.IsFlatpak = isFlatpak
	_config.SteamPath = steamPaths[0]
	_config.UserConfigPath = userConfigPaths[0]
	_config.UserArtworksPath = userConfigPaths[0] + "/grid"
	_config.ControllerTemplatesPath = controllerTemplatesPaths[0]

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
		err = json.Unmarshal(content, &_config)
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

	// Check if library was loaded
	if _config.SteamPath == "" {
		err = fmt.Errorf("cannot save library because Steam was not detected")
		return err
	}

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

// Ensure that shortcut has the correct settings
func EnsureShortcut(shortcut *shortcuts.Shortcut) error {

	var err error

	// Check if library was loaded
	if _config.SteamPath == "" {
		err = fmt.Errorf("cannot process library shortcut because Steam was not properly detected")
		return err
	}

	// Check if Steam was installed via flatpak
	// If yes, then we need to append the flatpak-spawn wrapper
	// This is necessary to have access to host commands
	if _config.IsFlatpak {
		shortcut.Exe = "/usr/bin/flatpak-spawn --host " + shortcut.Exe
	}

	// Determine appID and artworks path
	shortcut.AppID = shortcuts.GenerateShortcutID(shortcut.Exe, shortcut.AppName)
	artworksPath := _config.UserArtworksPath
	remove := []string{}

	// Logo: ${APPID}_logo.png
	logoPng := fmt.Sprintf("%s/%v_logo.png", artworksPath, shortcut.AppID)

	if shortcut.LogoURL != "" {
		shortcut.Logo = logoPng
	} else {
		shortcut.Logo = ""
		remove = append(remove, logoPng)
	}

	// Icon: ${APPID}_icon.ico || ${APPID}_icon.png
	iconPng := fmt.Sprintf("%s/%v_icon.png", artworksPath, shortcut.AppID)
	iconIco := fmt.Sprintf("%s/%v_icon.ico", artworksPath, shortcut.AppID)

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
	coverJpg := fmt.Sprintf("%s/%vp.jpg", artworksPath, shortcut.AppID)

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
	bannerJpg := fmt.Sprintf("%s/%v.jpg", artworksPath, shortcut.AppID)

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
	heroJpg := fmt.Sprintf("%s/%v_hero.jpg", artworksPath, shortcut.AppID)

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

	if _config.SteamPath == "" {
		err = fmt.Errorf("cannot add library shortcut because Steam was not properly detected")
		return err
	}

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

	if _config.SteamPath == "" {
		err = fmt.Errorf("cannot remove library shortcut because Steam was not properly detected")
		return err
	}

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
