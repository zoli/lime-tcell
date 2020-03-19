package main

import (
	"fmt"

	"github.com/limetext/backend/render"

	"github.com/gdamore/tcell"
)

func color(c render.Colour) tcell.Color {
	hex := fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
	return tcell.GetColor(hex)
}
