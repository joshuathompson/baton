package cmd

import (
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

var pauseOptions api.Options

func pausePlayer(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(&pauseOptions)

	if err != nil {
		log.Fatal(err)
	}

	if ctx.IsPlaying {
		err = api.PausePlayback(&pauseOptions)

		if err != nil {
			fmt.Printf("Failed to pause\n")
		} else {
			fmt.Printf("Spotify has been paused\n")
		}
	} else {
		err = api.StartPlayback(nil)

		if err != nil {
			fmt.Printf("Failed to unpause\n")
		} else {
			fmt.Printf("Spotify has been unpaused\n")
		}
	}
}

func init() {
	rootCmd.AddCommand(pauseCmd)

	pauseCmd.Flags().StringVarP(&pauseOptions.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Toggle spotify pause state",
	Long:  `Toggle spotify pause state`,
	Run:   pausePlayer,
}
