package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/spf13/cobra"
)

func setRepeatMode(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		err := api.SetRepeatMode(args[0])

		if err != nil {
			fmt.Printf("Failed to set repeat mode\n")
		} else {
			fmt.Printf("Repeat mode set to %s\n", args[0])
		}
	} else {
		ctx, err := api.GetCurrentPlaybackInformation()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Repeat mode is currently set to %s\n", ctx.RepeatState)
	}
}

func init() {
	rootCmd.AddCommand(repeatCmd)
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
