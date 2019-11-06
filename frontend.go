package main

import (
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
	f.screen.loop()
}

func (f *frontend) render(v *backend.View) {
	w, h := f.screen.Size()
	x, y := 0, 0

	runes := v.Substr(text.Region{0, v.Size()})
	for _, r := range runes {
		if r == '\n' {
			f.screen.SetContent(x, y, ' ', nil, defaultStyle)
		} else if r == '\t' {
			f.screen.SetContent(x, y, ' ', []rune{' ', ' ', ' ', ' ', ' ', ' ', ' '}, defaultStyle)
		} else {
			f.screen.SetContent(x, y, r, nil, defaultStyle)
		}

		x++
		if r == '\t' {
			x += 7
		}
		if r == '\n' || x > w {
			y++
			x = 0
		}
		if y > h {
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
