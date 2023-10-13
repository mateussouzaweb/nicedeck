package api

import "net/url"

// SearchByTermResult struct
type SearchByTermResult struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
	Data    []struct {
		Success bool `json:"success"`
		Data    DetailsResult
	} `json:"data"`
}

// Search game or application by term or name
func SearchByTerm(term string) (*SearchByTermResult, error) {

	endpoint := baseUrl + "/search/autocomplete/" + url.QueryEscape(term)
	result := SearchByTermResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}
