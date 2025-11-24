package insertinterval

// insert places newInterval into intervals and merges overlaps to keep ordering.
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals))
	newStart, newEnd := newInterval[0], newInterval[1]

	i := 0
	for i < len(intervals) && intervals[i][1] < newStart {
		result = append(result, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newEnd {
		if intervals[i][0] < newStart {
			newStart = intervals[i][0]
		}

		if intervals[i][1] > newEnd {
			newEnd = intervals[i][1]
		}
		i++
	}

	result = append(result, []int{newStart, newEnd})

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
