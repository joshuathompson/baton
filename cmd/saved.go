package cmd

import (
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/ui"
	"github.com/spf13/cobra"
)

func browseSavedTracks(cmd *cobra.Command, args []string) {
	res, err := api.GetSavedTracks(&searchOptions)

	if err != nil {
		fmt.Printf("Couldn't properly search Spotify. Have you authenticated with the 'auth' command?\n")
		fmt.Println("err", res, err)
		return
	}

	at := ui.NewSavedTrackTable(res)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(savedCmd)
	savedCmd.AddCommand(savedTracksCmd)
}

var savedCmd = &cobra.Command{
	Use:   "saved",
	Short: "Browse tracks or albums saved you've saved",
	Long:  `Browse tracks or albums saved you've saved`,
}

var savedTracksCmd = &cobra.Command{
	Use:   `tracks`,
	Short: "Browse saved tracks",
	Long:  `Browse saved tracks`,
	Run:   browseSavedTracks,
}
