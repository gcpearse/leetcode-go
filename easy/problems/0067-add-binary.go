package easy

func AddBinary(a, b string) string {
	sum := ""

	if len(a) < len(b) {
		a, b = b, a
	}

	for i := len(a) - len(b) - 1; i >= 0; i-- {
		b = "0" + b
	}

	carry := 0

	for i := len(a) - 1; i >= 0; i-- {
		switch carry {
		case 0:
			if a[i] == '0' && b[i] == '0' {
				sum = "0" + sum
			} else if a[i] == '1' && b[i] == '1' {
				sum = "0" + sum
				carry = 1
			} else {
				sum = "1" + sum
			}

		case 1:
			if a[i] == '0' && b[i] == '0' {
				sum = "1" + sum
				carry = 0
			} else if a[i] == '1' && b[i] == '1' {
				sum = "1" + sum
			} else {
				sum = "0" + sum
			}
		}
	}

	if carry == 1 {
		sum = "1" + sum
	}

	return sum
}

/*

Given two binary strings a and b, return their sum as a binary string.

Example 1:

Input: a = "11", b = "1"
Output: "100"

Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

*/
