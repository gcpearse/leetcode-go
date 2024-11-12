package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, test := range tests {
		if got := easy.IsPalindrome(test.x); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
