func isValid(s string) bool {
    stack := []rune{}
	toClose := map[rune]rune{')' : '(', '}':'{', ']':'['}

	for _, sym := range s {
		if open, exists := toClose[sym]; exists {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if top != open {
					return false
				}
			} else { return false }
		} else {
			stack = append(stack, sym)
		}
	}

	return len(stack) == 0
}
