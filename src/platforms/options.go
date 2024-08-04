package platforms

// Options struct
type Options struct {
	Platforms   []string `json:"platforms"`
	Preferences []string `json:"preferences"`
	Rebuild     bool     `json:"rebuild"`
}

// Transform values into valid options
func ToOptions(platforms []string, preferences []string, rebuild bool) *Options {

	options := Options{
		Platforms:   platforms,
		Rebuild:     rebuild,
		Preferences: preferences,
	}

	return &options
}
