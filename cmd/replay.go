package cmd

import (
	"fmt"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func replayTrack(cmd *cobra.Command, args []string) {
	err := api.SeekToPosition(0, &options)

	if err != nil {
		fmt.Printf("Couldn't seek to chosen position. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	ps, err := api.GetPlayerState(&options)

	if err != nil {
		fmt.Printf("Replaying current song\n")
	} else {
		var artistNames []string
		for _, artist := range ps.Item.Artists {
			artistNames = append(artistNames, artist.Name)
		}

		fmt.Printf("Replaying '%s' by %s from album %s\n", ps.Item.Name, strings.Join(artistNames, ", "), ps.Item.Album.Name)
	}
}

func init() {
	rootCmd.AddCommand(replayTrackCmd)

	replayTrackCmd.Flags().StringVarP(&options.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var replayTrackCmd = &cobra.Command{
	Use:   "replay",
	Short: "Replay current track from the beginning",
	Long:  `Replay current track from the beginning`,
	Run:   replayTrack,
}
