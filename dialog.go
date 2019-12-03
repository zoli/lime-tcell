package main

import "github.com/limetext/backend/keys"

type dialog struct {
	layout
	msg              string
	okBtn, cancelBtn string
	ch               chan bool
	ok               bool
}

func newDialog(msg string, ch chan bool, lay layout) *dialog {
	return &dialog{
		msg:   msg,
		ch:    ch,
		okBtn: "Ok", cancelBtn: "Cancel",
		layout: lay,
	}
}

func (d *dialog) HandleInput(kp keys.KeyPress) {
	switch kp.Key {
	case keys.Left, keys.Right, keys.Down, keys.Up:
		d.Toggle()
	case keys.Enter:
		d.Select()
		return
	}

	d.Render()
}

func (d *dialog) Render() {
	x, y := d.Position()
	style := defStyle

	fe.screen.Clear()

	for _, r := range d.msg {
		fe.screen.setContent(&x, &y, r, style)
	}
	x = 0
	y += 2

	if d.ok {
		style = style.Reverse(true)
	}
	for _, r := range d.okBtn {
		fe.screen.setContent(&x, &y, r, style)
	}
	x += 4

	style = defStyle
	if !d.ok {
		style = style.Reverse(true)
	}
	for _, r := range d.cancelBtn {
		fe.screen.setContent(&x, &y, r, style)
	}

	fe.screen.Show()
}

func (d *dialog) Toggle() {
	d.ok = !d.ok
}

func (d *dialog) Select() {
	d.ch <- d.ok
}
