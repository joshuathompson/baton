package api

import (
	"net/url"
	"strconv"
)

func GetDevices() (d devices, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/devices", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &d)

	return d, err
}

func GetCurrentPlaybackInformation() (ctx currentlyPlayingContext, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &ctx)

	return ctx, err
}

func GetRecentlyPlayedTracks() (rpt recentlyPlayedTracks, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/recently-played", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &rpt)

	return rpt, err
}

func GetCurrentlyPlayingTrack() (c currentlyPlayingTrack, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/currently-playing", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &c)

	return c, err
}

func SetRepeatMode(mode string) error {
	v := url.Values{}
	v.Add("state", mode)

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/repeat", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func SetVolume(volume int) error {
	v := url.Values{}
	v.Add("volume_percent", strconv.Itoa(volume))

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/volume", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func PausePlayback() error {
	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/pause", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func SeekToPosition(pos int) error {
	v := url.Values{}
	v.Add("position_ms", strconv.Itoa(pos))

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/seek", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func StartPlayback() error {
	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/play", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func SkipToNext() error {
	t := getAccessToken()

	r := buildRequest("POST", apiURLBase+"me/player/next", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func SkipToPrevious() error {
	t := getAccessToken()

	r := buildRequest("POST", apiURLBase+"me/player/previous", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}

func ToggleShuffle(state bool) error {
	v := url.Values{}
	v.Add("state", strconv.FormatBool(state))

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/shuffle", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err := makeRequest(r, nil)

	return err
}
