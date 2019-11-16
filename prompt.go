package main

import "github.com/limetext/backend/keys"

type prompt struct {
	dir                   string
	lines, selected       []string
	submitBtn, discardBtn string
}

func (p *prompt) Select() {

}

func (p *prompt) Submit() {

}

func (p *prompt) Discard() {

}

func (p *prompt) HandleInput(kp keys.KeyPress) {

}
