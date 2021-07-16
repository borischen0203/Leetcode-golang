package leetcode

func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}
