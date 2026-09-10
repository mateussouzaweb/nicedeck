package launcher

// Scanner function type
type Scanner func(options *Options) ([]*Game, error)

// Platform struct
type Platform struct {
	Name      string  `json:"name"`
	Launcher  string  `json:"launcher"`
	Program   string  `json:"program"`
	Available bool    `json:"available"`
	Installed bool    `json:"installed"`
	Scan      Scanner `json:"-"`
}

// Retrieve launcher platform specs
func GetPlatforms() ([]*Platform, error) {

	platforms := []*Platform{}

	// Amazon Games
	amazon, err := AmazonPlatform()
	if err != nil {
		return platforms, err
	} else {
		platforms = append(platforms, amazon)
	}

	// Battle.net
	battleNet, err := BattleNetPlatform()
	if err != nil {
		return platforms, err
	} else {
		platforms = append(platforms, battleNet)
	}

	// EA Games
	ea, err := EAPlatform()
	if err != nil {
		return platforms, err
	} else {
		platforms = append(platforms, ea)
	}

	// Epic Games
	epic, err := EpicPlatform()
	if err != nil {
		return platforms, err
	} else {
		platforms = append(platforms, epic)
	}

	// GOG
	gog, err := GOGPlatform()
	if err != nil {
		return platforms, err
	} else {
		platforms = append(platforms, gog)
	}

	// Ubisoft Games
	ubisoft, err := UbisoftPlatform()
	if err != nil {
		return platforms, err
	} else {
		platforms = append(platforms, ubisoft)
	}

	return platforms, nil
}
