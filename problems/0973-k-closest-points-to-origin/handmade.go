package kclosestpointstoorigin

import "sort"

const debugKClosest = false

// kClosest sorts the points by squared distance to the origin and returns the first k.
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return squreDist(points[i]) < squreDist(points[j])
	})

	return points[:k]
}

func squreDist(p1 []int) int {
	return p1[0]*p1[0] + p1[1]*p1[1]
}
