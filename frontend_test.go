package main

import (
	"testing"

	"github.com/gdamore/tcell"
)

func TestCreateNewViewOnCharacterInsert(t *testing.T) {
	newFe(t)
	defer fe.shutDown()

	if l := len(fe.views); l > 0 {
		t.Fatalf("Expected 0 views at first but got %d", l)
	}

	keyEv := tcell.NewEventKey(tcell.KeyRune, 's', tcell.ModNone)
	fe.screen.PostEventWait(keyEv)
	postQuitEvent()

	fe.loop()
	if l := len(fe.views); l != 1 {
		t.Fatalf("Expected 1 views after insert but got %d", l)
	}
}

func newFe(t *testing.T) {
	var err error
	fe, err = newFrontend()
	if err != nil {
		t.Fatal(err)
	}
	fe.screen.Screen = tcell.NewSimulationScreen("")

	err = fe.init()
	if err != nil {
		t.Fatal(err)
	}
}

func postQuitEvent() {
	ev := tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModCtrl)
	fe.screen.PostEventWait(ev)
}
