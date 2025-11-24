package mergeintervals

import (
	"reflect"
	"testing"
)

func cloneIntervals(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = append([]int(nil), src[i]...)
	}
	return dst
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{name: "example1", intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "example2_touching", intervals: [][]int{{1, 4}, {4, 5}}, want: [][]int{{1, 5}}},
		{name: "example3_unsorted", intervals: [][]int{{4, 7}, {1, 4}}, want: [][]int{{1, 7}}},
		{name: "single", intervals: [][]int{{0, 1}}, want: [][]int{{0, 1}}},
		{name: "nested_ranges", intervals: [][]int{{1, 10}, {2, 3}, {4, 5}}, want: [][]int{{1, 10}}},
		{name: "disjoint_after_sort", intervals: [][]int{{5, 6}, {1, 2}, {3, 4}}, want: [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{name: "chain_merge", intervals: [][]int{{1, 4}, {2, 3}, {3, 8}}, want: [][]int{{1, 8}}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := merge(cloneIntervals(tc.intervals))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("merge(%v) = %v, want %v", tc.intervals, got, tc.want)
			}
		})
	}
}
