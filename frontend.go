package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/text"

	"github.com/gdamore/tcell"
)

type frontend struct {
	screen  *screen
	editor  *backend.Editor
	views   map[*backend.View]*view
	overlay widget
}

func newFrontend() (*frontend, error) {
	f := new(frontend)
	f.views = make(map[*backend.View]*view)

	s, err := newScreen()
	if err != nil {
		return nil, err
	}
	f.screen = s

	return f, nil
}

func (f *frontend) init() error {
	f.editor = initEditor()
	f.editor.SetFrontend(f)

	return f.screen.Init()
}

func (f *frontend) shutDown() {
	f.screen.Fini()
}

func (f *frontend) VisibleRegion(bv *backend.View) text.Region {
	if w := f.views[bv]; w != nil {
		return w.VisibleRegion()
	}
	return text.Region{}
}

func (f *frontend) Show(bv *backend.View, r text.Region) {
	if v := f.views[bv]; v != nil {
		v.Render(r)
	}
}

func (f *frontend) StatusMessage(s string) {

}

func (f *frontend) ErrorMessage(s string) {

}

func (f *frontend) MessageDialog(s string) {

}

func (f *frontend) OkCancelDialog(msg string, okname string) bool {
	w, h := f.screen.Size()
	ch := make(chan bool)
	d := newDialog(msg, ch, newLayout(0, 0, w, h))
	d.okBtn = okname

	f.overlay = d
	f.overlay.Render()

	ret := <-ch
	f.overlay = nil
	if w := f.editor.ActiveWindow(); w != nil && w.ActiveView() != nil {
		f.Render(w.ActiveView())
	} else {
		f.screen.Clear(defStyle)
		f.screen.Show()
	}

	return ret
}

func (f *frontend) Prompt(title, folder string, flags int) []string {
	w, h := f.screen.Size()
	ch := make(chan []string)
	p := newPrompt(folder, ch, newLayout(0, 0, w, h))

	f.overlay = p
	f.overlay.Render()

	s := <-ch
	f.overlay = nil
	if w := f.editor.ActiveWindow(); w != nil && w.ActiveView() != nil {
		f.Render(w.ActiveView())
	} else {
		f.screen.Clear(defStyle)
		f.screen.Show()
	}

	return s
}

func (f *frontend) loop() {
	for {
		switch ev := f.screen.PollEvent().(type) {
		case *tcell.EventKey:
			kp := keyPress(ev)

			if kp.Ctrl && kp.Key == 'q' {
				return
			}

			if f.overlay != nil {
				f.overlay.HandleInput(kp)
			} else {
				if kp.IsCharacter() && len(f.views) == 0 {
					f.editor.ActiveWindow().NewFile()
				}
				f.editor.HandleInput(kp)
			}
		}
	}
}

func (f *frontend) newView(bv *backend.View) {
	w, h := f.screen.Size()
	v := newView(bv, newLayout(0, 0, w, h))

	f.views[bv] = v
}

func (f *frontend) closeView(bv *backend.View) {
	delete(f.views, bv)

	if len(f.views) == 0 {
		f.screen.Clear(defStyle)
		f.screen.Show()
		return
	}

	w := f.editor.ActiveWindow()
	index, vs := 0, w.Views()
	for i, v := range vs {
		if v == bv {
			index = i
			break
		}
	}
	index--
	if index == -1 {
		index = 1
	}
	w.SetActiveView(vs[index])
}

func (f *frontend) Render(bv *backend.View) {
	r := text.Region{}
	if v := f.views[bv]; v != nil {
		r = v.VisibleRegion()
	}

	f.Show(bv, r)
}
