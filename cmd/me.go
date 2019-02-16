package cmd

import (
	"fmt"
	"log"

	"baton/api"
	"baton/ui"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(meCmd)
	meCmd.AddCommand(myPlaylistsCmd)

}

var meCmd = &cobra.Command{
	Use:     "me",
	Short:   "Your playlists and tracks you've saved",
	Long:    `Your playlists and tracks you've saved`,
	Aliases: []string{"my"},
}

func browsePlayLists(cmd *cobra.Command, args []string) {

	res, err := api.GetMyPlaylists()

	if err != nil {

		fmt.Printf("Couldn't get your playlists from spotify. Have you authenticated with the 'auth' command?\n")
		fmt.Println("err", res, err)
		return
	}

	at := ui.NewPlaylistTable(res)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

var myPlaylistsCmd = &cobra.Command{
	Use:   `playlists`,
	Short: "Browse your playlists",
	Long:  `Browse your playlists`,
	Run:   browsePlayLists,
}
