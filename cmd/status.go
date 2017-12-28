package cmd

import (
	"fmt"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/spf13/cobra"
)

func reportStatus(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(nil)

	if err != nil {
		fmt.Printf("Couldn't get the player state. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	if ctx.Item != nil {
		track := ctx.Item.Name
		album := ctx.Item.Album.Name
		var playingState string
		var artistNames []string

		if ctx.IsPlaying {
			playingState = "Playing"
		} else {
			playingState = "Paused"
		}

		for _, artist := range ctx.Item.Artists {
			artistNames = append(artistNames, artist.Name)
		}

		progress := utils.MillisecondsToFormattedTime(ctx.ProgressMs)
		duration := utils.MillisecondsToFormattedTime(ctx.Item.DurationMs)

		fmt.Printf("Track: %s\n", track)
		fmt.Printf("Artist: %s\n", strings.Join(artistNames, ", "))
		fmt.Printf("Album: %s\n", album)
		fmt.Printf("Time Elapsed: %s - %s\n", progress, duration)
		fmt.Printf("State: %s\n", playingState)
	} else {
		fmt.Printf("Couldn't find information about the current status\n")
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show information about the current track",
	Long:  `Show information about the current track`,
	Run:   reportStatus,
}
