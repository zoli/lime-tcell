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
