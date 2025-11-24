package searchinrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "example target found", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{name: "example target missing", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{name: "single element missing", nums: []int{1}, target: 0, want: -1},
		{name: "single element present", nums: []int{1}, target: 1, want: 0},
		{name: "rotated at middle", nums: []int{6, 7, 1, 2, 3, 4, 5}, target: 2, want: 3},
		{name: "rotated at middle end", nums: []int{6, 7, 1, 2, 3, 4, 5}, target: 6, want: 0},
		{name: "no rotation", nums: []int{1, 2, 3, 4, 5, 6}, target: 4, want: 3},
		{name: "rotation of two", nums: []int{5, 6, 1, 2, 3, 4}, target: 6, want: 1},
		{name: "length two present", nums: []int{3, 1}, target: 1, want: 1},
		{name: "length two missing", nums: []int{3, 1}, target: 2, want: -1},
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
