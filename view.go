package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/backend/keys"
	"github.com/limetext/text"
)

type view struct {
	basicLayout
	bv *backend.View
	vr text.Region
}

func newView(bv *backend.View, x, y, w, h int) *view {
	return &view{bv: bv, basicLayout: createLayout(x, y, w, h)}
}

func (v *view) HandleInput(kp keys.KeyPress) {
	backend.GetEditor().HandleInput(kp)
}

func (v *view) Render(r text.Region) {
	fe.screen.Clear()
	v.setVisibleRegion(r)

	x, y := v.Position()
	_, h := v.Dimension()
	style := defStyle
	runes, sel := v.bv.Substr(v.VisibleRegion()), v.bv.Sel()
	for i, r := range runes {
		style = defStyle
		if sel.Contains(text.Region{i, i}) {
			style = style.Reverse(true)
		}

		fe.screen.setContent(&x, &y, r, style)

		if y > h-1 {
			break
		}
	}

	fe.screen.Show()
}

func (v *view) setVisibleRegion(r text.Region) {
	_, h := v.Dimension()
	row, _ := v.bv.RowCol(r.Begin())

	begin := v.bv.TextPoint(row, 0)

	endPoint := v.bv.TextPoint(row+h, 0)
	endLineR := v.bv.Line(endPoint)
	end := endLineR.End()

	v.vr = text.Region{begin, end}
}

func (v *view) VisibleRegion() text.Region {
	return v.vr
}
