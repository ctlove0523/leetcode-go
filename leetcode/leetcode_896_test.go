package leetcode

import "testing"

func TestIsMonotonic1(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	if IsMonotonic(nums) != true {
		t.Errorf("expected %t", true)
	}
}

func TestIsMonotonic2(t *testing.T) {
	nums := []int{6, 5, 4, 4}
	if IsMonotonic(nums) != true {
		t.Errorf("expected %t", true)
	}
}

func TestIsMonotonic3(t *testing.T) {
	nums := []int{1, 3, 2}
	if IsMonotonic(nums) != false {
		t.Errorf("expected %t", false)
	}
}

func TestIsMonotonic4(t *testing.T) {
	nums := []int{1, 2, 4, 5}
	if IsMonotonic(nums) != true {
		t.Errorf("expected %t", true)
	}
}

func TestIsMonotonic5(t *testing.T) {
	nums := []int{1, 1, 1}
	if IsMonotonic(nums) != true {
		t.Errorf("expected %t", true)
	}
}
