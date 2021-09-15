package leetcode

func majorityElement(nums []int) int {
	m := make(map[int]int)
	n := len(nums) / 2

	for _, value := range nums {
		m[value]++
		if m[value] > n {
			return value
		}
	}
	return -1
}
