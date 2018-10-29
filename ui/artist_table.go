package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/jroimartin/gocui"
)

// ArtistTable implements the Table interface for "Full" Artist objects as defined by the Spotify Web API
type ArtistTable struct {
	artists *api.FullArtistsPaged
}

// NewArtistTable creates a new instance of ArtistTable
func NewArtistTable(artistsPaged *api.FullArtistsPaged) *ArtistTable {
	return &ArtistTable{
		artists: artistsPaged,
	}
}

func (a *ArtistTable) getColumnWidths(maxX int) map[string]int {
	m := make(map[string]int)
	m["name"] = maxX / 3
	m["genre"] = maxX / 2
	m["popularity"] = maxX - m["name"] - m["genre"]

	return m
}

func (a *ArtistTable) renderHeader(v *gocui.View, maxX int) {
	columnWidths := a.getColumnWidths(maxX)

	namesHeader := utils.LeftPaddedString("NAME", columnWidths["name"], 2)
	genresHeader := utils.LeftPaddedString("GENRES", columnWidths["genre"], 2)
	popularitiesHeader := utils.LeftPaddedString("POPULARITY", columnWidths["popularity"], 2)

	fmt.Fprintf(v, "\u001b[1m%s[0m\n", utils.LeftPaddedString("ARTISTS", maxX, 2))
	fmt.Fprintf(v, "\u001b[1m%s %s %s\u001b[0m\n", namesHeader, genresHeader, popularitiesHeader)
}

func (a *ArtistTable) render(v *gocui.View, maxX int) {
	columnWidths := a.getColumnWidths(maxX)

	for _, artist := range a.artists.Items {
		name := utils.LeftPaddedString(artist.Name, columnWidths["name"], 2)
		genre := utils.LeftPaddedString(strings.Join(artist.Genres, ", "), columnWidths["genre"], 2)
		popularity := utils.LeftPaddedString(strconv.Itoa(artist.Popularity), columnWidths["popularity"], 2)

		fmt.Fprintf(v, "\n%s %s %s", name, genre, popularity)
	}
}

func (a *ArtistTable) renderFooter(v *gocui.View, maxX int) {
	fmt.Fprintf(v, "\u001b[1m%s\u001b[0m\n", utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d artists", len(a.artists.Items), a.artists.Total), maxX, 2))
}

func (a *ArtistTable) getTableLength() int {
	return len(a.artists.Items)
}

func (a *ArtistTable) loadNextRecords() error {
	if a.artists.Next != "" {
		res, err := api.GetNextSearchResults(a.artists.Next)

		if err != nil {
			return err
		}

		nextArtists := res.Artists

		a.artists.Href = nextArtists.Href
		a.artists.Offset = nextArtists.Offset
		a.artists.Next = nextArtists.Next
		a.artists.Previous = nextArtists.Previous
		a.artists.Items = append(a.artists.Items, nextArtists.Items...)
	}
	return nil
}

func (a *ArtistTable) playSelected(selectedIndex int) (string, error) {
	artist := a.artists.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: artist.URI,
	}

	chosenItem := fmt.Sprintf("Now playing top songs from artist: %s\n", artist.Name)

	return chosenItem, api.StartPlayback(&playerOptions)
}

func (a *ArtistTable) newTableFromSelection(selectedIndex int) (Table, error) {
	artist := a.artists.Items[selectedIndex]
	albumsPaged, err := api.GetAlbumsForArtist(artist.ID)

	if err != nil {
		return nil, err
	}

	return NewAlbumTable(&albumsPaged), nil
}

func (a *ArtistTable) handleSaveKey(selectedIndex int) error {
	return nil
}
