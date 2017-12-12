package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

var replayOptions api.Options

func replayTrack(cmd *cobra.Command, args []string) {
	err := api.SeekToPosition(0, &replayOptions)

	if err != nil {
		fmt.Printf("Failed to restart current track\n")
	} else {
		fmt.Printf("Replaying current track\n")
	}
}

func init() {
	rootCmd.AddCommand(replayTrackCmd)

	replayTrackCmd.Flags().StringVarP(&replayOptions.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var replayTrackCmd = &cobra.Command{
	Use:   "replay",
	Short: "Replay current track from the beginning",
	Long:  `Replay current track from the beginning`,
	Run:   replayTrack,
}
