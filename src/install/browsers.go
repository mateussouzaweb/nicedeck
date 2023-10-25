package install

// Installer for Firefox
func Firefox() *Program {
	return &Program{
		ID:           "firefox",
		Name:         "Firefox",
		Description:  "Web browser",
		Tags:         []string{"Utilities"},
		FlatpakAppID: "org.mozilla.firefox",
		IconURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/f968fdc88852a4a3a27a81fe3f57bfc5.ico",
		LogoURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/43285a8b542fcdc35377439e05dcb04f.png",
		CoverURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/4529f985441a035ae4a107b8862ba4dd.png",
		BannerURL:    "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/9384fe92aef7ea0128be2c916ed07cea.png",
		HeroURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/a318166b8539611449bf21ddc297a783.png",
	}
}

// Installer for Google Chrome
func GoogleChrome() *Program {
	return &Program{
		ID:           "google-chrome",
		Name:         "Google Chrome",
		Description:  "Web browser",
		Tags:         []string{"Utilities"},
		FlatpakAppID: "com.google.Chrome",
		IconURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3941c4358616274ac2436eacf67fae05.ico",
		LogoURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/3b049d0f6cbf5421d399f156807b8657.png",
		CoverURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d45c26607db83f6f14b09dd70123913b.png",
		BannerURL:    "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d40c243072a2d2957b3484e775f1f925.png",
		HeroURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/cae83cfcb1d8a2a4bb17bd1446fb1cee.png",
	}
}
