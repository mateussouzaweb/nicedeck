package launcher

// Options struct
type Options struct {
	Platforms   []*Platform `json:"platforms"`
	Include     []string    `json:"include"`
	Preferences []string    `json:"preferences"`
}

// Transform values into valid options
func ToOptions(include []string, preferences []string) (*Options, error) {

	options := &Options{
		Include:     include,
		Preferences: preferences,
		Platforms:   []*Platform{},
	}

	platforms, err := GetPlatforms()
	if err != nil {
		return options, err
	} else {
		options.Platforms = platforms
	}

	return options, nil
}
