package validanagram

import "fmt"

// isAnagram returns true if s and t contain the same multiset of letters.
func isAnagram(s, t string) bool {
	fmt.Printf("\n")
	if len(s) != len(t) {
		return false
	}

	aChar := rune('a')
	freq := make([]int, 26)
	// for _, c := range s {
	for i := 0; i < len(s); i++ {
		freq[rune(s[i])-aChar]++
		freq[rune(t[i])-aChar]--
	}

	for _, count := range freq {
		fmt.Printf("count:%d\n", count)
		if count != 0 {
			return false
		}
	}

	// fmt.Printf("\n")
	return true
}
