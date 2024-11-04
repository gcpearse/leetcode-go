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

/*

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false

*/
