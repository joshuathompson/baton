package appui

import (
	"github.com/gizak/termui"
	"github.com/joshuathompson/baton/api"
)

type PlaylistView struct {
	ViewProperties
	playlists *api.SimplePlaylistsPaged
}

func (p PlaylistView) PlaySelectedItem() error {
	return nil
}

func (p PlaylistView) GetViewForSelectedItem() (s *View, err error) {
	return nil, nil
}

func (p PlaylistView) NextPage() error {
	return nil
}

func (p PlaylistView) PreviousPage() error {
	return nil
}

func (p PlaylistView) GetGrid() *termui.Grid {
	return nil
}
