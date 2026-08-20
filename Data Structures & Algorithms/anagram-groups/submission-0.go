func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		s := [26]int{}
		for _, c := range str {
			s[c-'a']++
		}
		m[s] = append(m[s], str)
	}

	res := make([][]string, 0, len(m))
	for _, group := range m {
		res = append(res, group)
	}

	return res
}
