package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type addBinaryTest struct {
	a, b, want string
}

var addBinaryTests = []addBinaryTest{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
}

func TestAddBinary(t *testing.T) {
	for _, test := range addBinaryTests {
		if got := easy.AddBinary(test.a, test.b); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
