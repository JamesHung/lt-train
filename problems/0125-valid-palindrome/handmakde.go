package validpalindrome

import "unicode"

// isPalindrome returns true if s is a palindrome ignoring non-alphanumerics and case.
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNum(rune(s[left])) {
			left++
		}
		for left < right && !isAlphaNum(rune(s[right])) {
			right--
		}

		if left >= right {
			break
		}

		lChar := unicode.ToUpper(rune(s[left]))
		rChar := unicode.ToUpper(rune(s[right]))

		if lChar != rChar {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNum(r rune) bool {
	if '0' <= r && r <= '9' {
		return true
	}

	if 'a' <= r && r <= 'z' {
		return true
	}

	if 'A' <= r && r <= 'Z' {
		return true
	}

	return false
}
