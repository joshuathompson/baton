package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func transferDevice(cmd *cobra.Command, args []string) {
	devices, err := api.GetDevices()
	p := api.PlayerOptions{DeviceID: args[0]}
	err = api.StartPlayback(&p)

	if err != nil {
		fmt.Printf("Failed to transfer playback\n")
	} else {
		deviceName := ""
		deviceType := ""

		for _, d := range devices {
			if d.ID == args[0] {
				deviceName = d.Name
				deviceType = d.Type
			}
		}

		fmt.Printf("Transfered playback to %s '%s'\n", deviceType, deviceName)
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
