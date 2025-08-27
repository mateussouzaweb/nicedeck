package shortcuts

// Shortcut struct
type Shortcut struct {
	ID             string   `json:"id"`
	Platform       string   `json:"platform"`
	Program        string   `json:"program"`
	Layer          string   `json:"layer"`
	Type           string   `json:"type"`
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

// Perform action when creating the shortcut
func (s *Shortcut) OnCreate() error {

	// Ensure to have a valid shortcut ID
	if s.ID == "" || s.ID == "0" {
		s.ID = GenerateID(s.Name, s.Executable)
	}

	return nil
}

// Perform action when updating the shortcut
func (s *Shortcut) OnUpdate() error {
	return nil
}

// Perform action when removing the shortcut
func (s *Shortcut) OnRemove() error {
	return nil
}
