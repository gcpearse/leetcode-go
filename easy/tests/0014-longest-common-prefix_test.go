package easy_test

import (
	"reflect"
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, test := range tests {
		if got := easy.LongestCommonPrefix(test.strs); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
