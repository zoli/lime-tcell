package main

import (
	"github.com/gdamore/tcell"
)

var (
	defStyle = tcell.StyleDefault
)

type screen struct {
	tcell.Screen
}

func newScreen() (*screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	return &screen{Screen: s}, nil
}

func (s *screen) setContent(xp, yp *int, r rune, w int, style tcell.Style) {
	x, y := *xp, *yp

	ch := r
	if r == '\t' || r == '\n' {
		ch = '\x00'
	}

	l := 1
	if r == '\t' {
		l = 4
	}

	for i := 0; i < l; i++ {
		s.SetContent(x, y, ch, nil, style)
		x++
		if x > w-1 {
			x = 0
			y++
		}
	}

	if r == '\n' {
		y++
		x = 0
	}

	*xp, *yp = x, y
}

func (s *screen) Clear(style tcell.Style) {
	s.Screen.Fill(' ', style)
}
