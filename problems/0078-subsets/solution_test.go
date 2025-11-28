package subsets

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "example1", nums: []int{1, 2, 3}, want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{name: "example2", nums: []int{0}, want: [][]int{{}, {0}}},
		{name: "twoElems", nums: []int{4, -1}, want: [][]int{{}, {4}, {-1}, {4, -1}}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := subsets(tc.nums)
			normalize := func(sets [][]int) {
				for i := range sets {
					sort.Ints(sets[i])
				}
				sort.Slice(sets, func(i, j int) bool {
					a, b := sets[i], sets[j]
					if len(a) != len(b) {
						return len(a) < len(b)
					}
					for k := 0; k < len(a); k++ {
						if a[k] != b[k] {
							return a[k] < b[k]
						}
					}
					return false
				})
			}
			normalize(got)
			normalize(tc.want)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("subsets(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
