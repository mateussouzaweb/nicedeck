package scraper

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/steamgriddb/api"
)

// ScrapeInfo struct
type ScrapeInfo struct {
	Name       string   `json:"name"`
	ScraperID  int64    `json:"scraperId"`
	ShortcutID uint32   `json:"shortcutId"`
	IconURLs   []string `json:"iconUrls"`
	LogoURLs   []string `json:"logoUrls"`
	CoverURLs  []string `json:"coverUrls"`
	BannerURLs []string `json:"bannerUrls"`
	HeroURLs   []string `json:"heroUrls"`
}

// Scrape information such as images from given app or game name
func ScrapeFromName(name string) (*ScrapeInfo, error) {

	var result ScrapeInfo

	// Find reference and correct name
	search, err := api.SearchByTerm(name)
	if err != nil {
		return &result, err
	}

	if search.Success && len(search.Data) > 0 {
		searchResult := search.Data[0]
		if searchResult.ID != 0 {
			result.ScraperID = searchResult.ID
			result.Name = strings.Trim(searchResult.Name, " ")
		}
	}

	// Cancel reaming actions if not found
	if result.ScraperID == 0 {
		return &result, nil
	}

	// Find icon
	icon, err := api.GetImagesByID(
		"icon",
		fmt.Sprintf("%v", result.ScraperID),
		&api.ImagesParams{
			Dimensions: []string{"24", "32", "40", "48", "56", "64", "72", "80", "96", "100", "144", "192"},
			Mimes:      []string{"image/png", "image/vnd.microsoft.icon"},
			Types:      []string{"static"},
			Nsfw:       "false",
			Humor:      "false",
			Epilepsy:   "false",
		},
	)
	if err != nil {
		return &result, err
	}
	if icon.Success && len(icon.Data) > 0 {
		for _, item := range icon.Data {
			result.IconURLs = append(result.IconURLs, item.URL)
		}
	}

	// Find logo
	logo, err := api.GetImagesByID(
		"logo",
		fmt.Sprintf("%v", result.ScraperID),
		&api.ImagesParams{
			Mimes:    []string{"image/png"},
			Types:    []string{"static"},
			Nsfw:     "false",
			Humor:    "false",
			Epilepsy: "false",
		},
	)
	if err != nil {
		return &result, err
	}
	if logo.Success && len(logo.Data) > 0 {
		for _, item := range logo.Data {
			result.LogoURLs = append(result.LogoURLs, item.URL)
		}
	}

	// Find cover
	cover, err := api.GetImagesByID(
		"cover",
		fmt.Sprintf("%v", result.ScraperID),
		&api.ImagesParams{
			Mimes:      []string{"image/png", "image/jpeg"},
			Types:      []string{"static"},
			Dimensions: []string{"600x900"},
			Nsfw:       "false",
			Humor:      "false",
			Epilepsy:   "false",
		},
	)
	if err != nil {
		return &result, err
	}
	if cover.Success && len(cover.Data) > 0 {
		for _, item := range cover.Data {
			result.CoverURLs = append(result.CoverURLs, item.URL)
		}
	}

	// Find banner
	banner, err := api.GetImagesByID(
		"banner",
		fmt.Sprintf("%v", result.ScraperID),
		&api.ImagesParams{
			Mimes:      []string{"image/png", "image/jpeg"},
			Types:      []string{"static"},
			Dimensions: []string{"920x430", "460x215"},
			Nsfw:       "false",
			Humor:      "false",
			Epilepsy:   "false",
		},
	)
	if err != nil {
		return &result, err
	}
	if banner.Success && len(banner.Data) > 0 {
		for _, item := range banner.Data {
			result.BannerURLs = append(result.BannerURLs, item.URL)
		}
	}

	// Find hero
	hero, err := api.GetImagesByID(
		"hero",
		fmt.Sprintf("%v", result.ScraperID),
		&api.ImagesParams{
			Mimes:    []string{"image/png", "image/jpeg"},
			Types:    []string{"static"},
			Nsfw:     "false",
			Humor:    "false",
			Epilepsy: "false",
		},
	)
	if err != nil {
		return &result, err
	}
	if hero.Success && len(hero.Data) > 0 {
		for _, item := range hero.Data {
			result.HeroURLs = append(result.HeroURLs, item.URL)
		}
	}

	return &result, nil
}
