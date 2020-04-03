package main

import (
	"github.com/gdamore/tcell"
	"github.com/limetext/backend/log"
)

func init() {
	log.AddFilter("stdout", log.FINEST, log.NewFileLogWriter("debug.log", false))
	newFe()
}

func newFe() {
	var err error
	fe, err = newFrontend()
	if err != nil {
		panic(err)
	}

	scrn := tcell.NewSimulationScreen("")

	fe.scrn.Screen = scrn
	err = fe.init()
	if err != nil {
		panic(err)
	}

	scrn.SetSize(150, 23)
}

func postQuitEvent() {
	ev := tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModCtrl)
	fe.scrn.PostEventWait(ev)
}

func closeAll() {
	ws := fe.ed.Windows()
	for _, w := range ws {
		vs := w.Views()
		for _, v := range vs {
			v.SetScratch(true)
		}

		w.Close()
	}

	fe.ed.NewWindow()
}

func getContent(x, y int) rune {
	r, _, _, _ := fe.scrn.GetContent(x, y)
	return r
}
