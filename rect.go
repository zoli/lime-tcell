package main

type (
	rect interface {
		Position() (int, int)
		Dimension() (int, int)
		SetPosition(int, int)
		SetDimension(int, int)
		ZIndex() int
		SetZIndex(int)
	}

	basicRect struct {
		x, y, w, h, zi int
	}
)

func newRect(x, y, w, h int) rect {
	return &basicRect{x: x, y: y, w: w, h: h}
}

func (br *basicRect) Position() (int, int) {
	return br.x, br.y
}

func (br *basicRect) Dimension() (int, int) {
	return br.w, br.h
}

func (br *basicRect) SetPosition(x, y int) {
	br.x, br.y = x, y
}

func (br *basicRect) SetDimension(w, h int) {
	br.w, br.h = w, h
}

func (br *basicRect) ZIndex() int {
	return br.zi
}

func (br *basicRect) SetZIndex(i int) {
	br.zi = i
}
