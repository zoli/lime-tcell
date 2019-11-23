package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/backend/keys"
	"github.com/limetext/text"

	"github.com/gdamore/tcell"
)

type (
	frontend struct {
		screen  *screen
		editor  *backend.Editor
		widgets map[*backend.View]widget
		overlay widget
	}

	widget interface {
		HandleInput(keys.KeyPress)
		Render(text.Region)
		VisibleRegion() text.Region
		layout
	}
)

func newFrontend() (*frontend, error) {
	f := new(frontend)
	f.widgets = make(map[*backend.View]widget)

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

func (f *frontend) VisibleRegion(bv *backend.View) text.Region {
	if w := f.widgets[bv]; w != nil {
		return w.VisibleRegion()
	}
	return text.Region{}
}

func (f *frontend) Show(bv *backend.View, r text.Region) {
	if w := f.widgets[bv]; w != nil {
		w.Render(r)
	}
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

	hold := f.overlay
	f.overlay = p
	f.overlay.Render(text.Region{})

	s := <-ch
	f.overlay = hold

	return []string{s}
}

func (f *frontend) loop() {
	for {
		switch ev := f.screen.PollEvent().(type) {
		case *tcell.EventKey:
			kp := keyPress(ev)
			switch ev.Key() {
			case tcell.KeyCtrlQ:
				return
			default:
				if f.overlay != nil {
					f.overlay.HandleInput(kp)
				} else {
					f.editor.HandleInput(kp)
				}
			}
		}
	}
}

func (f *frontend) newView(bv *backend.View) {
	w, h := f.screen.Size()
	v := newView(bv, 0, 0, w, h)

	f.widgets[bv] = v
	f.Render(bv)
}

func (f *frontend) Render(bv *backend.View) {
	r := text.Region{}
	if w := f.widgets[bv]; w != nil {
		r = w.VisibleRegion()
	}

	f.Show(bv, r)
}
