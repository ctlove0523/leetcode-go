package leetcode_go

import "testing"

func TestMin(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Errorf("min value bettern 1 and 2 must be 1")
	}
}
