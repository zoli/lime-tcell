package main

import (
	"github.com/gdamore/tcell"
	"github.com/limetext/backend"
	"github.com/limetext/text"
)

type frontend struct {
	screen *screen
	editor *backend.Editor
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
	initEditor()
	f.editor = backend.GetEditor()
	f.editor.SetFrontend(f)
}

func (f *frontend) shutDown() {
	for _, w := range f.editor.Windows() {
		w.CloseAllViews()
		w.Close()
	}

	f.screen.Fini()
}

func (f *frontend) loop() {
	for {
		ev := f.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlQ:
				return
			default:
				kp := keyPress(ev)
				f.editor.HandleInput(kp)
			}
		}
	}
}

func (f *frontend) render(v *backend.View) {
	_, h := f.screen.Size()
	x, y := 0, 0
	style, reverseStyle := defaultStyle, defaultStyle.Reverse(true)

	runes := v.Substr(text.Region{0, v.Size()})
	sel := v.Sel()
	for i, r := range runes {
		style = defaultStyle
		if sel.Contains(text.Region{i, i}) {
			style = reverseStyle
		}

		f.screen.setContent(&x, &y, r, style)

		if y > h-1 {
			break
		}
	}

	f.screen.Show()
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
	return nil
}
