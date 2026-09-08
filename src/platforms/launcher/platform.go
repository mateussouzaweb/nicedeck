package launcher

// Scanner function type
type Scanner func() ([]*Game, error)

// Platform struct
type Platform struct {
	Name     string  `json:"name"`
	Launcher string  `json:"launcher"`
	Program  string  `json:"program"`
	Scan     Scanner `json:"-"`
}

// Retrieve launcher platform specs
func GetPlatforms() ([]*Platform, error) {

	platforms := []*Platform{}
	platforms = append(platforms, AmazonPlatform())
	platforms = append(platforms, EAPlatform())
	platforms = append(platforms, EpicPlatform())
	platforms = append(platforms, GOGPlatform())

	return platforms, nil
}
