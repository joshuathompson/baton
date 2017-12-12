package api

type PlaylistTrackLinks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type FullPlaylist struct {
	Collaborative bool              `json:"collaborative"`
	Description   string            `json:"description"`
	ExternalUrls  map[string]string `json:"external_urls"`
	Href          string            `json:"href"`
	ID            string            `json:"id"`
	Images        []Image           `json:"images"`
	Name          string            `json:"name"`
	Owner         *User             `json:"owner"`
	Public        bool              `json:"public"`
	SnapshotID    string            `json:"snapshot_id"`
	Tracks        []FullTracksPaged `json:"tracks"`
	Type          string            `json:"type"`
	URI           string            `json:"uri"`
}

type SimplePlaylist struct {
	Collaborative bool                `json:"collaborative"`
	ExternalUrls  map[string]string   `json:"external_urls"`
	Href          string              `json:"href"`
	ID            string              `json:"id"`
	Images        []Image             `json:"images"`
	Name          string              `json:"name"`
	Owner         *User               `json:"owner"`
	Public        bool                `json:"public"`
	SnapshotID    string              `json:"snapshot_id"`
	Tracks        *PlaylistTrackLinks `json:"tracks"`
	Type          string              `json:"type"`
	URI           string              `json:"uri"`
}

type SimplePlaylistsPaged struct {
	Href    string           `json:"href"`
	Items   []SimplePlaylist `json:"items"`
	Limit   int              `json:"limit"`
	Next    string           `json:"next"`
	Cursors *Cursor          `json:"cursors"`
	Total   int              `json:"total"`
}
