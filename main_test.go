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
	fe.screen.Screen = tcell.NewSimulationScreen("")

	err = fe.init()
	if err != nil {
		panic(err)
	}
}

func postQuitEvent() {
	ev := tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModCtrl)
	fe.screen.PostEventWait(ev)
}

func closeAll() {
	ws := fe.editor.Windows()
	for _, w := range ws {
		vs := w.Views()
		for _, v := range vs {
			v.SetScratch(true)
		}

		w.Close()
	}

	fe.editor.NewWindow()
}

func getContent(x, y int) rune {
	r, _, _, _ := fe.screen.GetContent(x, y)
	return r
}
