package main

import "github.com/limetext/backend"

func initEditor() {
	setCallBacks()

	ed := backend.GetEditor()
	ed.NewWindow()
}

func setCallBacks() {
	backend.OnNew.Add(fe.render)
	backend.OnModified.Add(fe.render)
	backend.OnSelectionModified.Add(fe.render)
}
