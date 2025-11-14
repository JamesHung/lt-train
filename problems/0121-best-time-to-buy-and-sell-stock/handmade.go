package besttimetobuyandsellstock

import (
	"math"
)

const debugMaxProfit = false

// maxProfit returns the maximum profit achievable by buying once and selling later.
func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	best := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		if profit := price - minPrice; profit > best {
			best = profit
		}

	}

	return best
}
