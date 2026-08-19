func isAnagram(s string, t string) bool {
	m := make(map[byte]int)

	for _, v := range []byte(s) {
		m[v]++
	}

	for _, v := range []byte(t) {
		if val, ok := m[v]; ok {
			if val == 1 {
				delete(m, v)
			} else { 
				m[v]-- 
			}

			continue
		}

		return false
	}

	return len(m) == 0
}
