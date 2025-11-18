package productoftwoexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "example2",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "containsZeros",
			nums: []int{0, 2, 0, 4},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := productExceptSelf(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("productExceptSelf(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
