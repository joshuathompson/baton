package ui

import (
	"fmt"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/jroimartin/gocui"
)

type AlbumTable struct {
	albums *api.SimpleAlbumsPaged
}

func NewAlbumTable(albumsPaged *api.SimpleAlbumsPaged) *AlbumTable {
	return &AlbumTable{
		albums: albumsPaged,
	}
}

func (a *AlbumTable) getColumnWidths(maxX int) map[string]int {
	m := make(map[string]int)
	m["artists"] = maxX / 3
	m["name"] = maxX - m["artists"]

	return m
}

func (a *AlbumTable) renderHeader(v *gocui.View, maxX int) {
	columnWidths := a.getColumnWidths(maxX)

	namesHeader := utils.LeftPaddedString("NAME", columnWidths["name"], 2)
	artistsHeader := utils.LeftPaddedString("ARTISTS", columnWidths["artists"], 2)

	fmt.Fprintf(v, "\u001b[1m%s[0m\n\n", utils.LeftPaddedString("ALBUMS", maxX, 2))
	fmt.Fprintf(v, "\u001b[1m%s %s\u001b[0m\n", namesHeader, artistsHeader)
}

func (a *AlbumTable) render(v *gocui.View, maxX int) {
	columnWidths := a.getColumnWidths(maxX)

	for _, album := range a.albums.Items {
		name := utils.LeftPaddedString(album.Name, columnWidths["name"], 2)
		var artistNames []string
		for _, artist := range album.Artists {
			artistNames = append(artistNames, artist.Name)
		}
		artists := utils.LeftPaddedString(strings.Join(artistNames, ", "), columnWidths["artists"], 2)

		fmt.Fprintf(v, "\n%s %s", name, artists)
	}
}

func (a *AlbumTable) getTableLength() int {
	return len(a.albums.Items)
}

func (a *AlbumTable) loadNextRecords() error {
	return nil
}

func (a *AlbumTable) playSelected(selectedIndex int) error {
	album := a.albums.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: album.URI,
	}
	return api.StartPlayback(&playerOptions)
}

func (a *AlbumTable) newTableFromSelection(selectedIndex int) (Table, error) {
	album := a.albums.Items[selectedIndex]
	tracksPaged, err := api.GetTracksForAlbum(album.ID)

	if err != nil {
		return nil, err
	}

	return NewSimpleTrackTable(&tracksPaged, &album), nil
}
