package easy_test

import (
	"reflect"
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type twoSumTest struct {
	nums   []int
	target int
	want   []int
}

var twoSumTests = []twoSumTest{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
	{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range twoSumTests {
		if got := easy.TwoSum(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
