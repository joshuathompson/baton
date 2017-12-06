package spotify

import (
	"time"
)

type Tokens struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type"`
	ExpiresIn    time.Duration `json:"expires_in"`
	ExpiresDate  time.Time     `json:"expires_date"`
	RefreshToken string        `json:"refresh_token"`
	Scope        string        `json:"scope"`
}
