package cmd

import (
	"fmt"

	"baton/api"
	"github.com/spf13/cobra"
)

func toggleShuffle(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(&options)

	if err != nil {
		fmt.Printf("Couldn't get the player state to retrieve shuffle status. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	err = api.ToggleShuffle(!ctx.ShuffleState, &options)

	if err != nil {
		fmt.Printf("Failed to toggle shuffle\n")
		return
	}

	if ctx.ShuffleState {
		fmt.Printf("Shuffle has been toggled off\n")
	} else {
		fmt.Printf("Shuffle has been toggled on\n")
	}
}

func init() {
	rootCmd.AddCommand(shuffleCmd)

	shuffleCmd.Flags().StringVarP(&options.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var shuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Toggle shuffle on/off",
	Long:  `Toggle shuffle on/off`,
	Run:   toggleShuffle,
}
