package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func playUri(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		err := api.StartPlayback(args[0], 0)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Playing uri: %s\n", args[0])
	} else {
		err := api.StartPlayback("", 0)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Resuming Spotify playback\n")
	}
}

func playArtist(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "artist", 1, 0)

	if err != nil {
		log.Fatal(err)
	}

	uri := res.Artists.Items[0].URI

	err = api.StartPlayback(uri, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Playing top songs for artist: %s\n", res.Artists.Items[0].Name)
}

func playAlbum(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "album", 1, 0)

	if err != nil {
		log.Fatal(err)
	}

	uri := res.Albums.Items[0].URI

	err = api.StartPlayback(uri, 0)

	if err != nil {
		log.Fatal(err)
	}

	var artistNames []string

	for _, artist := range res.Albums.Items[0].Artists {
		artistNames = append(artistNames, artist.Name)
	}

	fmt.Printf("Playing: %s by %s\n", res.Albums.Items[0].Name, strings.Join(artistNames, ", "))
}

func playPlaylist(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "playlist", 1, 0)

	if err != nil {
		log.Fatal(err)
	}

	uri := res.Playlists.Items[0].URI

	err = api.StartPlayback(uri, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Playing playlist: %s by user %s\n", res.Playlists.Items[0].Name, res.Playlists.Items[0].Owner.DisplayName)
}

func playTrack(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "track", 1, 0)

	if err != nil {
		log.Fatal(err)
	}

	track := res.Tracks.Items[0]

	if track.Album != nil {
		err = api.StartPlayback(track.Album.URI, track.TrackNumber)
	} else {
		err = api.StartPlayback(track.URI, 0)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Playing track: %s\n", res.Tracks.Items[0].Name)
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.AddCommand(playArtistCmd)
	playCmd.AddCommand(playAlbumCmd)
	playCmd.AddCommand(playPlaylistCmd)
	playCmd.AddCommand(playTrackCmd)
}

var playCmd = &cobra.Command{
	Use:   "play [uri]",
	Short: "Play specified artist, album, playlist, track, or uri",
	Long:  `Play specified artist, album, playlist, track, or uri`,
	Args:  cobra.MaximumNArgs(1),
	Run:   playUri,
}

var playArtistCmd = &cobra.Command{
	Use:   `artist "artist name"`,
	Short: "Play specified artist",
	Long:  `Play specified artist`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playArtist,
}

var playAlbumCmd = &cobra.Command{
	Use:   `album "album name"`,
	Short: "Play specified album",
	Long:  `Play specified album`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playAlbum,
}

var playPlaylistCmd = &cobra.Command{
	Use:   `playlist "playlist name"`,
	Short: "Play specified playlist",
	Long:  `Play specified playlist`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playPlaylist,
}

var playTrackCmd = &cobra.Command{
	Use:   `track "track name"`,
	Short: "Play specified track",
	Long:  `Play specified track`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playTrack,
}
