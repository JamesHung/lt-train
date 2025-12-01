package longestpalindromicsubstring

func findLongestPalindromicString(s string) string {
	if len(s) < 2 {
		return s
	}

	bestLeft, bestRight := 0, 0
	for center := 0; center < len(s); center++ {
		oddLeft, oddRight := longestPalindromicString(s, center, center)
		evenLeft, evenRight := longestPalindromicString(s, center, center+1)

		if oddRight-oddLeft > bestRight-bestLeft {
			bestLeft = oddLeft
			bestRight = oddRight
		}

		if evenRight-evenLeft > bestRight-bestLeft {
			bestLeft = evenLeft
			bestRight = evenRight
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
