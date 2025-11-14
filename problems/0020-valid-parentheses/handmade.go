package validparentheses

// isValid returns true when s is a well-formed parentheses string.
func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	matching := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		default:

			if len(stack) == 0 || stack[len(stack)-1] != matching[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
