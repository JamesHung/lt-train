package wordbreak

// wordBreak returns true if s can be segmented into dictionary words using DP.
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dict := make(map[string]struct{}, len(wordDict))
	maxLen := 0

	for _, word := range wordDict {
		if len(word) > maxLen {
			maxLen = len(word)
		}
		dict[word] = struct{}{}
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		if !dp[i] {
			continue
		}

		limit := i + maxLen
		if limit > n {
			limit = n
		}
		for j := i + 1; j <= limit; j++ {
			if _, ok := dict[s[i:j]]; ok {
				dp[j] = true
			}
		}
	}

	return dp[n]
}
