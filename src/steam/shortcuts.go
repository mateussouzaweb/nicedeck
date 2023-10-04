package steam

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Shortcut struct
type Shortcut struct {
	AppID               uint     `json:"appId" vdf:"appid"`
	AppName             string   `json:"appName" vdf:"AppName"`
	Exe                 string   `json:"exe" vdf:"Exe"`
	StartDir            string   `json:"startDir" vdf:"StartDir"`
	Icon                string   `json:"icon" vdf:"icon"`
	ShortcutPath        string   `json:"shortcutPath" vdf:"ShortcutPath"`
	LaunchOptions       string   `json:"launchOptions" vdf:"LaunchOptions"`
	IsHidden            uint     `json:"isHidden" vdf:"IsHidden"`
	AllowDesktopConfig  uint     `json:"allowDesktopConfig" vdf:"AllowDesktopConfig"`
	AllowOverlay        uint     `json:"allowOverlay" vdf:"AllowOverlay"`
	OpenVR              uint     `json:"openVr" vdf:"OpenVR"`
	Devkit              uint     `json:"devkit" vdf:"Devkit"`
	DevkitGameID        string   `json:"devkitGameId" vdf:"DevkitGameID"`
	DevkitOverrideAppID uint     `json:"devkitOverrideAppId" vdf:"DevkitOverrideAppID"`
	FlatpakAppID        string   `json:"flatpakAppId" vdf:"FlatpakAppID"`
	LastPlayTime        uint     `json:"lastPlayTime" vdf:"LastPlayTime"`
	Tags                []string `json:"tags" vdf:"tags"`
	IconURL             string   `json:"iconUrl" vdf:"IconUrl"`
	LogoURL             string   `json:"logoUrl" vdf:"LogoUrl"`
	CoverURL            string   `json:"coverUrl" vdf:"CoverUrl"`
	BannerURL           string   `json:"bannerUrl" vdf:"BannerUrl"`
	HeroURL             string   `json:"heroUrl" vdf:"HeroUrl"`
}

// Add non steam game to the steam shortcuts library
func AddToShotcuts(shortcut *Shortcut) error {

	// Determine appId
	shortcut.AppID = GenerateShortcutID(shortcut.Exe, shortcut.AppName)

	// Set icon path
	shortcut.Icon = fmt.Sprintf("%s/%v.ico", _config.ArtworksPath, shortcut.AppID)

	// Download artworks images
	// Required format
	// Icon: ${APPID}.ico
	// Logo: ${APPID}_logo.png
	// Cover: ${APPID}p.png
	// Banner: ${APPID}.png
	// Hero: ${APPID}_hero.png
	err := cli.Command(fmt.Sprintf(`
		# Make sure folder exist
		mkdir -p %s

		# Download images
		[ "%s" != "" ] && wget -q -O %s/%v.ico %s
		[ "%s" != "" ] && wget -q -O %s/%v_logo.png %s
		[ "%s" != "" ] && wget -q -O %s/%vp.png %s
		[ "%s" != "" ] && wget -q -O %s/%v.png %s
		[ "%s" != "" ] && wget -q -O %s/%v_hero.png %s
		`,
		_config.ArtworksPath,
		shortcut.IconURL, _config.ArtworksPath, shortcut.AppID, shortcut.IconURL,
		shortcut.LogoURL, _config.ArtworksPath, shortcut.AppID, shortcut.LogoURL,
		shortcut.CoverURL, _config.ArtworksPath, shortcut.AppID, shortcut.CoverURL,
		shortcut.BannerURL, _config.ArtworksPath, shortcut.AppID, shortcut.BannerURL,
		shortcut.HeroURL, _config.ArtworksPath, shortcut.AppID, shortcut.HeroURL,
	)).Run()

	if err != nil {
		return err
	}

	// Check if already exist an app with the same reference
	found := false
	for index, item := range _config.Shortcuts {
		if item.AppID == shortcut.AppID {

			// Keep current value for some keys
			shortcut.IsHidden = item.IsHidden
			shortcut.AllowDesktopConfig = item.AllowDesktopConfig
			shortcut.AllowOverlay = item.AllowOverlay
			shortcut.OpenVR = item.OpenVR
			shortcut.Devkit = item.Devkit
			shortcut.DevkitGameID = item.DevkitGameID
			shortcut.DevkitOverrideAppID = item.DevkitOverrideAppID
			shortcut.FlatpakAppID = item.FlatpakAppID
			shortcut.LastPlayTime = item.LastPlayTime
			shortcut.Tags = item.Tags

			// Replace with new object data
			_config.Shortcuts[index] = shortcut

			found = true
			break
		}
	}

	// Append to the list if not exist
	if !found {
		_config.Shortcuts = append(_config.Shortcuts, shortcut)
	}

	return nil
}
