package cmd

import (
	"fmt"
	"strings"

	"baton/api"
	"github.com/spf13/cobra"
)

func reportDevices(cmd *cobra.Command, args []string) {
	devices, err := api.GetDevices()

	if err != nil {
		fmt.Printf("Couldn't retrieve devices. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
	} else if len(devices) > 0 {
		var o []string
		for _, d := range devices {
			s := fmt.Sprintf("Name: %s\nType: %s\nID: %s\nActive: %v\n", d.Name, d.Type, d.ID, d.IsActive)
			o = append(o, s)
		}
		fmt.Print(strings.Join(o, "\n"))
	} else {
		fmt.Printf("No devices currently available\n")
	}
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List all available playback devices",
	Long:  `List all available playback devices`,
	Run:   reportDevices,
}
