package native

// Options struct
type Options struct {
	Platforms   []*Platform `json:"platforms"`
	Preferences []string    `json:"preferences"`
}

// Transform values into valid options
func ToOptions(preferences []string) (*Options, error) {

	options := &Options{
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
