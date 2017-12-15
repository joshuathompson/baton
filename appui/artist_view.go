package appui

import (
	"github.com/gizak/termui"
	"github.com/joshuathompson/baton/api"
)

type ArtistView struct {
	ViewProperties
	artists *api.FullArtistsPaged
}

func (a ArtistView) PlaySelectedItem() error {
	return nil
}

func (a ArtistView) GetViewForSelectedItem() (s *View, err error) {
	return nil, nil
}

func (a ArtistView) NextPage() error {
	return nil
}

func (a ArtistView) PreviousPage() error {
	return nil
}

func (a ArtistView) GetGrid() *termui.Grid {
	return nil
}
