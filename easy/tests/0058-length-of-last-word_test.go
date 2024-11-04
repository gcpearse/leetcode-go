package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

type lengthOfLastWordTest struct {
	s    string
	want int
}

var lengthOfLastWordTests = []lengthOfLastWordTest{
	{"Hello World", 5},
	{"   fly me   to   the moon  ", 4},
	{"luffy is still joyboy", 6},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, test := range lengthOfLastWordTests {
		if got := easy.LengthOfLastWord(test.s); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
