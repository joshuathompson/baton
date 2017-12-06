package cmd

import (
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func changeVolume(cmd *cobra.Command, args []string) {
	fmt.Println(args[0])
	err := api.SetVolume(args[0])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Volume changed to %s\n", args[0])
}

func init() {
	rootCmd.AddCommand(volumeCmd)
}

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "changes volume",
	Long:  `changes volume`,
	Run:   changeVolume,
}
