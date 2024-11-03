package medium

func IntToRoman(num int) string {
	type pair struct {
		roman   string
		integer int
	}

	pairs := []pair{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	result := ""

	for _, pair := range pairs {
		for num >= pair.integer {
			result += pair.roman
			num -= pair.integer
		}
	}

	return result
}
