package cmd

import (
	"fmt"
	"log"
	"strings"

	"baton/api"
	"baton/ui"
	"github.com/spf13/cobra"
)

func searchForArtists(cmd *cobra.Command, args []string) {
	res, err := api.Search(strings.Join(args, " "), "artist", &searchOptions)

	if err != nil {
		fmt.Printf("Couldn't properly search Spotify. Have you authenticated with the 'auth' command?\n")
		return
	}

	at := ui.NewArtistTable(res.Artists)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func searchForPlaylists(cmd *cobra.Command, args []string) {
	res, err := api.Search(strings.Join(args, " "), "playlist", &searchOptions)

	if err != nil {
		fmt.Printf("Couldn't properly search Spotify. Have you authenticated with the 'auth' command?\n")
		return
	}

	at := ui.NewPlaylistTable(res.Playlists)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func searchForAlbums(cmd *cobra.Command, args []string) {
	res, err := api.Search(strings.Join(args, " "), "album", &searchOptions)

	if err != nil {
		fmt.Printf("Couldn't properly search Spotify. Have you authenticated with the 'auth' command?\n")
		return
	}

	at := ui.NewAlbumTable(res.Albums)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func searchForTracks(cmd *cobra.Command, args []string) {
	res, err := api.Search(strings.Join(args, " "), "track", &searchOptions)

	if err != nil {
		fmt.Printf("Couldn't properly search Spotify. Have you authenticated with the 'auth' command?\n")
		return
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
	Short: "Search for specified artist, album, playlist, or track and select via interactive TUI",
	Long:  `Search for specified artist, album, playlist, or track and select via interactive TUI`,
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
