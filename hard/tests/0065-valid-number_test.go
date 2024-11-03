package hard_test

import (
	"testing"

	hard "github.com/gcpearse/leetcode-go/hard/problems"
)

type isNumberTest struct {
	s    string
	want bool
}

var isNumberTests = []isNumberTest{
	{"0", true},
	{"e", false},
	{".", false},
}

func TestIsNumber(t *testing.T) {
	for _, test := range isNumberTests {
		if got := hard.IsNumber(test.s); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
