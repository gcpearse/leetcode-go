package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type searchInsertTest struct {
	nums   []int
	target int
	want   int
}

var searchInsertTests = []searchInsertTest{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
}

func TestSearchInsert(t *testing.T) {
	for _, test := range searchInsertTests {
		if got := easy.SearchInsert(test.nums, test.target); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
