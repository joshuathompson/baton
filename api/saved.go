package api

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

// GetSavedTracks returns a list of all the songs the user has saved
func GetSavedTracks(opts *SearchOptions) (result *SavedTracksPaged, err error) {
	v, err := query.Values(opts)

	if err != nil {
		return result, err
	}

	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/tracks", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &result)

	return result, err
}

// GetNextSavedTracks takes in the Next fields from the paging objects returned from Saved and moves forward through the results
func GetNextSavedTracks(url string) (sr *SavedTracksPaged, err error) {
	t := getAccessToken()

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &sr)

	return sr, err
}
