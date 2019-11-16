package main

import "github.com/limetext/backend"

func initEditor() *backend.Editor {
	setCallBacks()

	ed := backend.GetEditor()
	ed.AddPackagesPath("./packages")
	ed.Init()
	ed.SetDefaultPath("./packages/Default")
	ed.SetUserPath("./packages/User")
	ed.NewWindow()

	return ed
}

func setCallBacks() {
	backend.OnNew.Add(fe.render)
	backend.OnModified.Add(fe.render)
	backend.OnSelectionModified.Add(fe.render)
}
