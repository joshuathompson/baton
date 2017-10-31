package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type ChannelRoot struct {
	Channels []Channel `json:"channels"`
}

type Channel struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DJ          string     `json:"dj"`
	DJMail      string     `json:"djmail"`
	Genre       string     `json:"genre"`
	Image       string     `json:"image"`
	LargeImage  string     `json:"largeimage"`
	XLImage     string     `json:"xlimage"`
	Twitter     string     `json:"twitter"`
	Updated     uint64     `json:"updated,string"`
	Playlists   []Playlist `json:"playlists"`
	Listeners   uint32     `json:"listeners,string"`
	LastPlaying string     `json:"lastPlaying"`
}

type Playlist struct {
	URL     string `json:"url"`
	Format  string `json:"format"`
	Quality string `json:"quality"`
}

func main() {
	channels, err := getChannels()

	if err != nil {
		log.Fatal(err)
	}

	playlistsArg := buildPlaylistsArg(channels)
	shell := os.Getenv("SHELL")
	mpv := fmt.Sprintf("mpv --idle --input-ipc-server=%s --playlist=%s --volume 30 &", "/tmp/somafm", playlistsArg)

	cmd := exec.Command(shell, "-c", mpv)

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mpv successfully started")
}

func buildPlaylistsArg(channels []Channel) string {
	var playlists []string

	for _, channel := range channels {
		for _, playlist := range channel.Playlists {
			playlists = append(playlists, playlist.URL)
		}
	}

	return strings.Join(playlists, " ")
}

func getChannels() ([]Channel, error) {
	r, err := http.Get("https://somafm.com/channels.json")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	var root ChannelRoot
	err = json.NewDecoder(r.Body).Decode(&root)

	return root.Channels, err
}
