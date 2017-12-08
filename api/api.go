package api

import (
	"bytes"
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

var client *http.Client

func init() {
	client = &http.Client{}
}

func buildBody(j map[string]interface{}) io.Reader {
	b, err := json.Marshal(j)

	if err != nil {
		log.Fatal(err)
	}

	return bytes.NewBuffer(b)
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
		log.Fatal(err)
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
