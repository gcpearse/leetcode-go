package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type romanToIntTest struct {
	s    string
	want int
}

var romanToIntTests = []romanToIntTest{
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"D", 500},
}

func TestRomanToInt(t *testing.T) {
	for _, test := range romanToIntTests {
		if got := easy.RomanToInt(test.s); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
