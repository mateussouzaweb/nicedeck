package scraper

// Options struct
type Options struct {
	Search string `json:"search"`
	Icon   bool   `json:"icon"`
	Logo   bool   `json:"logo"`
	Cover  bool   `json:"cover"`
	Banner bool   `json:"banner"`
	Hero   bool   `json:"hero"`
}

// Transform values into valid options
func ToOptions(search string, icon bool, logo bool, cover bool, banner bool, hero bool) *Options {

	options := Options{
		Search: search,
		Icon:   icon,
		Logo:   logo,
		Cover:  cover,
		Banner: banner,
		Hero:   hero,
	}

	return &options
}
