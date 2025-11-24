package permutations

// permute returns all permutations of distinct integers nums.
func permute(nums []int) [][]int {
	targetLen := len(nums)
	result := [][]int{}

	path := make([]int, 0, targetLen)
	used := make([]bool, targetLen)
	var dfs func()
	dfs = func() {
		if len(path) == targetLen {
			output := make([]int, targetLen)
			copy(output, path)
			result = append(result, output)
			return
		}

		for index, value := range nums {
			if used[index] {
				continue
			}

			used[index] = true
			path = append(path, value)
			dfs()
			used[index] = false
			path = path[:len(path)-1]
		}
	}

	dfs()

	return result
}
