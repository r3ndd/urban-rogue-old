package gui

import (
	"github.com/jroimartin/gocui"
)

var layouts []gocui.Manager = []gocui.Manager{}

func Run() {
	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		panic(err)
	}

	defer g.Close()

	g.SetManager(layouts...)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func AddLayout(layout gocui.Manager) {
	layouts = append(layouts, layout)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
