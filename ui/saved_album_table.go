package ui

import (
	"fmt"
	"strings"

	"baton/api"
	"baton/utils"
	"github.com/jroimartin/gocui"
)

// SavedAlbumTable implements the Table interface for "Saved" Album objects as defined by the Spotify Web API
type SavedAlbumTable struct {
	albums *api.SavedAlbumsPaged
}

// NewSavedAlbumTable creates a new instance of SavedAlbumTable
func NewSavedAlbumTable(albumsPaged *api.SavedAlbumsPaged) *SavedAlbumTable {
	return &SavedAlbumTable{
		albums: albumsPaged,
	}
}

func (a *SavedAlbumTable) getColumnWidths(maxX int) map[string]int {
	m := make(map[string]int)
	m["artists"] = maxX / 3
	m["name"] = maxX - m["artists"]

	return m
}

func (a *SavedAlbumTable) renderHeader(v *gocui.View, maxX int) {
	columnWidths := a.getColumnWidths(maxX)

	namesHeader := utils.LeftPaddedString("NAME", columnWidths["name"], 2)
	artistsHeader := utils.LeftPaddedString("ARTISTS", columnWidths["artists"], 2)

	fmt.Fprintf(v, "\u001b[1m%s[0m\n", utils.LeftPaddedString("ALBUMS", maxX, 2))
	fmt.Fprintf(v, "\u001b[1m%s %s\u001b[0m\n", namesHeader, artistsHeader)
}

func (a *SavedAlbumTable) render(v *gocui.View, maxX int) {
	columnWidths := a.getColumnWidths(maxX)

	for _, album := range a.albums.Items {
		name := utils.LeftPaddedString(album.Album.Name, columnWidths["name"], 2)
		var artistNames []string
		for _, artist := range album.Album.Artists {
			artistNames = append(artistNames, artist.Name)
		}
		artists := utils.LeftPaddedString(strings.Join(artistNames, ", "), columnWidths["artists"], 2)

		fmt.Fprintf(v, "\n%s %s", name, artists)
	}
}

func (a *SavedAlbumTable) renderFooter(v *gocui.View, maxX int) {
	fmt.Fprintf(v, "\u001b[1m%s\u001b[0m\n", utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d albums", len(a.albums.Items), a.albums.Total), maxX, 2))
}

func (a *SavedAlbumTable) getTableLength() int {
	return len(a.albums.Items)
}

func (a *SavedAlbumTable) loadNextRecords() error {
	if a.albums.Next != "" {

		res, err := api.GetNextSavedAlbums(a.albums.Next)

		if err != nil {
			return err
		}

		nextAlbums := res

		a.albums.Href = nextAlbums.Href
		a.albums.Offset = nextAlbums.Offset
		a.albums.Next = nextAlbums.Next
		a.albums.Previous = nextAlbums.Previous
		a.albums.Items = append(a.albums.Items, nextAlbums.Items...)

	}
	return nil
}

func (a *SavedAlbumTable) playSelected(selectedIndex int) (string, error) {
	album := a.albums.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: album.Album.URI,
	}

	var artistNames []string

	for _, artist := range album.Album.Artists {
		artistNames = append(artistNames, artist.Name)
	}

	chosenItem := fmt.Sprintf("Now playing the album: %s by %s\n", album.Album.Name, strings.Join(artistNames, ", "))

	return chosenItem, api.StartPlayback(&playerOptions)
}

func (a *SavedAlbumTable) newTableFromSelection(selectedIndex int) (Table, error) {
	album := a.albums.Items[selectedIndex]
	tracksPaged, err := api.GetTracksForAlbum(album.Album.ID)

	if err != nil {
		return nil, err
	}

	return NewSimpleTrackTable(&tracksPaged, &album.Album), nil
}

func (a *SavedAlbumTable) handleSaveKey(selectedIndex int) error {
	album := a.albums.Items[selectedIndex]
	err := api.RemoveSavedAlbum(album.Album.ID)
	if err != nil {
		return err
	}

	a.albums.Items = append(a.albums.Items[:selectedIndex], a.albums.Items[selectedIndex+1:]...)

	return nil
}
