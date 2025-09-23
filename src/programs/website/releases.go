package website

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get download URL from the latest release available
func GetDownloadURL(pageURL string, prefix string, search string) (string, error) {

	cli.Debug("Requesting %s\n", pageURL)

	// Make request to get HTML of target page
	res, err := http.Get(pageURL)
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

	// Isolate http links with unique lines
	// Process is required on minified HTML responses
	body = bytes.ReplaceAll(body, []byte("http"), []byte("\nhttp"))

	// Find first download link matching format
	search = strings.ReplaceAll(search, "*", "(.+)")
	searchRegex := regexp.MustCompile("(?i)" + search)
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

// Return packaging source from direct release link
func Link(url string) *packaging.Source {
	return &packaging.Source{
		Format: packaging.FindFormat(url),
		Resolver: func() (string, error) {
			return url, nil
		},
	}
}
