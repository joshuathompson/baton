package api

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

// The SearchResults struct describes the potential results of any search against the Spotify API
type SearchResults struct {
	Artists   *FullArtistsPaged     `json:"artists"`
	Albums    *SimpleAlbumsPaged    `json:"albums"`
	Tracks    *FullTracksPaged      `json:"tracks"`
	Playlists *SimplePlaylistsPaged `json:"playlists"`
}

// The SearchOptions struct describes the possible optional arguments for the Search function
type SearchOptions struct {
	Market string `json:"market,omitempty" url:"market,omitempty"`
	Limit  int    `json:"limit,omitempty" url:"limit,omitempty"`
	Offset int    `json:"offset,omitempty" url:"offset,omitempty"`
}

// Search queries the Spotify API based on the given query and options and returns the results wrapped in paging objects
func Search(q, types string, opts *SearchOptions) (sr SearchResults, err error) {
	v, err := query.Values(opts)

	if err != nil {
		return sr, err
	}

	v.Add("q", q)
	v.Add("type", types)

	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"search", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &sr)

	return sr, err
}

// GetNextSearchResults takes in the Next fields from the paging objects returned from Search and allows you to move forward through the results
func GetNextSearchResults(url string) (sr *SearchResults, err error) {
	t := getAccessToken()

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &sr)

	return sr, err
}
