package api

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/google/go-querystring/query"
)

type PlayerContext struct {
	Type         string            `json:"type"`
	Href         string            `json:"href"`
	ExternalUrls map[string]string `json:"external_urls"`
	URI          string            `json:"uri"`
}

type PlayerState struct {
	Device       *Device        `json:"device"`
	RepeatState  string         `json:"repeat_state"`
	ShuffleState bool           `json:"shuffle_state"`
	Context      *PlayerContext `json:"context"`
	Timestamp    int            `json:"timestamp"`
	ProgressMs   int            `json:"progress_ms"`
	IsPlaying    bool           `json:"is_playing"`
	Item         *FullTrack     `json:"item"`
}

type Options struct {
	DeviceID string `json:"device_id,omitempty" url:"device_id,omitempty"`
	Market   string `json:"market,omitempty" url:"market,omitempty"`
}

type PlayerOptions struct {
	DeviceID   string               `json:"device_id,omitempty" url:"device_id,omitempty"`
	ContextURI string               `json:"context_uri,omitempty" url:"context_uri,omitempty"`
	URIs       []string             `json:"uris,omitempty" url:"uris,omitempty"`
	Offset     *PlayerOffsetOptions `json:"offset,omitempty" url:"offset,omitempty"`
}

type PlayerOffsetOptions struct {
	Position int    `json:"position,omitempty" url:"position,omitempty"`
	URI      string `json:"uri,omitempty" url:"uri,omitempty"`
}

func GetPlayerState(opts *Options) (ps PlayerState, err error) {
	v, err := query.Values(opts)

	if err != nil {
		return ps, err
	}

	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &ps)

	return ps, err
}

func SetRepeatMode(state string, opts *Options) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	v.Add("state", state)

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/repeat", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

func SetVolume(vol int, opts *Options) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	v.Add("volume_percent", strconv.Itoa(vol))

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/volume", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

func PausePlayback(opts *Options) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/pause", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

func SeekToPosition(pos int, opts *Options) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	v.Add("position_ms", strconv.Itoa(pos))

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/seek", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

func StartPlayback(opts *PlayerOptions) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	opts.DeviceID = ""
	j, err := json.Marshal(opts)

	if err != nil {
		log.Fatal(err)
	}

	b := bytes.NewBuffer(j)

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/play", v, b)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

func SkipToNext(opts *Options) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	t := getAccessToken()

	r := buildRequest("POST", apiURLBase+"me/player/next", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

func SkipToPrevious(opts *Options) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	t := getAccessToken()

	r := buildRequest("POST", apiURLBase+"me/player/previous", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

func ToggleShuffle(state bool, opts *Options) error {
	v, err := query.Values(opts)

	if err != nil {
		return err
	}

	v.Add("state", strconv.FormatBool(state))

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/player/shuffle", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}
