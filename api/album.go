package api

type Copyright struct {
	Text string `json:"text"`
	Type int    `json:"type"`
}

type FullAlbum struct {
	AlbumType            string              `json:"album"`
	Artists              []SimpleArtist      `json:"artists"`
	AvailableMarkets     []string            `json:"available_markets"`
	Copyrights           []Copyright         `json:"copyrights"`
	ExternalIDs          map[string]string   `json:"external_ids"`
	ExternalUrls         map[string]string   `json:"external_urls"`
	Genres               []string            `json:"genres"`
	Href                 string              `json:"href"`
	ID                   string              `json:"id"`
	Images               []Image             `json:"images"`
	Label                string              `json:"label"`
	Name                 string              `json:"name"`
	Popularity           int                 `json:"popularity"`
	ReleaseDate          string              `json:"release_date"`
	ReleaseDatePrecision string              `json:"release_date_precision"`
	Tracks               []SimpleTracksPaged `json:"tracks"`
	Type                 string              `json:"type"`
	URI                  string              `json:"uri"`
}

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
	Href    string        `json:"href"`
	Items   []SimpleAlbum `json:"items"`
	Limit   int           `json:"limit"`
	Next    string        `json:"next"`
	Cursors *Cursor       `json:"cursors"`
	Total   int           `json:"total"`
}
