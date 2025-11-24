package maximumsubarray

// maxSubArray finds the contiguous range with the largest sum using Kadane's scan.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	best := nums[0]
	current := best
	for i := 1; i < len(nums); i++ {
		next := nums[i]
		if current+next < next {
			current = next
		} else {
			current += next
		}

		if current > best {
			best = current
		}
	}

	return best
}
