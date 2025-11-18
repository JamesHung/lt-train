package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "duplicates",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := threeSum(append([]int(nil), tc.nums...))
			if !reflect.DeepEqual(normalizeTriplets(got), normalizeTriplets(tc.want)) {
				t.Fatalf("threeSum(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func normalizeTriplets(in [][]int) [][]int {
	for _, triplet := range in {
		sort.Ints(triplet)
	}
	sort.Slice(in, func(i, j int) bool {
		for idx := 0; idx < 3 && idx < len(in[i]) && idx < len(in[j]); idx++ {
			if in[i][idx] != in[j][idx] {
				return in[i][idx] < in[j][idx]
			}
		}
		return len(in[i]) < len(in[j])
	})
	return in
}
