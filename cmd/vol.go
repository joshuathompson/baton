package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joshuathompson/baton/api"
	"github.com/joshuathompson/baton/utils"
	"github.com/spf13/cobra"
)

func increaseVolume(cmd *cobra.Command, args []string) {
	ctx, err := api.GetCurrentPlaybackInformation()

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Device != nil {
		if utils.StringInSlice(ctx.Device.Type, []string{"CastVideo", "Phone"}) {
			fmt.Printf("Can't get/set volume for %s '%s', this type of device doesn't support volume commands in the web api beta\n", ctx.Device.Type, ctx.Device.Name)
		} else {
			v := ctx.Device.VolumePercent + 10

			if v > 100 {
				v = 100
			}

			err = api.SetVolume(v)

			if err != nil {
				fmt.Printf("Failed to set volume\n")
			} else {
				fmt.Printf("Volume increased to %d%%\n", v)
			}
		}
	} else {
		fmt.Printf("No device currently playing\n")
	}
}

func decreaseVolume(cmd *cobra.Command, args []string) {
	ctx, err := api.GetCurrentPlaybackInformation()

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Device != nil {
		if utils.StringInSlice(ctx.Device.Type, []string{"CastVideo", "Phone"}) {
			fmt.Printf("Can't get/set volume for %s '%s', this type of device doesn't support volume commands in the web api beta\n", ctx.Device.Type, ctx.Device.Name)
		} else {
			v := ctx.Device.VolumePercent - 10

			if v < 0 {
				v = 0
			}

			err = api.SetVolume(v)

			if err != nil {
				fmt.Printf("Failed to set volume\n")
			} else {
				fmt.Printf("Volume decreased to %d%%.\n", v)
			}
		}
	} else {
		fmt.Printf("No device currently playing\n")
	}
}

func getSetVolume(cmd *cobra.Command, args []string) {
	ctx, err := api.GetCurrentPlaybackInformation()

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Device != nil {
		if utils.StringInSlice(ctx.Device.Type, []string{"CastVideo", "Phone"}) {
			fmt.Printf("Can't get/set volume for %s '%s', this type of device doesn't support volume commands in the web api beta\n", ctx.Device.Type, ctx.Device.Name)
		} else {
			if len(args) > 0 {
				p, err := strconv.Atoi(args[0])

				if err != nil {
					fmt.Printf("Volume number must be an integer between 0-100\n")
					return
				}

				err = api.SetVolume(p)

				if err != nil {
					fmt.Printf("Failed to set volume\n")
				} else {
					fmt.Printf("Volume changed to %s%% for %s '%s'\n", args[0], ctx.Device.Type, ctx.Device.Name)
				}
			} else {
				fmt.Printf("Volume for %s %s is %d%%\n", ctx.Device.Type, ctx.Device.Name, ctx.Device.VolumePercent)
			}
		}
	} else {
		fmt.Printf("No device currently playing")
	}
}

func init() {
	rootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(volumeUpCmd)
	volumeCmd.AddCommand(volumeDownCmd)
}

var volumeCmd = &cobra.Command{
	Use:           "vol [0-100]",
	Short:         "Get/Set volume",
	Long:          `Get/Set volume`,
	Args:          cobra.MaximumNArgs(1),
	Run:           getSetVolume,
}

var volumeUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Increase volume by 10%",
	Long:  `Increase volume by 10%`,
	Run:   increaseVolume,
}

var volumeDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Decrease volume by 10%",
	Long:  `Decrease volume by 10%`,
	Run:   decreaseVolume,
}
