package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Table interface {
	render(v *gocui.View, maxX int)
	renderHeader(v *gocui.View, maxX int)
	getTableLength() int
	loadNextRecords() error
	playSelected() error
	newTableFromSelection() (Table, error)
}

var tables []Table
var currentTable Table

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	_, y := v.Cursor()
	if y < currentTable.getTableLength()-1 {
		v.MoveCursor(0, 1, false)
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	v.MoveCursor(0, -1, false)
	return nil
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

		_, err = g.SetCurrentView("table")

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

		fmt.Fprintf(v, "[q] Quit [j] Down [k] Up [p] Play [enter] Load items for selected [backspace] Go back")
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	err := g.SetKeybinding("", 'q', gocui.ModNone, quit)

	if err != nil {
		return err
	}

	err = g.SetKeybinding("table", 'j', gocui.ModNone, cursorDown)

	if err != nil {
		return err
	}

	err = g.SetKeybinding("table", 'k', gocui.ModNone, cursorUp)

	if err != nil {
		return err
	}

	return nil
}

func Run(initialTable Table) error {
	tables = append(tables, initialTable)
	currentTable = initialTable

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
