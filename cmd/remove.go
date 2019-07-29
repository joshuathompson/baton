package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func removeTrack(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(nil)

	if err != nil {
		fmt.Printf("Couldn't get the player state. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	if ctx.Item == nil {
		fmt.Printf("Couldn't find information about the current status\n")
		return
	}

	err = api.RemoveSavedTrack(ctx.Item.ID)
	if err != nil {
		fmt.Printf("Couldn't remove the track from saved. %s\n", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(removeTrackCmd)
}

var removeTrackCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove current playing track from saved",
	Long:  `Remove current playing track from saved`,
	Run:   removeTrack,
}
