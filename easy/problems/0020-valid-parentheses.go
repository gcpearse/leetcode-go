package easy

func IsValid(s string) bool {
	if len(s) < 2 || len(s)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}

	for _, char := range s {
		if _, ok := pairs[char]; ok {
			stack = append(stack, char)
		} else if len(stack) > 0 && pairs[stack[len(stack)-1]] == char {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
