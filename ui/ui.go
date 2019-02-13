package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// The Table interface describes the logic necessary to draw, update, and respond to a list of table items with subtables
type Table interface {
	render(v *gocui.View, maxX int)
	renderHeader(v *gocui.View, maxX int)
	renderFooter(v *gocui.View, maxX int)
	getTableLength() int
	loadNextRecords() error
	playSelected(selectedIndex int) (string, error)
	newTableFromSelection(selectedIndex int) (Table, error)
	handleSaveKey(selectedIndex int) error
}

var (
	currentTable    Table
	previousTables  []Table
	previousCursors []int
	previousOrigins []int
	chosenItem      string
)

func printNowPlaying() {
	if chosenItem != "" {
		fmt.Printf("%s", chosenItem)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func getSelectedY(v *gocui.View) int {
	_, y := v.Cursor()
	_, oy := v.Origin()

	return y + oy
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	y := getSelectedY(v)
	if y < currentTable.getTableLength()-1 {
		v.MoveCursor(0, 1, false)
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	v.MoveCursor(0, -1, false)
	return nil
}

func playSelected(g *gocui.Gui, v *gocui.View) error {
	y := getSelectedY(v)
	_, err := currentTable.playSelected(y)
	return err
}

func saveSelected(g *gocui.Gui, v *gocui.View) error {
	y := getSelectedY(v)
	err := currentTable.handleSaveKey(y)
	return err
}

func playSelectedAndExit(g *gocui.Gui, v *gocui.View) error {
	y := getSelectedY(v)
	selected, err := currentTable.playSelected(y)

	if err != nil {
		return err
	}

	chosenItem = selected

	return gocui.ErrQuit
}

func pushTable(g *gocui.Gui, v *gocui.View) error {
	y := getSelectedY(v)
	nt, err := currentTable.newTableFromSelection(y)

	if err != nil {
		return err
	}

	if nt != nil {
		_, cy := v.Cursor()
		_, oy := v.Origin()
		previousCursors = append(previousCursors, cy)
		previousOrigins = append(previousOrigins, oy)
		previousTables = append(previousTables, currentTable)
		currentTable = nt
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
	}
	return nil
}

func popTable(g *gocui.Gui, v *gocui.View) error {
	if len(previousTables) > 0 {
		lastIndex := previousCursors[len(previousCursors)-1]
		lastOrigin := previousOrigins[len(previousOrigins)-1]
		currentTable = previousTables[len(previousTables)-1]

		previousCursors = previousCursors[:len(previousCursors)-1]
		previousOrigins = previousOrigins[:len(previousOrigins)-1]
		previousTables = previousTables[:len(previousTables)-1]

		err := v.SetCursor(0, lastIndex)
		err = v.SetOrigin(0, lastOrigin)

		if err != nil {
			return err
		}
	}
	return nil
}

func loadNextRecords(g *gocui.Gui, v *gocui.View) error {
	return currentTable.loadNextRecords()
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("header", -1, -1, maxX, 3)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
	}

	v.Clear()
	currentTable.renderHeader(v, maxX)

	v, err = g.SetView("table", -1, 1, maxX, maxY-2)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack

		_, err = g.SetCurrentView("table")

		if err != nil {
			return err
		}
	}

	v.Clear()
	currentTable.render(v, maxX)

	v, err = g.SetView("footer", -1, maxY-3, maxX, maxY)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
	}

	v.Clear()
	currentTable.renderFooter(v, maxX)

	v, err = g.SetView("instructions", -1, maxY-2, maxX, maxY)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
		v.BgColor = gocui.ColorBlue

		fmt.Fprintf(v, "[q] Quit [h] Go back [j] Down [k] Up [l] Go forward [m] Load Additional [s] Save selected song/album to library [p] Play [enter] Play and Exit")
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	err := g.SetKeybinding("", 'q', gocui.ModNone, quit)
	err = g.SetKeybinding("table", 'h', gocui.ModNone, popTable)
	err = g.SetKeybinding("table", gocui.KeyArrowLeft, gocui.ModNone, popTable)
	err = g.SetKeybinding("table", 'j', gocui.ModNone, cursorDown)
	err = g.SetKeybinding("table", gocui.KeyArrowDown, gocui.ModNone, cursorDown)
	err = g.SetKeybinding("table", 'k', gocui.ModNone, cursorUp)
	err = g.SetKeybinding("table", gocui.KeyArrowUp, gocui.ModNone, cursorUp)
	err = g.SetKeybinding("table", 'l', gocui.ModNone, pushTable)
	err = g.SetKeybinding("table", gocui.KeyArrowRight, gocui.ModNone, pushTable)
	err = g.SetKeybinding("table", 'p', gocui.ModNone, playSelected)
	err = g.SetKeybinding("table", gocui.KeyEnter, gocui.ModNone, playSelectedAndExit)
	err = g.SetKeybinding("table", 'm', gocui.ModNone, loadNextRecords)
	err = g.SetKeybinding("table", 's', gocui.ModNone, saveSelected)

	if err != nil {
		return err
	}

	return nil
}

// Run starts the TUI for a struct that implements the table interface.  The prebuilt tables are for artists, albums, tracks, and playlists.
// Run will block indefinitely until returning an error or nil
func Run(initialTable Table) error {
	currentTable = initialTable
	defer printNowPlaying()

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		return err
	}

	defer g.Close()

	g.SetManagerFunc(layout)

	err = keybindings(g)

	if err != nil {
		return err
	}

	err = g.MainLoop()

	if err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
