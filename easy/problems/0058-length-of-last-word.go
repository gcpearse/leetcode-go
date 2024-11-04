package easy

func LengthOfLastWord(s string) int {
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if count > 0 {
			break
		}
	}

	return count
}
