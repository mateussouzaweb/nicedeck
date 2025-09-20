package native

// Options struct
type Options struct {
	Platforms   []string `json:"platforms"`
	Preferences []string `json:"preferences"`
}

// Transform values into valid options
func ToOptions(platforms []string, preferences []string) *Options {

	options := Options{
		Platforms:   platforms,
		Preferences: preferences,
	}

	return &options
}
