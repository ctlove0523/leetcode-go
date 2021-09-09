package leetcode

import "testing"

func TestTribonacci(t *testing.T) {
	if Tribonacci(4) != 4 {
		t.Errorf("excepted %d", 4)
	}
}

func TestTribonacci1(t *testing.T) {
	if Tribonacci(25) != 1389537 {
		t.Errorf("excepted %d", 4)
	}
}
