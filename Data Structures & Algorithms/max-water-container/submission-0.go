func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	m := 0

	for l < r {
		s := min(heights[l], heights[r]) * (r-l)
		if m < s { m = s }

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return m
}
