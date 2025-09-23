package github

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get asset direct download URL from the latest release available
func GetAssetURL(repository string, search string) (string, error) {

	repository = strings.ReplaceAll(repository, "https://github.com/", "")
	repository = strings.Trim(repository, "/")

	domain := "https://api.github.com"
	endpoint := fmt.Sprintf("%s/repos/%s/releases", domain, repository)

	// Response struct
	var releases []struct {
		Name       string `json:"name"`
		PreRelease bool   `json:"prerelease"`
		Assets     []struct {
			Name        string `json:"name"`
			ContentType string `json:"content_type"`
			URL         string `json:"url"`
			DownloadURL string `json:"browser_download_url"`
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
		for _, asset := range release.Assets {
			if searchRegex.MatchString(asset.Name) {
				return asset.DownloadURL, nil
			}
		}
	}

	return "", fmt.Errorf("could not retrieve latest release asset")
}

// Return packaging source from release
func Release(repository string, search string) *packaging.Source {
	return &packaging.Source{
		Format: packaging.FindFormat(search),
		Resolver: func() (string, error) {
			return GetAssetURL(repository, search)
		},
	}
}
