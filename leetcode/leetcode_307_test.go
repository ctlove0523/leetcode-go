package leetcode

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	array := Constructor([]int{1, 3, 5})
	//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]

	sumRange := array.SumRange(0, 2)
	if sumRange != 9 {
		t.Errorf("excepte 9 but is %d", sumRange)
	}

	array.Update(1,2)
	sumRange = array.SumRange(0,2)
	if sumRange != 8 {
		t.Errorf("excepte 8 but is %d", sumRange)
	}
}

func TestNumArray_SumRange2(t *testing.T) {
	array := Constructor([]int{-1})

	sumRange := array.SumRange(0, 0)
	if sumRange != -1 {
	t.Errorf("excepte -1 but is %d", sumRange)
	}

	array.Update(0,1)
	sumRange = array.SumRange(0,0)
	if sumRange != 1 {
	t.Errorf("excepte 1 but is %d", sumRange)
	}
}
