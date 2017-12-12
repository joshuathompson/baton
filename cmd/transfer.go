package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func transferDevice(cmd *cobra.Command, args []string) {
	p := api.PlayerOptions{DeviceID: args[0]}
	err := api.StartPlayback(&p)

	if err != nil {
		fmt.Println("Failed to transfer playback")
	} else {
		fmt.Println("Successfully transfered playback!")
	}
}

func init() {
	rootCmd.AddCommand(transferCmd)
}

var transferCmd = &cobra.Command{
	Use:   "transfer [device]",
	Short: "Transfer playback to another device by id",
	Long:  `Transfer playback to another device by id`,
	Run:   transferDevice,
	Args:  cobra.ExactArgs(1),
}
