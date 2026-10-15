func sortColors(nums []int) {
    col := make([]int, 3)

	for _, n := range nums {
		col[n]++
	}

	j := 0
	for i, c := range col {
		for range c {
			nums[j] = i
			j++
		}
	}
}
