package api

import (
	"fmt"
	"net/url"
	"strings"
)

// ImagesParams struct
type ImagesParams struct {
	Styles     []string `query:"styles"`
	Dimensions []string `query:"dimensions"`
	Mimes      []string `query:"mimes"`
	Types      []string `query:"types"`
	Nsfw       string   `query:"nsfw"`
	Humor      string   `query:"humor"`
	Epilepsy   string   `query:"epilepsy"`
	OneOfTag   []string `query:"oneoftag"`
	Page       int      `query:"page"`
}

// Retrieve formated query params
func (i *ImagesParams) getQueryParams() string {

	query := url.Values{}

	if len(i.Styles) > 0 {
		query.Add("styles", strings.Join(i.Styles, ","))
	}
	if len(i.Dimensions) > 0 {
		query.Add("dimensions", strings.Join(i.Dimensions, ","))
	}
	if len(i.Mimes) > 0 {
		query.Add("mimes", strings.Join(i.Mimes, ","))
	}
	if len(i.Types) > 0 {
		query.Add("types", strings.Join(i.Types, ","))
	}
	if i.Nsfw != "" {
		query.Add("nsfw", i.Nsfw)
	}
	if i.Humor != "" {
		query.Add("humor", i.Humor)
	}
	if i.Epilepsy != "" {
		query.Add("epilepsy", i.Epilepsy)
	}
	if len(i.OneOfTag) > 0 {
		query.Add("oneoftag", strings.Join(i.OneOfTag, ","))
	}
	if i.Page > 0 {
		query.Add("page", fmt.Sprintf("%d", i.Page))
	}

	return query.Encode()
}

// ImageAuthor struct
type ImageAuthor struct {
	Name    string `json:"name"`
	Steam64 string `json:"steam64"`
	Avatar  string `json:"avatar"`
}

// ImageResult struct
type ImageResult struct {
	ID     int64       `json:"id"`
	Score  int64       `json:"score"`
	Style  string      `json:"style"`
	URL    string      `json:"url"`
	Thumb  string      `json:"thumb"`
	Tags   []string    `json:"tags"`
	Author ImageAuthor `json:"author"`
}

// ImagesByIdResult struct
type ImagesByIdResult struct {
	Success bool          `json:"success"`
	Data    []ImageResult `json:"data"`
}

// ImagesByPlatformAndIdResult struct
type ImagesByPlatformAndIdResult struct {
	Success bool          `json:"success"`
	Data    []ImageResult `json:"data"`
}

// ImagesByPlatformAndMultipleIdsResult struct
type ImagesByPlatformAndMultipleIdsResult struct {
	Success bool `json:"success"`
	Data    []struct {
		Success bool          `json:"success"`
		Status  int64         `json:"status"`
		Data    []ImageResult `json:"data"`
	}
}

// Retrieve images for game or application by ID
func GetImagesById(format string, id string, params *ImagesParams) (*ImagesByIdResult, error) {

	endpoint := getImagesBaseUrl(format) + "/game/" + id + "?" + params.getQueryParams()
	result := ImagesByIdResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}

// Retrieve images for game or application by platform and ID
func GetImagesByPlatformAndId(format string, platform string, id string, params *ImagesParams) (*ImagesByPlatformAndIdResult, error) {

	endpoint := getImagesBaseUrl(format) + "/" + platform + "/" + id + "?" + params.getQueryParams()
	result := ImagesByPlatformAndIdResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}

// Retrieve images for games or applications by platform and given IDs
func GetImagesByPlatformAndMultipleIds(format string, platform string, ids []string, params *ImagesParams) (*ImagesByPlatformAndMultipleIdsResult, error) {

	endpoint := getImagesBaseUrl(format) + "/" + platform + "/" + strings.Join(ids, ",") + "?" + params.getQueryParams()
	result := ImagesByPlatformAndMultipleIdsResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}

// Retrieve base URL for images targeted for image format
func getImagesBaseUrl(format string) string {

	// Options are: banners, covers, grids, heroes, logos, icons
	// Fallback to grids if could not detect the format
	switch format {
	case "banners":
	case "banner":
	case "covers":
	case "cover":
	case "grids":
	case "grid":
		return baseUrl + "/grids"
	case "heroes":
	case "hero":
		return baseUrl + "/heroes"
	case "logos":
	case "logo":
		return baseUrl + "/logos"
	case "icons":
	case "icon":
		return baseUrl + "/icons"
	default:
	}

	return baseUrl + "/grids"
}
