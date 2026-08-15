package platforms

// Options struct
type Options struct {
	Include     []string `json:"include"`
	Preferences []string `json:"preferences"`
}

// Transform values into valid options
func ToOptions(include []string, preferences []string) *Options {

	options := Options{
		Include:     include,
		Preferences: preferences,
	}

	return &options
}
