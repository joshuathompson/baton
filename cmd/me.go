package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(meCmd)

}

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Your playlists and tracks you've saved",
	Long:  `Your playlists and tracks you've saved`,
}

// var savedTracksCmd = &cobra.Command{
// 	Use:   `tracks`,
// 	Short: "Browse saved tracks",
// 	Long:  `Browse saved tracks`,
// 	Run:   browseSavedTracks,
// }
