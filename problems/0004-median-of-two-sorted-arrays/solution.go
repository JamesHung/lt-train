package medianoftwosortedarrays

import (
	"fmt"
	"math"
)

const debugMedianLogs = true

// findMedianSortedArrays finds the median of two sorted arrays in O(log(min(m,n))) time.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fmt.Printf("\n")
	if debugMedianLogs {
		fmt.Printf("invoke: nums1=%v nums2=%v\n", nums1, nums2)
	}
	if len(nums1) > len(nums2) {
		if debugMedianLogs {
			fmt.Println("swapping arrays to keep first shorter")
		}
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2
	lo, hi := 0, m
	for lo <= hi {
		partition1 := (lo + hi) / 2
		partition2 := totalLeft - partition1

		if debugMedianLogs {
			fmt.Printf("[loop] lo=%d hi=%d partition1=%d partition2=%d\n", lo, hi, partition1, partition2)
		}

		maxLeft1 := math.MinInt
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}
		minRight1 := math.MaxInt
		if partition1 < m {
			minRight1 = nums1[partition1]
		}

		maxLeft2 := math.MinInt
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}
		minRight2 := math.MaxInt
		if partition2 < n {
			minRight2 = nums2[partition2]
		}

		if debugMedianLogs {
			fmt.Printf("maxLeft1=%d minRight1=%d | maxLeft2=%d minRight2=%d\n",
				maxLeft1, minRight1, maxLeft2, minRight2)
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 0 {
				leftMax := max(maxLeft1, maxLeft2)
				rightMin := min(minRight1, minRight2)
				median := float64(leftMax+rightMin) / 2
				if debugMedianLogs {
					fmt.Printf("even total -> leftMax=%d rightMin=%d median=%.5f\n", leftMax, rightMin, median)
				}
				return median
			}
			median := float64(max(maxLeft1, maxLeft2))
			if debugMedianLogs {
				fmt.Printf("odd total -> median=%.5f\n", median)
			}
			return median
		}

		if maxLeft1 > minRight2 {
			if debugMedianLogs {
				fmt.Println("maxLeft1 > minRight2 -> move hi left")
			}
			hi = partition1 - 1
		} else {
			if debugMedianLogs {
				fmt.Println("maxLeft2 > minRight1 -> move lo right")
			}
			lo = partition1 + 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
