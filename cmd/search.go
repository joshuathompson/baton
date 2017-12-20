package cmd

import (
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/ui"
	"github.com/spf13/cobra"
)

func searchForArtists(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "artist", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewArtistTable(res.Artists)

	ui.Run(at)
}

func searchForPlaylists(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "playlist", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewPlaylistTable(res.Playlists)

	ui.Run(at)
}

func searchForAlbums(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "album", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewAlbumTable(res.Albums)

	ui.Run(at)
}

func searchForTracks(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "track", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewTrackTable(res.Tracks)

	ui.Run(at)
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchArtistsCmd)
	searchCmd.AddCommand(searchPlaylistsCmd)
	searchCmd.AddCommand(searchAlbumsCmd)
	searchCmd.AddCommand(searchTracksCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "STUFF",
	Long:  `THINGS`,
}

var searchArtistsCmd = &cobra.Command{
	Use:   `artist "artist name"`,
	Short: "Search specified artists",
	Long:  `Search specified artists`,
	Args:  cobra.MinimumNArgs(1),
	Run:   searchForArtists,
}

var searchPlaylistsCmd = &cobra.Command{
	Use:   `playlist "playlist name"`,
	Short: "Search specified playlists",
	Long:  `Search specified playlists`,
	Args:  cobra.MinimumNArgs(1),
	Run:   searchForPlaylists,
}

var searchAlbumsCmd = &cobra.Command{
	Use:   `album "album name"`,
	Short: "Search specified albums",
	Long:  `Search specified albums`,
	Args:  cobra.MinimumNArgs(1),
	Run:   searchForAlbums,
}

var searchTracksCmd = &cobra.Command{
	Use:   `track "track name"`,
	Short: "Search specified tracks",
	Long:  `Search specified tracks`,
	Args:  cobra.MinimumNArgs(1),
	Run:   searchForTracks,
}
