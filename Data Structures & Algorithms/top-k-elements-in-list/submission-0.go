func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	m := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		m[num]++
	}

	for num, count := range m {
		freq[count] = append(freq[count], num)
	}

	for i := len(freq) - 1; i > 0; i-- {
		for _, t := range freq[i] {
			res = append(res, t)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
