package questions

import "testing"

func TestFindTargetSumWays1(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	if FindTargetSumWays(nums, target) != 5 {
		t.Errorf("expected result is %d ", 5)
	}
}

func TestFindTargetSumWays2(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	target := 1
	if FindTargetSumWays(nums, target) != 256 {
		t.Errorf("expected result is %d ", 256)
	}
}

func TestFindTargetSumWays3(t *testing.T) {
	nums := []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}
	target := 0
	if FindTargetSumWays(nums, target) != 0 {
		t.Errorf("expected result is %d", 256)
	}
}
