package medium

import "strings"

// 滑动窗口
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	i, j, max := 0, 1, 1
	l := len(s)
	for i < l && j < l {
		if !strings.Contains(s[i:j], string(s[j])) {
			if n := j-i+1; n > max {
				max = n
			}
			j++
		} else {
			i++
		}
	}

	return max
}