package questions

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	sort.Ints(arr)

	res := make([]int, k)
	for i := 0; i < k && i < len(arr); i++ {
		res[i] = arr[i]
	}

	return res
}
