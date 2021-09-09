package leetcode

type NumArray struct {
	nums []int
	sum  []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+2)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	sum[len(sum)-1] = sum[len(sum)-2]

	return NumArray{
		nums: nums,
		sum:  sum,
	}
}

func (this *NumArray) Update(index int, val int) {
	old := this.nums[index]
	this.nums[index] = val
	for i := index+1; i < len(this.sum); i++ {
		this.sum[i] += val - old
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}
