package cmd

import (
	"github.com/spf13/cobra"
)

func playURI(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play [uri]",
	Short: "Play specified artist, album, playlist, or song",
	Long:  `Play specified artist, album, playlist, or song`,
	Args:  cobra.MaximumNArgs(1),
	Run:   playURI,
}
