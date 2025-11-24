package permutations

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name: "example 2",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "example 3 single",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "includes negative",
			nums: []int{-1, 2, 0},
			want: [][]int{
				{-1, 2, 0}, {-1, 0, 2},
				{2, -1, 0}, {2, 0, -1},
				{0, -1, 2}, {0, 2, -1},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := permute(tc.nums)
			if !samePerms(got, tc.want) {
				t.Fatalf("permute(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func samePerms(a, b [][]int) bool {
	return reflect.DeepEqual(normalizePerms(a), normalizePerms(b))
}

func normalizePerms(perms [][]int) [][]int {
	out := make([][]int, len(perms))
	for i, p := range perms {
		cp := append([]int(nil), p...)
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
