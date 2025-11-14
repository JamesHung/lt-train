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
