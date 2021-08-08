package leetcode

func IsMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	bigZero := 0
	smallZero := 0
	zero := 0

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 0 {
			zero++
		} else if nums[i]-nums[i-1] > 0 {
			bigZero++
		} else {
			smallZero++
		}
		if bigZero !=0 && smallZero !=0 {
			return false
		}
	}

	return true
}
