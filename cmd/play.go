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
		playerOptions.ContextURI = args[0]
		err := api.StartPlayback(&playerOptions)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Playing uri: %s\n", args[0])
	} else {
		err := api.StartPlayback(&playerOptions)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Resuming Spotify playback\n")
	}
}

func playArtist(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "artist", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	playerOptions.ContextURI = res.Artists.Items[0].URI

	err = api.StartPlayback(&playerOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Playing top songs for artist: %s\n", res.Artists.Items[0].Name)
}

func playAlbum(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "album", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	playerOptions.ContextURI = res.Albums.Items[0].URI

	err = api.StartPlayback(&playerOptions)

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
	res, err := api.Search(args[0], "playlist", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	playerOptions.ContextURI = res.Playlists.Items[0].URI

	err = api.StartPlayback(&playerOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Playing playlist: %s by user %s\n", res.Playlists.Items[0].Name, res.Playlists.Items[0].Owner.DisplayName)
}

func playTrack(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "track", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	track := res.Tracks.Items[0]

	if track.Album != nil {
		playerOptions.ContextURI = track.Album.URI
		playerOptions.Offset = &api.PlayerOffsetOptions{URI: track.URI}
		err = api.StartPlayback(&playerOptions)
	} else {
		playerOptions.ContextURI = track.URI
		err = api.StartPlayback(&playerOptions)
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

	playCmd.PersistentFlags().StringVarP(&playerOptions.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var playCmd = &cobra.Command{
	Use:   "play [uri]",
	Short: "Play top result for specified artist, album, playlist, track, or uri",
	Long:  `Play top result for specified artist, album, playlist, track, or uri`,
	Args:  cobra.MaximumNArgs(1),
	Run:   playUri,
}

var playArtistCmd = &cobra.Command{
	Use:   `artist "artist name"`,
	Short: "Play top result for specified artist",
	Long:  `Play top result for specified artist`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playArtist,
}

var playAlbumCmd = &cobra.Command{
	Use:   `album "album name"`,
	Short: "Play top result for specified album",
	Long:  `Play top result for specified album`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playAlbum,
}

var playPlaylistCmd = &cobra.Command{
	Use:   `playlist "playlist name"`,
	Short: "Play top result for specified playlist",
	Long:  `Play top result for specified playlist`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playPlaylist,
}

var playTrackCmd = &cobra.Command{
	Use:   `track "track name"`,
	Short: "Play top result for specified track",
	Long:  `Play top result for specified track`,
	Args:  cobra.MinimumNArgs(1),
	Run:   playTrack,
}
