package climbingstairs

// climbStairs returns the number of distinct ways to reach step n.
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	prev1, prev2 := 1, 2
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev1, prev2 = prev2, current
	}

	return prev2
}

// n = 1, 1 p1
// n = 2, 2 p2 p1
// n = 3, 3 p2 p1
// n = 4, 5 p2
