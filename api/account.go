package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

const (
	accountsURLBase = "accounts.spotify.com/"
)

// The Tokens struct describes a combination of the items returned from Spotify's API Authorization process as well as Baton-created fields to store in your config directory 
type Tokens struct {
	AccessToken    string        `json:"access_token"`
	TokenType      string        `json:"token_type"`
	ExpiresIn      time.Duration `json:"expires_in"`
	ExpirationDate time.Time     `json:"expiration_date"`
	ClientID       string        `json:"client_id"`
	ClientSecret   string        `json:"client_secret"`
	RefreshToken   string        `json:"refresh_token"`
	Scope          string        `json:"scope"`
}

// GetAuthorizationURL builds an Authorization URL for the user to navigate to from their ClientID
func GetAuthorizationURL(id string) string {
	v := url.Values{}
	v.Set("client_id", id)
	v.Set("response_type", "code")
	v.Set("redirect_uri", "http://localhost:15298/callback")
	v.Set("scope", "playlist-read-private user-top-read user-library-read user-library-modify user-read-currently-playing user-read-recently-played user-modify-playback-state user-read-playback-state user-follow-read playlist-read-collaborative")

	r := buildRequest("GET", accountsURLBase+"authorize", v, nil)
	return r.URL.String()
}

// AuthorizeWithCode completes the Authorization process and stores your refresh and current access tokens in the config directory for Baton
func AuthorizeWithCode(id, secret, code string) {
	var t Tokens
	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("code", code)
	v.Set("redirect_uri", "http://localhost:15298/callback")

	r := buildRequest("POST", accountsURLBase+"api/token", v, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.SetBasicAuth(id, secret)

	err := makeRequest(r, &t)

	if err != nil {
		log.Fatal(err)
	}

	t.ExpirationDate = time.Now().Add((t.ExpiresIn - 30) * time.Second)
	t.ClientID = id
	t.ClientSecret = secret

	writeTokensToConfig(t)
}

func getAccessToken() string {
	var t Tokens

	rt := viper.GetString("refresh_token")
	id := viper.GetString("client_id")
	secret := viper.GetString("client_secret")
	expiration := viper.GetTime("expiration_date")

	if rt == "" {
		log.Fatal("No valid token found, please run `baton auth` to authenticate")
	}

	if expiration.Before(time.Now()) {
		v := url.Values{}
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", rt)

		r := buildRequest("POST", accountsURLBase+"api/token", v, nil)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.SetBasicAuth(id, secret)

		err := makeRequest(r, &t)

		if err != nil {
			log.Fatal(err)
		}

		t.ExpirationDate = time.Now().Add((t.ExpiresIn - 30) * time.Second)
		t.ClientID = id
		t.ClientSecret = secret
		t.RefreshToken = rt

		writeTokensToConfig(t)

		return t.AccessToken
	}

	return viper.GetString("access_token")
}

func writeTokensToConfig(t Tokens) {
	ts, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(viper.ConfigFileUsed(), ts, 0666)

	if err != nil {
		log.Fatal(err)
	}
}
