package coinchange

const debugCoinChange = false

// coinChange returns the minimum coins to reach amount using bottom-up DP.
func coinChange(coins []int, amount int) int {
	const maxInt = int(^uint(0) >> 1)

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = maxInt
		for _, coin := range coins {
			//  5 = 7-2
			delta := i - coin
			if i >= coin && dp[delta] != maxInt && dp[delta]+1 < dp[i] {
				dp[i] = dp[delta] + 1
			}
		}
	}

	if dp[amount] == maxInt {
		return -1
	}

	return dp[amount]
}
