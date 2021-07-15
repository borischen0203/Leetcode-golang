package leetcode

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	pre := 0

	for _, val := range nums {
		if nums[pre] != val {
			pre++
			nums[pre] = val
		}
	}
	return pre + 1

}
