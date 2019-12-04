package main

import (
	"testing"
	"time"

	"github.com/gdamore/tcell"
)

func TestCreateNewViewOnCharacterInsert(t *testing.T) {
	defer closeAll()

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

	time.Sleep(500 * time.Millisecond)
	r, _, _, _ := fe.screen.GetContent(0, 0)
	if r != 's' {
		t.Errorf("Expected character 's' in (0, 0) but got %q", r)
	}
}

func TestNewViewClearsPage(t *testing.T) {
	defer closeAll()

	fe.screen.SetContent(0, 0, 'a', nil, defStyle)
	fe.editor.ActiveWindow().NewFile()

	_, h := fe.screen.Size()
	for i := 0; i < h; i++ {
		r, _, _, _ := fe.screen.GetContent(0, i)
		if r != ' ' {
			t.Errorf("Expected screen be clear but got %q in (0, %d)",
				r, i)
		}
	}
}
