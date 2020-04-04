package main

import "github.com/limetext/backend"

type pane struct {
	rect
	views map[*backend.View]*view
}

func newPane(rec rect) *pane {
	return &pane{rect: rec}
}

func (p *pane) newView(bv *backend.View) {
	p.views[bv] = newView(bv, p.rect)
}

func (p *pane) closeView(bv *backend.View) {
	delete(p.views, bv)
}
