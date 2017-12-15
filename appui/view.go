package appui

import (
	"reflect"

	"github.com/gizak/termui"
	"github.com/joshuathompson/baton/api"
)

type View interface {
	PlaySelectedItem() error
	NextPage() error
	PreviousPage() error
	SetSelectedIndex()
	GetSelectedIndex() int
	GetViewForSelectedItem() (*View, error)
	GetGrid() *termui.Grid
}

type ViewProperties struct {
	SelectedIndex int
	pageSize      int
}

func (v ViewProperties) GetSelectedIndex() int {
	return v.SelectedIndex
}

func (v ViewProperties) SetSelectedIndex(i index) {
	v.SelectedIndex = i
}

func NewView(i interface{}) View {
	v := reflect.ValueOf(i)

	switch v.Type().Elem().Name() {
	case "FullArtistsPaged":
		a := v.Interface().(*api.FullArtistsPaged)
		return ArtistView{
			ViewProperties: ViewProperties{
				pageSize: 15,
			},
			artists: a,
		}
	case "SimpleAlbumsPaged":
		a := v.Interface().(*api.SimpleAlbumsPaged)
		return AlbumView{
			ViewProperties: ViewProperties{
				pageSize: 15,
			},
			albums: a,
		}
	case "SimplePlaylistsPaged":
		p := v.Interface().(*api.SimplePlaylistsPaged)
		return PlaylistView{
			ViewProperties: ViewProperties{
				pageSize: 15,
			},
			playlists: p,
		}
	case "FullTracksPaged":
		t := v.Interface().(*api.FullTracksPaged)
		return TrackView{
			ViewProperties: ViewProperties{
				pageSize: 15,
			},
			tracks: t,
		}
	case "SimpleTracksPaged":
		t := v.Interface().(*api.SimpleTracksPaged)
		return SimpleTrackView{
			ViewProperties: ViewProperties{
				pageSize: 15,
			},
			tracks: t,
		}
	default:
		return nil
	}
}
