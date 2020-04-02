package easy

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome(-121){
		t.Fail()
	}

	if !IsPalindrome(121){
		t.Fail()
	}

	if IsPalindrome(100) {
		t.Fail()
	}

}
