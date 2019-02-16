package cmd

import (
	"fmt"

	"baton/api"

	"github.com/spf13/cobra"
)

func pausePlayer(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(&options)

	if err != nil {
		fmt.Printf("Couldn't get pause information from the spotify player. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	if ctx.IsPlaying {
		err = api.PausePlayback(&options)

		if err != nil {
			fmt.Printf("Failed to pause\n")
		} else {
			fmt.Printf("Spotify has been paused\n")
		}
	} else {
		err = api.StartPlayback(&playerOptions)

		if err != nil {
			fmt.Printf("Failed to unpause\n")
		} else {
			fmt.Printf("Spotify has been unpaused\n")
		}
	}
}

func init() {
	rootCmd.AddCommand(pauseCmd)

	pauseCmd.Flags().StringVarP(&options.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Toggle spotify pause state",
	Long:  `Toggle spotify pause state`,
	Run:   pausePlayer,
}
