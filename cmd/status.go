package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func reportStatus(cmd *cobra.Command, args []string) {
	ctx, err := api.GetCurrentPlaybackInformation()

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Item != nil {
		song := ctx.Item.Name
		album := ctx.Item.Album.Name
		var artistNames []string

		for _, artist := range ctx.Item.Artists {
			artistNames = append(artistNames, artist.Name)
		}

		fmt.Printf("%s -- %s -- %s\n", song, strings.Join(artistNames, ","), album)
	} else {
		fmt.Printf("No currently playing song\n")
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show currently playing song/artist/album",
	Long:  `Show currently playing song/artist/album`,
	Run:   reportStatus,
}
