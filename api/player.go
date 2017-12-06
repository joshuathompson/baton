package api

import (
	"net/url"
)

func SetVolume(volume string) error {
	v := url.Values{}
	v.Add("volume_percent", volume)

	r := buildRequest("PUT", apiURLBase+"me/player/volume", v)
	r.Header.Add("Authorization", "Bearer "+getAccessToken())

	err := makeRequest(r, nil)

	return err
}
