package main

import (
	"github.com/limetext/backend/keys"
	"github.com/limetext/text"
)

type widget interface {
	rect
	Render(text.Region)
	VisibleRegion() text.Region
	HandleInput(keys.KeyPress)
}
