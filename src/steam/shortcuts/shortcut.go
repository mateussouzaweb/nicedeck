package shortcuts

// Shortcut struct
type Shortcut struct {
	AppID               uint     `json:"appId"`
	AppName             string   `json:"appName"`
	Exe                 string   `json:"exe"`
	StartDir            string   `json:"startDir"`
	Icon                string   `json:"icon"`
	IconURL             string   `json:"iconUrl"`
	Logo                string   `json:"logo"`
	LogoURL             string   `json:"logoUrl"`
	Cover               string   `json:"cover"`
	CoverURL            string   `json:"coverUrl"`
	Banner              string   `json:"banner"`
	BannerURL           string   `json:"bannerUrl"`
	Hero                string   `json:"hero"`
	HeroURL             string   `json:"heroUrl"`
	ShortcutPath        string   `json:"shortcutPath"`
	LaunchOptions       string   `json:"launchOptions"`
	IsHidden            uint     `json:"isHidden"`
	AllowDesktopConfig  uint     `json:"allowDesktopConfig"`
	AllowOverlay        uint     `json:"allowOverlay"`
	OpenVR              uint     `json:"openVr"`
	Devkit              uint     `json:"devkit"`
	DevkitGameID        string   `json:"devkitGameId"`
	DevkitOverrideAppID uint     `json:"devkitOverrideAppId"`
	FlatpakAppID        string   `json:"flatpakAppId"`
	LastPlayTime        uint     `json:"lastPlayTime"`
	Tags                []string `json:"tags"`
}
