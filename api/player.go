package api

import (
	"net/url"
	"strconv"
)

func GetDevices() (d devices, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/devices", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &d)

	return d, err
}

func GetCurrentPlaybackInformation() (ctx currentlyPlayingContext, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &ctx)

	return ctx, err
}

func GetRecentlyPlayedTracks() (rpt recentlyPlayedTracks, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/recently-played", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &rpt)

	return rpt, err
}

func GetCurrentlyPlayingTrack() (c currentlyPlayingTrack, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/currently-playing", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &c)

	return c, err
}

func SetVolume(volume string) error {
	v := url.Values{}
	v.Add("volume_percent", volume)

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/volume", v)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func GetCurrentlyPlaying() error {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/currently-playing", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func PausePlayback() error {
	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/pause", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func SkipToTrack() error {
	t := getAccessToken()

	r := buildRequest("POST", apiURLBase+"me/player/next", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func SkipToPrevious() error {
	t := getAccessToken()

	r := buildRequest("POST", apiURLBase+"me/player/previous", nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func ToggleShuffle(state bool) error {
	v := url.Values{}
	v.Add("state", strconv.FormatBool(state))

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/shuffle", v)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}
