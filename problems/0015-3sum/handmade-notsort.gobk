package threesum

import "sort"

const debugThreeSum = false

// threeSum returns all unique triplets whose sum is zero.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0, len(nums))

	for i, n := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -n
		left, right := i+1, len(nums)-1

		for left < right {
			test := nums[left] + nums[right]
			if test == target {
				val := []int{n, nums[left], nums[right]}
				result = append(result, val)
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if test < target {
				left++
			} else if test > target {
				right--
			}

		}
	}

	return result
}
