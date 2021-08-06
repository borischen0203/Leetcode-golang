package leetcode

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, value := range nums {
		if _, ok := m[value]; !ok {
			m[value] = 1
		} else {
			delete(m, value)
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}
