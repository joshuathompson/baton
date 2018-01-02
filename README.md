# Baton
> Baton is a CLI tool to manage Spotify playback and includes a CUI for searches  

[![asciicast](https://asciinema.org/a/RgR4iT9wz2J3bjCx0p7Wj9Rnn.png)](https://asciinema.org/a/RgR4iT9wz2J3bjCx0p7Wj9Rnn)

## NOTE:  You can't change songs, volume, etc without Spotify premium due to a limit on their web API

## Install
Fetch the release for your platform [from the following page](https://github.com/joshuathompson/baton/releases).  Use `chmod` to set permissions on your binary and move it to `/usr/local/bin` or equivalent.

## Setup
Use the `baton auth` command to perform an initial setup.  The command will take you through the process but you will need to login to the [Spotify API dashboard](https://beta.developer.spotify.com/dashboard/login), create an app, and set it up with a redirect URL of http://localhost:15298/callback.

This process will generate a long-lasting refresh token and ideally will never have to be repeated.

## Usage

### CLI Commands

Command              | Description
---------------------|---------------------------------------
auth                 | authorize Baton to access the Spotify Web API on your behalf
devices              | list all available playback devices
help                 | help about any command
next                 | skip to next track
pause                | toggle Spotify pause state
play                 | play top result for specified artist, album, playlist, track, or uri
prev                 | skip to previous track
repeat               | get/set repeat mode
replay               | replay current track from the beginning
search               | search for specified artist, album, playlist, or track and select via interactive CUI
seek                 | skip to a specific time (seconds) of the current track
share                | get uri and url for current track
shuffle              | toggle shuffle on/off
status               | show information about the current track
transfer             | transfer playback to another device by id
vol                  | get/set volume

### CUI Keybinds

Keybind              | Description
---------------------|---------------------------------------
<kbd>h</kbd>         | go back one screen
<kbd>j</kbd>         | move the cursor down a line
<kbd>k</kbd>         | move the cursor up a line
<kbd>l</kbd>         | go into playlist, album, or artist
<kbd>p</kbd>         | play selected item
<kbd>Enter</kbd>     | play selected item and quit
<kbd>m</kbd>         | load additional pages from search query
<kbd>q</kbd>         | quit

## License
MIT 

## Credits
Baton is built using:
* [Cobra](https://github.com/spf13/cobra)
* [Gocui](https://github.com/jroimartin/gocui)
