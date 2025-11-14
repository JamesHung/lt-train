package topkfrequentelements

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "example1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "example2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "example3",
			nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "kEqualsUniqueCount",
			nums: []int{4, 4, 5, 6},
			k:    3,
			want: []int{4, 5, 6},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := topKFrequent(tc.nums, tc.k)
			if len(got) != len(tc.want) {
				t.Fatalf("topKFrequent(%v, %d) length=%d want=%d", tc.nums, tc.k, len(got), len(tc.want))
			}
			slices.Sort(got)
			want := slices.Clone(tc.want)
			slices.Sort(want)
			for i := range want {
				if got[i] != want[i] {
					t.Fatalf("topKFrequent(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
				}
			}
		})
	}
}
