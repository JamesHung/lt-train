package maximumsubarray

// maxSubArray finds the contiguous range with the largest sum using Kadane's scan.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// if len(nums) == 1 {
	// 	return nums[0]
	// }
	current := nums[0]
	best := current
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if current+n < n {
			current = n
		} else {
			current += n
		}

		if current > best {
			best = current
		}
	}

	return best
}
