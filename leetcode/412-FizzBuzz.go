package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	var result []string = make([]string, n)
	for i := 1; i <= n; i++ {
		isFizz := i%3 == 0
		isBuzz := i%5 == 0
		var str = ""
		if isFizz {
			str += "Fizz"
		}
		if isBuzz {
			str += "Buzz"
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		result[i-1] = str
	}
	return result
}

func fizzBuzz2(n int) []string {
	var result []string = make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
