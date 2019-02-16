package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"baton/api"

	appdir "github.com/ProtonMail/go-appdir"
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
	cfgDir := appdir.New("baton").UserConfig()
	viper.AddConfigPath(cfgDir)
	viper.SetConfigName("baton")
	cfgFile := filepath.Join(cfgDir, "baton.json")

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		if err := os.MkdirAll(cfgDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(cfgFile, []byte("{}"), 0666); err != nil {
			log.Fatal(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
