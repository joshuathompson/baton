package cmd

import (
	"fmt"
	"strconv"

	"baton/api"

	"github.com/spf13/cobra"
)

func seekToPosition(cmd *cobra.Command, args []string) {
	pos, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf("Time specified could not be converted to seconds\n")
		return
	}

	err = api.SeekToPosition(pos*1000, &options)

	if err != nil {
		fmt.Printf("Failed to skip to entered position\n")
	} else {
		fmt.Printf("Skipping to %d seconds\n", pos)
	}
}

func init() {
	rootCmd.AddCommand(seekCmd)

	seekCmd.Flags().StringVarP(&options.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var seekCmd = &cobra.Command{
	Use:   "seek [pos]",
	Short: "Skip to a specific time (seconds) of the current track",
	Long:  `Skip to a specific time (seconds) of the current track`,
	Run:   seekToPosition,
}
