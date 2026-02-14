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
	RelativePath   string   `json:"relativePath"`
	IconPath       string   `json:"iconPath"`
	LogoPath       string   `json:"logoPath"`
	CoverPath      string   `json:"coverPath"`
	BannerPath     string   `json:"bannerPath"`
	HeroPath       string   `json:"heroPath"`
	Tags           []string `json:"tags"`
	Timestamp      int64    `json:"timestamp"`
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
	if source.RelativePath != "" {
		s.RelativePath = source.RelativePath
	}
	if source.IconPath != "" {
		s.IconPath = source.IconPath
	}
	if source.LogoPath != "" {
		s.LogoPath = source.LogoPath
	}
	if source.CoverPath != "" {
		s.CoverPath = source.CoverPath
	}
	if source.BannerPath != "" {
		s.BannerPath = source.BannerPath
	}
	if source.HeroPath != "" {
		s.HeroPath = source.HeroPath
	}
	if len(source.Tags) > 0 {
		s.Tags = source.Tags
	}
	if source.Timestamp > s.Timestamp {
		s.Timestamp = source.Timestamp
	}
}
