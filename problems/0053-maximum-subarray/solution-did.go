package maximumsubarray

import "math"

// maxSubArrayDivideConquer solves the maximum subarray using divide & conquer.
func maxSubArrayDivideConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return divideMax(nums, 0, len(nums)-1)
}

func divideMax(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)/2
	leftBest := divideMax(nums, left, mid)
	rightBest := divideMax(nums, mid+1, right)
	crossBest := maxCrossingSum(nums, left, mid, right)
	return max3(leftBest, rightBest, crossBest)
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	sum := 0
	bestLeft := math.MinInt
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > bestLeft {
			bestLeft = sum
		}
	}

	sum = 0
	bestRight := math.MinInt
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > bestRight {
			bestRight = sum
		}
	}

	return bestLeft + bestRight
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= c {
		return b
	}
	return c
}
