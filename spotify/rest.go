package spotify

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

var client *http.Client

const (
	scheme  = "https"
	baseURL = "accounts.spotify.com/"
)

func init() {
	client = &http.Client{}
}

func BuildRequest(method, endpoint string, query url.Values) *http.Request {
	if query == nil {
		query = url.Values{}
	}

	u := &url.URL{
		Scheme:   scheme,
		Path:     baseURL + endpoint,
		RawQuery: query.Encode(),
	}

	r, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return r
}

func AddHeaders(r *http.Request, headers map[string]string) {
	for k, v := range headers {
		r.Header.Add(k, v)
	}
}

func MakeRequest(r *http.Request, d interface{}) error {
	res, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(d)

	return err
}
