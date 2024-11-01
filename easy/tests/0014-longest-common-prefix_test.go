package easy_test

import (
	"reflect"
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type longestCommonPrefixTest struct {
	strs []string
	want string
}

var longestCommonPrefixTests = []longestCommonPrefixTest{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range longestCommonPrefixTests {
		if got := easy.LongestCommonPrefix(test.strs); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
