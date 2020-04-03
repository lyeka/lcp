package medium

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if LengthOfLongestSubstring("pwwkew") != 3 {
		t.Fail()
	}
}
