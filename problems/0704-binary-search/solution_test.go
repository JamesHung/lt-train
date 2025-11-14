package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example hit",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "example miss",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "single element hit",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "single element miss",
			nums:   []int{5},
			target: 3,
			want:   -1,
		},
		{
			name:   "left boundary",
			nums:   []int{1, 3, 5, 7, 9},
			target: 1,
			want:   0,
		},
		{
			name:   "right boundary",
			nums:   []int{1, 3, 5, 7, 9},
			target: 9,
			want:   4,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.want {
				t.Fatalf("search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
