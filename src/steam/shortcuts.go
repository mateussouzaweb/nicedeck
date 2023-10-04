package steam

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Shortcut struct
type Shortcut struct {
	AppId              string   `json:"appid" vdf:"appid"`
	AppName            string   `json:"appName" vdf:"AppName"`
	Exe                string   `json:"exe" vdf:"Exe"`
	StartDir           string   `json:"startDir" vdf:"StartDir"`
	Icon               string   `json:"icon" vdf:"icon"`
	ShortcutPath       string   `json:"shortcutPath" vdf:"ShortcutPath"`
	LaunchOptions      string   `json:"launchOptions" vdf:"LaunchOptions"`
	IsHidden           string   `json:"isHidden" vdf:"IsHidden"`
	AllowDesktopConfig string   `json:"allowDesktopConfig" vdf:"AllowDesktopConfig"`
	AllowOverlay       string   `json:"allowOverlay" vdf:"AllowOverlay"`
	OpenVR             string   `json:"openVR" vdf:"OpenVR"`
	Devkit             string   `json:"devkit" vdf:"Devkit"`
	DevkitGameID       string   `json:"devkitGameID" vdf:"DevkitGameID"`
	LastPlayTime       string   `json:"lastPlayTime" vdf:"LastPlayTime"`
	Tags               []string `json:"tags" vdf:"tags"`
	IconURL            string   `json:"iconUrl" vdf:"IconUrl"`
	LogoURL            string   `json:"logoUrl" vdf:"LogoUrl"`
	CoverURL           string   `json:"coverUrl" vdf:"CoverUrl"`
	BannerURL          string   `json:"bannerUrl" vdf:"BannerUrl"`
	HeroURL            string   `json:"heroUrl" vdf:"HeroUrl"`
}

// Shortcuts struct
type Shortcuts = struct {
	Shortcuts []*Shortcut `json:"shortcuts" vdf:"shortcuts"`
}

// Add non steam game to the steam shortcuts library
func AddToShotcuts(shortcut *Shortcut) error {

	// Determine appId
	shortcut.AppId = GenerateShortcutId(shortcut.Exe + shortcut.AppName)

	// Set icon path
	shortcut.Icon = _config.ArtworksPath + "/" + shortcut.AppId + ".ico"

	// Download artworks images
	// Required format
	// Icon: ${APPID}.ico
	// Logo: ${APPID}_logo.png
	// Cover: ${APPID}p.png
	// Banner: ${APPID}.png
	// Hero: ${APPID}_hero_.png
	err := cli.Command(fmt.Sprintf(`
		# Make sure folder exist
		mkdir -p %s

		# Download images
		[ "%s" != "" ] && wget -q -O %s/%s.ico %s
		[ "%s" != "" ] && wget -q -O %s/%s_logo.png %s
		[ "%s" != "" ] && wget -q -O %s/%sp.png %s
		[ "%s" != "" ] && wget -q -O %s/%s.png %s
		[ "%s" != "" ] && wget -q -O %s/%s_hero.png %s
		`,
		_config.ArtworksPath,
		shortcut.IconURL, _config.ArtworksPath, shortcut.AppId, shortcut.IconURL,
		shortcut.LogoURL, _config.ArtworksPath, shortcut.AppId, shortcut.LogoURL,
		shortcut.CoverURL, _config.ArtworksPath, shortcut.AppId, shortcut.CoverURL,
		shortcut.BannerURL, _config.ArtworksPath, shortcut.AppId, shortcut.BannerURL,
		shortcut.HeroURL, _config.ArtworksPath, shortcut.AppId, shortcut.HeroURL,
	)).Run()

	if err != nil {
		return err
	}

	// Check if already exist an app with the same reference
	found := false
	for index, item := range _config.Shortcuts.Shortcuts {
		if item.AppId == shortcut.AppId {

			// Keep current value for some keys
			shortcut.IsHidden = item.IsHidden
			shortcut.AllowDesktopConfig = item.AllowDesktopConfig
			shortcut.AllowOverlay = item.AllowOverlay
			shortcut.OpenVR = item.OpenVR
			shortcut.Devkit = item.Devkit
			shortcut.DevkitGameID = item.DevkitGameID
			shortcut.LastPlayTime = item.LastPlayTime
			shortcut.Tags = item.Tags

			// Replace with new object data
			_config.Shortcuts.Shortcuts[index] = shortcut

			found = true
			break
		}
	}

	// Append to the list if not exist
	if !found {
		_config.Shortcuts.Shortcuts = append(_config.Shortcuts.Shortcuts, shortcut)
	}

	return nil
}
