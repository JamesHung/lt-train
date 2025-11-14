package insertinterval

const debugInsertInterval = false

// insert places newInterval into intervals and merges overlaps to keep ordering.
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	start, end := newInterval[0], newInterval[1]
	i := 0

	for i < len(intervals) && intervals[i][1] < start {
		result = append(result, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= end {
		if intervals[i][0] < start {
			start = intervals[i][0]
		}
		if intervals[i][1] > end {
			end = intervals[i][1]
		}
		i++
	}
	result = append(result, []int{start, end})

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
