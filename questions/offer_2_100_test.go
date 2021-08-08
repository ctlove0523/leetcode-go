package questions

import "testing"

func TestMinimumTotal(t *testing.T) {
	nums := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	if MinimumTotal(nums) != 11 {
		t.Errorf("excepte %d", 11)
	}
}
