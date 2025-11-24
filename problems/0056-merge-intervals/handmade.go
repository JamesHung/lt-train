package mergeintervals

import (
	"fmt"
	"sort"
)

// merge should combine overlapping intervals; implementation intentionally left blank for practice.
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	fmt.Printf("%+v\n", intervals)

	guardStart, guardEnd := intervals[0][0], intervals[0][1]
	intervals = intervals[1:]
	result := [][]int{}

	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if start <= guardEnd {
			if end > guardEnd {
				guardEnd = end
			}
			continue
		}

		result = append(result, []int{guardStart, guardEnd})
		guardStart, guardEnd = start, end
	}

	result = append(result, []int{guardStart, guardEnd})

	return result
}
