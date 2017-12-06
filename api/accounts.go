package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

func GetAuthorizationURL(id string) string {
	v := url.Values{}
	v.Set("client_id", id)
	v.Set("response_type", "code")
	v.Set("redirect_uri", "http://localhost:15298/callback")
	v.Set("scope", "playlist-read-private user-top-read user-library-read user-read-currently-playing user-read-recently-played user-modify-playback-state user-read-playback-state user-follow-read playlist-read-collaborative")

	r := buildRequest("GET", accountsBase+"authorize", v)
	return r.URL.String()
}

func AuthorizeWithCode(id, secret, code string) {
	var t tokens
	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("code", code)
	v.Set("redirect_uri", "http://localhost:15298/callback")

	r := buildRequest("POST", accountsBase+"api/token", v)
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
	var t tokens

	rt := viper.GetString("refresh_token")
	id := viper.GetString("client_id")
	secret := viper.GetString("client_secret")
	expiration := viper.GetTime("expiration_date")

	if expiration.Before(time.Now()) {
		v := url.Values{}
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", rt)

		r := buildRequest("POST", accountsBase+"api/token", v)
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

func writeTokensToConfig(t tokens) {
	ts, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(viper.ConfigFileUsed(), ts, 0666)

	if err != nil {
		log.Fatal(err)
	}
}
