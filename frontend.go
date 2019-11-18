package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/text"

	"github.com/gdamore/tcell"
)

type frontend struct {
	screen       *screen
	editor       *backend.Editor
	widgets      []widget
	activeWidget widget
}

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

	activeLay := f.activeWidget
	f.addWidget(p)
	f.activeWidget.Render(text.Region{})
	s := <-ch

	f.activeWidget = activeLay

	return []string{s}
}

func (f *frontend) addWidget(w widget) {
	f.widgets = append(f.widgets, w)
	f.activeWidget = w
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
				if f.activeWidget != nil {
					f.activeWidget.HandleInput(kp)
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

	f.addWidget(v)
	f.activeWidget.Render(text.Region{})
}
