package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	preSum := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		preSum += nums[i]
		if nums[i] >= preSum {
			preSum = nums[i]
		}
		if preSum >= result {
			result = preSum
		}
	}
	return result
}
