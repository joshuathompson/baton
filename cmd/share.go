package cmd

import (
	"fmt"
	"log"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func getURIAndURL(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(nil)

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Item != nil {
		fmt.Printf("URI: %s\n", ctx.Item.URI)
		fmt.Printf("URL: %s\n", ctx.Item.Href)
	} else {
		fmt.Printf("There doesn't appear to be a track playing currently\n")
	}
}

func getURI(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(nil)

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Item != nil {
		fmt.Printf("%s\n", ctx.Item.URI)
	} else {
		fmt.Printf("There doesn't appear to be a track playing currently\n")
	}
}

func getURL(cmd *cobra.Command, args []string) {
	ctx, err := api.GetPlayerState(nil)

	if err != nil {
		log.Fatal(err)
	}

	if ctx.Item != nil {
		fmt.Printf("%s\n", ctx.Item.Href)
	} else {
		fmt.Printf("There doesn't appear to be a track playing currently\n")
	}
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.AddCommand(shareURICmd)
	shareCmd.AddCommand(shareURLCmd)
}

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Get URI and URL for current track",
	Long:  `Get URI and URL for current track`,
	Run:   getURIAndURL,
}

var shareURICmd = &cobra.Command{
	Use:   "uri",
	Short: "Get URI for current track",
	Long:  `Get URI for current track`,
	Run:   getURI,
}

var shareURLCmd = &cobra.Command{
	Use:   "url",
	Short: "Get URL for the current track",
	Long:  `Get URL for the current track`,
	Run:   getURL,
}
