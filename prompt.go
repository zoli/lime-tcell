package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/limetext/backend/keys"
	"github.com/limetext/backend/log"
)

type prompt struct {
	basicLayout
	dir              string
	submBtn, discBtn string
	ch               chan []string
	lines            []string
	selected         int
}

func newPrompt(dir string, ch chan []string, x, y, w, h int) *prompt {
	return &prompt{dir: dir, ch: ch, basicLayout: createLayout(x, y, w, h)}
}

func (p *prompt) HandleInput(kp keys.KeyPress) {
	switch kp.Key {
	case keys.Up:
		p.MoveUp()
	case keys.Down:
		p.MoveDown()
	case keys.Enter:
		p.Select()
	case keys.Escape:
		p.Discard()
	}
}

func (p *prompt) Select() {
	name := path.Join(p.dir, p.lines[p.selected])
	fi, err := os.Stat(name)
	isDir := err == nil && fi.IsDir()

	if isDir {
		p.dir = name
		p.selected = 0
	} else {
		p.Submit()
	}
}

func (p *prompt) Submit() {
	name := path.Join(p.dir, p.lines[p.selected])
	p.ch <- []string{name}
}

func (p *prompt) Discard() {
	p.ch <- nil
}

func (p *prompt) MoveUp() {
	if p.selected == 0 {
		p.selected = len(p.lines) - 1
	} else {
		p.selected--
	}

	p.Render()
}

func (p *prompt) MoveDown() {
	if p.selected == len(p.lines)-1 {
		p.selected = 0
	} else {
		p.selected++
	}

	p.Render()
}

func (p *prompt) Render() {
	p.init()
	x, y := p.Position()
	style := defStyle

	fe.screen.Clear()
	for i, l := range p.lines {
		style = defStyle
		if i == p.selected {
			style = style.Reverse(true)
		}

		runes := []rune(l)
		for _, r := range runes {
			fe.screen.setContent(&x, &y, r, style)
		}

		y++
		x = p.x
		if y > p.h-1 {
			break
		}
	}

	fe.screen.Show()
}

func (p *prompt) init() {
	fis, err := ioutil.ReadDir(p.dir)
	if err != nil {
		log.Error(err)
	}

	p.lines = make([]string, 0)
	for _, fi := range fis {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		p.lines = append(p.lines, fi.Name())
	}
}
