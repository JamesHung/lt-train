package searchinrotatedsortedarray

// search should return the index of target in a rotated sorted array or -1 if absent.
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return mid

		case nums[lo] <= nums[mid]:
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		case nums[mid] <= nums[hi]:
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
