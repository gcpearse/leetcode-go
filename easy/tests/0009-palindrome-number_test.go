package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type isPalindromeTest struct {
	x    int
	want bool
}

var isPalindromeTests = []isPalindromeTest{
	{121, true},
	{-121, false},
	{10, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range isPalindromeTests {
		if got := easy.IsPalindrome(test.x); got != test.want {
			t.Errorf("x: %v", test.x)
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
