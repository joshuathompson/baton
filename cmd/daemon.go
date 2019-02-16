// +build !windows

package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"baton/daemon"

	appdir "github.com/ProtonMail/go-appdir"
	"github.com/spf13/cobra"
)

func runDaemon(cmd *cobra.Command, args []string) {
	if isDaemon {
		daemon.Run(logFile, outFile, pipeFile)
	} else {
		pid, err := daemon.Start(pipeFile, logFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Daemon running with PID=%d, pipe=%s\n", pid, pipeFile)
	}
}

func init() {
	dataDir := appdir.New("baton").UserData()
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.PersistentFlags().StringVar(&logFile, "log-file", filepath.Join(dataDir, "daemon.log"), "Daemon log file")
	daemonCmd.PersistentFlags().StringVar(&outFile, "out-file", filepath.Join(dataDir, "stdout.txt"), "Daemon output file")
	daemonCmd.PersistentFlags().StringVar(&pipeFile, "pipe-file", filepath.Join(dataDir, "baton.pipe"), "Daemon pipe file")
	daemonCmd.PersistentFlags().BoolVar(&isDaemon, "x", false, "Don't use this flag")
}

var (
	isDaemon = false
	logFile  string
	outFile  string
	pipeFile string
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run in the background",
	Long:  "Run in the background",
	Args:  cobra.MaximumNArgs(1),
	Run:   runDaemon,
}
