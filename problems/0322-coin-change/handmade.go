package coinchange

const debugCoinChange = false

// coinChange returns the minimum coins to reach amount using bottom-up DP.
func coinChange(coins []int, amount int) int {
	const maxInt = int(^uint(0) >> 1)

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = maxInt
		for _, coin := range coins {

			if i < coin {
				continue
			}
			//  5 = 7-2
			delta := i - coin
			// 這裡的+1 是指一個該coin 加上原本dp[delta]
			if dp[delta] != maxInt && dp[delta]+1 < dp[i] {
				dp[i] = dp[delta] + 1
			}
		}
	}

	if dp[amount] == maxInt {
		return -1
	}

	return dp[amount]
}
