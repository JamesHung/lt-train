package maximumsubarray

const debugMaxSubArray = false

// maxSubArray finds the contiguous range with the largest sum using Kadane's scan.
func maxSubArray(nums []int) int {
	best := nums[0]
	current := nums[0]

	for i := 1; i < len(nums); i++ {
		if current < 0 {
			current = nums[i]
		} else {
			current += nums[i]
		}
		if current > best {
			best = current
		}
	}

	return best
}
