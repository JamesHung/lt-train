package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "example2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "example3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if len(got) != 2 || got[0] != tc.want[0] || got[1] != tc.want[1] {
				t.Fatalf("twoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
