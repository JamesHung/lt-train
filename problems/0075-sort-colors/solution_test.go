package sortcolors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "example2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			name: "alreadySorted",
			nums: []int{0, 0, 1, 1, 2, 2},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "allSame",
			nums: []int{1, 1, 1},
			want: []int{1, 1, 1},
		},
		{
			name: "single",
			nums: []int{2},
			want: []int{2},
		},
		{
			name: "zigzag",
			nums: []int{0, 2, 0, 2, 1, 1, 0, 2},
			want: []int{0, 0, 0, 1, 1, 2, 2, 2},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int{}, tc.nums...)
			sortColors(nums)
			if !reflect.DeepEqual(nums, tc.want) {
				t.Fatalf("sortColors(%v) = %v, want %v", tc.nums, nums, tc.want)
			}
		})
	}
}
