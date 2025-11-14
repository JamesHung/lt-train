package longestpalindromicsubstring

import "fmt"

const debugLongestPalindrome = true

// longestPalindrome returns the longest palindromic substring using expand-around-center.
func longestPalindrome(s string) string {
	if debugLongestPalindrome {
		fmt.Printf("input=%q\n", s)
	}

	if len(s) < 2 {
		return s
	}

	bestStart, bestEnd := 0, 0
	for center := 0; center < len(s); center++ {
		leftOdd, rightOdd := expandAroundCenter(s, center, center)
		leftEven, rightEven := expandAroundCenter(s, center, center+1)
		fmt.Printf("leftOdd: %d, rightOdd %d\n", leftOdd, rightOdd)
		fmt.Printf("leftEven: %d, rightEven %d\n", leftEven, rightEven)

		if rightOdd-leftOdd > bestEnd-bestStart {
			bestStart, bestEnd = leftOdd, rightOdd
			logPalindrome(center, leftOdd, rightOdd, s)
		}
		if rightEven-leftEven > bestEnd-bestStart {
			bestStart, bestEnd = leftEven, rightEven
			logPalindrome(center, leftEven, rightEven, s)
		}
	}

	result := s[bestStart : bestEnd+1]
	if debugLongestPalindrome {
		fmt.Printf("result=%q\n", result)
	}
	return result
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}

func logPalindrome(center, start, end int, s string) {
	if debugLongestPalindrome {
		fmt.Printf(
			"center=%d new best [%d,%d]=%q length=%d\n",
			center,
			start,
			end,
			s[start:end+1],
			end-start+1,
		)
	}
}
