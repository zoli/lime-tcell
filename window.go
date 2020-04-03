package main

import "github.com/limetext/backend"

type window struct {
	bw           *backend.Window
	pane         *pane
	widgets      []widget
	activeWidget widget
}

func newWindow(bw *backend.Window) *window {
	return &window{bw: bw, widgets: make([]widget, 0)}
}

func (w *window) okCancelDialog(msg string, okname string) bool {
	return true
}

func (w *window) prompt(title, folder string, flags int) []string {
	return nil
}
