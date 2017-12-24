package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joshuathompson/baton/api"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var options api.Options
var playerOptions api.PlayerOptions
var searchOptions api.SearchOptions

var rootCmd = &cobra.Command{
	Use:   "baton",
	Short: "A CLI tool to orchestrate your Spotify",
	Long:  `A CLI tool to orchestrate your Spotify`,
}

// Execute is the entrypoint for the CLI called from the main function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home + "/.config")
	viper.SetConfigName("baton")
	cfgFile := home + "/.config/baton.json"

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		err := ioutil.WriteFile(cfgFile, []byte("{}"), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}
