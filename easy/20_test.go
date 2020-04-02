package easy

import (
	"testing"
)

func TestIsValid(t *testing.T) {

	if !IsValid("{()}") {
		t.Fail()
	}

	if IsValid("{{})") {
		t.Fail()
	}
}