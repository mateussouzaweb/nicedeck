package steam

import (
	"strings"
)

type App struct {
	AppId              string   `json:"appid"`
	AppName            string   `json:"AppName"`
	Exe                string   `json:"Exe"`
	StartDir           string   `json:"StartDir"`
	Icon               string   `json:"icon"`
	ShortcutPath       string   `json:"ShortcutPath"`
	LaunchOptions      string   `json:"LaunchOptions"`
	IsHidden           string   `json:"IsHidden"`
	AllowDesktopConfig string   `json:"AllowDesktopConfig"`
	AllowOverlay       string   `json:"AllowOverlay"`
	OpenVR             string   `json:"OpenVR"`
	Devkit             string   `json:"Devkit"`
	DevkitGameID       string   `json:"DevkitGameID"`
	LastPlayTime       string   `json:"LastPlayTime"`
	Tags               []string `json:"tags"`
	IconURL            string   `json:"IconUrl"`
	LogoURL            string   `json:"LogoUrl"`
	CoverURL           string   `json:"CoverUrl"`
	BannerURL          string   `json:"BannerUrl"`
	HeroURL            string   `json:"HeroUrl"`
}

// Add non steam game to the steam library
func AddToSteam(app *App) error {

	// Get user path
	path, err := GetUserDataPath()
	if err != nil {
		return err
	}

	// Determine appId
	app.AppId = GenerateShortcutId(app.Exe + app.AppName)

	// Set icon path
	app.Icon = "{USERDATAPATH}/config/grid/{APPID}.ico"
	app.Icon = strings.Replace(app.Icon, "{USERDATAPATH}", path, 1)
	app.Icon = strings.Replace(app.Icon, "{APPID}", app.AppId, 1)

	// Download artworks
	err = DownloadArtworks(
		app.AppId,
		app.IconURL,
		app.LogoURL,
		app.CoverURL,
		app.BannerURL,
		app.HeroURL,
	)
	if err != nil {
		return err
	}

	// // Check if shortcuts file exist
	// file := path + "/config/shortcuts.vdf"
	// if !cli.ExistFile(file) {
	// 	return fmt.Errorf("shortcuts file does not exist: %s", file)
	// }

	// // Read file content
	// content, err := ReadVDF(file)
	// if err != nil {
	// 	return err
	// }

	// // Write new file content
	// err = WriteVDF(file, content)
	// return err

	return nil
}
