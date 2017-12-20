package api

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type SearchResults struct {
	Artists   *FullArtistsPaged     `json:"artists"`
	Albums    *SimpleAlbumsPaged    `json:"albums"`
	Tracks    *FullTracksPaged      `json:"tracks"`
	Playlists *SimplePlaylistsPaged `json:"playlists"`
}

type SearchOptions struct {
	Market string `json:"market,omitempty" url:"market,omitempty"`
	Limit  int    `json:"limit,omitempty" url:"limit,omitempty"`
	Offset int    `json:"offset,omitempty" url:"offset,omitempty"`
}

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

func GetNextSearchResults(url string) (sr *SearchResults, err error) {
	t := getAccessToken()

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &sr)

	return sr, err
}
