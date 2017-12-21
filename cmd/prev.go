package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func skipToPrev(cmd *cobra.Command, args []string) {
	err := api.SkipToPrevious(&options)

	if err != nil {
		fmt.Printf("Couldn't skip to previous track. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		return
	}

	time.Sleep(time.Millisecond * 150)
	ps, err := api.GetPlayerState(&options)

	if err != nil {
		fmt.Printf("Skipped to previous track\n")
	} else {
		var artistNames []string
		for _, artist := range ps.Item.Artists {
			artistNames = append(artistNames, artist.Name)
		}
		fmt.Printf("Skipped to previous track, '%s' by %s from album %s\n", ps.Item.Name, strings.Join(artistNames, ", "), ps.Item.Album.Name)
	}
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
