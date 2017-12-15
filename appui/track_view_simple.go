package appui

import (
	"github.com/gizak/termui"
	"github.com/joshuathompson/baton/api"
)

type SimpleTrackView struct {
	ViewProperties
	tracks *api.SimpleTracksPaged
}

func (t SimpleTrackView) PlaySelectedItem() error {
	return nil
}

func (t SimpleTrackView) GetViewForSelectedItem() (s *View, err error) {
	return nil, nil
}

func (t SimpleTrackView) NextPage() error {
	return nil
}

func (t SimpleTrackView) PreviousPage() error {
	return nil
}

func (t SimpleTrackView) GetGrid() *termui.Grid {
	return nil
}
