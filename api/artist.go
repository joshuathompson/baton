package api

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

type SimpleArtist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type FullArtistsPaged struct {
	Href     string       `json:"href"`
	Items    []FullArtist `json:"items"`
	Limit    int          `json:"limit"`
	Next     string       `json:"next"`
	Offset   int          `json:"offset"`
	Previous string       `json:"previous"`
	Total    int          `json:"total"`
}

func GetAlbumsForArtist(artistID string) (pa SimpleAlbumsPaged, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"artists/"+artistID+"/albums", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pa)

	return pa, err
}
