package management

import (
	"github.com/mateussouzaweb/nicedeck/src/scraper"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
)

// Scrape information such as images from given app or game name
func ScrapeData(options *scraper.Options) (*scraper.ScrapeResult, error) {
	return scraper.Scrape(options)
}

// Scrape data from shortcut and return if was found
func ScrapeShortcut(shortcut *shortcuts.Shortcut) (bool, error) {

	// Scrape additional ROM information
	options := scraper.ToOptions(shortcut.Name, true, true, true, true, true)
	scrape, err := scraper.Scrape(options)
	if err != nil {
		return false, err
	}

	// Skip if scrape not found anything...
	if scrape.Name == "" {
		return false, nil
	}

	// Determine best name and images for the shortcut
	shortcut.Name = scrape.Name

	if len(scrape.IconURLs) > 0 {
		shortcut.IconPath = scrape.IconURLs[0]
	}
	if len(scrape.LogoURLs) > 0 {
		shortcut.LogoPath = scrape.LogoURLs[0]
	}
	if len(scrape.CoverURLs) > 0 {
		shortcut.CoverPath = scrape.CoverURLs[0]
	}
	if len(scrape.BannerURLs) > 0 {
		shortcut.BannerPath = scrape.BannerURLs[0]
	}
	if len(scrape.HeroURLs) > 0 {
		shortcut.HeroPath = scrape.HeroURLs[0]
	}

	return true, nil
}
