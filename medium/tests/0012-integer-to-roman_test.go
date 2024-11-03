package medium_test

import (
	"testing"

	medium "github.com/gcpearse/leetcode-go/medium/problems"
)

type intToRomanTest struct {
	num  int
	want string
}

var intToRomanTests = []intToRomanTest{
	{3749, "MMMDCCXLIX"},
	{58, "LVIII"},
	{1994, "MCMXCIV"},
}

func TestIntToRoman(t *testing.T) {
	for _, test := range intToRomanTests {
		if got := medium.IntToRoman(test.num); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
