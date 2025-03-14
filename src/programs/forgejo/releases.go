package forgejo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get asset direct download URL from the latest release available
func GetAssetURL(domain string, repository string, search string) (string, error) {

	domain = strings.Trim(domain, "/")
	repository = strings.ReplaceAll(repository, domain, "")
	repository = strings.Trim(repository, "/")

	// Request latest releases
	endpoint := fmt.Sprintf("%s/api/v1/repos/%s/releases", domain, repository)
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Decode response into struct
	var releases []struct {
		Name   string `json:"name"`
		Assets []struct {
			Name        string `json:"name"`
			Type        string `json:"type"`
			DownloadURL string `json:"browser_download_url"`
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
			if searchRegex.MatchString(asset.Name) {
				return asset.DownloadURL, nil
			}
		}
	}

	return "", fmt.Errorf("could not retrieve latest release asset")
}

// Return packaging source from release
func Release(domain string, repository string, search string) *packaging.Source {
	return &packaging.Source{
		Format: packaging.FindFormat(search),
		Resolver: func() (string, error) {
			return GetAssetURL(domain, repository, search)
		},
	}
}
