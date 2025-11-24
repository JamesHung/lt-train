package combinationsum

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "example 3 none",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "unsorted input with multiple combos",
			candidates: []int{7, 2, 6, 3},
			target:     9,
			want:       [][]int{{2, 2, 2, 3}, {2, 7}, {3, 3, 3}, {3, 6}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := combinationSum(tc.candidates, tc.target)
			if !sameCombos(got, tc.want) {
				t.Fatalf("combinationSum(%v, %d) = %v, want %v", tc.candidates, tc.target, got, tc.want)
			}
		})
	}
}

func sameCombos(a, b [][]int) bool {
	na := normalizeCombos(a)
	nb := normalizeCombos(b)
	return reflect.DeepEqual(na, nb)
}

func normalizeCombos(list [][]int) [][]int {
	out := make([][]int, len(list))
	for i, combo := range list {
		cp := append([]int(nil), combo...)
		sort.Ints(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		return lexLess(out[i], out[j])
	})
	return out
}

func lexLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
