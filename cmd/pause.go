package cmd

import (
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func pausePlayer(cmd *cobra.Command, args []string) {
	ctx, err := api.GetCurrentPlaybackInformation()

	if err != nil {
		log.Fatal(err)
	}

	if ctx.IsPlaying {
		err = api.PausePlayback()

		if err != nil {
			fmt.Printf("Failed to pause\n")
		} else {
			fmt.Printf("Spotify has been paused\n")
		}
	} else {
		err = api.StartPlayback("", 0)

		if err != nil {
			fmt.Printf("Failed to unpause\n")
		} else {
			fmt.Printf("Spotify has been unpaused\n")
		}
	}
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Toggle spotify pause state",
	Long:  `Toggle spotify pause state`,
	Run:   pausePlayer,
}
