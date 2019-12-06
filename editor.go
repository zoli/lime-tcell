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
	backend.OnNew.Add(fe.newView)

	backend.OnClose.Add(fe.closeView)

	backend.OnActivated.Add(fe.Render)
	backend.OnModified.Add(fe.Render)
	backend.OnSelectionModified.Add(fe.Render)
}
