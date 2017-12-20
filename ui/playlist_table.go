package ui

import (
	"fmt"
	"strconv"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/jroimartin/gocui"
)

type PlaylistTable struct {
	playlists *api.SimplePlaylistsPaged
}

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

	loadedLength := maxX / 3
	loadedHeader := utils.LeftPaddedString(fmt.Sprintf("Showing %d of %d playlists", len(p.playlists.Items), p.playlists.Total), loadedLength, 2)
	titleLength := maxX - loadedLength

	fmt.Fprintf(v, "\u001b[1m%s %s[0m\n\n", utils.LeftPaddedString("PLAYLISTS", titleLength, 2), loadedHeader)
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

func (p *PlaylistTable) getTableLength() int {
	return len(p.playlists.Items)
}

func (p *PlaylistTable) loadNextRecords() error {
	if p.playlists.Next != "" {
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
	}

	return nil
}

func (p *PlaylistTable) playSelected(selectedIndex int) error {
	playlist := p.playlists.Items[selectedIndex]
	playerOptions := api.PlayerOptions{
		ContextURI: playlist.URI,
	}
	return api.StartPlayback(&playerOptions)
}

func (p *PlaylistTable) newTableFromSelection(selectedIndex int) (Table, error) {
	playlist := p.playlists.Items[selectedIndex]
	tracksPaged, err := api.GetTracksForPlaylist(playlist.Owner.ID, playlist.ID)

	if err != nil {
		return nil, err
	}

	return NewPlaylistTrackTable(&tracksPaged, &playlist), nil
}
