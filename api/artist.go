package api

import "net/http"

// The FullArtist struct describes a "Full" Artist object as defined by the Spotify Web API
type FullArtist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Followers    *Followers        `json:"followers"`
	Genres       []string          `json:"genres"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Images       []Image           `json:"images"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

// The SimpleArtist struct describes a "Simple" Artist object as defined by the Spotify Web API
type SimpleArtist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

// The FullArtistsPaged struct is a slice of FullArtist objects wrapped in a Spotify paging object
type FullArtistsPaged struct {
	Href     string       `json:"href"`
	Items    []FullArtist `json:"items"`
	Limit    int          `json:"limit"`
	Next     string       `json:"next"`
	Offset   int          `json:"offset"`
	Previous string       `json:"previous"`
	Total    int          `json:"total"`
}

// GetAlbumsForArtist returns a list of "Simple" Album objects in a paging object for the given artist
func GetAlbumsForArtist(artistID string) (pa SimpleAlbumsPaged, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"artists/"+artistID+"/albums", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pa)

	return pa, err
}

// GetNextAlbumsForArtist takes in the Next field from the paging objects returned from GetAlbumsForArtist and allows you to move forward through the albums
func GetNextAlbumsForArtist(url string) (pa SimpleAlbumsPaged, err error) {
	t := getAccessToken()

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pa)

	return pa, err
}
