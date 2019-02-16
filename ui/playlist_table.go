package ui

import (
	"fmt"
	"strconv"
	"strings"

	"baton/api"
	"baton/utils"

	"github.com/jroimartin/gocui"
)

// PlaylistTable implements the Table interface for "Simple" Playlist objects as defined by the Spotify Web API
type PlaylistTable struct {
	playlists *api.SimplePlaylistsPaged
}

// NewPlaylistTable creates a new instance of PlaylistTable
func NewPlaylistTable(playlistsPaged *api.SimplePlaylistsPaged) *PlaylistTable {
	return &PlaylistTable{
		playlists: playlistsPaged,
	}
}

func (p *PlaylistTable) getColumnWidths(maxX int) map[string]int {
	m := make(map[string]int)
	m["owner"] = maxX / 3
	m["total"] = maxX / 6
	m["collaborative"] = maxX / 8
	m["name"] = maxX - m["owner"] - m["total"] - m["collaborative"]

	return m
}

func (p *PlaylistTable) renderHeader(v *gocui.View, maxX int) {
	columnWidths := p.getColumnWidths(maxX)

	nameHeader := utils.LeftPaddedString("NAME", columnWidths["name"], 2)
	ownerHeader := utils.LeftPaddedString("OWNER", columnWidths["owner"], 2)
	collaborativeHeader := utils.LeftPaddedString("COLLABORATIVE", columnWidths["collaborative"], 2)
	totalHeader := utils.LeftPaddedString("TOTAL", columnWidths["total"], 2)

	fmt.Fprintf(v, "\u001b[1m%s[0m\n", utils.LeftPaddedString("PLAYLISTS", maxX, 2))
	fmt.Fprintf(v, "\u001b[1m%s %s %s %s\u001b[0m\n", nameHeader, ownerHeader, totalHeader, collaborativeHeader)
}

func (p *PlaylistTable) render(v *gocui.View, maxX int) {
	columnWidths := p.getColumnWidths(maxX)

	for _, playlist := range p.playlists.Items {
		name := utils.LeftPaddedString(playlist.Name, columnWidths["name"], 2)
		owner := utils.LeftPaddedString(playlist.Owner.DisplayName, columnWidths["owner"], 2)
		collaborative := utils.LeftPaddedString(strconv.FormatBool(playlist.Collaborative), columnWidths["owner"], 2)
		total := utils.LeftPaddedString(strconv.Itoa(playlist.Tracks.Total), columnWidths["total"], 2)

		fmt.Fprintf(v, "\n%s %s %s %s", name, owner, total, collaborative)
	}
}

func (p *PlaylistTable) renderFooter(v *gocui.View, maxX int) {
	fmt.Fprintf(v, "\u001b[1m%s\u001b[0m\n", utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d playlists", len(p.playlists.Items), p.playlists.Total), maxX, 2))
}

func (p *PlaylistTable) getTableLength() int {
	return len(p.playlists.Items)
}

func (p *PlaylistTable) loadNextRecords() error {
	if p.playlists.Next != "" {
		if strings.Contains(p.playlists.Next, "api.spotify.com/v1/search") {
			res, err := api.GetNextSearchResults(p.playlists.Next)

			if err != nil {
				return err
			}

			nextPlaylists := res.Playlists

			p.playlists.Href = nextPlaylists.Href
			p.playlists.Offset = nextPlaylists.Offset
			p.playlists.Next = nextPlaylists.Next
			p.playlists.Previous = nextPlaylists.Previous
			p.playlists.Items = append(p.playlists.Items, nextPlaylists.Items...)
		} else {
			res, err := api.GetNextMyPlaylists(p.playlists.Next)

			if err != nil {
				return err
			}

			nextPlaylists := res

			p.playlists.Href = nextPlaylists.Href
			p.playlists.Offset = nextPlaylists.Offset
			p.playlists.Next = nextPlaylists.Next
			p.playlists.Previous = nextPlaylists.Previous
			p.playlists.Items = append(p.playlists.Items, nextPlaylists.Items...)
		}

	}

	return nil
}

func (p *PlaylistTable) playSelected(selectedIndex int) (string, error) {
	playlist := p.playlists.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: playlist.URI,
	}

	chosenItem := fmt.Sprintf("Now playing the playlist: %s by %s\n", playlist.Name, playlist.Owner.DisplayName)

	return chosenItem, api.StartPlayback(&playerOptions)
}

func (p *PlaylistTable) newTableFromSelection(selectedIndex int) (Table, error) {
	playlist := p.playlists.Items[selectedIndex]
	tracksPaged, err := api.GetTracksForPlaylist(playlist.Owner.ID, playlist.ID)

	if err != nil {
		return nil, err
	}

	return NewPlaylistTrackTable(&tracksPaged, &playlist), nil
}

func (p *PlaylistTable) handleSaveKey(selectedIndex int) error {
	return nil
}
