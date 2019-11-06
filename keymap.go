package main

import (
	"github.com/gdamore/tcell"
	"github.com/limetext/backend/keys"
	"github.com/limetext/backend/log"
)

var keymap = map[tcell.Key]keys.Key{
	tcell.KeyF1:     keys.F1,
	tcell.KeyF2:     keys.F2,
	tcell.KeyF3:     keys.F3,
	tcell.KeyF4:     keys.F4,
	tcell.KeyF5:     keys.F5,
	tcell.KeyF6:     keys.F6,
	tcell.KeyF7:     keys.F7,
	tcell.KeyF8:     keys.F8,
	tcell.KeyF9:     keys.F9,
	tcell.KeyF10:    keys.F10,
	tcell.KeyF11:    keys.F11,
	tcell.KeyF12:    keys.F12,
	tcell.KeyInsert: keys.Insert,
	tcell.KeyPgUp:   keys.PageUp,
	tcell.KeyPgDn:   keys.PageDown,
	tcell.KeyHome:   keys.Home,
	tcell.KeyEnd:    keys.End,
}

func keyPress(ek *tcell.EventKey) keys.KeyPress {
	var kp keys.KeyPress

	mods := ek.Modifiers()
	kp.Alt = mods == tcell.ModAlt
	kp.Ctrl = mods == tcell.ModCtrl
	kp.Shift = mods == tcell.ModShift

	if key := ek.Key(); key == tcell.KeyRune {
		kp.Text = string(ek.Rune())
	} else if k, ok := keymap[key]; ok {
		kp.Key = k
	} else {
		log.Warn("unrecognized key %s %s", ek, kp)
	}

	return kp
}
