package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type strStrTest struct {
	haystack string
	needle   string
	want     int
}

var strStrTests = []strStrTest{
	{"sadbutsad", "sad", 0},
	{"leetcode", "leeto", -1},
}

func TestStrStr(t *testing.T) {
	for _, test := range strStrTests {
		if got := easy.StrStr(test.haystack, test.needle); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
