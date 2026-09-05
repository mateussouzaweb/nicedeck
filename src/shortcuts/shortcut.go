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
	Reference      string   `json:"reference"`
	IconPath       string   `json:"iconPath"`
	LogoPath       string   `json:"logoPath"`
	CoverPath      string   `json:"coverPath"`
	BannerPath     string   `json:"bannerPath"`
	HeroPath       string   `json:"heroPath"`
	Tags           []string `json:"tags"`
	Timestamp      int64    `json:"timestamp"`

	// @Deprecated: RelativePath is deprecated and will be removed in future versions. Use Reference instead.
	RelativePath string `json:"relativePath"`
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
	if source.Reference != "" {
		s.Reference = source.Reference
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
