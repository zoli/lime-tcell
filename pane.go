package main

import "github.com/limetext/backend"

type pane struct {
	views map[*backend.View]*view
}

func (p *pane) newView(bv *backend.View) {

}

func (p *pane) closeView(bv *backend.View) {

}
