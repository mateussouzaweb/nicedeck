package website

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get download URL from the latest release available
func GetDownloadURL(pageURL string, prefix string, search string) (string, error) {

	// Make request to get HTML of target page
	resp, err := http.Get(pageURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Find first download link matching format
	search = strings.ReplaceAll(search, "*", "(.+)")
	searchRegex := regexp.MustCompile(search)
	matches := searchRegex.FindSubmatch(body)

	url := ""
	if len(matches) > 0 {
		url = string(matches[0])
	}

	// Append prefix when necessary
	if prefix != "" && url != "" {
		url = prefix + url
	}

	// Return final result
	if url != "" {
		return url, nil
	}

	return "", fmt.Errorf("could not retrieve latest release")
}

// Return packaging source from release
func Release(pageURL string, prefix string, search string) *packaging.Source {
	return &packaging.Source{
		Format: packaging.FindFormat(search),
		Resolver: func() (string, error) {
			return GetDownloadURL(pageURL, prefix, search)
		},
	}
}
