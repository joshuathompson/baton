package cmd

import (
	"errors"
	"fmt"

	"baton/api"
	"baton/utils"

	"github.com/spf13/cobra"
)

func setRepeatMode(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		err := api.SetRepeatMode(args[0], &options)

		if err != nil {
			fmt.Printf("Couldn't set repeat mode. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		} else {
			fmt.Printf("Repeat mode set to %s\n", args[0])
		}
	} else {
		ctx, err := api.GetPlayerState(&options)

		if err != nil {
			fmt.Printf("Couldn't get information about the spotify player. Is Spotify active on a device?  Have you authenticated with the 'auth' command?\n")
		} else {
			fmt.Printf("Repeat mode is currently set to %s\n", ctx.RepeatState)
		}
	}
}

func init() {
	rootCmd.AddCommand(repeatCmd)

	repeatCmd.Flags().StringVarP(&options.DeviceID, "device", "d", "", "id of the device this command is targeting")
}

var repeatCmd = &cobra.Command{
	Use:   "repeat [track|context|off]",
	Short: "Get/Set repeat mode",
	Long:  `Get/Set repeat mode`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 && !utils.StringInSlice(args[0], []string{"track", "context", "off"}) {
			return errors.New("Mode must be track, context, or off")
		}
		return nil
	},
	Run: setRepeatMode,
}
