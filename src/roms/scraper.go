package roms

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/steamgriddb/api"
)

// ScrapeInfo struct
type ScrapeInfo struct {
	Name       string `json:"name"`
	ScraperId  int64  `json:"scraperId"`
	ShortcutId uint32 `json:"shortcutId"`
	IconURL    string `json:"iconUrl"`
	LogoURL    string `json:"logoUrl"`
	CoverURL   string `json:"coverUrl"`
	BannerURL  string `json:"bannerUrl"`
	HeroURL    string `json:"heroUrl"`
}

// Scrape additional ROM information such as images
func ScrapeROM(rom *ROM) (*ScrapeInfo, error) {

	var result ScrapeInfo

	// Find reference and correct name
	search, err := api.SearchByTerm(rom.Name)
	if err != nil {
		return &result, err
	}

	if search.Success && len(search.Data) > 0 {
		searchResult := search.Data[0]
		if searchResult.Success && searchResult.Data.ID != 0 {
			result.ScraperId = searchResult.Data.ID
			result.Name = searchResult.Data.Name
		}
	}

	// Cancel reaming actions if not found
	if result.ScraperId == 0 {
		return &result, nil
	}

	// Find icon
	icon, err := api.GetImagesByPlatformAndId(
		"icon",
		"steam",
		fmt.Sprintf("%v", result.ScraperId),
		&api.ImagesParams{
			Nsfw:  "false",
			Humor: "false",
			Page:  1,
		},
	)
	if err != nil {
		return &result, err
	}
	if icon.Success && len(icon.Data) > 0 {
		result.IconURL = icon.Data[0].URL
	}

	// Find logo
	logo, err := api.GetImagesByPlatformAndId(
		"logo",
		"steam",
		fmt.Sprintf("%v", result.ScraperId),
		&api.ImagesParams{
			Nsfw:  "false",
			Humor: "false",
			Page:  1,
		},
	)
	if err != nil {
		return &result, err
	}
	if logo.Success && len(logo.Data) > 0 {
		result.LogoURL = logo.Data[0].URL
	}

	// Find cover
	cover, err := api.GetImagesByPlatformAndId(
		"cover",
		"steam",
		fmt.Sprintf("%v", result.ScraperId),
		&api.ImagesParams{
			Nsfw:  "false",
			Humor: "false",
			Page:  1,
		},
	)
	if err != nil {
		return &result, err
	}
	if cover.Success && len(cover.Data) > 0 {
		result.CoverURL = cover.Data[0].URL
	}

	// Find banner
	banner, err := api.GetImagesByPlatformAndId(
		"banner",
		"steam",
		fmt.Sprintf("%v", result.ScraperId),
		&api.ImagesParams{
			Nsfw:  "false",
			Humor: "false",
			Page:  1,
		},
	)
	if err != nil {
		return &result, err
	}
	if banner.Success && len(banner.Data) > 0 {
		result.BannerURL = banner.Data[0].URL
	}

	// Find hero
	hero, err := api.GetImagesByPlatformAndId(
		"hero",
		"steam",
		fmt.Sprintf("%v", result.ScraperId),
		&api.ImagesParams{
			Nsfw:  "false",
			Humor: "false",
			Page:  1,
		},
	)
	if err != nil {
		return &result, err
	}
	if hero.Success && len(hero.Data) > 0 {
		result.HeroURL = hero.Data[0].URL
	}

	return &result, nil
}
