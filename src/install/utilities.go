package install

// Installer for ProtonPlus
func ProtonPlus() *Program {
	return &Program{
		ID:           "protonplus",
		Name:         "ProtonPlus",
		Description:  "Wine and Proton-based compatiblity tools manager",
		Tags:         []string{"Gaming", "Utilities"},
		FlatpakAppID: "com.vysp3r.ProtonPlus",
		IconURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/fe13849d9b9437c5a61a1760ada2a5a6.png",
		LogoURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/4d8c150eb82579842e2d5dc5faa07999.png",
		CoverURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/7901f04bfecd29119dfcce1c708108b1.png",
		BannerURL:    "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/f38705891f01bda4bd16551f42ff7c0a.png",
		HeroURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/bc6f714aa3dfeef9320a838b79515c2d.png",
	}
}
