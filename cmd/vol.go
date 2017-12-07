package cmd

import (
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func changeVolume(cmd *cobra.Command, args []string) {
	err := api.SetVolume(args[0])

	if err != nil {
		fmt.Println("Failed to set volume.")
	} else {
		fmt.Printf("Volume changed to %s\n", args[0])
	}
}

func asd(cmd *cobra.Command, args []string) {
	ctx, err := api.GetCurrentlyPlayingTrack()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", ctx)
}

func init() {
	rootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(volumeSubCmd)
}

var volumeCmd = &cobra.Command{
	Use:   "vol",
	Short: "changes volume",
	Long:  `changes volume`,
	Run:   changeVolume,
	Args:  cobra.ExactArgs(1),
}

var volumeSubCmd = &cobra.Command{
	Use:   "up",
	Short: "changes volume",
	Long:  `changes volume`,
	Run:   asd,
}
