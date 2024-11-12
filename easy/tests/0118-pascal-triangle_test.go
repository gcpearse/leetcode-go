package easy_test

import (
	"reflect"
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
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

	for _, test := range tests {
		if got := easy.Generate(test.numRows); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
