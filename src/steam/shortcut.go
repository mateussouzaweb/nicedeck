package steam

// Shortcut struct
type Shortcut struct {
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
	DevKit              uint     `json:"devKit"`
	DevKitGameID        string   `json:"devKitGameId"`
	DevKitOverrideAppID uint     `json:"devKitOverrideAppId"`
	LastPlayTime        uint     `json:"lastPlayTime"`
	Tags                []string `json:"tags"`
	Timestamp           int64    `json:"timestamp"`

	// Extended specs
	// @deprecated and will be removed in future versions
	RelativePath string `json:"relativePath"`
	Description  string `json:"description"`
	IconURL      string `json:"iconUrl"`
	LogoURL      string `json:"logoUrl"`
	CoverURL     string `json:"coverUrl"`
	BannerURL    string `json:"bannerUrl"`
	HeroURL      string `json:"heroUrl"`
}
