package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "example1", nums: []int{1, 2, 3, 1}, want: true},
		{name: "example2", nums: []int{1, 2, 3, 4}, want: false},
		{name: "example3", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
		{name: "single", nums: []int{42}, want: false},
		{name: "negatives", nums: []int{-7, -1, -7, 3}, want: true},
		{name: "strictlyIncreasing", nums: []int{-3, -2, -1, 0, 1, 2}, want: false},
	}

	for _, tc := range tests {
		// capture range variable
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := containsDuplicate(tc.nums)
			if got != tc.want {
				t.Fatalf("containsDuplicate(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
