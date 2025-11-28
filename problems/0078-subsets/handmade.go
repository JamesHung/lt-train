package subsets

const debugSubsets = false

// subsets returns all possible subsets (the power set) of nums using DFS/backtracking.
func subsets(nums []int) [][]int {
	result := make([][]int, 0, 1<<len(nums))
	var cur []int

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			clone := make([]int, len(cur))
			copy(clone, cur)
			result = append(result, clone)
			return
		}
		// Not take nums[idx]
		dfs(idx + 1)

		// Take nums[idx]
		cur = append(cur, nums[idx])
		dfs(idx + 1)
		cur = cur[:len(cur)-1]
	}

	dfs(0)
	return result
}
