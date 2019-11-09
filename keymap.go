package main

import (
	"github.com/gdamore/tcell"
	"github.com/limetext/backend/keys"
	"github.com/limetext/backend/log"
)

var keymap = map[tcell.Key]keys.KeyPress{
	tcell.KeyUp:     keys.KeyPress{Key: keys.Up},
	tcell.KeyDown:   keys.KeyPress{Key: keys.Down},
	tcell.KeyRight:  keys.KeyPress{Key: keys.Right},
	tcell.KeyLeft:   keys.KeyPress{Key: keys.Left},
	tcell.KeyPgUp:   keys.KeyPress{Key: keys.PageUp},
	tcell.KeyPgDn:   keys.KeyPress{Key: keys.PageDown},
	tcell.KeyHome:   keys.KeyPress{Key: keys.Home},
	tcell.KeyEnd:    keys.KeyPress{Key: keys.End},
	tcell.KeyInsert: keys.KeyPress{Key: keys.Insert},
	tcell.KeyDelete: keys.KeyPress{Key: keys.Delete},
	tcell.KeyF1:     keys.KeyPress{Key: keys.F1},
	tcell.KeyF2:     keys.KeyPress{Key: keys.F2},
	tcell.KeyF3:     keys.KeyPress{Key: keys.F3},
	tcell.KeyF4:     keys.KeyPress{Key: keys.F4},
	tcell.KeyF5:     keys.KeyPress{Key: keys.F5},
	tcell.KeyF6:     keys.KeyPress{Key: keys.F6},
	tcell.KeyF7:     keys.KeyPress{Key: keys.F7},
	tcell.KeyF8:     keys.KeyPress{Key: keys.F8},
	tcell.KeyF9:     keys.KeyPress{Key: keys.F9},
	tcell.KeyF10:    keys.KeyPress{Key: keys.F10},
	tcell.KeyF11:    keys.KeyPress{Key: keys.F11},
	tcell.KeyF12:    keys.KeyPress{Key: keys.F12},

	tcell.KeyCtrlSpace:     keys.KeyPress{Ctrl: true, Key: ' '},
	tcell.KeyCtrlA:         keys.KeyPress{Ctrl: true, Key: 'a'},
	tcell.KeyCtrlB:         keys.KeyPress{Ctrl: true, Key: 'b'},
	tcell.KeyCtrlC:         keys.KeyPress{Ctrl: true, Key: 'c'},
	tcell.KeyCtrlD:         keys.KeyPress{Ctrl: true, Key: 'd'},
	tcell.KeyCtrlE:         keys.KeyPress{Ctrl: true, Key: 'e'},
	tcell.KeyCtrlF:         keys.KeyPress{Ctrl: true, Key: 'f'},
	tcell.KeyCtrlG:         keys.KeyPress{Ctrl: true, Key: 'g'},
	tcell.KeyCtrlH:         keys.KeyPress{Ctrl: true, Key: 'h'},
	tcell.KeyCtrlI:         keys.KeyPress{Ctrl: true, Key: 'i'},
	tcell.KeyCtrlJ:         keys.KeyPress{Ctrl: true, Key: 'j'},
	tcell.KeyCtrlK:         keys.KeyPress{Ctrl: true, Key: 'k'},
	tcell.KeyCtrlL:         keys.KeyPress{Ctrl: true, Key: 'l'},
	tcell.KeyCtrlM:         keys.KeyPress{Ctrl: true, Key: 'm'},
	tcell.KeyCtrlN:         keys.KeyPress{Ctrl: true, Key: 'n'},
	tcell.KeyCtrlO:         keys.KeyPress{Ctrl: true, Key: 'o'},
	tcell.KeyCtrlP:         keys.KeyPress{Ctrl: true, Key: 'p'},
	tcell.KeyCtrlQ:         keys.KeyPress{Ctrl: true, Key: 'q'},
	tcell.KeyCtrlR:         keys.KeyPress{Ctrl: true, Key: 'r'},
	tcell.KeyCtrlS:         keys.KeyPress{Ctrl: true, Key: 's'},
	tcell.KeyCtrlT:         keys.KeyPress{Ctrl: true, Key: 't'},
	tcell.KeyCtrlU:         keys.KeyPress{Ctrl: true, Key: 'u'},
	tcell.KeyCtrlV:         keys.KeyPress{Ctrl: true, Key: 'v'},
	tcell.KeyCtrlW:         keys.KeyPress{Ctrl: true, Key: 'w'},
	tcell.KeyCtrlX:         keys.KeyPress{Ctrl: true, Key: 'x'},
	tcell.KeyCtrlY:         keys.KeyPress{Ctrl: true, Key: 'y'},
	tcell.KeyCtrlZ:         keys.KeyPress{Ctrl: true, Key: 'z'},
	tcell.KeyCtrlLeftSq:    keys.KeyPress{Ctrl: true, Key: keys.Escape},
	tcell.KeyCtrlBackslash: keys.KeyPress{Ctrl: true, Key: keys.Backspace},
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
		kp = k
	} else {
		log.Warn("unrecognized key key: %v", key)
	}

	return kp
}
