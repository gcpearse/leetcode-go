package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, test := range tests {
		if got := easy.ClimbStairs(test.n); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
