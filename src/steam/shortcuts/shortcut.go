package shortcuts

// Shortcut struct
type Shortcut struct {
	// Default specs
	AppID               uint     `json:"appId"`
	AppName             string   `json:"appName"`
	StartDir            string   `json:"startDir"`
	Exe                 string   `json:"exe"`
	LaunchOptions       string   `json:"launchOptions"`
	ShortcutPath        string   `json:"shortcutPath"`
	Icon                string   `json:"icon"`
	IsHidden            uint     `json:"isHidden"`
	AllowDesktopConfig  uint     `json:"allowDesktopConfig"`
	AllowOverlay        uint     `json:"allowOverlay"`
	OpenVR              uint     `json:"openVr"`
	Devkit              uint     `json:"devkit"`
	DevkitGameID        string   `json:"devkitGameId"`
	DevkitOverrideAppID uint     `json:"devkitOverrideAppId"`
	LastPlayTime        uint     `json:"lastPlayTime"`
	Tags                []string `json:"tags"`

	// Extended specs
	IconURL      string `json:"iconUrl"`
	Logo         string `json:"logo"`
	LogoURL      string `json:"logoUrl"`
	Cover        string `json:"cover"`
	CoverURL     string `json:"coverUrl"`
	Banner       string `json:"banner"`
	BannerURL    string `json:"bannerUrl"`
	Hero         string `json:"hero"`
	HeroURL      string `json:"heroUrl"`
	Platform     string `json:"platform"`
	RelativePath string `json:"relativePath"`
}
