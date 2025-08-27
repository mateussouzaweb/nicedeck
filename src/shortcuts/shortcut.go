package shortcuts

// Shortcut struct
type Shortcut struct {
	ID             string   `json:"id"`
	Program        string   `json:"program"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	StartDirectory string   `json:"startDirectory"`
	Executable     string   `json:"executable"`
	LaunchOptions  string   `json:"launchOptions"`
	ShortcutPath   string   `json:"shortcutPath"`
	RelativePath   string   `json:"relativePath"`
	IconPath       string   `json:"iconPath"`
	IconURL        string   `json:"iconUrl"`
	LogoPath       string   `json:"logoPath"`
	LogoURL        string   `json:"logoUrl"`
	CoverPath      string   `json:"coverPath"`
	CoverURL       string   `json:"coverUrl"`
	BannerPath     string   `json:"bannerPath"`
	BannerURL      string   `json:"bannerUrl"`
	HeroPath       string   `json:"heroPath"`
	HeroURL        string   `json:"heroUrl"`
	Tags           []string `json:"tags"`
}
