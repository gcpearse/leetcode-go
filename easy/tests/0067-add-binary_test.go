package easy_test

import (
	"testing"

	easy "github.com/gcpearse/leetcode-go/easy/problems"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b, want string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, test := range tests {
		if got := easy.AddBinary(test.a, test.b); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
