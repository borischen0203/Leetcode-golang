package leetcode

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	first := 1
	second := 2
	next := 0

	for i := 3; i <= n; i++ {
		next = first + second
		first = second
		second = next
	}

	return next
}
