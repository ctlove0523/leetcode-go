package leetcode_go

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func MinValueOfSlice(a []int) int {
	if len(a) == 0 {
		return 0
	}
	maxValue := a[0]
	for i := 1; i < len(a); i++ {
		if maxValue > a[i] {
			maxValue = a[i]
		}
	}

	return maxValue
}
