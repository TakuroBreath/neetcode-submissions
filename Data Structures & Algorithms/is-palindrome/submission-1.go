func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	s = strings.ToLower(s)
	
	for l < r {
		for l < r && !isAvailable(s[l]) {
			l++
		}

		for l < r && !isAvailable(s[r]) {
			r--
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func isAvailable(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c))
}