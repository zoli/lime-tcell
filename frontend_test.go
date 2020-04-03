package main

import (
	"testing"
	"time"

	"github.com/gdamore/tcell"
)

func TestCreateNewViewOnCharacterInsert(t *testing.T) {
	defer closeAll()

	if l := len(fe.activeWindow().pane.views); l > 0 {
		t.Fatalf("Expected 0 views at first but got %d", l)
	}

	keyEv := tcell.NewEventKey(tcell.KeyRune, 's', tcell.ModNone)
	fe.scrn.PostEventWait(keyEv)

	postQuitEvent()

	fe.loop()
	if l := len(fe.activeWindow().pane.views); l != 1 {
		t.Fatalf("Expected 1 views after insert but got %d", l)
	}

	time.Sleep(500 * time.Millisecond)
	r := getContent(0, 0)
	if r != 's' {
		t.Errorf("Expected character 's' in (0, 0) but got %q", r)
	}
}

func TestNewViewClearsPage(t *testing.T) {
	defer closeAll()

	fe.scrn.SetContent(0, 0, 'a', nil, defStyle)
	fe.ed.ActiveWindow().NewFile()

	_, h := fe.scrn.Size()
	for i := 0; i < h; i++ {
		if r := getContent(0, i); r != ' ' {
			t.Errorf("Expected screen be clear but got %q in (0, %d)",
				r, i)
		}
	}
}

func TestCloseViewShowsBeforeView(t *testing.T) {
	defer closeAll()

	v1 := fe.ed.ActiveWindow().NewFile()
	edit := v1.BeginEdit()
	v1.Insert(edit, 0, "a")
	v1.EndEdit(edit)
	t.Logf("(0, 0) %q", getContent(0, 0))

	v2 := fe.ed.ActiveWindow().NewFile()
	edit = v2.BeginEdit()
	v2.Insert(edit, 0, "b")
	v2.EndEdit(edit)
	t.Logf("(0, 0) %q", getContent(0, 0))

	v2.SetScratch(true)
	v2.Close()

	r := getContent(0, 0)
	if r != 'a' {
		t.Errorf("Expected character 'a' in (0, 0) but got %q", r)
	}
}
