package api

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/google/go-querystring/query"
)

// The PlayerContext struct describes the current context of what is playing on the active device.  ex. The context could be an "album" which can then be derived from the URI
// The can be used to determine that "One More Time by Daft Punk" that the user is listening to is actually in a user created playlist and that the context is NOT the album versiom, for instance 
type PlayerContext struct {
	Type         string            `json:"type"`
	Href         string            `json:"href"`
	ExternalUrls map[string]string `json:"external_urls"`
	URI          string            `json:"uri"`
}

// The PlayerState struct describes the current playback state of Spotify
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

// The Options struct describes options that can be used by the majority of API endpoints
type Options struct {
	DeviceID string `json:"device_id,omitempty" url:"device_id,omitempty"`
	Market   string `json:"market,omitempty" url:"market,omitempty"`
}

// The PlayerOptions struct describes options that are specific to the /me/player endpoints
type PlayerOptions struct {
	DeviceID   string               `json:"device_id,omitempty" url:"device_id,omitempty"`
	ContextURI string               `json:"context_uri,omitempty" url:"context_uri,omitempty"`
	URIs       []string             `json:"uris,omitempty" url:"uris,omitempty"`
	Offset     *PlayerOffsetOptions `json:"offset,omitempty" url:"offset,omitempty"`
}

// The PlayerOffsetOptions describes how to set the offset within a context when controlling playback
// For example, you can use Position to specify track number within an album OR you can use the URI to point to that same track directly  
type PlayerOffsetOptions struct {
	Position int    `json:"position,omitempty" url:"position,omitempty"`
	URI      string `json:"uri,omitempty" url:"uri,omitempty"`
}


// GetDevices returns a list of available playback devices
func GetDevices() (d []Device, err error) {
	var ds Devices
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/player/devices", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &ds)

	return ds.Devices, err
}


// GetPlayerState returns the active device, whether the player is paused, progress of current song, and other playback information
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

// SetRepeatMode allows you to set the Repeat Mode of the current device
// Allowed values are track, context, and off
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

// SetVolume allows you to control the volume percentage of the device
// Allowed values are 0-100
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

// PausePlayback pauses playback on the current device
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

// SeekToPosition skips to a position defined in seconds for the current playback
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

// StartPlayback can resume playback or change playback to a new URI/context
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

// SkipToNext skips to the next song within the current context
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

// SkipToPrevious skips to the previous song within the current context
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

// ToggleShuffle toggles the shuffle state on/off
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
