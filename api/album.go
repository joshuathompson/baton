package api

import (
	"net/http"
	"time"
)

// The SimpleAlbum struct describes a "Simple" Album object as defined by the Spotify Web API
type SimpleAlbum struct {
	AlbumType        string            `json:"album"`
	Artists          []SimpleArtist    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	Images           []Image           `json:"images"`
	Name             string            `json:"name"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

// The SavedAlbum struct describes a Saved Track object as defined by the Spotify Web API
type SavedAlbum struct {
	AddedAt *time.Time  `json:"added_at"`
	Album   SimpleAlbum `json:"album"`
}

// The SimpleAlbumsPaged struct is a slice of SimpleAlbum objects wrapped in a Spotify paging object
type SimpleAlbumsPaged struct {
	Href     string        `json:"href"`
	Items    []SimpleAlbum `json:"items"`
	Limit    int           `json:"limit"`
	Next     string        `json:"next"`
	Offset   int           `json:"offset"`
	Previous string        `json:"previous"`
	Total    int           `json:"total"`
}

// The SavedAlbumsPaged struct is a slice of SavedAlbum objects wrapped in a Spotify paging object
type SavedAlbumsPaged struct {
	Href     string       `json:"href"`
	Items    []SavedAlbum `json:"items"`
	Limit    int          `json:"limit"`
	Next     string       `json:"next"`
	Offset   int          `json:"offset"`
	Previous string       `json:"previous"`
	Total    int          `json:"total"`
}

// GetTracksForAlbum returns a list of "Simple" Track objects in a paging object for the given album
func GetTracksForAlbum(albumID string) (pt SimpleTracksPaged, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"albums/"+albumID+"/tracks", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pt)

	return pt, err
}

// GetNextTracksForAlbum takes in the Next field from the paging objects returned from GetTracksForAlbum and allows you to move forward through the tracks
func GetNextTracksForAlbum(url string) (pt SimpleTracksPaged, err error) {
	t := getAccessToken()

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pt)

	return pt, err
}
