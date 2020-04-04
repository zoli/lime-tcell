package main

import "github.com/limetext/backend"

type window struct {
	rect
	bw           *backend.Window
	pane         *pane
	widgets      []widget
	activeWidget widget
}

func newWindow(bw *backend.Window, w, h int) *window {
	return &window{
		rect:    newRect(0, 0, w, h),
		bw:      bw,
		pane:    newPane(newRect(0, 0, w, h)),
		widgets: make([]widget, 0),
	}
}

func (w *window) okCancelDialog(msg string, okname string) bool {
	w, h := fe.scrn.Size()
	ch := make(chan bool)
	d := newDialog(msg, ch, newLayout(0, 0, w, h))
	d.okBtn = okname

	fe.overlay = d
	fe.overlay.Render()

	ret := <-ch
	fe.overlay = nil
	if w := fe.editor.ActiveWindow(); w != nil && w.ActiveView() != nil {
		fe.Render(w.ActiveView())
	} else {
		fe.scrn.Clear(defStyle)
		fe.scrn.Show()
	}

	return ret
}

func (w *window) prompt(title, folder string, flags int) []string {
	w, h := fe.scrn.Size()
	ch := make(chan []string)
	p := newPrompt(folder, ch, newLayout(0, 0, w, h))

	fe.overlay = p
	fe.overlay.Render()

	s := <-ch
	fe.overlay = nil
	if w := fe.editor.ActiveWindow(); w != nil && w.ActiveView() != nil {
		fe.Render(w.ActiveView())
	} else {
		fe.scrn.Clear(defStyle)
		fe.scrn.Show()
	}

	return s
}
