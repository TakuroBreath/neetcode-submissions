func countStudents(students []int, sandwiches []int) int {
    res := len(students)
	cnt := make([]int, 2)

	for _, s := range students {
		cnt[s]++
	}

	for _, sand := range sandwiches {
		if cnt[sand] > 0 {
			res--
			cnt[sand]--
		} else {
			break
		}
	}

	return res
}