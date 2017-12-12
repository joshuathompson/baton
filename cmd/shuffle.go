package cmd

import (
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

var shuffleOptions api.Options

func toggleShuffle(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(&shuffleOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = api.ToggleShuffle(!ctx.ShuffleState, &shuffleOptions)

	if err != nil {
		fmt.Printf("Failed to toggle shuffle\n")
	}

	if ctx.ShuffleState {
		fmt.Printf("Shuffle has been toggled off\n")
	} else {
		fmt.Printf("Shuffle has been toggled on\n")
	}
}

func init() {
	rootCmd.AddCommand(shuffleCmd)

	shuffleCmd.Flags().StringVarP(&shuffleOptions.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var shuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Toggle shuffle on/off",
	Long:  `Toggle shuffle on/off`,
	Run:   toggleShuffle,
}
