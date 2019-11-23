package main

type (
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

func createLayout(x, y, w, h int) basicLayout {
	return basicLayout{x, y, w, h}
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
