package sortcolors

import "fmt"

const debugSortColors = false

// sortColors reorders nums in-place to 0s, then 1s, then 2s using Dutch National Flag.
func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			if debugSortColors {
				fmt.Printf("swap 0: low=%d mid=%d -> %v\n", low, mid, nums)
			}
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			if debugSortColors {
				fmt.Printf("swap 2: mid=%d high=%d -> %v\n", mid, high, nums)
			}
			high--
		}
	}
}
