package main

import "github.com/gdamore/tcell"

var (
	defaultStyle = tcell.StyleDefault
)

type screen struct {
	tcell.Screen
}

func newScreen() (*screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err = s.Init(); err != nil {
		return nil, err
	}

	return &screen{Screen: s}, nil
}

func (s *screen) loop() {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlQ:
				return
			}
		}
	}
}

func (s *screen) setContent(xp, yp *int, r rune, style tcell.Style) {
	w, _ := s.Size()
	x, y := *xp, *yp

	ch := r
	if r == '\t' || r == '\n' {
		ch = '\x00'
	}

	l := 1
	if r == '\t' {
		l = 8
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
