package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type climbStairsTest struct {
	n, want int
}

var climbStairsTests = []climbStairsTest{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 8},
}

func TestClimbStairs(t *testing.T) {
	for _, test := range climbStairsTests {
		if got := easy.ClimbStairs(test.n); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
