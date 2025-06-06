package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get asset direct download URL from the latest release available
func GetAssetURL(domain string, projectId string, search string) (string, error) {

	domain = strings.Trim(domain, "/")

	// Request latest releases
	endpoint := fmt.Sprintf("%s/api/v4/projects/%s/releases", domain, projectId)
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer func() {
		errors.Join(err, res.Body.Close())
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Decode response into struct
	var releases []struct {
		Name            string `json:"name"`
		UpcomingRelease bool   `json:"upcoming_release"`
		Assets          []struct {
			Links []struct {
				Name           string `json:"name"`
				URL            string `json:"url"`
				DirectAssetUrl string `json:"direct_asset_url"`
			} `json:"links"`
		} `json:"assets"`
	}

	err = json.Unmarshal(body, &releases)
	if err != nil {
		return "", err
	}

	// Create rule from search
	search = strings.ReplaceAll(search, "*", "(.+)")
	searchRegex := regexp.MustCompile(search)

	// Check for matching asset
	for _, release := range releases {
		for _, asset := range release.Assets {
			for _, link := range asset.Links {
				if searchRegex.MatchString(link.Name) {
					return link.URL, nil
				}
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
