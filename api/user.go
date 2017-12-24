package api

// The User struct describes a User object as defined by the Spotify Web API
type User struct {
	DisplayName  string            `json:"display_name"`
	ExternalUrls map[string]string `json:"external_urls"`
	Followers    *Followers        `json:"followers"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Images       []Image           `json:"images"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}