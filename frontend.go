package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/backend/keys"
	"github.com/limetext/text"

	"github.com/gdamore/tcell"
)

type (
	frontend struct {
		screen       *screen
		editor       *backend.Editor
		layouts      []layout
		activeLayout layout
	}

	layout interface {
		HandleInput(keys.KeyPress)
		Render(text.Region)
		Position() (int, int)
		Dimension() (int, int)
		SetPosition(int, int)
		SetDimension(int, int)
	}
)

func newFrontend() (*frontend, error) {
	f := new(frontend)

	s, err := newScreen()
	if err != nil {
		return nil, err
	}
	f.screen = s

	return f, nil
}

func (f *frontend) init() {
	f.editor = initEditor()
	f.editor.SetFrontend(f)
}

func (f *frontend) shutDown() {
	for _, w := range f.editor.Windows() {
		w.CloseAllViews()
		w.Close()
	}

	f.screen.Fini()
}

func (f *frontend) VisibleRegion(v *backend.View) text.Region {
	return text.Region{}
}

func (f *frontend) Show(v *backend.View, r text.Region) {

}

func (f *frontend) StatusMessage(s string) {

}

func (f *frontend) ErrorMessage(s string) {

}

func (f *frontend) MessageDialog(s string) {

}

func (f *frontend) OkCancelDialog(msg string, okname string) bool {
	return false
}

func (f *frontend) Prompt(title, folder string, flags int) []string {
	w, h := f.screen.Size()
	ch := make(chan string)
	p := newPrompt(folder, ch, 0, 0, w, h)

	activeLay := f.activeLayout
	f.addLayout(p)
	f.activeLayout.Render(text.Region{})
	s := <-ch

	f.activeLayout = activeLay

	return []string{s}
}

func (f *frontend) addLayout(lay layout) {
	f.layouts = append(f.layouts, lay)
	f.activeLayout = lay
}

func (f *frontend) loop() {
	for {
		ev := f.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			kp := keyPress(ev)
			switch ev.Key() {
			case tcell.KeyCtrlQ:
				return
			default:
				if f.activeLayout != nil {
					f.activeLayout.HandleInput(kp)
				} else {
					f.editor.HandleInput(kp)
				}
			}
		}
	}
}
