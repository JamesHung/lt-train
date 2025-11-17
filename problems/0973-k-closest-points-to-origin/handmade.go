package kclosestpointstoorigin

import "sort"

const debugKClosest = false

// kClosest sorts the points by squared distance to the origin and returns the first k.
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return squaredDist(points[i]) < squaredDist(points[j])
	})

	return points[0:k]
}

func squaredDist(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}
