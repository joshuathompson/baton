package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/spf13/cobra"
)

func reportStatus(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(nil)

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Item != nil {
		track := ctx.Item.Name
		album := ctx.Item.Album.Name
		var artistNames []string

		for _, artist := range ctx.Item.Artists {
			artistNames = append(artistNames, artist.Name)
		}

		progress := utils.MillisecondsToFormattedTime(ctx.ProgressMs)
		duration := utils.MillisecondsToFormattedTime(ctx.Item.DurationMs)

		fmt.Printf("Track: %s\n", track)
		fmt.Printf("Artist: %s\n", strings.Join(artistNames, ", "))
		fmt.Printf("Album: %s\n", album)
		fmt.Printf("Time Elapsed: %s - %s\n", progress, duration)
	} else {
		fmt.Printf("No currently playing track\n")
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show information about the current track",
	Long:  `Show information about the current track`,
	Run:   reportStatus,
}
