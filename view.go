package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/backend/log"
	"github.com/limetext/text"
)

type view struct {
	layout
	bv *backend.View
	vr text.Region
}

func newView(bv *backend.View, lay layout) *view {
	return &view{bv: bv, layout: lay}
}

func (v *view) Render(r text.Region) {
	v.calcVisibleRegion(r)
	log.Finest("Rendering %s, %s", v.bv, v.vr)

	fe.screen.Clear()

	w, _ := v.Dimension()
	x, y := v.Position()

	runes, sel := v.bv.Substr(v.VisibleRegion()), v.bv.Sel()
	for i, r := range runes {
		style := defStyle
		if sel.Contains(text.Region{i, i}) {
			style = style.Reverse(true)
		}

		fe.screen.setContent(&x, &y, r, w, style)
	}

	fe.screen.Show()
}

func (v *view) calcVisibleRegion(r text.Region) {
	_, h := v.Dimension()

	row, _ := v.bv.RowCol(r.Begin())
	begin := v.bv.TextPoint(row, 0)

	endPoint := v.bv.TextPoint(row+h-1, 0)
	endLineR := v.bv.Line(endPoint)
	end := endLineR.End()

	v.vr = text.Region{begin, end}
}

func (v *view) VisibleRegion() text.Region {
	return v.vr
}
