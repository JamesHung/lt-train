package combinationsum

import (
	"fmt"
	"sort"
)

const debugCombinationSum = true

// combinationSum returns all unique combinations that sum to target; numbers may be reused.
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := make([][]int, 0, len(candidates))
	temp := make([]int, 0, len(candidates))

	var dfs func(start int, remains int)
	dfs = func(start, remains int) {
		if remains == 0 {
			finalCombs := make([]int, len(temp))
			copy(finalCombs, temp)
			result = append(result, finalCombs)
			return
		}

		for i := start; i < len(candidates); i++ {
			candidate := candidates[i]
			if candidate > remains {
				break
			}

			temp = append(temp, candidate)
			if debugCombinationSum {
				fmt.Printf("push %d -> temp %v (rem %d)\n", candidate, temp, remains-candidate)
			}
			dfs(i, remains-candidate)
			if debugCombinationSum {
				fmt.Printf("return -> temp %v (next pop %d)\n", temp, candidate)
			}
			temp = temp[:len(temp)-1]
			if debugCombinationSum {
				fmt.Printf("pop %d -> temp %v\n", candidate, temp)
			}
		}

	}

	dfs(0, target)

	return result
}
