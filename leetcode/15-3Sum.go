package leetcode

import "sort"

type void struct{}

var member void

func threeSum(nums []int) [][]int {
	var result [][]int
	if nums == nil || len(nums) < 3 {
		return result
	}
	set := make(map[[3]int]void)

	sort.Ints(nums)
	var start int
	var end int
	var sum int
	for i := 0; i < len(nums); i++ {
		target := i
		start = i + 1
		end = len(nums) - 1
		for start < end {
			sum = nums[start] + nums[end]
			if sum == -nums[target] {
				add := [3]int{nums[target], nums[start], nums[end]}
				if _, exist := set[add]; !exist {
					set[add] = member
					result = append(result, []int{nums[target], nums[start], nums[end]})
				}
				start++
				end--

			} else if sum < -nums[target] {
				start++
			} else if sum > -nums[target] {
				end--
			}
		}
	}
	return result
}
