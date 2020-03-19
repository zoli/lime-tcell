package main

import (
	"github.com/limetext/backend"
	"github.com/limetext/backend/log"
	"github.com/limetext/text"

	"github.com/gdamore/tcell"
)

type view struct {
	layout
	bv    *backend.View
	vr    text.Region
	style tcell.Style
}

func newView(bv *backend.View, lay layout) *view {
	v := &view{bv: bv, layout: lay}
	v.bv.Settings().AddOnChange("color_scheme", func(key string) {
		if key != "color_scheme" {
			return
		}

		v.loadStyle()

		fe.Render(v.bv)
	})
	v.bv.Settings().AddOnChange("syntac_update", func(key string) {
		if key != "lime.syntax.updated" {

		}

		fe.Render(v.bv)
	})
	v.loadStyle()
	return v
}

func (v *view) Render(r text.Region) {
	v.calcVisibleRegion(r)
	log.Finest("Rendering %s, %s", v.bv, v.vr)

	fe.screen.Clear(v.style)

	w, _ := v.Dimension()
	x, y := v.Position()
	vr, sel := v.VisibleRegion(), v.bv.Sel()
	recipe := v.bv.Transform(vr).Transcribe()

	runes := v.bv.SubstrR(vr)
	for i, r := range runes {
		point := vr.Begin() + i
		style := v.style

		for j := 0; j < len(recipe) && point >= recipe[j].Region.Begin(); j++ {
			if point > recipe[j].Region.End() {
				continue
			}

			fg := color(recipe[j].Flavour.Foreground)
			bg := color(recipe[j].Flavour.Background)
			style = style.Foreground(fg)
			style = style.Background(bg)
		}

		if sel.Contains(text.Region{point, point}) {
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

func (v *view) loadStyle() {
	v.style = defStyle

	cs := v.bv.Settings().String("color_scheme", "")
	scheme := backend.GetEditor().GetColorScheme(cs)

	fg := color(scheme.GlobalSettings().Foreground)
	bg := color(scheme.GlobalSettings().Background)
	v.style = v.style.Foreground(fg)
	v.style = v.style.Background(bg)
}
