package leetcode

import "strings"

/**
 * Given a pattern and a string s, find if s follows the same pattern.
 *	Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
 *	Example 1:
 *	Input: pattern = "abba", s = "dog cat cat dog"
 *	Output: true
**/

func wordPattern(pattern string, s string) bool {
	stringArr := strings.Fields(s)
	store := make(map[rune]string)
	appearance := make(map[string]bool)
	if len(stringArr) != len(pattern) {
		return false
	}
	for index, char := range pattern {
		_, exist := store[char]
		if !exist {
			if _, appear := appearance[stringArr[index]]; appear {
				return false
			} else {
				store[char] = stringArr[index]
				appearance[stringArr[index]] = true
			}
		} else {
			if store[char] != stringArr[index] {
				return false
			}
		}
	}
	return true
}
