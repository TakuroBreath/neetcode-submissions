func isPalindrome(s string) bool {
	var filtered []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, r)
		}
	}

	for i := 0; i < len(filtered)/2; i++ {
		if filtered[i] != filtered[len(filtered)-1-i] {
			return false
		}
	}

	return true
}