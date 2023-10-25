package api

import "net/url"

// SearchByTermResult struct
type SearchByTermResult struct {
	Success bool            `json:"success"`
	Errors  []string        `json:"errors"`
	Data    []DetailsResult `json:"data"`
}

// Search game or application by term or name
func SearchByTerm(term string) (*SearchByTermResult, error) {

	endpoint := baseURL + "/search/autocomplete/" + url.QueryEscape(term)
	result := SearchByTermResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}
