package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	apiURLBase = "api.spotify.com/v1/"
)

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Cursor struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

var client *http.Client

func init() {
	client = &http.Client{}
}

func buildRequest(method, path string, query url.Values, b io.Reader) *http.Request {
	if query == nil {
		query = url.Values{}
	}

	u := &url.URL{
		Scheme:   "https",
		Path:     path,
		RawQuery: query.Encode(),
	}

	r, err := http.NewRequest(method, u.String(), b)

	if err != nil {
		log.Fatal(err)
	}

	return r
}

func makeRequest(r *http.Request, d interface{}) error {
	res, err := client.Do(r)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusCreated &&
		res.StatusCode != http.StatusAccepted &&
		res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}

	if d != nil {
		return json.NewDecoder(res.Body).Decode(d)
	}

	return nil
}
