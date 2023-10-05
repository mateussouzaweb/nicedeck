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
func AddToShortcuts(shortcut *Shortcut) error {

	// Get destination for images
	artworksPath := _config.ArtworksPath

	// Determine appId
	shortcut.AppID = GenerateShortcutID(shortcut.Exe, shortcut.AppName)

	// Set icon path
	shortcut.Icon = fmt.Sprintf("%s/%v.ico", artworksPath, shortcut.AppID)

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
		artworksPath,
		shortcut.IconURL, artworksPath, shortcut.AppID, shortcut.IconURL,
		shortcut.LogoURL, artworksPath, shortcut.AppID, shortcut.LogoURL,
		shortcut.CoverURL, artworksPath, shortcut.AppID, shortcut.CoverURL,
		shortcut.BannerURL, artworksPath, shortcut.AppID, shortcut.BannerURL,
		shortcut.HeroURL, artworksPath, shortcut.AppID, shortcut.HeroURL,
	)).Run()

	if err != nil {
		return err
	}

	return _config.AddShortcut(shortcut)
}
