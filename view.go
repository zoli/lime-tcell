package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/backend/keys"
	"github.com/limetext/text"
)

type view struct {
	layout
	bv *backend.View
}

func newView(bv *backend.View, x, y, w, h int) *view {
	return &view{bv: bv, layout: createLayout(x, y, w, h)}
}

func (v *view) HandleInput(kp keys.KeyPress) {
	backend.GetEditor().HandleInput(kp)
}

func (v *view) Render(r text.Region) {
	fe.screen.Clear()
	x, y := v.x, v.y
	style, reverseStyle := defStyle, defStyle.Reverse(true)

	runes := v.bv.Substr(text.Region{0, v.bv.Size()})
	sel := v.bv.Sel()
	for i, r := range runes {
		style = defStyle
		if sel.Contains(text.Region{i, i}) {
			style = reverseStyle
		}

		fe.screen.setContent(&x, &y, r, style)

		if y > v.h-1 {
			break
		}
	}

	fe.screen.Show()
}
