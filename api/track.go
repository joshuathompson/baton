package api

import "time"

// The TrackLink struct describes a TrackLink object as defined by the Spotify Web API
type TrackLink struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

// The FullTrack struct describes a "Full" Track object as defined by the Spotify Web API
type FullTrack struct {
	Album            *SimpleAlbum      `json:"album"`
	Artists          []SimpleArtist    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMs       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalIDs      map[string]string `json:"external_ids"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       *TrackLink        `json:"linked_from"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	PreviewURL       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

// The SimpleTrack struct describes a "Simple" Track object as defined by the Spotify Web API
type SimpleTrack struct {
	Artists          []SimpleArtist    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMs       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       *TrackLink        `json:"linked_from"`
	Name             string            `json:"name"`
	PreviewURL       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

// The PlaylistTrack struct describes a Playlist Track object as defined by the Spotify Web API
type PlaylistTrack struct {
	AddedAt *time.Time `json:"added_at"`
	AddedBy *User      `json:"added_by"`
	IsLocal bool       `json:"is_local"`
	Track   FullTrack  `json:"track"`
}

// The SimpleTracksPaged struct is a slice of SimpleTrack objects wrapped in a Spotify paging object
type SimpleTracksPaged struct {
	Href     string        `json:"href"`
	Items    []SimpleTrack `json:"items"`
	Limit    int           `json:"limit"`
	Next     string        `json:"next"`
	Offset   int           `json:"offset"`
	Previous string        `json:"previous"`
	Total    int           `json:"total"`
}

// The FullTracksPaged struct is a slice of FullTrack objects wrapped in a Spotify paging object
type FullTracksPaged struct {
	Href     string      `json:"href"`
	Items    []FullTrack `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

// The PlaylistTracksPaged struct is a slice of PlaylistTrack objects wrapped in a Spotify paging object
type PlaylistTracksPaged struct {
	Href     string          `json:"href"`
	Items    []PlaylistTrack `json:"items"`
	Limit    int             `json:"limit"`
	Next     string          `json:"next"`
	Offset   int             `json:"offset"`
	Previous string          `json:"previous"`
	Total    int             `json:"total"`
}
