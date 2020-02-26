package main

import (
	"testing"

	"github.com/limetext/text"
)

func TestCalcVisibleRegion(t *testing.T) {
	defer closeAll()

	bv := fe.editor.ActiveWindow().OpenFile("testdata/file", 0644)
	v := fe.views[bv]

	tests := []struct {
		vr, in, exp text.Region
	}{
		{
			text.Region{0, 365},
			text.Region{10, 20},
			text.Region{0, 365},
		},
		{
			text.Region{0, 365},
			text.Region{364, 366},
			text.Region{343, 739},
		},
		{
			text.Region{148, 621},
			text.Region{0, 0},
			text.Region{0, 365},
		},
	}
	for i, test := range tests {
		v.vr = test.vr
		v.calcVisibleRegion(test.in)

		if v.vr.String() != test.exp.String() {
			t.Errorf("Test %d: expected visible region %s, but got %s",
				i, test.exp, v.vr)
		}
	}
}
