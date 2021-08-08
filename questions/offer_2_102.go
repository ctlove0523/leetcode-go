package questions

func FindTargetSumWays(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target || nums[0] == -target{
			if target == 0 {
				return 2
			} else {
				return 1
			}
		} else {
			return 0
		}
	}

	return FindTargetSumWays(nums[:len(nums)-1], target+nums[len(nums)-1]) +
		FindTargetSumWays(nums[:len(nums)-1], target-nums[len(nums)-1])
}
