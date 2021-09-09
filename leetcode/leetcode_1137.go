package leetcode

func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	t0 := 0
	t1 := 1
	t2 := 1

	for i := 3; i <= n; i++ {
		v := t0 + t1 + t2
		t0 = t1
		t1 = t2
		t2 = v
	}

	return t2
}
