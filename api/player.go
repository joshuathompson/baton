package api

import (
	"fmt"
	"net/url"
)

func SetVolume(volume string) error {
	v := url.Values{}
	v.Add("volume_percent", volume)

	fmt.Println(getAccessToken())
	r := buildRequest("PUT", apiBase+"me/player/volume", v)
	r.Header.Add("Authorization", "Bearer "+getAccessToken())

	err := makeRequest(r, nil)

	return err
}
