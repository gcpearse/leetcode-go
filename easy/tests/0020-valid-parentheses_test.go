package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type isValidTest struct {
	s    string
	want bool
}

var isValidTests = []isValidTest{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([])", true},
	{"]", false},
	{"){", false},
}

func TestIsValid(t *testing.T) {
	for _, test := range isValidTests {
		if got := easy.IsValid(test.s); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
