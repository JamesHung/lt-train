package twosum

func twoSum(nums []int, target int) []int {
	checked := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := checked[target-num]; ok {
			return []int{j, i}
		}

		checked[num] = i
	}

	return nil
}

// func twoSum(nums []int, target int) []int {
// 	check := make(map[int]int)

// 	for i, n := range nums {
// 		delta := target - n
// 		if ni, ok := check[n]; ok {
// 			return []int{ni, i}
// 		}
// 		check[delta] = i
// 	}

// 	return nil
// }
