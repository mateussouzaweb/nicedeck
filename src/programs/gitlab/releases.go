package gitlab

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get asset direct download URL from the latest release available
func GetAssetURL(domain string, projectId string, search string) (string, error) {

	domain = strings.Trim(domain, "/")
	endpoint := fmt.Sprintf("%s/api/v4/projects/%s/releases", domain, projectId)

	// Response struct
	var releases []struct {
		Name            string `json:"name"`
		UpcomingRelease bool   `json:"upcoming_release"`
		Assets          struct {
			Links []struct {
				Name           string `json:"name"`
				URL            string `json:"url"`
				DirectAssetUrl string `json:"direct_asset_url"`
			} `json:"links"`
		} `json:"assets"`
	}

	// Request latest releases
	err := fs.RetrieveJSON(endpoint, &releases)
	if err != nil {
		return "", err
	}

	// Create rule from search
	search = strings.ReplaceAll(search, "*", "(.+)")
	searchRegex := regexp.MustCompile("(?i)" + search)

	// Check for matching asset
	for _, release := range releases {
		for _, link := range release.Assets.Links {
			if searchRegex.MatchString(link.Name) {
				return link.URL, nil
			}
			if searchRegex.MatchString(link.URL) {
				return link.URL, nil
			}
		}
	}

	return "", fmt.Errorf("could not retrieve latest release asset")
}

// Return packaging source from release
func Release(domain string, projectId string, search string) *packaging.Source {
	return &packaging.Source{
		Format: packaging.FindFormat(search),
		Resolver: func() (string, error) {
			return GetAssetURL(domain, projectId, search)
		},
	}
}
