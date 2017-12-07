package api

import "time"

type tokens struct {
	AccessToken    string        `json:"access_token"`
	TokenType      string        `json:"token_type"`
	ExpiresIn      time.Duration `json:"expires_in"`
	ExpirationDate time.Time     `json:"expiration_date"`
	ClientID       string        `json:"client_id"`
	ClientSecret   string        `json:"client_secret"`
	RefreshToken   string        `json:"refresh_token"`
	Scope          string        `json:"scope"`
}

type image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type copyright struct {
	Text string `json:"text"`
	Type int    `json:"type"`
}

type cursor struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

type album struct {
	AlbumType            string            `json:"album"`
	Artists              []artistSimple    `json:"artists"`
	AvailableMarkets     []string          `json:"available_markets"`
	Copyrights           []copyright       `json:"copyrights"`
	ExternalIDs          map[string]string `json:"external_ids"`
	ExternalUrls         map[string]string `json:"external_urls"`
	Genres               []string          `json:"genres"`
	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []image           `json:"images"`
	Label                string            `json:"label"`
	Name                 string            `json:"name"`
	Popularity           int               `json:"popularity"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Tracks               []tracksInCursor  `json:"tracks"`
	Type                 string            `json:"type"`
	URI                  string            `json:"uri"`
}

type albumSimple struct {
	AlbumType        string            `json:"album"`
	Artists          []artistSimple    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	Images           []image           `json:"images"`
	Name             string            `json:"name"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

type artist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Followers    followers         `json:"followers"`
	Genres       []string          `json:"genres"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Images       []image           `json:"images"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type artistSimple struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type trackLink struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type track struct {
	Album            albumSimple       `json:"album"`
	Artists          []artistSimple    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMs       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalIDs      map[string]string `json:"external_ids"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       trackLink         `json:"linked_from"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	PreviewURL       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

type trackSimple struct {
	Artists          []artistSimple    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMs       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       trackLink         `json:"linked_from"`
	Name             string            `json:"name"`
	PreviewURL       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

type tracksInCursor struct {
	Href    string        `json:"href"`
	Items   []trackSimple `json:"items"`
	Limit   int           `json:"limit"`
	Next    string        `json:"next"`
	Cursors cursor        `json:"cursors"`
	Total   int           `json:"total"`
}

type playHistory struct {
	Track    trackSimple `json:"track"`
	PlayedAt time.Time   `json:"played_at"`
	Context  context     `json:"context"`
}

type recentlyPlayedTracks struct {
	Href    string        `json:"href"`
	Items   []playHistory `json:"items"`
	Limit   int           `json:"limit"`
	Next    string        `json:"next"`
	Cursors cursor        `json:"cursors"`
	Total   int           `json:"total"`
}

type device struct {
	ID            string `json:"id"`
	IsActive      bool   `json:"is_active"`
	IsRestricted  bool   `json:"is_restricted"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	VolumePercent int    `json:"volume_percent"`
}

type devices struct {
	Devices []device `json:"devices"`
}

type context struct {
	Type         string            `json:"type"`
	Href         string            `json:"href"`
	ExternalUrls map[string]string `json:"external_urls"`
	URI          string            `json:"uri"`
}

type currentlyPlayingContext struct {
	Device       device  `json:"device"`
	RepeatState  string  `json:"repeat_state"`
	ShuffleState bool    `json:"shuffle_state"`
	Context      context `json:"context"`
	Timestamp    int     `json:"timestamp"`
	ProgressMs   int     `json:"progress_ms"`
	IsPlaying    bool    `json:"is_playing"`
	Item         track   `json:"item"`
}

type currentlyPlayingTrack struct {
	Context    context `json:"context"`
	Timestamp  int     `json:"timestamp"`
	ProgressMs int     `json:"progress_ms"`
	IsPlaying  bool    `json:"is_playing"`
	Item       track   `json:"item"`
}
