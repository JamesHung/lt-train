package lengthoflongestsubstring

import "fmt"

const debugLogs = true

// lengthOfLongestSubstring returns the longest substring length without repeating characters.
// Uses sliding window with last seen positions to keep O(n) time.
func lengthOfLongestSubstring(s string) int {
	length, longestStr := longestSubstringWindow(s)
	fmt.Printf("longestStr: %s\n", longestStr)
	return length
}

// longestSubstringWindow also returns the actual substring achieving the best length.
func longestSubstringWindow(s string) (int, string) {
	runes := []rune(s)
	seen := make(map[rune]int, len(runes))
	start, best, bestStart := 0, 0, 0

	for i, r := range runes {
		if prev, ok := seen[r]; ok && prev >= start {
			start = prev + 1
		}

		seen[r] = i
		if current := i - start + 1; current > best {
			best = current
			bestStart = start
		}
	}

	return best, string(runes[bestStart : bestStart+best])
}
