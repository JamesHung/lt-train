package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example1", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "example2", nums: []int{1}, want: 1},
		{name: "example3", nums: []int{5, 4, -1, 7, 8}, want: 23},
		{name: "allNegative", nums: []int{-5, -2, -3}, want: -2},
		{name: "resetAfterDip", nums: []int{8, -19, 5, -4, 20}, want: 21},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := maxSubArray(tc.nums); got != tc.want {
				t.Fatalf("maxSubArray(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
