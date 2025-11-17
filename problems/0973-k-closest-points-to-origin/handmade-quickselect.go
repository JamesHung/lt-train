package kclosestpointstoorigin

// kClosestQuickselect selects the k closest points using in-place quickselect.
func kClosestQuickselect(points [][]int, k int) [][]int {
	if k >= len(points) {
		return points
	}

	selectK(points, 0, len(points)-1, k)
	return points[:k]
}

func selectK(points [][]int, left, right, k int) {
	for left <= right {
		pivotIdx := partition(points, left, right)
		if pivotIdx == k {
			return
		}
		if pivotIdx < k {
			left = pivotIdx + 1
		} else {
			right = pivotIdx - 1
		}
	}
}

func partition(points [][]int, left, right int) int {
	pivotIdx := left + (right-left)/2
	pivotDist := squaredDist(points[pivotIdx])
	points[pivotIdx], points[right] = points[right], points[pivotIdx]

	store := left
	for i := left; i < right; i++ {
		if squaredDist(points[i]) < pivotDist {
			points[i], points[store] = points[store], points[i]
			store++
		}
	}

	points[store], points[right] = points[right], points[store]
	return store
}
