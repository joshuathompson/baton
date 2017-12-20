package api

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

type SimpleAlbumsPaged struct {
	Href     string        `json:"href"`
	Items    []SimpleAlbum `json:"items"`
	Limit    int           `json:"limit"`
	Next     string        `json:"next"`
	Offset   int           `json:"offset"`
	Previous string        `json:"previous"`
	Total    int           `json:"total"`
}

func GetTracksForAlbum(albumID string) (pt SimpleTracksPaged, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"albums/"+albumID+"/tracks", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pt)

	return pt, err
}
