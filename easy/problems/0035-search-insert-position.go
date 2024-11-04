package easy

func SearchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}

	return len(nums)
}
