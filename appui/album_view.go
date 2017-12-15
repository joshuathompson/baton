package appui

import (
	"github.com/gizak/termui"
	"github.com/joshuathompson/baton/api"
)

type AlbumView struct {
	ViewProperties
	albums *api.SimpleAlbumsPaged
}

func (a AlbumView) PlaySelectedItem() error {
	return nil
}

func (a AlbumView) GetViewForSelectedItem() (s *View, err error) {
	return nil, nil
}

func (a AlbumView) NextPage() error {
	return nil
}

func (a AlbumView) PreviousPage() error {
	return nil
}

func (a AlbumView) GetGrid() *termui.Grid {
	return nil
}
