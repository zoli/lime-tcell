package main

import "github.com/limetext/backend/keys"

type (
	widget interface {
		HandleInput(keys.KeyPress)
		Render()
		layout
	}

	layout interface {
		Position() (int, int)
		Dimension() (int, int)
		SetPosition(int, int)
		SetDimension(int, int)
	}

	basicLayout struct {
		x, y, w, h int
	}
)

func newLayout(x, y, w, h int) layout {
	return &basicLayout{x, y, w, h}
}

func (bl *basicLayout) Position() (int, int) {
	return bl.x, bl.y
}

func (bl *basicLayout) Dimension() (int, int) {
	return bl.w, bl.h
}

func (bl *basicLayout) SetPosition(x, y int) {
	bl.x, bl.y = x, y
}

func (bl *basicLayout) SetDimension(w, h int) {
	bl.w, bl.h = w, h
}
