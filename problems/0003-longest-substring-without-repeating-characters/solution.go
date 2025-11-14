package lengthoflongestsubstring

import "fmt"

const debugLogs = true

// lengthOfLongestSubstring returns the longest substring length without repeating characters.
// Uses sliding window with last seen positions to keep O(n) time.
func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	lastSeen := make(map[rune]int, len(runes))
	start := 0
	best := 0

	for i, ch := range runes {
		if prev, ok := lastSeen[ch]; ok && prev >= start {
			if debugLogs {
				fmt.Printf("duplicate %q at %d; moving start from %d to %d\n", ch, i, start, prev+1)
			}
			start = prev + 1
		}
		lastSeen[ch] = i
		if curr := i - start + 1; curr > best {
			if debugLogs {
				fmt.Printf("new best %d for window [%d,%d] => %q\n", curr, start, i, string(runes[start:i+1]))
			}
			best = curr
		}
		if debugLogs {
			fmt.Printf("i=%d char=%q window=[%d,%d] best=%d\n", i, ch, start, i, best)
		}
	}

	return best
}
