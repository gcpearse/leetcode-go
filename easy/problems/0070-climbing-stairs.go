package easy

func ClimbStairs(n int) int {
	x, y := 0, 1

	for range n - 1 {
		x, y = y, x+y
	}

	return x + y
}
