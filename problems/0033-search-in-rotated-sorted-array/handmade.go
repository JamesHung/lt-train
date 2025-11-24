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

// func search(nums []int, target int) int {

// 	var binarySearch func(low int, high int) int
// 	binarySearch = func(low int, high int) int {
// 		if low > high {
// 			return -1
// 		}

// 		mid := low + (high-low)/2

// 		switch {
// 		case nums[mid] == target:
// 			return mid
// 		case nums[low] <= nums[mid]:
// 			if nums[low] <= target && target < nums[mid] {
// 				high = mid - 1
// 			} else {
// 				low = mid + 1
// 			}
// 		default:
// 			if nums[mid] < target && target <= nums[high] {
// 				low = mid + 1
// 			} else {
// 				high = mid - 1
// 			}
// 		}

// 		return binarySearch(low, high)
// 	}

// 	return binarySearch(0, len(nums)-1)
// }
