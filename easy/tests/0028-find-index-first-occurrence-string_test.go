package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}

	for _, test := range tests {
		if got := easy.StrStr(test.haystack, test.needle); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
