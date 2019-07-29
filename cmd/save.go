package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func saveTrack(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(nil)

	if err != nil {
		fmt.Printf("Couldn't get the player state. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	if ctx.Item == nil {
		fmt.Printf("Couldn't find information about the current status\n")
		return
	}

	err = api.SaveTrack(ctx.Item.ID)
	if err != nil {
		fmt.Printf("Couldn't save the track. %s\n", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(saveTrackCmd)
}

var saveTrackCmd = &cobra.Command{
	Use:   "save",
	Short: "Save current playing track",
	Long:  `Save current playing track`,
	Run:   saveTrack,
}
