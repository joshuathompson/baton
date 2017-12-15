package appui

import (
	"github.com/gizak/termui"
	"github.com/joshuathompson/baton/api"
)

type TrackView struct {
	ViewProperties
	tracks *api.FullTracksPaged
}

func (t TrackView) PlaySelectedItem() error {
	return nil
}

func (t TrackView) GetViewForSelectedItem() (s *View, err error) {
	return nil, nil
}

func (t TrackView) NextPage() error {
	return nil
}

func (t TrackView) PreviousPage() error {
	return nil
}

func (t TrackView) GetGrid() *termui.Grid {
	return nil
}
