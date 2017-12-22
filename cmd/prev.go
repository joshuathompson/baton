package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func skipToPrev(cmd *cobra.Command, args []string) {
	err := api.SkipToPrevious(&options)

	if err != nil {
		fmt.Printf("Couldn't skip to previous track. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	fmt.Printf("Skipped to previous track\n")
}

func init() {
	rootCmd.AddCommand(skipToPrevCmd)

	skipToPrevCmd.Flags().StringVarP(&options.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var skipToPrevCmd = &cobra.Command{
	Use:     "prev",
	Short:   "Skip to previous track",
	Long:    `Skip to previous track`,
	Run:     skipToPrev,
	Aliases: []string{"previous"},
}
