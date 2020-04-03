package medium

import "testing"

const (
	s1 = "pwwkew"
	s2 = "abcabcbb"
	s3 = ""
)

func TestLengthOfLongestSubstring(t *testing.T) {

	if LengthOfLongestSubstring(s1) != 3 {
		t.Fatal(s1)
	}

	if LengthOfLongestSubstring(s3) != 0 {
		t.Fatal(s3)
	}

}

func TestLengthOfLongestSubstring2(t *testing.T) {
	if LengthOfLongestSubstring2(s1) != 3 {
		t.Fatal(s1)
	}

	if LengthOfLongestSubstring2(s2) != 3 {
		t.Fatal(s2)
	}

	if LengthOfLongestSubstring2(s3) != 0 {
		t.Fatal(s3)
	}
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	if LengthOfLongestSubstring3(s1) != 3 {
		t.Fatal(s1)
	}

	if LengthOfLongestSubstring3(s2) != 3 {
		t.Fatal(s2)
	}

	if LengthOfLongestSubstring3(s3) != 0 {
		t.Fatal(s3)
	}
}
