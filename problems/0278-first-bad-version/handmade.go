package firstbadversion

import "fmt"

// isBadVersion simulates the LeetCode-provided API and is set in tests.
var isBadVersion func(version int) bool

// firstBadVersion finds the earliest bad release using binary search.
func firstBadVersion(n int) int {
	if n == 1 {
		return n
	}

	lowest, highest := 1, n
	for lowest < highest {
		fmt.Printf("lowest=%d highest=%d\n", lowest, highest)
		fmt.Printf("(highest-lowest)/2=+%v\n", (highest-lowest)/2)
		mid := lowest + (highest-lowest)/2

		if isBadVersion(mid) {
			highest = mid
		} else {
			lowest = mid + 1
		}
	}
	return lowest
}
