package main

import "github.com/limetext/backend/keys"

type dialog struct {
	rect
	msg              string
	okBtn, cancelBtn string
	ch               chan bool
	ok               bool
}

func newDialog(msg string, ch chan bool, rec rect) *dialog {
	return &dialog{
		msg:   msg,
		ch:    ch,
		okBtn: "Ok", cancelBtn: "Cancel",
		rect: rec,
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
	w, _ := d.Dimension()
	x, y := d.Position()
	style := defStyle

	fe.scrn.Clear(defStyle)

	for _, r := range d.msg {
		fe.scrn.setContent(&x, &y, r, w, style)
	}
	x = 0
	y += 2

	if d.ok {
		style = style.Reverse(true)
	}
	for _, r := range d.okBtn {
		fe.scrn.setContent(&x, &y, r, w, style)
	}
	x += 4

	style = defStyle
	if !d.ok {
		style = style.Reverse(true)
	}
	for _, r := range d.cancelBtn {
		fe.scrn.setContent(&x, &y, r, w, style)
	}

	fe.scrn.Show()
}

func (d *dialog) Toggle() {
	d.ok = !d.ok
}

func (d *dialog) Select() {
	d.ch <- d.ok
}
