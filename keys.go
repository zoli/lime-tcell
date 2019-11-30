package main

import (
	"github.com/gdamore/tcell"
	"github.com/limetext/backend/keys"
	"github.com/limetext/backend/log"
)

var (
	keymap = map[tcell.Key]keys.Key{
		tcell.KeyUp:         keys.Up,
		tcell.KeyDown:       keys.Down,
		tcell.KeyRight:      keys.Right,
		tcell.KeyLeft:       keys.Left,
		tcell.KeyPgUp:       keys.PageUp,
		tcell.KeyPgDn:       keys.PageDown,
		tcell.KeyHome:       keys.Home,
		tcell.KeyEnd:        keys.End,
		tcell.KeyInsert:     keys.Insert,
		tcell.KeyDelete:     keys.Delete,
		tcell.KeyF1:         keys.F1,
		tcell.KeyF2:         keys.F2,
		tcell.KeyF3:         keys.F3,
		tcell.KeyF4:         keys.F4,
		tcell.KeyF5:         keys.F5,
		tcell.KeyF6:         keys.F6,
		tcell.KeyF7:         keys.F7,
		tcell.KeyF8:         keys.F8,
		tcell.KeyF9:         keys.F9,
		tcell.KeyF10:        keys.F10,
		tcell.KeyF11:        keys.F11,
		tcell.KeyF12:        keys.F12,
		tcell.KeyBackspace:  keys.Backspace,
		tcell.KeyBackspace2: keys.Backspace,
		tcell.KeyTab:        '\t',
		tcell.KeyEscape:     keys.Escape,
		tcell.KeyEnter:      keys.Enter,
	}

	ctrlKeyMap = map[tcell.Key]keys.Key{
		tcell.KeyCtrlSpace:     ' ',
		tcell.KeyCtrlA:         'a',
		tcell.KeyCtrlB:         'b',
		tcell.KeyCtrlC:         'c',
		tcell.KeyCtrlD:         'd',
		tcell.KeyCtrlE:         'e',
		tcell.KeyCtrlF:         'f',
		tcell.KeyCtrlG:         'g',
		tcell.KeyCtrlH:         'h',
		tcell.KeyCtrlI:         'i',
		tcell.KeyCtrlJ:         'j',
		tcell.KeyCtrlK:         'k',
		tcell.KeyCtrlL:         'l',
		tcell.KeyCtrlM:         'm',
		tcell.KeyCtrlN:         'n',
		tcell.KeyCtrlO:         'o',
		tcell.KeyCtrlP:         'p',
		tcell.KeyCtrlQ:         'q',
		tcell.KeyCtrlR:         'r',
		tcell.KeyCtrlS:         's',
		tcell.KeyCtrlT:         't',
		tcell.KeyCtrlU:         'u',
		tcell.KeyCtrlV:         'v',
		tcell.KeyCtrlW:         'w',
		tcell.KeyCtrlX:         'x',
		tcell.KeyCtrlY:         'y',
		tcell.KeyCtrlZ:         'z',
		tcell.KeyCtrlLeftSq:    keys.Escape,
		tcell.KeyCtrlBackslash: '\\',
	}
)

func keyPress(ek *tcell.EventKey) keys.KeyPress {
	var kp keys.KeyPress

	mods := ek.Modifiers()
	kp.Alt = mods == tcell.ModAlt
	kp.Ctrl = mods == tcell.ModCtrl
	kp.Shift = mods == tcell.ModShift

	if key := ek.Key(); key == tcell.KeyRune {
		kp.Text = string(ek.Rune())
		kp.Key = keys.Key(ek.Rune())
	} else if k, ok := ctrlKeyMap[key]; kp.Ctrl && ok {
		kp.Key = k
	} else if k, ok := keymap[key]; ok {
		kp.Key = k
	} else {
		log.Warn("unrecognized key key: %v", key)
	}

	return kp
}
