package partitionequalsubsetsum

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"example1", []int{1, 5, 11, 5}, true},
		{"example2", []int{1, 2, 3, 5}, false},
		{"single", []int{2}, false},
		{"twoEqual", []int{3, 3}, true},
		{"balancedWithRepeats", []int{2, 2, 3, 3}, true},
		{"cannot", []int{2, 2, 3, 5, 7}, false},
		{"larger", []int{1, 2, 5}, false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := canPartition(tc.nums)
			if got != tc.want {
				t.Fatalf("canPartition(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
