package longestpalindromicsubstring

func findLongestPalindromicString(s string) string {
	if len(s) < 2 {
		return s
	}

	bestLeft := 0
	bestRight := 0
	for center := 0; center < len(s); center++ {
		oddStringLeft, oddStringRight := longestPalindromicString(s, center, center)
		evenStringLeft, eventStringRight := longestPalindromicString(s, center, center+1)

		if oddStringRight-oddStringLeft > bestRight-bestLeft {
			bestLeft = oddStringLeft
			bestRight = oddStringRight
		}

		if eventStringRight-evenStringLeft > bestRight-bestLeft {
			bestLeft = evenStringLeft
			bestRight = eventStringRight
		}
	}

	return s[bestLeft : bestRight+1]
}

func longestPalindromicString(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}
