package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/jroimartin/gocui"
)

// TrackTable implements the Table interface for "Full" Track objects as defined by the Spotify Web API
type TrackTable struct {
	tracks *api.FullTracksPaged
	title  string
}

// NewTrackTable creates a new instance of TrackTable
func NewTrackTable(fullTracksPaged *api.FullTracksPaged) *TrackTable {
	return &TrackTable{
		tracks: fullTracksPaged,
	}
}

func (t *TrackTable) getColumnWidths(maxX int) map[string]int {
	m := make(map[string]int)
	m["length"] = maxX / 8
	m["artist"] = maxX / 4
	m["album"] = maxX / 5
	m["popularity"] = maxX / 10
	m["name"] = maxX - m["track_number"] - m["length"] - m["artist"] - m["album"] - m["popularity"]

	return m
}

func (t *TrackTable) renderHeader(v *gocui.View, maxX int) {
	columnWidths := t.getColumnWidths(maxX)

	namesHeader := utils.LeftPaddedString("NAME", columnWidths["name"], 2)
	artistHeader := utils.LeftPaddedString("ARTIST", columnWidths["artist"], 2)
	albumHeader := utils.LeftPaddedString("ALBUM", columnWidths["album"], 2)
	lengthHeader := utils.LeftPaddedString("LENGTH", columnWidths["length"], 2)
	popularityHeader := utils.LeftPaddedString("POPULARITY", columnWidths["popularity"], 2)

	fmt.Fprintf(v, "\u001b[1m%s[0m\n", utils.LeftPaddedString("TRACKS", maxX, 2))
	fmt.Fprintf(v, "\u001b[1m%s %s %s %s %s\u001b[0m\n", namesHeader, artistHeader, albumHeader, lengthHeader, popularityHeader)
}

func (t *TrackTable) render(v *gocui.View, maxX int) {

	columnWidths := t.getColumnWidths(maxX)

	for _, track := range t.tracks.Items {
		name := utils.LeftPaddedString(track.Name, columnWidths["name"], 2)
		var artistNames []string
		for _, artist := range track.Artists {
			artistNames = append(artistNames, artist.Name)
		}
		artists := utils.LeftPaddedString(strings.Join(artistNames, ", "), columnWidths["artist"], 2)
		album := utils.LeftPaddedString(track.Album.Name, columnWidths["album"], 2)
		length := utils.LeftPaddedString(utils.MillisecondsToFormattedTime(track.DurationMs), columnWidths["length"], 2)
		popularity := utils.LeftPaddedString(strconv.Itoa(track.Popularity), columnWidths["popularity"], 2)

		fmt.Fprintf(v, "\n%s %s %s %s %s", name, artists, album, length, popularity)
	}
}

func (t *TrackTable) renderFooter(v *gocui.View, maxX int) {
	fmt.Fprintf(v, "\u001b[1m%s\u001b[0m\n", utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d tracks", len(t.tracks.Items), t.tracks.Total), maxX, 2))
}

func (t *TrackTable) getTableLength() int {
	return len(t.tracks.Items)
}

func (t *TrackTable) loadNextRecords() error {
	if t.tracks.Next != "" {
		res, err := api.GetNextSearchResults(t.tracks.Next)

		if err != nil {
			return err
		}

		nextTracks := res.Tracks

		t.tracks.Href = nextTracks.Href
		t.tracks.Offset = nextTracks.Offset
		t.tracks.Next = nextTracks.Next
		t.tracks.Previous = nextTracks.Previous
		t.tracks.Items = append(t.tracks.Items, nextTracks.Items...)
	}

	return nil
}

func (t *TrackTable) playSelected(selectedIndex int) (string, error) {
	track := t.tracks.Items[selectedIndex]

	var artistNames []string

	for _, artist := range track.Artists {
		artistNames = append(artistNames, artist.Name)
	}

	if track.Album != nil {
		playerOptions := api.PlayerOptions{
			ContextURI: track.Album.URI,
			Offset: &api.PlayerOffsetOptions{
				URI: track.URI,
			},
		}

		chosenItem := fmt.Sprintf("Now playing: '%s' by %s from the album %s\n", track.Name, strings.Join(artistNames, ", "), track.Album.Name)

		return chosenItem, api.StartPlayback(&playerOptions)
	}

	playerOptions := api.PlayerOptions{
		ContextURI: track.URI,
	}

	chosenItem := fmt.Sprintf("Now playing: '%s' by %s\n", track.Name, strings.Join(artistNames, ", "))

	return chosenItem, api.StartPlayback(&playerOptions)
}

func (t *TrackTable) newTableFromSelection(selectedIndex int) (Table, error) {
	track := t.tracks.Items[selectedIndex]
	if track.Album != nil {
		playerOptions := api.PlayerOptions{
			ContextURI: track.Album.URI,
			Offset: &api.PlayerOffsetOptions{
				URI: track.URI,
			},
		}
		return nil, api.StartPlayback(&playerOptions)
	}

	playerOptions := api.PlayerOptions{
		ContextURI: track.URI,
	}
	return nil, api.StartPlayback(&playerOptions)
}

func (t *TrackTable) handleSaveKey(selectedIndex int) error {
	track := t.tracks.Items[selectedIndex]
	err := api.SaveTrack(track.ID)
	if err != nil {
		return err
	}
	return nil
}
