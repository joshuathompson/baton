package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/jroimartin/gocui"
)

type SimpleTrackTable struct {
	tracks *api.SimpleTracksPaged
	album  *api.SimpleAlbum
}

func NewSimpleTrackTable(simpleTracksPaged *api.SimpleTracksPaged, album *api.SimpleAlbum) *SimpleTrackTable {
	return &SimpleTrackTable{
		album:  album,
		tracks: simpleTracksPaged,
	}
}

func (t *SimpleTrackTable) getColumnWidths(maxX int) map[string]int {
	m := make(map[string]int)
	m["track_number"] = maxX / 10
	m["length"] = maxX / 8
	m["artist"] = maxX / 4
	m["album"] = maxX / 5
	m["name"] = maxX - m["track_number"] - m["length"] - m["artist"] - m["album"]

	return m
}

func (t *SimpleTrackTable) renderHeader(v *gocui.View, maxX int) {
	columnWidths := t.getColumnWidths(maxX)

	trackNumberHeader := utils.LeftPaddedString("#", columnWidths["track_number"], 2)
	namesHeader := utils.LeftPaddedString("NAME", columnWidths["name"], 2)
	artistHeader := utils.LeftPaddedString("ARTIST", columnWidths["artist"], 2)
	albumHeader := utils.LeftPaddedString("ALBUM", columnWidths["album"], 2)
	lengthHeader := utils.LeftPaddedString("LENGTH", columnWidths["length"], 2)

	loadedLength := maxX / 3
	loadedHeader := utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d tracks", len(t.tracks.Items), t.tracks.Total), loadedLength, 2)
	titleLength := maxX - loadedLength

	fmt.Fprintf(v, "\u001b[1m%s %s[0m\n\n", utils.LeftPaddedString("TRACKS", titleLength, 2), loadedHeader)
	fmt.Fprintf(v, "\u001b[1m%s %s %s %s %s\u001b[0m\n", trackNumberHeader, namesHeader, artistHeader, albumHeader, lengthHeader)
}

func (t *SimpleTrackTable) render(v *gocui.View, maxX int) {
	columnWidths := t.getColumnWidths(maxX)

	for _, track := range t.tracks.Items {
		trackNumber := utils.LeftPaddedString(strconv.Itoa(track.TrackNumber), columnWidths["track_number"], 2)
		name := utils.LeftPaddedString(track.Name, columnWidths["name"], 2)
		var artistNames []string
		for _, artist := range t.album.Artists {
			artistNames = append(artistNames, artist.Name)
		}
		artists := utils.LeftPaddedString(strings.Join(artistNames, ", "), columnWidths["artist"], 2)
		album := utils.LeftPaddedString(t.album.Name, columnWidths["album"], 2)
		length := utils.LeftPaddedString(utils.MillisecondsToFormattedTime(track.DurationMs), columnWidths["length"], 2)

		fmt.Fprintf(v, "\n%s %s %s %s %s", trackNumber, name, artists, album, length)
	}
}

func (t *SimpleTrackTable) getTableLength() int {
	return len(t.tracks.Items)
}

func (t *SimpleTrackTable) loadNextRecords() error {
	if t.tracks.Next != "" {
		nextTracks, err := api.GetNextTracksForAlbum(t.tracks.Next)

		if err != nil {
			return err
		}

		t.tracks.Href = nextTracks.Href
		t.tracks.Offset = nextTracks.Offset
		t.tracks.Next = nextTracks.Next
		t.tracks.Previous = nextTracks.Previous
		t.tracks.Items = append(t.tracks.Items, nextTracks.Items...)
	}

	return nil
}

func (t *SimpleTrackTable) playSelected(selectedIndex int) error {
	track := t.tracks.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: t.album.URI,
		Offset: &api.PlayerOffsetOptions{
			URI: track.URI,
		},
	}
	return api.StartPlayback(&playerOptions)
}

func (t *SimpleTrackTable) newTableFromSelection(selectedIndex int) (Table, error) {
	track := t.tracks.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: t.album.URI,
		Offset: &api.PlayerOffsetOptions{
			URI: track.URI,
		},
	}
	err := api.StartPlayback(&playerOptions)
	return nil, err
}
