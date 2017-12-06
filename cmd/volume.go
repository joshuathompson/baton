package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func changeVolume(cmd *cobra.Command, args []string) {
	key := viper.Get("access_token")
	fmt.Printf("%+v", key)
}

func init() {
	rootCmd.AddCommand(volumeCmd)
}

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "changes volume",
	Long:  `changes volume`,
	Run:   changeVolume,
}
