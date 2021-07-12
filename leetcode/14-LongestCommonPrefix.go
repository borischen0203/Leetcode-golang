package leetcode

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.HasPrefix(strs[i], pre) != true {
			pre = pre[0 : len(pre)-1]
		}
	}
	return pre
}
