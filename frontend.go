package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/text"

	"github.com/gdamore/tcell"
)

type frontend struct {
	scrn    *screen
	ed      *backend.Editor
	windows map[*backend.Window]*window
}

func newFrontend() (*frontend, error) {
	f := &frontend{windows: make(map[*backend.Window]*window)}

	s, err := newScreen()
	if err != nil {
		return nil, err
	}
	f.scrn = s

	return f, nil
}

func (f *frontend) init() error {
	f.setCallBacks()
	f.initEditor()
	f.ed.SetFrontend(f)

	return f.scrn.Init()
}

func (f *frontend) setCallBacks() {
	backend.OnNewWindow.Add(fe.addWindow)

	backend.OnNew.Add(fe.newView)
	backend.OnClose.Add(fe.closeView)
	backend.OnActivated.Add(fe.Render)
	backend.OnModified.Add(fe.Render)
	backend.OnSelectionModified.Add(fe.Render)
}

func (f *frontend) initEditor() {
	ed := backend.GetEditor()
	ed.Init()
	ed.SetDefaultPath("./packages/Default")
	ed.SetUserPath("./packages/User")
	ed.AddPackagesPath("./packages")
	ed.NewWindow()

	f.ed = ed
}

func (f *frontend) shutDown() {
	f.scrn.Fini()
}

func (f *frontend) VisibleRegion(bv *backend.View) text.Region {
	if v := f.view(bv); v != nil {
		return v.VisibleRegion()
	}
	return text.Region{}
}

func (f *frontend) Show(bv *backend.View, r text.Region) {
	if v := f.view(bv); v != nil {
		v.Render(r)
	}
}

func (f *frontend) StatusMessage(s string) {}

func (f *frontend) ErrorMessage(s string) {}

func (f *frontend) MessageDialog(s string) {}

func (f *frontend) OkCancelDialog(msg string, okname string) bool {
	return f.activeWindow().okCancelDialog(msg, okname)
}

func (f *frontend) Prompt(title, folder string, flags int) []string {
	return f.activeWindow().prompt(title, folder, flags)
}

func (f *frontend) loop() {
	for {
		switch ev := f.scrn.PollEvent().(type) {
		case *tcell.EventKey:
			kp := keyPress(ev)
			if kp.Ctrl && kp.Key == 'q' {
				return
			}

			if w := f.activeWindow().activeWidget; w != nil {
				w.HandleInput(kp)
				break
			}

			if kp.IsCharacter() && len(f.activeWindow().pane.views) == 0 {
				f.ed.ActiveWindow().NewFile()
			}
			f.ed.HandleInput(kp)
		}
	}
}

func (f *frontend) Render(bv *backend.View) {
	r := text.Region{}
	if v := f.view(bv); v != nil {
		r = v.VisibleRegion()
	}

	f.Show(bv, r)
}

func (f *frontend) addWindow(bw *backend.Window) {
	w := newWindow(bw)
	f.windows[bw] = w
}

func (f *frontend) window(bw *backend.Window) *window {
	return f.windows[bw]
}

func (f *frontend) view(bv *backend.View) *view {
	return f.window(bv.Window()).pane.views[bv]
}

func (f *frontend) activeWindow() *window {
	return f.windows[f.ed.ActiveWindow()]
}

func (f *frontend) newView(bv *backend.View) {
	f.window(bv.Window()).pane.newView(bv)
}

func (f *frontend) closeView(bv *backend.View) {
	f.window(bv.Window()).pane.closeView(bv)
}
