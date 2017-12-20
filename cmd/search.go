package cmd

import (
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/ui"
	"github.com/spf13/cobra"
)

func searchForArtist(cmd *cobra.Command, args []string) {
	res, err := api.Search(args[0], "artist", &searchOptions)

	if err != nil {
		log.Fatal(err)
	}

	at := ui.NewArtistTable(res.Artists)

	ui.Run(at)
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchArtistCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "STUFF",
	Long:  `THINGS`,
}

var searchArtistCmd = &cobra.Command{
	Use:   `artist "artist name"`,
	Short: "Search specified artist",
	Long:  `Search specified artist`,
	Args:  cobra.MinimumNArgs(1),
	Run:   searchForArtist,
}
