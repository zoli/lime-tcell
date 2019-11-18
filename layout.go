package main

import (
	"github.com/limetext/backend/keys"
	"github.com/limetext/text"
)

type (
	widget interface {
		HandleInput(keys.KeyPress)
		Render(text.Region)
		Position() (int, int)
		Dimension() (int, int)
		SetPosition(int, int)
		SetDimension(int, int)
	}

	layout struct {
		x, y, w, h int
	}
)

func createLayout(x, y, w, h int) layout {
	return layout{x, y, w, h}
}

func (l *layout) Position() (int, int) {
	return l.x, l.y
}

func (l *layout) Dimension() (int, int) {
	return l.w, l.h
}

func (l *layout) SetPosition(x, y int) {
	l.x, l.y = x, y
}

func (l *layout) SetDimension(w, h int) {
	l.w, l.h = w, h
}
