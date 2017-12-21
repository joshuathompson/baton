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
		fmt.Printf("Couldn't transfer playback. Is the device ID supplied correct? Is Spotify active on that device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	devices, err := api.GetDevices()

	if err != nil {
		fmt.Printf("Transferred playback successfully\n")
	} else {
		for _, d := range devices {
			if d.ID == args[0] {
				fmt.Printf("Transfered playback to %s '%s'\n", d.Type, d.Name)
				return
			}
		}
		fmt.Printf("Transferred playback successfully\n")
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
