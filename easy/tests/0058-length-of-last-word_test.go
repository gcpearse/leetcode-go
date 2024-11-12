package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, test := range tests {
		if got := easy.LengthOfLastWord(test.s); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
