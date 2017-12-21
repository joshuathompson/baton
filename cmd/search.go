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

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func searchForPlaylists(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "playlist", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewPlaylistTable(res.Playlists)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func searchForAlbums(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "album", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewAlbumTable(res.Albums)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func searchForTracks(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "track", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewTrackTable(res.Tracks)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
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
	Short: "Search via interactive TUI",
	Long:  `Search via interactive TUI`,
}

var searchArtistsCmd = &cobra.Command{
	Use:   `artists "artist name"`,
	Short: "Search specified artists",
	Long:  `Search specified artists`,
	Args:  cobra.ExactArgs(1),
	Run:   searchForArtists,
}

var searchPlaylistsCmd = &cobra.Command{
	Use:   `playlists "playlist name"`,
	Short: "Search specified playlists",
	Long:  `Search specified playlists`,
	Args:  cobra.ExactArgs(1),
	Run:   searchForPlaylists,
}

var searchAlbumsCmd = &cobra.Command{
	Use:   `albums "album name"`,
	Short: "Search specified albums",
	Long:  `Search specified albums`,
	Args:  cobra.ExactArgs(1),
	Run:   searchForAlbums,
}

var searchTracksCmd = &cobra.Command{
	Use:   `tracks "track name"`,
	Short: "Search specified tracks",
	Long:  `Search specified tracks`,
	Args:  cobra.ExactArgs(1),
	Run:   searchForTracks,
}
