package insertinterval

const debugInsertInterval = false

// insert places newInterval into intervals and merges overlaps to keep ordering.
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	merged := []int{newInterval[0], newInterval[1]}
	mergedPlaced := false

	for _, interval := range intervals {
		switch {
		case interval[1] < merged[0]:
			result = append(result, interval)
		case interval[0] > merged[1]:
			if !mergedPlaced {
				result = append(result, merged)
				mergedPlaced = true
			}
			result = append(result, interval)
		default:
			merged[0] = min(interval[0], merged[0])
			merged[1] = max(interval[1], merged[1])
		}
	}

	if !mergedPlaced {
		result = append(result, merged)
	}

	return result
}
