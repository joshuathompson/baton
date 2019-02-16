package daemon

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// Start starts the daemon and returns the new process's PID.
func Start(pipeFile, logFile string) (int, error) {
	if err := os.MkdirAll(filepath.Dir(pipeFile), os.ModePerm); err != nil {
		return 0, err
	}
	if err := syscall.Mkfifo(pipeFile, 0666); os.IsExist(err) {
		return 0, fmt.Errorf("File %s already exists, make sure no other baton daemons are running then delete it", pipeFile)
	} else if err != nil {
		return 0, err
	}

	cmd := exec.Command(os.Args[0], "daemon", "--x", "--log-file", logFile, "--pipe-file", pipeFile)
	if err := cmd.Start(); err != nil {
		return 0, err
	}

	return cmd.Process.Pid, nil
}

func Run(pipeFile, logFile string) {
	if err := os.MkdirAll(filepath.Dir(logFile), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	w, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(w)

	defer os.Remove(pipeFile)
	for {
		b, err := ioutil.ReadFile(pipeFile)
		if err == io.EOF {
			log.Println("Encountered EOF")
			break
		} else if err != nil {
			log.Println("Error reading from pipe:", err)
			continue
		}

		cmd := strings.TrimSpace(string(b))
		log.Println("Received command:", cmd)

		if cmd == "exit" {
			break
		} else {
			// TODO
		}
	}

	log.Println("Exiting")
}
