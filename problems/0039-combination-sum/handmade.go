package combinationsum

import (
	"fmt"
	"sort"
)

// combinationSum returns all unique combinations that sum to target; numbers may be reused.
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	fmt.Printf("%+v\n", candidates)
	length := len(candidates)
	result := make([][]int, 0, length)
	temp := make([]int, 0, length)

	var dfs func(index int, remains int)
	dfs = func(index int, remains int) {
		fmt.Printf("remains: %d\n", remains)
		if remains == 0 {
			comb := make([]int, len(temp))
			copy(comb, temp)
			result = append(result, comb)
			return
		}

		for i := index; i < length; i++ {
			candi := candidates[i]
			fmt.Printf("%d\n", candi)
			if candi > remains {
				break
			}

			temp = append(temp, candi)
			fmt.Printf("append %+v\n", temp)
			dfs(i, remains-candi)
			temp = temp[:len(temp)-1]
			fmt.Printf("pop %+v\n", temp)
		}
	}

	dfs(0, target)

	return result
}
