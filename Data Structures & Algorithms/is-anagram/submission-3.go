func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var abc = [26]int{}

	for i := 0; i < len(s); i++ {
		abc[s[i]-'a']++
		abc[t[i]-'a']--
	}

	for _, c := range abc {
		if c != 0 {
			return false
		}
	}

	return true
}
