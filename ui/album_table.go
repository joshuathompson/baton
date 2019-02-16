package ui

import (
	"fmt"
	"strings"

	"baton/api"
	"baton/utils"
	"github.com/jroimartin/gocui"
)

// AlbumTable implements the Table interface for "Simple" Album objects as defined by the Spotify Web API
type AlbumTable struct {
	albums *api.SimpleAlbumsPaged
}

// NewAlbumTable creates a new instance of AlbumTable
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

	fmt.Fprintf(v, "\u001b[1m%s[0m\n", utils.LeftPaddedString("ALBUMS", maxX, 2))
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

func (a *AlbumTable) renderFooter(v *gocui.View, maxX int) {
	fmt.Fprintf(v, "\u001b[1m%s\u001b[0m\n", utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d albums", len(a.albums.Items), a.albums.Total), maxX, 2))
}

func (a *AlbumTable) getTableLength() int {
	return len(a.albums.Items)
}

func (a *AlbumTable) loadNextRecords() error {
	if a.albums.Next != "" {
		if strings.Contains(a.albums.Next, "api.spotify.com/v1/search") {
			res, err := api.GetNextSearchResults(a.albums.Next)

			if err != nil {
				return err
			}

			nextAlbums := res.Albums

			a.albums.Href = nextAlbums.Href
			a.albums.Offset = nextAlbums.Offset
			a.albums.Next = nextAlbums.Next
			a.albums.Previous = nextAlbums.Previous
			a.albums.Items = append(a.albums.Items, nextAlbums.Items...)
		} else {
			nextAlbums, err := api.GetNextAlbumsForArtist(a.albums.Next)

			if err != nil {
				return err
			}

			a.albums.Href = nextAlbums.Href
			a.albums.Offset = nextAlbums.Offset
			a.albums.Next = nextAlbums.Next
			a.albums.Previous = nextAlbums.Previous
			a.albums.Items = append(a.albums.Items, nextAlbums.Items...)
		}
	}
	return nil
}

func (a *AlbumTable) playSelected(selectedIndex int) (string, error) {
	album := a.albums.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: album.URI,
	}

	var artistNames []string

	for _, artist := range album.Artists {
		artistNames = append(artistNames, artist.Name)
	}

	chosenItem := fmt.Sprintf("Now playing the album: %s by %s\n", album.Name, strings.Join(artistNames, ", "))

	return chosenItem, api.StartPlayback(&playerOptions)
}

func (a *AlbumTable) newTableFromSelection(selectedIndex int) (Table, error) {
	album := a.albums.Items[selectedIndex]
	tracksPaged, err := api.GetTracksForAlbum(album.ID)

	if err != nil {
		return nil, err
	}

	return NewSimpleTrackTable(&tracksPaged, &album), nil
}

func (a *AlbumTable) handleSaveKey(selectedIndex int) error {
	album := a.albums.Items[selectedIndex]

	err := api.SaveAlbum(album.ID)
	if err != nil {
		return err
	}
	return nil
}
