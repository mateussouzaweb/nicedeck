package packaging

// Program struct
type Program struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Folders     []string `json:"folders"`
	Flags       []string `json:"flags"`
	Website     string   `json:"website"`
	IconURL     string   `json:"iconUrl"`
	LogoURL     string   `json:"logoUrl"`
	CoverURL    string   `json:"coverUrl"`
	BannerURL   string   `json:"bannerUrl"`
	HeroURL     string   `json:"heroUrl"`
	Package     Package  `json:"-"`
}
