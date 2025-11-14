package majorityelement

func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}

		if major == num {
			count++
		} else {
			count--
		}
	}

	return major

}

// #hashmap
// majorityElement is left empty intentionally for handwriting practice.
// func majorityElement(nums []int) int {

// 	freq := make(map[int]int, len(nums))
// 	for _, n := range nums {
// 		freq[n]++
// 	}

// 	for num, count := range freq {
// 		if count > len(nums)/2 {
// 			return num
// 		}
// 	}

// 	return 0
// }
