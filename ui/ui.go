package ui

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/joshuathompson/baton/api"
	"github.com/jroimartin/gocui"
)

type Table interface {
	render(v *gocui.View, maxX int)
	renderHeader(v *gocui.View, maxX int)
	getTableLength() int
	loadNextRecords() error
	playSelected(selectedIndex int) error
	newTableFromSelection(selectedIndex int) (Table, error)
}

var currentTable Table
var previousTables []Table
var previousCursors []int
var selectAndExit bool

func reportResult() {
	if selectAndExit {
		time.Sleep(time.Millisecond * 50)
		opts := api.Options{}
		ps, err := api.GetPlayerState(&opts)

		if err != nil {
			log.Fatal()
		}

		var artistNames []string
		for _, artist := range ps.Item.Artists {
			artistNames = append(artistNames, artist.Name)
		}

		fmt.Printf("Now playing '%s' by %s from the album %s\n", ps.Item.Name, strings.Join(artistNames, ", "), ps.Item.Album.Name)
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
	return currentTable.playSelected(y)
}

func playSelectedAndExit(g *gocui.Gui, v *gocui.View) error {
	y := getSelectedY(v)
	err := currentTable.playSelected(y)

	if err != nil {
		return err
	}

	selectAndExit = true

	return gocui.ErrQuit
}

func pushTable(g *gocui.Gui, v *gocui.View) error {
	y := getSelectedY(v)
	nt, err := currentTable.newTableFromSelection(y)

	if err != nil {
		return err
	}

	if nt != nil {
		previousCursors = append(previousCursors, y)
		previousTables = append(previousTables, currentTable)
		currentTable = nt
		v.SetCursor(0, 0)
	}
	return nil
}

func popTable(g *gocui.Gui, v *gocui.View) error {
	if len(previousTables) > 0 {
		lastIndex := previousCursors[len(previousCursors)-1]
		currentTable = previousTables[len(previousTables)-1]

		previousCursors = previousCursors[:len(previousCursors)-1]
		previousTables = previousTables[:len(previousTables)-1]

		err := v.SetCursor(0, lastIndex)

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

	v, err = g.SetView("table", -1, 2, maxX, maxY-1)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack

		err = g.SetCurrentView("table")

		if err != nil {
			return err
		}
	}

	v.Clear()
	currentTable.render(v, maxX)

	v, err = g.SetView("statusbar", -1, maxY-2, maxX, maxY)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
		v.BgColor = gocui.ColorBlue

		fmt.Fprintf(v, "[q] Quit [h] Go back [j] Down [k] Up [l] Go forward [m] Load Additional [p] Play [enter] Play and Exit")
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	err := g.SetKeybinding("", 'q', gocui.ModNone, quit)
	err = g.SetKeybinding("table", 'h', gocui.ModNone, popTable)
	err = g.SetKeybinding("table", 'j', gocui.ModNone, cursorDown)
	err = g.SetKeybinding("table", 'k', gocui.ModNone, cursorUp)
	err = g.SetKeybinding("table", 'l', gocui.ModNone, pushTable)
	err = g.SetKeybinding("table", 'p', gocui.ModNone, playSelected)
	err = g.SetKeybinding("table", gocui.KeyEnter, gocui.ModNone, playSelectedAndExit)
	err = g.SetKeybinding("table", 'm', gocui.ModNone, loadNextRecords)

	if err != nil {
		return err
	}

	return nil
}

func Run(initialTable Table) error {
	currentTable = initialTable
	defer reportResult()

	g := gocui.NewGui()
	g.Init()
	defer g.Close()

	g.SetLayout(layout)

	err := keybindings(g)

	if err != nil {
		return err
	}

	err = g.MainLoop()

	if err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
