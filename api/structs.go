package api

import (
	"time"
)

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
