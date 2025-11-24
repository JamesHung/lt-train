package combinationsum

import "sort"

// combinationSum returns all unique combinations that sum to target; numbers may be reused.
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	temp := []int{}
	result := [][]int{}

	var dfs func(index int, remains int)
	dfs = func(index int, remains int) {
		if remains == 0 {
			finals := make([]int, len(temp))
			copy(finals, temp)
			result = append(result, finals)
			return
		}

		for i := index; i < len(candidates); i++ {
			candidate := candidates[i]
			if candidate > remains {
				break
			}

			temp = append(temp, candidate)
			dfs(i, remains-candidate)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0, target)

	return result
}
