package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func skipToNext(cmd *cobra.Command, args []string) {
	err := api.SkipToNext(&options)

	if err != nil {
		fmt.Printf("Failed to skip to next track\n")
	} else {
		time.Sleep(time.Millisecond * 150)
		ps, err := api.GetPlayerState(&options)

		if err != nil {
			fmt.Printf("Skipped to next track\n")
		} else {
			var artistNames []string
			for _, artist := range ps.Item.Artists {
				artistNames = append(artistNames, artist.Name)
			}
			fmt.Printf("Skipped to next track, '%s' by %s from album %s\n", ps.Item.Name, strings.Join(artistNames, ", "), ps.Item.Album.Name)
		}
	}
}

func init() {
	rootCmd.AddCommand(skipToNextCmd)

	skipToNextCmd.Flags().StringVarP(&options.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var skipToNextCmd = &cobra.Command{
	Use:   "next",
	Short: "Skip to next track",
	Long:  `Skip to next track`,
	Run:   skipToNext,
}
