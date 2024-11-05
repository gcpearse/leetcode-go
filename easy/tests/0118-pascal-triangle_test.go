package easy_test

import (
	"reflect"
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type generateTest struct {
	numRows int
	want    [][]int
}

var generateTests = []generateTest{
	{5, [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}},
	{1, [][]int{
		{1},
	}},
}

func TestGenerate(t *testing.T) {
	for _, test := range generateTests {
		if got := easy.Generate(test.numRows); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
