package medianoftwosortedarrays

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "example1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "example2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "single array non-empty",
			nums1: []int{},
			nums2: []int{2, 3},
			want:  2.5,
		},
		{
			name:  "different lengths",
			nums1: []int{0, 0},
			nums2: []int{0, 0, 0},
			want:  0.0,
		},
		{
			name:  "negatives",
			nums1: []int{-5, -3, -1},
			nums2: []int{-2, 4, 6, 8},
			want:  -1.0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := findMedianSortedArrays(tc.nums1, tc.nums2)
			if math.Abs(got-tc.want) > 1e-9 {
				t.Fatalf("findMedianSortedArrays(%v, %v) = %v, want %v", tc.nums1, tc.nums2, got, tc.want)
			}
		})
	}
}
