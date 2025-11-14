package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example1", nums: []int{3, 2, 3}, want: 3},
		{name: "example2", nums: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
		{name: "all same", nums: []int{5, 5, 5}, want: 5},
		{name: "majority at end", nums: []int{1, 2, 1, 2, 1, 1, 1}, want: 1},
		{name: "majority at start", nums: []int{9, 9, 9, 1, 2, 3}, want: 9},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := majorityElement(tc.nums)
			if got != tc.want {
				t.Fatalf("majorityElement(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
