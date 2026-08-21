func longestConsecutive(nums []int) int {
	ans := 0
	m := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		m[v] = struct{}{}
	}

	for num := range m {
		if _, exist := m[num-1]; !exist {
			c := 1
			for {
				if _, ok := m[num+c]; ok {
					c++
				} else {
					if ans < c { 
						ans = c
					}
					break
				}
			}
		}
	}

	return ans
}
