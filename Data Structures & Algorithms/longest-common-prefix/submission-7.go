func longestCommonPrefix(strs []string) string {
    res := []byte{}

	for {
		if len(res) >= len(strs[0]) {
			return string(res)
		}
		
		t := strs[0][len(res)]
		for _, s := range strs {
			if len(res) >= len(s) || s[len(res)] != t {
				return string(res)
			}
		}
		res = append(res, t)
	}
}
