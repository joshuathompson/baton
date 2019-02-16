package cmd

import (
	"fmt"
	"log"

	"baton/api"
	"baton/ui"
	"github.com/spf13/cobra"
)

func browseSavedTracks(cmd *cobra.Command, args []string) {
	res, err := api.GetSavedTracks(&searchOptions)

	if err != nil {
		fmt.Printf("Couldn't get your saved tracks. Have you authenticated with the 'auth' command?\n")
		fmt.Println("err", res, err)
		return
	}

	at := ui.NewSavedTrackTable(res)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func browseSavedAlbums(cmd *cobra.Command, args []string) {
	res, err := api.GetSavedAlbums(&searchOptions)

	if err != nil {
		fmt.Printf("Couldn't get your saved albums. Have you authenticated with the 'auth' command?\n")
		fmt.Println("err", res, err)
		return
	}

	at := ui.NewSavedAlbumTable(res)

	err = ui.Run(at)

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	meCmd.AddCommand(savedCmd)
	savedCmd.AddCommand(savedTracksCmd)
	savedCmd.AddCommand(savedAlbumsCmd)
}

var savedCmd = &cobra.Command{
	Use:   "saved",
	Short: "Browse tracks or albums saved you've saved",
	Long:  `Browse tracks or albums saved you've saved`,
}

var savedTracksCmd = &cobra.Command{
	Use:     `tracks`,
	Aliases: []string{"songs"},
	Short:   "Browse saved tracks",
	Long:    `Browse saved tracks`,
	Run:     browseSavedTracks,
}

var savedAlbumsCmd = &cobra.Command{
	Use:   `albums`,
	Short: "Browse saved Albums",
	Long:  `Browse saved Albums`,
	Run:   browseSavedAlbums,
}
