package cmd

import (
	"fmt"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func skipToPrev(cmd *cobra.Command, args []string) {
	err := api.SkipToPrevious()

	if err != nil {
		fmt.Printf("Failed to skip to previous track\n")
	} else {
		fmt.Printf("Skipped to previous track\n")
	}
}

func init() {
	rootCmd.AddCommand(skipToPrevCmd)
}

var skipToPrevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Skip to previous track",
	Long:  `Skip to previous track`,
	Run:   skipToPrev,
}
