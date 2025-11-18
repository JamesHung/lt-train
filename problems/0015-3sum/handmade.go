package threesum

import "sort"

const debugThreeSum = false

// threeSum returns all unique triplets whose sum is zero.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
