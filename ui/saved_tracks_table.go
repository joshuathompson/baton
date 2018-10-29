package ui

import (
	"fmt"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/jroimartin/gocui"
)

// SavedTrackTable implements the Table interface for "Saved" Track objects as defined by the Spotify Web API
type SavedTrackTable struct {
	tracks *api.SavedTracksPaged
	title  string
}

// NewSavedTrackTable creates a new instance of SavedTrackTable
func NewSavedTrackTable(savedTracksPage *api.SavedTracksPaged) *SavedTrackTable {
	return &SavedTrackTable{
		tracks: savedTracksPage,
	}
}

func (t *SavedTrackTable) getColumnWidths(maxX int) map[string]int {
	m := make(map[string]int)
	m["length"] = maxX / 8
	m["artist"] = maxX / 4
	m["album"] = maxX / 5
	m["popularity"] = maxX / 10
	m["addedOn"] = maxX / 6
	m["name"] = maxX - m["track_number"] - m["length"] - m["artist"] - m["album"] - m["addedOn"]

	return m
}

func (t *SavedTrackTable) renderHeader(v *gocui.View, maxX int) {
	columnWidths := t.getColumnWidths(maxX)

	namesHeader := utils.LeftPaddedString("NAME", columnWidths["name"], 2)
	artistHeader := utils.LeftPaddedString("ARTIST", columnWidths["artist"], 2)
	albumHeader := utils.LeftPaddedString("ALBUM", columnWidths["album"], 2)
	lengthHeader := utils.LeftPaddedString("LENGTH", columnWidths["length"], 2)

	addedOnHeader := utils.LeftPaddedString("ADDED ON", columnWidths["addedOn"], 2)

	fmt.Fprintf(v, "\u001b[1m%s[0m\n", utils.LeftPaddedString("TRACKS", maxX, 2))
	fmt.Fprintf(v, "\u001b[1m%s %s %s %s %s\u001b[0m\n", namesHeader, artistHeader, albumHeader, lengthHeader, addedOnHeader)
}

func (t *SavedTrackTable) render(v *gocui.View, maxX int) {
	columnWidths := t.getColumnWidths(maxX)

	for _, track := range t.tracks.Items {
		name := utils.LeftPaddedString(track.Track.Name, columnWidths["name"], 2)
		var artistNames []string
		for _, artist := range track.Track.Artists {
			artistNames = append(artistNames, artist.Name)
		}
		artists := utils.LeftPaddedString(strings.Join(artistNames, ", "), columnWidths["artist"], 2)
		album := utils.LeftPaddedString(track.Track.Album.Name, columnWidths["album"], 2)
		length := utils.LeftPaddedString(utils.MillisecondsToFormattedTime(track.Track.DurationMs), columnWidths["length"], 2)
		// popularity := utils.LeftPaddedString(strconv.Itoa(track.Popularity), columnWidths["popularity"], 2)
		addedDate := utils.LeftPaddedString(track.AddedAt.Format("January 1 2006"), columnWidths["addedOn"], 2)

		fmt.Fprintf(v, "\n%s %s %s %s %s", name, artists, album, length, addedDate)
	}
}

func (t *SavedTrackTable) renderFooter(v *gocui.View, maxX int) {
	fmt.Fprintf(v, "\u001b[1m%s\u001b[0m\n", utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d tracks", len(t.tracks.Items), t.tracks.Total), maxX, 2))
}

func (t *SavedTrackTable) getTableLength() int {
	return len(t.tracks.Items)
}

func (t *SavedTrackTable) loadNextRecords() error {
	if t.tracks.Next != "" {
		res, err := api.GetNextSavedTracks(t.tracks.Next)

		if err != nil {
			return err
		}

		nextTracks := res

		t.tracks.Href = nextTracks.Href
		t.tracks.Offset = nextTracks.Offset
		t.tracks.Next = nextTracks.Next
		t.tracks.Previous = nextTracks.Previous
		t.tracks.Items = append(t.tracks.Items, nextTracks.Items...)
	}

	return nil
}

func (t *SavedTrackTable) playSelected(selectedIndex int) (string, error) {
	track := t.tracks.Items[selectedIndex]

	var artistNames []string

	for _, artist := range track.Track.Artists {
		artistNames = append(artistNames, artist.Name)
	}

	if track.Track.Album != nil {
		playerOptions := api.PlayerOptions{
			URIs: append([]string{track.Track.URI}, t.getArrayOfSavedSongURIs(selectedIndex)...),
			// ContextURI: track.Track.URI,
			Offset: &api.PlayerOffsetOptions{
				URI: track.Track.URI,
			},
		}

		chosenItem := fmt.Sprintf("Now playing: '%s' by %s from the album %s\n", track.Track.Name, strings.Join(artistNames, ", "), track.Track.Album.Name)

		return chosenItem, api.StartPlayback(&playerOptions)
	}

	playerOptions := api.PlayerOptions{
		URIs: append([]string{track.Track.URI}, t.getArrayOfSavedSongURIs(selectedIndex)...),
		// ContextURI: track.Track.URI,
		Offset: &api.PlayerOffsetOptions{
			URI: track.Track.URI,
		},
	}

	chosenItem := fmt.Sprintf("Now playing: '%s' by %s\n", track.Track.Name, strings.Join(artistNames, ", "))

	return chosenItem, api.StartPlayback(&playerOptions)
}

func (t *SavedTrackTable) newTableFromSelection(selectedIndex int) (Table, error) {
	track := t.tracks.Items[selectedIndex]
	if track.Track.Album != nil {
		playerOptions := api.PlayerOptions{
			URIs: append([]string{track.Track.URI}, t.getArrayOfSavedSongURIs(selectedIndex)...),
			// ContextURI: track.Track.URI,
			Offset: &api.PlayerOffsetOptions{
				URI: track.Track.URI,
			},
		}
		return nil, api.StartPlayback(&playerOptions)
	}

	playerOptions := api.PlayerOptions{
		URIs: append([]string{track.Track.URI}, t.getArrayOfSavedSongURIs(selectedIndex)...),
		// ContextURI: track.Track.URI,
		Offset: &api.PlayerOffsetOptions{
			URI: track.Track.URI,
		},
	}
	return nil, api.StartPlayback(&playerOptions)
}

func (t *SavedTrackTable) getArrayOfSavedSongURIs(currentIndex int) (songURIs []string) {
	for k, v := range t.tracks.Items {
		if k != currentIndex {
			songURIs = append(songURIs, v.Track.URI)
		}
	}
	return songURIs
}
