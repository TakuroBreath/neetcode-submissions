func isAnagram(s string, t string) bool {
	m := make(map[byte]int)

	for _, v := range []byte(s) {
		m[v]++
	}

	for _, v := range []byte(t) {
		if val, ok := m[v]; ok {
			m[v]--
			if val == 1 {
				delete(m, v)
			}

			continue
		}

		return false
	}

	if len(m) == 0 {
		return true
	}

	return false
}
