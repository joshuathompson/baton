package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

var prevOptions api.Options

func skipToPrev(cmd *cobra.Command, args []string) {
	err := api.SkipToPrevious(&prevOptions)

	if err != nil {
		fmt.Printf("Failed to skip to previous track\n")
	} else {
		fmt.Printf("Skipped to previous track\n")
	}
}

func init() {
	rootCmd.AddCommand(skipToPrevCmd)

	skipToPrevCmd.Flags().StringVarP(&prevOptions.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var skipToPrevCmd = &cobra.Command{
	Use:     "prev",
	Short:   "Skip to previous track",
	Long:    `Skip to previous track`,
	Run:     skipToPrev,
	Aliases: []string{"previous"},
}
