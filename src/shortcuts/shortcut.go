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

// Safely merge shortcuts
func (s *Shortcut) Merge(source *Shortcut) {
	if source.Program != "" {
		s.Program = source.Program
	}
	if source.Name != "" {
		s.Name = source.Name
	}
	if source.Description != "" {
		s.Description = source.Description
	}
	if source.StartDirectory != "" {
		s.StartDirectory = source.StartDirectory
	}
	if source.Executable != "" {
		s.Executable = source.Executable
	}
	if source.LaunchOptions != "" {
		s.LaunchOptions = source.LaunchOptions
	}
	if source.ShortcutPath != "" {
		s.ShortcutPath = source.ShortcutPath
	}
	if source.RelativePath != "" {
		s.RelativePath = source.RelativePath
	}
	if source.IconURL != "" {
		s.IconURL = source.IconURL
	}
	if source.LogoURL != "" {
		s.LogoURL = source.LogoURL
	}
	if source.CoverURL != "" {
		s.CoverURL = source.CoverURL
	}
	if source.BannerURL != "" {
		s.BannerURL = source.BannerURL
	}
	if source.HeroURL != "" {
		s.HeroURL = source.HeroURL
	}
	if len(source.Tags) > 0 {
		s.Tags = source.Tags
	}
}
